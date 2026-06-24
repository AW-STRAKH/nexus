package quic

import (
	"context"

	quicgo "github.com/quic-go/quic-go"

	transportpkg "github.com/awatansh/nexus/internal/transport"
)

func (t *Transport) Dial(
	ctx context.Context,
	address string,
) (transportpkg.Connection, error) {

	session, err := quicgo.DialAddr(
		ctx,
		address,
		generateClientTLSConfig(),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return NewConnection(session), nil
}
