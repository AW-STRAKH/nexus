package app

import (
	"context"
	"log"

	"github.com/awatansh/nexus/internal/identity"
	"github.com/awatansh/nexus/internal/transport"
)

type Node struct {
	identityService identity.Service
	transport       transport.Transport
	connectAddress  string
}

func NewNode(
	identityService identity.Service,
	transport transport.Transport,
	connectAddress string,
) *Node {

	return &Node{
		identityService: identityService,
		transport:       transport,
		connectAddress:  connectAddress,
	}
}

func (n *Node) Start(
	ctx context.Context,
) error {

	identity, err :=
		n.identityService.LoadOrCreate(ctx)

	if err != nil {
		return err
	}

	log.Printf(
		"Node started. PeerID=%s",
		identity.PeerID(),
	)

	if err := n.transport.Start(ctx); err != nil {
		return err
	}

	// Temporary code for testing outgoing connections.
	if n.connectAddress != "" {

		go func() {

			conn, err := n.transport.Dial(
				context.Background(),
				n.connectAddress,
			)

			if err != nil {
				log.Printf(
					"failed to connect: %v",
					err,
				)
				return
			}

			log.Printf(
				"connected to peer: %s",
				conn.ID(),
			)
		}()
	}

	return nil
}
