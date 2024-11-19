package main

import (
	"log"
	"my-first-api/internal/todo"
	"my-first-api/internal/transport"
)

func main() {

	svc := todo.NewService()
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
