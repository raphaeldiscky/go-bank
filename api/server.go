package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/raphaeldiscky/simple-bank.git/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store db.Store
	router *gin.Engine // send requests to the correct handler for processing
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
    
	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}


// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}