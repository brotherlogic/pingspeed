package main

import (
	"log"
	"os"
)

func (s *Server) Init() {
	data, err := os.ReadFile("/etc/machine-id")
	log.Printf("Read %v -> %v", string(data), err)
}
