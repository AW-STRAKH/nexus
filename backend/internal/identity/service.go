package identity

import (
	"context"
	"crypto/ed25519"
)

type Service interface {
	LoadOrCreate(ctx context.Context) (*Identity, error)

	Sign(
		ctx context.Context,
		data []byte,
	) ([]byte, error)

	Verify(
		ctx context.Context,
		data []byte,
		signature []byte,
		publicKey ed25519.PublicKey,
	) bool

	GetIdentity() *Identity
}
