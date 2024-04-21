package gapi

import (
	"fmt"

	db "github.com/raphaeldiscky/simple-bank/db/sqlc"
	pt "github.com/raphaeldiscky/simple-bank/pb"
	"github.com/raphaeldiscky/simple-bank/token"
	"github.com/raphaeldiscky/simple-bank/util"
	"github.com/raphaeldiscky/simple-bank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pt.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server and setup routing
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
