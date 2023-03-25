package main

import (
	"github.com/Calmantara/go-dts-07/server"
	_ "github.com/lib/pq"
)

func main() {
	// run http server
	server.NewHttpServer()
}
