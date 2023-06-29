// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"math"

	"github.com/lasthyphen/dijetsnode/codec"
	"github.com/lasthyphen/dijetsnode/codec/linearcodec"
	"github.com/lasthyphen/dijetsnode/utils/wrappers"
)

const codecVersion = 0

// Codec does serialization and deserialization for Teleporter messages.
var c codec.Manager

func init() {
	c = codec.NewManager(math.MaxInt)
	lc := linearcodec.NewCustomMaxLength(math.MaxInt32)

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&BitSetSignature{}),
		c.RegisterCodec(codecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
