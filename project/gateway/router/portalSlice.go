package router

import (
	"context"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pb "github.com/carefree/api/project/portal/slice/v1"
)

type portalSliceService struct {
	cli    pb.SliceServiceClient
	method string
	req    []byte
}

func (s *portalSliceService) handle(ctx context.Context) (proto.Message, error) {
	switch s.method {
	case "GetSlice":
		req := &pb.GetSliceRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.GetSlice(ctx, req)
	case "UpdateSlice":
		req := &pb.UpdateSliceRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.UpdateSlice(ctx, req)
	case "CreateSlice":
		req := &pb.CreateSliceRequest{}
		if err := protojson.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.CreateSlice(ctx, req)
	default:
		return nil, errors.New("unknown mehtod name")
	}
}
