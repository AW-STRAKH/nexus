package peer

import (
	"sync"
	"time"

	"github.com/awatansh/nexus/internal/transport"
)

type Manager struct {
	mu    sync.RWMutex
	peers map[string]*Peer
}

func NewManager() *Manager {

	return &Manager{
		peers: make(map[string]*Peer),
	}
}

func (m *Manager) Add(
	peerID string,
	connection transport.Connection,
	capabilities []string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()

	m.peers[peerID] = &Peer{
		PeerID:       peerID,
		Connection:   connection,
		Capabilities: capabilities,
		State:        StateConnected,
		ConnectedAt:  now,
		LastSeenAt:   now,
	}
}

func (m *Manager) Remove(
	peerID string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(
		m.peers,
		peerID,
	)
}

func (m *Manager) Get(
	peerID string,
) (*Peer, bool) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	peer, found :=
		m.peers[peerID]

	return peer, found
}

func (m *Manager) List() []*Peer {

	m.mu.RLock()
	defer m.mu.RUnlock()

	peers := make(
		[]*Peer,
		0,
		len(m.peers),
	)

	for _, peer := range m.peers {
		peers = append(
			peers,
			peer,
		)
	}

	return peers
}

func (m *Manager) Count() int {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.peers)
}

func (m *Manager) MarkSeen(
	peerID string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	peer, found :=
		m.peers[peerID]

	if !found {
		return
	}

	peer.LastSeenAt = time.Now()
}

func (m *Manager) MarkDisconnected(
	peerID string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	peer, found :=
		m.peers[peerID]

	if !found {
		return
	}

	peer.State = StateDisconnected
}

func (m *Manager) ConnectedPeers() []*Peer {

	m.mu.RLock()
	defer m.mu.RUnlock()

	var peers []*Peer

	for _, peer := range m.peers {

		if peer.State != StateConnected {
			continue
		}

		peers = append(
			peers,
			peer,
		)
	}

	return peers
}

func (m *Manager) GetConnection(
	peerID string,
) (transport.Connection, bool) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	peer, found :=
		m.peers[peerID]

	if !found {
		return nil, false
	}

	return peer.Connection, true
}
