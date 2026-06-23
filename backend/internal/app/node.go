package app

import (
	"context"
	"log"

	"github.com/awatansh/nexus/internal/identity"
)

type Node struct {
	identityService identity.Service
}

func NewNode(
	identityService identity.Service,
) *Node {

	return &Node{
		identityService: identityService,
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

	return nil
}
