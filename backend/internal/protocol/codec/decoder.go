package codec

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "github.com/awatansh/nexus/internal/protocol/proto"
)

func DecodeHello(
	env *pb.Envelope,
) (*pb.Hello, error) {

	// Validate message type first
	if env.Type != pb.MessageType_HELLO {
		return nil, fmt.Errorf(
			"unexpected message type: %v",
			env.Type,
		)
	}

	hello := &pb.Hello{}

	if err := proto.Unmarshal(
		env.Payload,
		hello,
	); err != nil {
		return nil, err
	}

	return hello, nil
}
