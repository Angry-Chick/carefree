package resource

import "github.com/golang/protobuf/proto"

// Resource is a named resource as described on
// https://cloud.google.com/apis/design/resources.
type Resource interface {
	proto.Message
	// GetName returns the relative resource name as described on
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name.
	GetName() string
}
