package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	db "github.com/raphaeldiscky/simple-bank/db/sqlc"
	"github.com/raphaeldiscky/simple-bank/token"
	"github.com/raphaeldiscky/simple-bank/util"
	"github.com/raphaeldiscky/simple-bank/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func NewTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, username string, role string, duration time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(username, role, duration)
	require.NoError(t, err)
	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}
	return metadata.NewIncomingContext(context.Background(), md)
}
