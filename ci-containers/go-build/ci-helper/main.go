package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Loading config...")
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Println("Payload:")
	fmt.Print(cfg.ExtractPayload())
}
