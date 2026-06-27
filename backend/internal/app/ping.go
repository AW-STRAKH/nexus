package app

import (
	"context"

	"github.com/awatansh/nexus/internal/protocol/heartbeat"
)

func (n *Node) SendPing(
	ctx context.Context,
	peerID string,
) error {

	conn, found :=
		n.peerManager.GetConnection(
			peerID,
		)

	if !found {
		return nil
	}

	manager := heartbeat.NewManager()

	return manager.SendPing(
		ctx,
		conn,
		peerID,
	)
}
