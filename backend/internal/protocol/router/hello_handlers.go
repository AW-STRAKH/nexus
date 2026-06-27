package router

import (
	"context"
	"log"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
	"google.golang.org/protobuf/proto"
)

func RegisterHelloHandlers(
	r *Router,
) {

	r.Register(
		nexusproto.MessageType_HELLO,
		HandlerFunc(
			func(
				ctx context.Context,
				envelope *nexusproto.Envelope,
			) error {

				hello := &nexusproto.Hello{}

				if err := proto.Unmarshal(
					envelope.Payload,
					hello,
				); err != nil {

					return err
				}

				log.Printf(
					"HELLO received from peer=%s client=%s capabilities=%v",
					hello.PeerId,
					hello.ClientName,
					hello.Capabilities,
				)

				return nil
			},
		),
	)

	r.Register(
		nexusproto.MessageType_HELLO_ACK,
		HandlerFunc(
			func(
				ctx context.Context,
				envelope *nexusproto.Envelope,
			) error {

				ack := &nexusproto.HelloAck{}

				if err := proto.Unmarshal(
					envelope.Payload,
					ack,
				); err != nil {

					return err
				}

				log.Printf(
					"HELLO_ACK received from peer=%s capabilities=%v",
					ack.PeerId,
					ack.Capabilities,
				)

				return nil
			},
		),
	)
}
