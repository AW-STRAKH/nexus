package signer

import (
	"crypto/ed25519"

	"google.golang.org/protobuf/proto"

	pb "github.com/awatansh/nexus/internal/protocol/proto"
)

func SignEnvelope(
	env *pb.Envelope,
	privateKey ed25519.PrivateKey,
) error {

	copyEnv := proto.Clone(env).(*pb.Envelope)

	// Signature field must not participate in signing
	copyEnv.Signature = nil

	data, err := proto.Marshal(copyEnv)
	if err != nil {
		return err
	}

	env.Signature = ed25519.Sign(
		privateKey,
		data,
	)

	return nil
}

func VerifyEnvelope(
	env *pb.Envelope,
	publicKey ed25519.PublicKey,
) (bool, error) {

	copyEnv := proto.Clone(env).(*pb.Envelope)

	signature := copyEnv.Signature

	copyEnv.Signature = nil

	data, err := proto.Marshal(copyEnv)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(
		publicKey,
		data,
		signature,
	), nil
}
