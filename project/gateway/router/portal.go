package router

import (
	"context"
	"encoding/json"
	"errors"

	pb "github.com/carefree/api/project/portal/v1"
)

type portalService struct {
	cli pb.PortalServiceClient
	req []byte
}

func (s *portalService) handle(ctx context.Context, method string) ([]byte, error) {
	switch method {
	case "SignUp":
		req := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		if err := json.Unmarshal(s.req, &req); err != nil {
			return nil, err
		}
		rs, err := s.cli.SignUp(ctx, &pb.SignUpRequest{
			Password: req.Password,
			Username: req.Username,
		})
		if err != nil {
			return nil, err
		}
		resp, err := json.Marshal(rs)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, errors.New("unknown method")
	}
}
