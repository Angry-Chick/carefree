package router

import (
	"context"
	"encoding/json"
	"errors"

	pb "github.com/carefree/api/project/portal/slice/v1"
)

type spaceSliceService struct {
	cli pb.SliceServiceClient
	req []byte
}

func (s *spaceSliceService) handle(ctx context.Context, method string) (resp []byte, err error) {
	var rs interface{}
	switch method {
	case "GetSlice":
		req := struct {
			Name string `json:"name"`
		}{}
		if err := json.Unmarshal(s.req, &req); err != nil {
			return nil, err
		}
		rs, err = s.cli.GetSlice(ctx, &pb.GetSliceRequest{
			Name: req.Name,
		})
	case "CreateSlice":
		req := struct {
			Space string `json:"space"`
			ID    string `json:"id"`
		}{}
		if err := json.Unmarshal(s.req, &req); err != nil {
			return nil, err
		}
		rs, err = s.cli.CreateSlice(ctx, &pb.CreateSliceRequest{
			Space: req.Space,
			Id:    req.ID,
			Slice: &pb.Slice{},
		})
	default:
		return nil, errors.New("unknown mehtod name")
	}
	if err != nil {
		return nil, err
	}
	resp, err = json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
