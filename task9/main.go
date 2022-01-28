package main

import (
	"fmt"
	"os"

	"github.com/Leonid-Nenkin/go-level1-hw/task9/configuration"
)

func main() {
	var cfg Configuration

	config, err := configuration.Load(cfg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
