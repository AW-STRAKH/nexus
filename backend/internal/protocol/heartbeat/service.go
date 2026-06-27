package heartbeat

import (
	"context"
	"log"
	"time"

	"github.com/awatansh/nexus/internal/peer"
)

type Sender interface {
	SendPing(
		ctx context.Context,
		peerID string,
	) error
}

type Service struct {
	peerManager *peer.Manager
	sender      Sender
	interval    time.Duration
}

func NewService(
	peerManager *peer.Manager,
	sender Sender,
) *Service {

	return &Service{
		peerManager: peerManager,
		sender:      sender,
		interval:    5 * time.Second, // shorter interval for testing
	}
}

func (s *Service) Start(
	ctx context.Context,
) {

	log.Printf(
		"heartbeat service started",
	)

	ticker := time.NewTicker(
		s.interval,
	)

	go func() {

		defer ticker.Stop()

		for {

			select {

			case <-ctx.Done():

				log.Printf(
					"heartbeat service stopped",
				)

				return

			case <-ticker.C:

				connectedPeers :=
					s.peerManager.ConnectedPeers()

				log.Printf(
					"heartbeat tick: connected peers=%d",
					len(connectedPeers),
				)

				for _, p := range connectedPeers {

					log.Printf(
						"sending ping to %s",
						p.PeerID,
					)

					if err := s.sender.SendPing(
						ctx,
						p.PeerID,
					); err != nil {

						log.Printf(
							"failed to ping %s: %v",
							p.PeerID,
							err,
						)
					}
				}
			}
		}
	}()
}
