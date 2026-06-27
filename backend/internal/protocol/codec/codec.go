package codec

import (
	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
	"google.golang.org/protobuf/proto"
)

func Encode(
	envelope *nexusproto.Envelope,
) ([]byte, error) {

	return proto.Marshal(
		envelope,
	)
}

func Decode(
	data []byte,
) (*nexusproto.Envelope, error) {

	envelope := &nexusproto.Envelope{}

	if err := proto.Unmarshal(
		data,
		envelope,
	); err != nil {

		return nil, err
	}

	return envelope, nil
}
