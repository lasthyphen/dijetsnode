// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gkeystore

import (
	"context"

	"google.golang.org/grpc"

	"github.com/lasthyphen/dijetsnode/api/keystore"
	"github.com/lasthyphen/dijetsnode/database"
	"github.com/lasthyphen/dijetsnode/database/rpcdb"
	"github.com/lasthyphen/dijetsnode/vms/rpcchainvm/grpcutils"

	keystorepb "github.com/lasthyphen/dijetsnode/proto/pb/keystore"
	rpcdbpb "github.com/lasthyphen/dijetsnode/proto/pb/rpcdb"
)

var _ keystorepb.KeystoreServer = (*Server)(nil)

// Server is a snow.Keystore that is managed over RPC.
type Server struct {
	keystorepb.UnsafeKeystoreServer
	ks keystore.BlockchainKeystore
}

// NewServer returns a keystore connected to a remote keystore
func NewServer(ks keystore.BlockchainKeystore) *Server {
	return &Server{
		ks: ks,
	}
}

func (s *Server) GetDatabase(
	_ context.Context,
	req *keystorepb.GetDatabaseRequest,
) (*keystorepb.GetDatabaseResponse, error) {
	db, err := s.ks.GetRawDatabase(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	closer := dbCloser{Database: db}

	// start the db server
	serverListener, err := grpcutils.NewListener()
	if err != nil {
		return nil, err
	}
	serverAddr := serverListener.Addr().String()

	go grpcutils.Serve(serverListener, func(opts []grpc.ServerOption) *grpc.Server {
		server := grpcutils.NewDefaultServer(opts)
		closer.closer.Add(server)
		db := rpcdb.NewServer(&closer)
		rpcdbpb.RegisterDatabaseServer(server, db)
		return server
	})
	return &keystorepb.GetDatabaseResponse{ServerAddr: serverAddr}, nil
}

type dbCloser struct {
	database.Database
	closer grpcutils.ServerCloser
}

func (db *dbCloser) Close() error {
	err := db.Database.Close()
	db.closer.Stop()
	return err
}
