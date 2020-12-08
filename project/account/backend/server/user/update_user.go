package user

import (
	"context"

	"github.com/carefree/project/account/datamodel/user"

	pb "github.com/carefree/api/project/account/user/v1"
)

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	ru := req.User
	rs := user.New(s.db)
	u, err := rs.Get(ru.Name)
	if err != nil {
		return nil, err
	}
	u.AvatarUrl = ru.AvatarUrl
	u.Email = ru.Email
	u.PhoneNumber = ru.PhoneNumber

	// TODO(ljy): 待修改 password 结构之后，修改密码使用单独的接口操作。
	u.Password = ru.Password
	return rs.Update(u)
}
