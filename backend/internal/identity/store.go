package identity

import (
	"context"
	"errors"
)

var ErrIdentityNotFound = errors.New("identity not found")

type Store interface {
	Load(
		ctx context.Context,
	) (*Identity, error)

	Save(
		ctx context.Context,
		identity *Identity,
	) error
}
