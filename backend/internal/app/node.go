package app

import (
	"context"
	"log"

	"github.com/awatansh/nexus/internal/identity"
	"github.com/awatansh/nexus/internal/peer"
	"github.com/awatansh/nexus/internal/protocol/handshake"
	"github.com/awatansh/nexus/internal/protocol/heartbeat"
	"github.com/awatansh/nexus/internal/protocol/router"
	"github.com/awatansh/nexus/internal/transport"
)

type Node struct {
	identityService identity.Service
	transport       transport.Transport
	peerManager     *peer.Manager
	router          *router.Router

	heartbeatService *heartbeat.Service

	connectAddress string
}

func NewNode(
	identityService identity.Service,
	transport transport.Transport,
	connectAddress string,
) *Node {

	r := router.NewRouter()

	router.RegisterDefaultHandlers(
		r,
	)

	router.RegisterHelloHandlers(
		r,
	)

	router.RegisterPingHandlers(
		r,
	)

	node := &Node{
		identityService: identityService,
		transport:       transport,
		peerManager:     peer.NewManager(),
		router:          r,
		connectAddress:  connectAddress,
	}

	node.heartbeatService = heartbeat.NewService(
		node.peerManager,
		node,
	)

	return node
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

			manager := handshake.NewManager()

			if err := manager.SendHello(
				context.Background(),
				conn,
				identity.PeerID(),
			); err != nil {

				log.Printf(
					"failed to send hello: %v",
					err,
				)
				return
			}

			log.Printf(
				"HELLO sent to %s",
				conn.ID(),
			)

			ack, err := manager.ReceiveHelloAck(
				context.Background(),
				conn,
			)

			if err != nil {
				log.Printf(
					"failed to receive hello ack: %v",
					err,
				)
				return
			}

			log.Printf(
				"received HELLO_ACK from peer %s capabilities=%v",
				ack.PeerId,
				ack.Capabilities,
			)

			n.peerManager.Add(
				ack.PeerId,
				conn,
				ack.Capabilities,
			)

			log.Printf(
				"peer added. total peers=%d",
				n.peerManager.Count(),
			)

			n.startReadLoop(
				context.Background(),
				conn,
			)
		}()
	}

	n.heartbeatService.Start(
		ctx,
	)

	return nil
}
