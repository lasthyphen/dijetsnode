// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package manager

import (
	"github.com/lasthyphen/dijetsnode/database"
	"github.com/lasthyphen/dijetsnode/utils"
	"github.com/lasthyphen/dijetsnode/version"
)

var _ utils.Sortable[*VersionedDatabase] = (*VersionedDatabase)(nil)

type VersionedDatabase struct {
	Database database.Database
	Version  *version.Semantic
}

// Close the underlying database
func (db *VersionedDatabase) Close() error {
	return db.Database.Close()
}

// Note this sorts in descending order (newest version --> oldest version)
func (db *VersionedDatabase) Less(other *VersionedDatabase) bool {
	return db.Version.Compare(other.Version) > 0
}
