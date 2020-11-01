package db

// This file defines methods to marshal proto.Messages to DB bytes, and
// unmarshal from DB bytes.

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// Marshal marshals a proto.Message to DB bytes.
func Marshal(m proto.Message) ([]byte, error) {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(a)
}

// Unmarshal unmarshals DB bytes to the given proto.Message.
func Unmarshal(b []byte, m proto.Message) error {
	var a any.Any
	if err := proto.Unmarshal(b, &a); err != nil {
		return err
	}
	return ptypes.UnmarshalAny(&a, m)
}

// FromBytes unmarshals DB bytes to a arbitrary registered proto.Message type.
// It's the caller's duty to type assert the returned proto.Message with desired
// types.
func FromBytes(b []byte) (proto.Message, error) {
	var a any.Any
	if err := proto.Unmarshal(b, &a); err != nil {
		return nil, err
	}
	var da ptypes.DynamicAny
	if err := ptypes.UnmarshalAny(&a, &da); err != nil {
		return nil, err
	}
	return da.Message, nil
}
