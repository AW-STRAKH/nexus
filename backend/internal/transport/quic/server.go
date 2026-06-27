package quic

import (
	"context"
	"log"

	"github.com/awatansh/nexus/internal/protocol/handshake"
)

func (t *Transport) acceptLoop(
	ctx context.Context,
) {

	handshakeManager := handshake.NewManager()

	for {

		session, err := t.listener.Accept(ctx)

		if err != nil {
			log.Printf(
				"failed to accept connection: %v",
				err,
			)
			return
		}

		connection := NewConnection(session)

		go func() {

			hello, err := handshakeManager.ReceiveHello(
				ctx,
				connection,
			)

			if err != nil {
				log.Printf(
					"failed to receive hello: %v",
					err,
				)
				return
			}

			log.Printf(
				"peer connected: PeerID=%s Client=%s Version=%s Capabilities=%v",
				hello.PeerId,
				hello.ClientName,
				hello.ClientVersion,
				hello.Capabilities,
			)

			if err := handshakeManager.SendHelloAck(
				ctx,
				connection,
				hello.PeerId,
			); err != nil {

				log.Printf(
					"failed to send hello ack: %v",
					err,
				)
				return
			}

			log.Printf(
				"HELLO_ACK sent to %s",
				connection.ID(),
			)
		}()
	}
}
