package main

import (
	"github.com/Runacy/distributed-service-go/proglog/sample-http-server/internal/server"
)

func main() {
	server.NewHTTPServer(":8880")
	// srv := server.NewHTTPServer(":8880")
	// log.Fatal(srv.ListenAndServe())
}
