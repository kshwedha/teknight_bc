package main

import (
	"flag"
	"fmt"
	"os"

	"teknight_bc/config"
	"teknight_bc/src/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Serve()
}
