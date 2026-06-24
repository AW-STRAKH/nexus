package quic

import "time"

type Config struct {
	ListenAddress      string
	ConnectionTimeout  time.Duration
	KeepAliveInterval  time.Duration
	MaxIncomingStreams int64
}

func DefaultConfig() Config {
	return Config{
		ListenAddress:      ":9000",
		ConnectionTimeout:  10 * time.Second,
		KeepAliveInterval:  30 * time.Second,
		MaxIncomingStreams: 100,
	}
}
