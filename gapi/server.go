package gapi

import (
	"fmt"

	db "github.com/raphaeldiscky/simple-bank.git/db/sqlc"
	pt "github.com/raphaeldiscky/simple-bank.git/pb"
	"github.com/raphaeldiscky/simple-bank.git/token"
	"github.com/raphaeldiscky/simple-bank.git/utils"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pt.UnimplementedSimpleBankServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server and setup routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
