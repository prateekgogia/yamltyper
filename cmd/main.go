package main

import (
	"github.com/prateekgogia/yamltyper/pkg/frontend"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting yamltyper")
	if err := frontend.Run(); err != nil {
		log.Fatalf("Failed to run err: %v", err)
	}
}
