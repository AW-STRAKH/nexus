package codec

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	pb "github.com/awatansh/nexus/internal/protocol/proto"
)

const CurrentProtocolVersion = 1

func EncodeEnvelope(
	senderPeerID string,
	messageType pb.MessageType,
	message proto.Message,
) (*pb.Envelope, error) {

	payload, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &pb.Envelope{
		MessageId:       uuid.NewString(),
		SenderPeerId:    senderPeerID,
		TimestampMs:     time.Now().UnixMilli(),
		Type:            messageType,
		ProtocolVersion: CurrentProtocolVersion,
		Payload:         payload,
	}, nil
}
