package router

import (
	"context"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pb "github.com/carefree/api/project/portal/user/v1"
)

type portalUserService struct {
	cli    pb.UserServiceClient
	method string
	req    []byte
}

func (s *portalUserService) handle(ctx context.Context) (proto.Message, error) {
	switch s.method {
	case "GetUser":
		req := &pb.GetUserRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.GetUser(ctx, req)
	default:
		return nil, errors.New("method not implemented")
	}
}
