package quic

import (
	"context"
	"log"
)

func (t *Transport) acceptLoop(
	ctx context.Context,
) {

	for {

		conn, err := t.listener.Accept(ctx)

		if err != nil {
			log.Printf(
				"failed to accept connection: %v",
				err,
			)
			return
		}

		connection := NewConnection(conn)

		log.Printf(
			"new peer connected: %s",
			connection.ID(),
		)
	}
}
