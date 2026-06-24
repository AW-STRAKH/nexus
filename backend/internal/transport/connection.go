package transport

import "context"

type Connection interface {
	ID() string

	Send(
		ctx context.Context,
		data []byte,
	) error

	Receive(
		ctx context.Context,
	) ([]byte, error)

	Close() error
}
