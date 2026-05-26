package main

import (
	"fmt"
	"log"

	"Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	if err := cfg.SetUser("caio"); err != nil {
		log.Fatalf("failed to set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("failed to re-read config: %v", err)
	}

	fmt.Printf("%+v\n", cfg)
}
