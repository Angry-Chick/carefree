package db

// This file defines methods to marshal proto.Messages to DB bytes, and
// unmarshal from DB bytes.

import (
	"google.golang.org/protobuf/proto"
)

func Marshal(m proto.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func Unmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}
