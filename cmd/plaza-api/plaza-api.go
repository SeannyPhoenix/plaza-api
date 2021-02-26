package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/seannyphoenix/plaza-api/internal/db"
	"github.com/seannyphoenix/plaza-api/internal/server"
)

var mongoURL = "mongodb://mongo.seannyphoenix.com:27017"
var databaseName = "plaza"

func main() {

	// Connect to the database
	client := db.Connect(mongoURL, databaseName)
	defer client.Disconnect(context.Background())

	// Create the server
	httpServer := server.Create()
	// fmt.Printf("%+v\n", httpServer)
	l, err := net.Listen("tcp", httpServer.Addr)
	// fmt.Printf("%+v\n", l)
	if err != nil {
		log.Fatal("Could not listen")
	} else {
		fmt.Printf("Server listening on port %v\n", httpServer.Addr)
	}
	log.Fatal(httpServer.Serve(l))
}
