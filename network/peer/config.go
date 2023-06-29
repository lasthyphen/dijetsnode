// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/lasthyphen/dijetsnode/ids"
	"github.com/lasthyphen/dijetsnode/message"
	"github.com/lasthyphen/dijetsnode/network/throttling"
	"github.com/lasthyphen/dijetsnode/snow/networking/router"
	"github.com/lasthyphen/dijetsnode/snow/networking/tracker"
	"github.com/lasthyphen/dijetsnode/snow/uptime"
	"github.com/lasthyphen/dijetsnode/snow/validators"
	"github.com/lasthyphen/dijetsnode/utils/logging"
	"github.com/lasthyphen/dijetsnode/utils/set"
	"github.com/lasthyphen/dijetsnode/utils/timer/mockable"
	"github.com/lasthyphen/dijetsnode/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize int
	Clock           mockable.Clock
	Metrics         *Metrics
	MessageCreator  message.Creator

	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	MySubnets            set.Set[ids.ID]
	Beacons              validators.Set
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64

	// Tracks CPU/disk usage caused by each peer.
	ResourceTracker tracker.ResourceTracker

	// Calculates uptime of peers
	UptimeCalculator uptime.Calculator

	// Signs my IP so I can send my signed IP address in the Version message
	IPSigner *IPSigner
}
