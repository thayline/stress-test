package config

import "flag"

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func Load() Config {
	cfg := Config{}

	flag.StringVar(&cfg.URL, "url", "", "URL to test")
	flag.IntVar(&cfg.Requests, "requests", 1, "Total number of requests")
	flag.IntVar(&cfg.Concurrency, "concurrency", 1, "Number of concurrent requests")

	flag.Parse()
	return cfg
}
