package router

import (
	"context"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pb "github.com/carefree/api/project/portal/space/v1"
)

type portalSpaceService struct {
	cli    pb.SpaceServiceClient
	method string
	req    []byte
}

func (s *portalSpaceService) handle(ctx context.Context) (proto.Message, error) {
	switch s.method {
	case "GetSpace":
		req := &pb.GetSpaceRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.GetSpace(ctx, req)
	case "CreateSpace":
		req := &pb.CreateSpaceRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.CreateSpace(ctx, req)
	default:
		return nil, errors.New("unknown mehtod name")
	}
}
