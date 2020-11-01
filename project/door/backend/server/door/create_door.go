package door

import (
	"context"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"
	"github.com/carefree/project/door/datamodel/door"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/door/v1/door"
)

func (s *Server) CreateDoor(ctx context.Context, req *pb.CreateDoorRequest) (*pb.Door, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDoor not implemented")
}

type createProc struct {
	req   *pb.CreateDoorRequest
	resp  *pb.Door
	doors *door.Resources
}

func (p *createProc) Reset(ctx context.Context, db *db.DB) error {
	p.doors = door.New(db)
	p.resp = nil
	return nil
}

func (p *createProc) Commit(db *db.DB) error {
	return db.Commit()
}

func (p *createProc) Do() (err error) {
	id := p.req.Id
	if id == "" {
		if id, err = objectid.Base64(); err != nil {
			return err
		}
	} else if err := door.CheckID(id); err != nil {
		return err
	}
	n := &pb.Door{
		Name: door.FullName(id),
	}
	resp, err := p.doors.Create(n)
	if err != nil {
		return err
	}
	p.resp = resp
	return
}
