package router

import (
	"context"
	"encoding/json"
	"errors"

	pb "github.com/carefree/api/project/portal/space/v1"
)

type portalSpaceService struct {
	cli pb.SpaceServiceClient
	req []byte
}

func (s *portalSpaceService) handle(ctx context.Context, method string) (resp []byte, err error) {
	var rs interface{}
	switch method {
	case "GetSpace":
		req := struct {
			Name string `json:"name"`
		}{}
		if err := json.Unmarshal(s.req, &req); err != nil {
			return nil, err
		}
		rs, err = s.cli.GetSpace(ctx, &pb.GetSpaceRequest{
			Name: req.Name,
		})
	case "CreateSpace":
		req := struct {
			ID string `json:"id"`
		}{}
		if err := json.Unmarshal(s.req, &req); err != nil {
			return nil, err
		}
		rs, err = s.cli.CreateSpace(ctx, &pb.CreateSpaceRequest{
			Id:    req.ID,
			Space: &pb.Space{},
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
