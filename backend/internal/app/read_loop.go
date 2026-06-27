package app

import (
	"context"
	"log"

	"github.com/awatansh/nexus/internal/protocol/codec"
	"github.com/awatansh/nexus/internal/transport"
)

func (n *Node) startReadLoop(
	ctx context.Context,
	conn transport.Connection,
) {

	go func() {

		for {

			data, err := conn.Receive(
				ctx,
			)

			if err != nil {

				log.Printf(
					"connection closed: %v",
					err,
				)

				return
			}

			envelope, err :=
				codec.Decode(data)

			if err != nil {

				log.Printf(
					"failed to decode envelope: %v",
					err,
				)

				continue
			}

			if err := n.router.Dispatch(
				ctx,
				envelope,
			); err != nil {

				log.Printf(
					"dispatch failed: %v",
					err,
				)
			}
		}
	}()
}
