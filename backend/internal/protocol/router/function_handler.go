package router

import (
	"context"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
)

type HandlerFunc func(
	ctx context.Context,
	envelope *nexusproto.Envelope,
) error

func (f HandlerFunc) Handle(
	ctx context.Context,
	envelope *nexusproto.Envelope,
) error {

	return f(
		ctx,
		envelope,
	)
}
