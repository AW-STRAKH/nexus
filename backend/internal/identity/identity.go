package identity

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
)

type Identity struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

func (i *Identity) PeerID() string {
	hash := sha256.Sum256(i.PublicKey)

	return hex.EncodeToString(hash[:])
}
