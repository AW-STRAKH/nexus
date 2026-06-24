package quic

import (
	"context"
	"log"
	"sync"

	quicgo "github.com/quic-go/quic-go"
)

type Transport struct {
	config Config

	listener *quicgo.Listener

	mu      sync.RWMutex
	started bool
}

func NewTransport(
	config Config,
) *Transport {

	return &Transport{
		config: config,
	}
}

func (t *Transport) Start(
	ctx context.Context,
) error {

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.started {
		return nil
	}

	tlsConfig, err := generateTLSConfig()
	if err != nil {
		return err
	}

	listener, err := quicgo.ListenAddr(
		t.config.ListenAddress,
		tlsConfig,
		nil,
	)

	if err != nil {
		return err
	}

	t.listener = listener
	t.started = true

	log.Printf(
		"QUIC server listening on %s",
		t.config.ListenAddress,
	)

	go t.acceptLoop(ctx)

	return nil
}

func (t *Transport) Stop(
	ctx context.Context,
) error {

	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.started {
		return nil
	}

	t.started = false

	if t.listener != nil {
		return t.listener.Close()
	}

	return nil
}
