package main

import (
	"flag"
	"fmt"
	"os"

	r "github.com/kshwedha/teknight_bc/"
)

func main() {
	// environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	// config.Init(*environment)
	server.Init()
}
