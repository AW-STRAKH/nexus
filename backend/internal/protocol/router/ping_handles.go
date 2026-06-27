package router

import (
	"context"
	"log"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
)

func RegisterPingHandlers(
	r *Router,
) {

	r.Register(
		nexusproto.MessageType_PING,
		HandlerFunc(
			func(
				ctx context.Context,
				envelope *nexusproto.Envelope,
			) error {

				log.Printf(
					"PING received from %s",
					envelope.SenderPeerId,
				)

				return nil
			},
		),
	)

	r.Register(
		nexusproto.MessageType_PONG,
		HandlerFunc(
			func(
				ctx context.Context,
				envelope *nexusproto.Envelope,
			) error {

				log.Printf(
					"PONG received from %s",
					envelope.SenderPeerId,
				)

				return nil
			},
		),
	)
}
