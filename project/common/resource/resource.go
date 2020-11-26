package resource

import "google.golang.org/protobuf/proto"

type Resource interface {
	proto.Message
	GetName() string
}
