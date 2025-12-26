package main

import (
	"log"
	"time"

	"github.com/thayline/stress-test/internal/config"
	"github.com/thayline/stress-test/internal/executor"
	"github.com/thayline/stress-test/internal/report"
)

func main() {
	cfg := config.Load()

	if cfg.URL == "" {
		log.Fatal("[FATAL] URL is required")
	}

	start := time.Now()
	results := executor.Run(cfg.URL, cfg.Requests, cfg.Concurrency)
	duration := time.Since(start)

	report.Print(results, duration)
}
