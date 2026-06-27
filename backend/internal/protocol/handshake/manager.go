package handshake

import (
	"context"
	"fmt"
	"time"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type Connection interface {
	Send(ctx context.Context, data []byte) error
	Receive(ctx context.Context) ([]byte, error)
}

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) SendHello(
	ctx context.Context,
	conn Connection,
	peerID string,
) error {

	hello := &nexusproto.Hello{
		PeerId:          peerID,
		ClientName:      "Nexus",
		ClientVersion:   "0.1.0",
		ProtocolVersion: 1,
		Capabilities: []string{
			"chat",
			"search",
			"file_transfer",
		},
	}

	payload, err := proto.Marshal(hello)
	if err != nil {
		return err
	}

	envelope := &nexusproto.Envelope{
		MessageId:       uuid.NewString(),
		SenderPeerId:    peerID,
		TimestampMs:     time.Now().UnixMilli(),
		Type:            nexusproto.MessageType_HELLO,
		ProtocolVersion: 1,
		Payload:         payload,
	}

	data, err := proto.Marshal(envelope)
	if err != nil {
		return err
	}

	return conn.Send(ctx, data)
}

func (m *Manager) ReceiveHello(
	ctx context.Context,
	conn Connection,
) (*nexusproto.Hello, error) {

	data, err := conn.Receive(ctx)
	if err != nil {
		return nil, err
	}

	envelope := &nexusproto.Envelope{}

	if err := proto.Unmarshal(data, envelope); err != nil {
		return nil, err
	}

	if envelope.Type != nexusproto.MessageType_HELLO {
		return nil, fmt.Errorf(
			"expected HELLO, got %v",
			envelope.Type,
		)
	}

	hello := &nexusproto.Hello{}

	if err := proto.Unmarshal(
		envelope.Payload,
		hello,
	); err != nil {
		return nil, err
	}

	return hello, nil
}

func (m *Manager) SendHelloAck(
	ctx context.Context,
	conn Connection,
	peerID string,
) error {

	ack := &nexusproto.HelloAck{
		PeerId:        peerID,
		ClientName:    "Nexus",
		ClientVersion: "0.1.0",
		Capabilities: []string{
			"chat",
			"search",
			"file_transfer",
		},
	}

	payload, err := proto.Marshal(ack)
	if err != nil {
		return err
	}

	envelope := &nexusproto.Envelope{
		MessageId:       uuid.NewString(),
		SenderPeerId:    peerID,
		TimestampMs:     time.Now().UnixMilli(),
		Type:            nexusproto.MessageType_HELLO_ACK,
		ProtocolVersion: 1,
		Payload:         payload,
	}

	data, err := proto.Marshal(envelope)
	if err != nil {
		return err
	}

	return conn.Send(ctx, data)
}

func (m *Manager) ReceiveHelloAck(
	ctx context.Context,
	conn Connection,
) (*nexusproto.HelloAck, error) {

	data, err := conn.Receive(ctx)
	if err != nil {
		return nil, err
	}

	envelope := &nexusproto.Envelope{}

	if err := proto.Unmarshal(data, envelope); err != nil {
		return nil, err
	}

	if envelope.Type != nexusproto.MessageType_HELLO_ACK {
		return nil, fmt.Errorf(
			"expected HELLO_ACK, got %v",
			envelope.Type,
		)
	}

	ack := &nexusproto.HelloAck{}

	if err := proto.Unmarshal(
		envelope.Payload,
		ack,
	); err != nil {
		return nil, err
	}

	return ack, nil
}
