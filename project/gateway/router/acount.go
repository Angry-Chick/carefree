package router

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/protobuf/proto"

	pb "github.com/carefree/api/project/account/v1"
)

type accountService struct {
	cli    pb.AccountServiceClient
	method string
	req    []byte
}

func (s *accountService) handle(ctx context.Context) (proto.Message, error) {
	switch s.method {
	case "BasicAuth":
		var req = &struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		if err := json.Unmarshal(s.req, req); err != nil {
			return nil, err
		}
		return s.cli.BasicAuth(ctx, &pb.BasicAuthRequest{
			Login: &pb.BasicAuthRequest_Username{
				Username: req.Username,
			},
			Password: req.Password,
		})
	default:
		return nil, errors.New("unknown")
	}
}
