package identity

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
)

type service struct {
	store Store

	identity *Identity
}

func NewService(
	store Store,
) Service {

	return &service{
		store: store,
	}
}

func (s *service) LoadOrCreate(
	ctx context.Context,
) (*Identity, error) {

	identity, err := s.store.Load(ctx)

	if err == nil {
		s.identity = identity
		return identity, nil
	}

	if !errors.Is(err, ErrIdentityNotFound) {
		return nil, err
	}

	publicKey, privateKey, err :=
		ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return nil, err
	}

	identity = &Identity{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	if err := s.store.Save(
		ctx,
		identity,
	); err != nil {
		return nil, err
	}

	s.identity = identity

	return identity, nil
}

func (s *service) GetIdentity() *Identity {
	return s.identity
}

func (s *service) Sign(
	ctx context.Context,
	data []byte,
) ([]byte, error) {

	if s.identity == nil {
		return nil,
			errors.New(
				"identity not initialized",
			)
	}

	signature := ed25519.Sign(
		s.identity.PrivateKey,
		data,
	)

	return signature, nil
}

func (s *service) Verify(
	ctx context.Context,
	data []byte,
	signature []byte,
	publicKey ed25519.PublicKey,
) bool {

	return ed25519.Verify(
		publicKey,
		data,
		signature,
	)
}
