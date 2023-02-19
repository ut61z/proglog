package main

import (
	"log"

	"github.com/ut61z/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
