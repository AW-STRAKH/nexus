package router

import (
	"context"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
)

type Handler interface {
	Handle(
		ctx context.Context,
		envelope *nexusproto.Envelope,
	) error
}
