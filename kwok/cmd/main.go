package main

import (
	"flag"

	"github.com/ray-project/kuberay/kwok/pkg/server"
)

func main() {
	port := flag.String("port", "8265", "HTTP server port")
	flag.Parse()

	server.SetupJobServer(*port)
}
