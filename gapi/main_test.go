package gapi

import (
	"testing"
	"time"

	db "github.com/raphaeldiscky/simple-bank/db/sqlc"
	"github.com/raphaeldiscky/simple-bank/utils"
	"github.com/raphaeldiscky/simple-bank/worker"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
