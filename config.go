package gotg

import "time"

type Config struct {
	Address string
	Timeout time.Duration
}

var DefaultConfig = Config{
	Address: "127.0.0.1:4458",
	Timeout: 10 * time.Second,
}
