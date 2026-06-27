package peer

import (
	"testing"
	"time"
)

func TestManagerAddAndGet(
	t *testing.T,
) {

	manager := NewManager()

	manager.Add(
		"peer-1",
		nil,
		[]string{
			"chat",
			"search",
		},
	)

	peer, found :=
		manager.Get(
			"peer-1",
		)

	if !found {
		t.Fatal(
			"expected peer to exist",
		)
	}

	if peer.PeerID != "peer-1" {
		t.Fatalf(
			"expected peer-1, got %s",
			peer.PeerID,
		)
	}
}

func TestManagerRemove(
	t *testing.T,
) {

	manager := NewManager()

	manager.Add(
		"peer-1",
		nil,
		nil,
	)

	manager.Remove(
		"peer-1",
	)

	_, found :=
		manager.Get(
			"peer-1",
		)

	if found {
		t.Fatal(
			"expected peer to be removed",
		)
	}
}

func TestManagerCount(
	t *testing.T,
) {

	manager := NewManager()

	manager.Add(
		"peer-1",
		nil,
		nil,
	)

	manager.Add(
		"peer-2",
		nil,
		nil,
	)

	if manager.Count() != 2 {
		t.Fatalf(
			"expected count=2, got %d",
			manager.Count(),
		)
	}

}

func TestMarkSeen(
	t *testing.T,
) {

	manager := NewManager()

	manager.Add(
		"peer-1",
		nil,
		nil,
	)

	peer, _ := manager.Get(
		"peer-1",
	)

	oldTime := peer.LastSeenAt

	time.Sleep(
		10 * time.Millisecond,
	)

	manager.MarkSeen(
		"peer-1",
	)

	peer, _ = manager.Get(
		"peer-1",
	)

	if !peer.LastSeenAt.After(
		oldTime,
	) {
		t.Fatal(
			"expected LastSeenAt to be updated",
		)
	}
}

func TestMarkDisconnected(
	t *testing.T,
) {

	manager := NewManager()

	manager.Add(
		"peer-1",
		nil,
		nil,
	)

	manager.MarkDisconnected(
		"peer-1",
	)

	peer, _ := manager.Get(
		"peer-1",
	)

	if peer.State != StateDisconnected {
		t.Fatal(
			"expected peer to be disconnected",
		)
	}
}
