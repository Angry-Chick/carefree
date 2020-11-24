package router

import (
	"context"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pb "github.com/carefree/api/project/portal/v1"
)

type portalService struct {
	cli    pb.PortalServiceClient
	method string
	req    []byte
}

func (s *portalService) handle(ctx context.Context) (proto.Message, error) {
	switch s.method {
	case "SignUp":
		req := &pb.SignUpRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.SignUp(ctx, req)

	default:
		return nil, errors.New("unknown method")
	}
}
