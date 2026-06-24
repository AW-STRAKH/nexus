package signer

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	pb "github.com/awatansh/nexus/internal/protocol/proto"
)

func TestSignAndVerify(
	t *testing.T,
) {

	pub, priv, err :=
		ed25519.GenerateKey(rand.Reader)

	if err != nil {
		t.Fatal(err)
	}

	env := &pb.Envelope{
		MessageId:       "msg-1",
		SenderPeerId:    "peer-1",
		ProtocolVersion: 1,
	}

	if err := SignEnvelope(
		env,
		priv,
	); err != nil {
		t.Fatal(err)
	}

	ok, err := VerifyEnvelope(
		env,
		pub,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal(
			"signature verification failed",
		)
	}
}

func TestVerifyFailsWhenEnvelopeTampered(
	t *testing.T,
) {

	pub, priv, err :=
		ed25519.GenerateKey(rand.Reader)

	if err != nil {
		t.Fatal(err)
	}

	env := &pb.Envelope{
		MessageId:    "msg-1",
		SenderPeerId: "peer-1",
	}

	if err := SignEnvelope(
		env,
		priv,
	); err != nil {
		t.Fatal(err)
	}

	// Attacker changes sender
	env.SenderPeerId = "evil-peer"

	ok, err := VerifyEnvelope(
		env,
		pub,
	)

	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Fatal(
			"verification should fail",
		)
	}
}
