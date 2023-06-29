// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/lasthyphen/dijetsnode/snow"
	"github.com/lasthyphen/dijetsnode/snow/uptime"
	"github.com/lasthyphen/dijetsnode/utils"
	"github.com/lasthyphen/dijetsnode/utils/timer/mockable"
	"github.com/lasthyphen/dijetsnode/vms/platformvm/config"
	"github.com/lasthyphen/dijetsnode/vms/platformvm/fx"
	"github.com/lasthyphen/dijetsnode/vms/platformvm/reward"
	"github.com/lasthyphen/dijetsnode/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Manager
	Rewards      reward.Calculator
	Bootstrapped *utils.Atomic[bool]
}
