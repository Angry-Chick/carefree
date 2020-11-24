package slice

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"
	"github.com/carefree/project/portal/datamodel/slice"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/portal/slice/v1"
)

// Server implement carefree Door service.
type Server struct {
	db *db.DB
}

// NewServer returns an Door service.
func NewServer(db *db.DB) *Server {
	return &Server{db: db}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterSliceServiceServer(svr.GRPC, s)
}

func (s *Server) CreateSlice(ctx context.Context, req *pb.CreateSliceRequest) (*pb.Slice, error) {
	rs := req.Slice
	var id string
	if id == "" {
		id = objectid.MustNew().Base64()
	}
	resp := &pb.Slice{
		Name:       slice.FullName(req.Space, id),
		Background: rs.Background,
		Bookmarks:  rs.Bookmarks,
	}
	slices := slice.New(s.db)
	slices.Create(resp)
	return resp, nil
}

func (s *Server) UpdateSlice(ctx context.Context, req *pb.UpdateSliceRequest) (*pb.Slice, error) {
	rs := req.Slice
	slices := slice.New(s.db)
	ns, err := slices.Get(rs.Name)
	if err != nil {
		return nil, err
	}
	ns.Background = rs.Background
	ns.Bookmarks = rs.Bookmarks
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlice not implemented")
}

func (s *Server) DeleteSlice(ctx context.Context, req *pb.DeleteSliceRequest) (*empty.Empty, error) {
	slices := slice.New(s.db)
	if err := slices.Delete(req.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetSlice(ctx context.Context, req *pb.GetSliceRequest) (*pb.Slice, error) {
	slices := slice.New(s.db)
	rs, err := slices.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
