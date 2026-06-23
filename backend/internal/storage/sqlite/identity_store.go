package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/awatansh/nexus/internal/identity"
)

type IdentityStore struct {
	db *sql.DB
}

func NewIdentityStore(
	db *sql.DB,
) *IdentityStore {

	return &IdentityStore{
		db: db,
	}
}

func (s *IdentityStore) Load(
	ctx context.Context,
) (*identity.Identity, error) {

	query := `
	SELECT
		public_key,
		private_key
	FROM identity
	WHERE id = 1
	`

	row := s.db.QueryRowContext(
		ctx,
		query,
	)

	var publicKey []byte
	var privateKey []byte

	err := row.Scan(
		&publicKey,
		&privateKey,
	)

	if err != nil {

		if errors.Is(
			err,
			sql.ErrNoRows,
		) {
			return nil,
				identity.ErrIdentityNotFound
		}

		return nil, err
	}

	return &identity.Identity{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

func (s *IdentityStore) Save(
	ctx context.Context,
	identityObj *identity.Identity,
) error {

	query := `
	INSERT OR REPLACE INTO identity (
		id,
		public_key,
		private_key,
		created_at
	)
	VALUES (
		1,
		?,
		?,
		CURRENT_TIMESTAMP
	)
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		identityObj.PublicKey,
		identityObj.PrivateKey,
	)

	return err
}
