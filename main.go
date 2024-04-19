package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/raphaeldiscky/simple-bank.git/api"
	db "github.com/raphaeldiscky/simple-bank.git/db/sqlc"
	"github.com/raphaeldiscky/simple-bank.git/gapi"
	pt "github.com/raphaeldiscky/simple-bank.git/pb"
	"github.com/raphaeldiscky/simple-bank.git/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create grpc server:", err)
	}
	grpcServer := grpc.NewServer()
	pt.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}

}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
