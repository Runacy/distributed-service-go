package main

import (
	"log"

	"github.com/Runacy/distributed-service-go/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8880")
	log.Fatal(srv.ListenAndServe())
}
