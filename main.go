package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gpaulo00/velocity-test/config"
	"github.com/gpaulo00/velocity-test/db"
	"github.com/gpaulo00/velocity-test/routes"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
