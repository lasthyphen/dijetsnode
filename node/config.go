// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"crypto/tls"
	"time"

	"github.com/lasthyphen/dijetsnode/chains"
	"github.com/lasthyphen/dijetsnode/genesis"
	"github.com/lasthyphen/dijetsnode/ids"
	"github.com/lasthyphen/dijetsnode/nat"
	"github.com/lasthyphen/dijetsnode/network"
	"github.com/lasthyphen/dijetsnode/snow/consensus/avalanche"
	"github.com/lasthyphen/dijetsnode/snow/networking/benchlist"
	"github.com/lasthyphen/dijetsnode/snow/networking/router"
	"github.com/lasthyphen/dijetsnode/snow/networking/sender"
	"github.com/lasthyphen/dijetsnode/snow/networking/tracker"
	"github.com/lasthyphen/dijetsnode/trace"
	"github.com/lasthyphen/dijetsnode/utils/crypto/bls"
	"github.com/lasthyphen/dijetsnode/utils/dynamicip"
	"github.com/lasthyphen/dijetsnode/utils/ips"
	"github.com/lasthyphen/dijetsnode/utils/logging"
	"github.com/lasthyphen/dijetsnode/utils/profiler"
	"github.com/lasthyphen/dijetsnode/utils/set"
	"github.com/lasthyphen/dijetsnode/utils/timer"
	"github.com/lasthyphen/dijetsnode/vms"
)

type IPCConfig struct {
	IPCAPIEnabled      bool     `json:"ipcAPIEnabled"`
	IPCPath            string   `json:"ipcPath"`
	IPCDefaultChainIDs []string `json:"ipcDefaultChainIDs"`
}

type APIAuthConfig struct {
	APIRequireAuthToken bool   `json:"apiRequireAuthToken"`
	APIAuthPassword     string `json:"-"`
}

type APIIndexerConfig struct {
	IndexAPIEnabled      bool `json:"indexAPIEnabled"`
	IndexAllowIncomplete bool `json:"indexAllowIncomplete"`
}

type HTTPConfig struct {
	APIConfig `json:"apiConfig"`
	HTTPHost  string `json:"httpHost"`
	HTTPPort  uint16 `json:"httpPort"`

	HTTPSEnabled bool   `json:"httpsEnabled"`
	HTTPSKey     []byte `json:"-"`
	HTTPSCert    []byte `json:"-"`

	APIAllowedOrigins []string `json:"apiAllowedOrigins"`

	ShutdownTimeout time.Duration `json:"shutdownTimeout"`
	ShutdownWait    time.Duration `json:"shutdownWait"`
}

type APIConfig struct {
	APIAuthConfig    `json:"authConfig"`
	APIIndexerConfig `json:"indexerConfig"`
	IPCConfig        `json:"ipcConfig"`

	// Enable/Disable APIs
	AdminAPIEnabled    bool `json:"adminAPIEnabled"`
	InfoAPIEnabled     bool `json:"infoAPIEnabled"`
	KeystoreAPIEnabled bool `json:"keystoreAPIEnabled"`
	MetricsAPIEnabled  bool `json:"metricsAPIEnabled"`
	HealthAPIEnabled   bool `json:"healthAPIEnabled"`
}

type IPConfig struct {
	IPPort           ips.DynamicIPPort `json:"ip"`
	IPUpdater        dynamicip.Updater `json:"-"`
	IPResolutionFreq time.Duration     `json:"ipResolutionFrequency"`
	// True if we attempted NAT traversal
	AttemptedNATTraversal bool `json:"attemptedNATTraversal"`
	// Tries to perform network address translation
	Nat nat.Router `json:"-"`
}

type StakingConfig struct {
	genesis.StakingConfig
	EnableStaking         bool            `json:"enableStaking"`
	StakingTLSCert        tls.Certificate `json:"-"`
	StakingSigningKey     *bls.SecretKey  `json:"-"`
	DisabledStakingWeight uint64          `json:"disabledStakingWeight"`
	StakingKeyPath        string          `json:"stakingKeyPath"`
	StakingCertPath       string          `json:"stakingCertPath"`
	StakingSignerPath     string          `json:"stakingSignerPath"`
}

type StateSyncConfig struct {
	StateSyncIDs []ids.NodeID `json:"stateSyncIDs"`
	StateSyncIPs []ips.IPPort `json:"stateSyncIPs"`
}

type BootstrapConfig struct {
	// Should Bootstrap be retried
	RetryBootstrap bool `json:"retryBootstrap"`

	// Max number of times to retry bootstrap before warning the node operator
	RetryBootstrapWarnFrequency int `json:"retryBootstrapWarnFrequency"`

	// Timeout before emitting a warn log when connecting to bootstrapping beacons
	BootstrapBeaconConnectionTimeout time.Duration `json:"bootstrapBeaconConnectionTimeout"`

	// Max number of containers in an ancestors message sent by this node.
	BootstrapAncestorsMaxContainersSent int `json:"bootstrapAncestorsMaxContainersSent"`

	// This node will only consider the first [AncestorsMaxContainersReceived]
	// containers in an ancestors message it receives.
	BootstrapAncestorsMaxContainersReceived int `json:"bootstrapAncestorsMaxContainersReceived"`

	// Max time to spend fetching a container and its
	// ancestors while responding to a GetAncestors message
	BootstrapMaxTimeGetAncestors time.Duration `json:"bootstrapMaxTimeGetAncestors"`

	BootstrapIDs []ids.NodeID `json:"bootstrapIDs"`
	BootstrapIPs []ips.IPPort `json:"bootstrapIPs"`
}

type DatabaseConfig struct {
	// Path to database
	Path string `json:"path"`

	// Name of the database type to use
	Name string `json:"name"`

	// Path to config file
	Config []byte `json:"-"`
}

// Config contains all of the configurations of an Avalanche node.
type Config struct {
	HTTPConfig          `json:"httpConfig"`
	IPConfig            `json:"ipConfig"`
	StakingConfig       `json:"stakingConfig"`
	genesis.TxFeeConfig `json:"txFeeConfig"`
	StateSyncConfig     `json:"stateSyncConfig"`
	BootstrapConfig     `json:"bootstrapConfig"`
	DatabaseConfig      `json:"databaseConfig"`

	// Genesis information
	GenesisBytes []byte `json:"-"`
	AvaxAssetID  ids.ID `json:"avaxAssetID"`

	// ID of the network this node should connect to
	NetworkID uint32 `json:"networkID"`

	// Health
	HealthCheckFreq time.Duration `json:"healthCheckFreq"`

	// Network configuration
	NetworkConfig network.Config `json:"networkConfig"`

	GossipConfig sender.GossipConfig `json:"gossipConfig"`

	AdaptiveTimeoutConfig timer.AdaptiveTimeoutConfig `json:"adaptiveTimeoutConfig"`

	BenchlistConfig benchlist.Config `json:"benchlistConfig"`

	ProfilerConfig profiler.Config `json:"profilerConfig"`

	LoggingConfig logging.Config `json:"loggingConfig"`

	PluginDir string `json:"pluginDir"`

	// File Descriptor Limit
	FdLimit uint64 `json:"fdLimit"`

	// Consensus configuration
	ConsensusParams avalanche.Parameters `json:"consensusParams"`

	// Metrics
	MeterVMEnabled bool `json:"meterVMEnabled"`

	// Router that is used to handle incoming consensus messages
	ConsensusRouter          router.Router       `json:"-"`
	RouterHealthConfig       router.HealthConfig `json:"routerHealthConfig"`
	ConsensusShutdownTimeout time.Duration       `json:"consensusShutdownTimeout"`
	// Gossip a container in the accepted frontier every [ConsensusGossipFrequency]
	ConsensusGossipFrequency time.Duration `json:"consensusGossipFreq"`

	TrackedSubnets set.Set[ids.ID] `json:"trackedSubnets"`

	SubnetConfigs map[ids.ID]chains.SubnetConfig `json:"subnetConfigs"`

	ChainConfigs map[string]chains.ChainConfig `json:"-"`
	ChainAliases map[ids.ID][]string           `json:"chainAliases"`

	VMManager vms.Manager `json:"-"`

	// Halflife to use for the processing requests tracker.
	// Larger halflife --> usage metrics change more slowly.
	SystemTrackerProcessingHalflife time.Duration `json:"systemTrackerProcessingHalflife"`

	// Frequency to check the real resource usage of tracked processes.
	// More frequent checks --> usage metrics are more accurate, but more
	// expensive to track
	SystemTrackerFrequency time.Duration `json:"systemTrackerFrequency"`

	// Halflife to use for the cpu tracker.
	// Larger halflife --> cpu usage metrics change more slowly.
	SystemTrackerCPUHalflife time.Duration `json:"systemTrackerCPUHalflife"`

	// Halflife to use for the disk tracker.
	// Larger halflife --> disk usage metrics change more slowly.
	SystemTrackerDiskHalflife time.Duration `json:"systemTrackerDiskHalflife"`

	CPUTargeterConfig tracker.TargeterConfig `json:"cpuTargeterConfig"`

	DiskTargeterConfig tracker.TargeterConfig `json:"diskTargeterConfig"`

	RequiredAvailableDiskSpace         uint64 `json:"requiredAvailableDiskSpace"`
	WarningThresholdAvailableDiskSpace uint64 `json:"warningThresholdAvailableDiskSpace"`

	TraceConfig trace.Config `json:"traceConfig"`

	// See comment on [MinPercentConnectedStakeHealthy] in platformvm.Config
	MinPercentConnectedStakeHealthy map[ids.ID]float64 `json:"minPercentConnectedStakeHealthy"`

	// See comment on [UseCurrentHeight] in platformvm.Config
	UseCurrentHeight bool `json:"useCurrentHeight"`

	// ProvidedFlags contains all the flags set by the user
	ProvidedFlags map[string]interface{} `json:"-"`

	// ChainDataDir is the root path for per-chain directories where VMs can
	// write arbitrary data.
	ChainDataDir string `json:"chainDataDir"`
}
