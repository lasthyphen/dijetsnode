// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/lasthyphen/dijetsnode/ids"
	"github.com/lasthyphen/dijetsnode/vms/components/avax"
)

type UTXOGetter interface {
	GetUTXO(utxoID ids.ID) (*avax.UTXO, error)
}

type UTXOAdder interface {
	AddUTXO(utxo *avax.UTXO)
}

type UTXODeleter interface {
	DeleteUTXO(utxoID ids.ID)
}
