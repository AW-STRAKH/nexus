package codec

import (
	"testing"

	"google.golang.org/protobuf/proto"

	pb "github.com/awatansh/nexus/internal/protocol/proto"
)

func TestDecodeHello(t *testing.T) {

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

	payload, err := proto.Marshal(hello)
	if err != nil {
		t.Fatalf("failed to marshal hello: %v", err)
	}

	env := &pb.Envelope{
		Type:    pb.MessageType_HELLO,
		Payload: payload,
	}

	decoded, err := DecodeHello(env)
	if err != nil {
		t.Fatalf("failed to decode hello: %v", err)
	}

	if decoded.PeerId != hello.PeerId {
		t.Fatalf(
			"expected peer id %s got %s",
			hello.PeerId,
			decoded.PeerId,
		)
	}
}

func TestDecodeHello_InvalidMessageType(
	t *testing.T,
) {

	env := &pb.Envelope{
		Type: pb.MessageType_CHAT,
	}

	_, err := DecodeHello(env)

	if err == nil {
		t.Fatal(
			"expected error for invalid message type",
		)
	}
}
