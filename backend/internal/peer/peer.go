package peer

import (
	"time"

	"github.com/awatansh/nexus/internal/transport"
)

type Peer struct {
	PeerID       string
	Connection   transport.Connection
	Capabilities []string

	State State

	ConnectedAt time.Time
	LastSeenAt  time.Time
}
