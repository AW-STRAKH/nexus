package quic

import (
	"context"

	quicgo "github.com/quic-go/quic-go"
)

type Connection struct {
	id      string
	session *quicgo.Conn
}

func NewConnection(
	session *quicgo.Conn,
) *Connection {

	return &Connection{
		id:      session.RemoteAddr().String(),
		session: session,
	}
}

func (c *Connection) ID() string {
	return c.id
}

func (c *Connection) Send(
	ctx context.Context,
	data []byte,
) error {

	stream, err := c.session.OpenStreamSync(ctx)
	if err != nil {
		return err
	}

	defer stream.Close()

	_, err = stream.Write(data)

	return err
}

func (c *Connection) Receive(
	ctx context.Context,
) ([]byte, error) {

	stream, err := c.session.AcceptStream(ctx)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 64*1024)

	n, err := stream.Read(buffer)

	// If we received bytes, return them even if the
	// stream was closed afterwards (EOF).
	if n > 0 {
		return buffer[:n], nil
	}

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Connection) Close() error {
	return c.session.CloseWithError(0, "")
}
