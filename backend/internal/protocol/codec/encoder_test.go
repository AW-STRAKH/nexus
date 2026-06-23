package codec

import (
	"testing"

	pb "github.com/awatansh/nexus/internal/protocol/proto"
)

func TestEncodeDecodeHello(
	t *testing.T,
) {

	hello := &pb.Hello{
		PeerId:          "peer-1",
		ClientName:      "Nexus",
		ClientVersion:   "0.1.0",
		ProtocolVersion: 1,
		Capabilities: []string{
			"chat",
			"search",
		},
	}

	env, err := EncodeEnvelope(
		"peer-1",
		pb.MessageType_HELLO,
		hello,
	)

	if err != nil {
		t.Fatalf(
			"failed to encode: %v",
			err,
		)
	}

	decoded, err := DecodeHello(env)

	if err != nil {
		t.Fatalf(
			"failed to decode: %v",
			err,
		)
	}

	if decoded.PeerId != hello.PeerId {
		t.Fatalf(
			"expected peer id %s got %s",
			hello.PeerId,
			decoded.PeerId,
		)
	}

	if decoded.ClientName != hello.ClientName {
		t.Fatalf(
			"expected client name %s got %s",
			hello.ClientName,
			decoded.ClientName,
		)
	}
}
