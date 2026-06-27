package heartbeat

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
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

func (m *Manager) SendPing(
	ctx context.Context,
	conn Connection,
	peerID string,
) error {

	ping := &nexusproto.Ping{
		TimestampMs: time.Now().UnixMilli(),
	}

	payload, err := proto.Marshal(
		ping,
	)
	if err != nil {
		return err
	}

	envelope := &nexusproto.Envelope{
		MessageId:       uuid.NewString(),
		SenderPeerId:    peerID,
		TimestampMs:     time.Now().UnixMilli(),
		Type:            nexusproto.MessageType_PING,
		ProtocolVersion: 1,
		Payload:         payload,
	}

	data, err := proto.Marshal(
		envelope,
	)
	if err != nil {
		return err
	}

	return conn.Send(
		ctx,
		data,
	)
}

func (m *Manager) SendPong(
	ctx context.Context,
	conn Connection,
	peerID string,
) error {

	pong := &nexusproto.Pong{
		TimestampMs: time.Now().UnixMilli(),
	}

	payload, err := proto.Marshal(
		pong,
	)
	if err != nil {
		return err
	}

	envelope := &nexusproto.Envelope{
		MessageId:       uuid.NewString(),
		SenderPeerId:    peerID,
		TimestampMs:     time.Now().UnixMilli(),
		Type:            nexusproto.MessageType_PONG,
		ProtocolVersion: 1,
		Payload:         payload,
	}

	data, err := proto.Marshal(
		envelope,
	)
	if err != nil {
		return err
	}

	return conn.Send(
		ctx,
		data,
	)
}

func (m *Manager) ReceivePing(
	ctx context.Context,
	conn Connection,
) (*nexusproto.Ping, error) {

	data, err := conn.Receive(ctx)
	if err != nil {
		return nil, err
	}

	envelope := &nexusproto.Envelope{}

	if err := proto.Unmarshal(
		data,
		envelope,
	); err != nil {
		return nil, err
	}

	if envelope.Type != nexusproto.MessageType_PING {
		return nil, fmt.Errorf(
			"expected PING got %v",
			envelope.Type,
		)
	}

	ping := &nexusproto.Ping{}

	if err := proto.Unmarshal(
		envelope.Payload,
		ping,
	); err != nil {
		return nil, err
	}

	return ping, nil
}

func (m *Manager) ReceivePong(
	ctx context.Context,
	conn Connection,
) (*nexusproto.Pong, error) {

	data, err := conn.Receive(ctx)
	if err != nil {
		return nil, err
	}

	envelope := &nexusproto.Envelope{}

	if err := proto.Unmarshal(
		data,
		envelope,
	); err != nil {
		return nil, err
	}

	if envelope.Type != nexusproto.MessageType_PONG {
		return nil, fmt.Errorf(
			"expected PONG got %v",
			envelope.Type,
		)
	}

	pong := &nexusproto.Pong{}

	if err := proto.Unmarshal(
		envelope.Payload,
		pong,
	); err != nil {
		return nil, err
	}

	return pong, nil
}
