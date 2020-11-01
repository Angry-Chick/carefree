package door

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/door/v1/door"
)

func (s *Server) UpdateDoor(ctx context.Context, req *pb.UpdateDoorRequest) (*pb.Door, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoor not implemented")
}
