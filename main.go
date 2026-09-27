package main

import (
	"fmt"
	"log"

	"gator/internal/config"
)

func main() {
	// 1. Read the initial config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// 2. Set current user and update config file on disk
	err = cfg.SetUser("Anas")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	// 3. Read the updated config file back and print it
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading updated config: %v", err)
	}

	fmt.Printf("Updated Config struct: %+v\n", updatedCfg)
}