package user

import (
	"fmt"
	"path"
	"regexp"

	"github.com/carefree/project/common/db"

	pb "github.com/carefree/api/project/home/admin/user/v1"
)

func ToResource(r *db.Row) (*pb.User, error) {
	var res pb.User
	if err := db.Unmarshal(r.Resource, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

const idPattern = `[\w-]{4,32}`

var idRE = regexp.MustCompile(fmt.Sprintf("^%s$", idPattern))

func CheckID(id string) error {
	if idRE.MatchString(id) {
		return nil
	}
	return fmt.Errorf("invalid BankAccount id: %q not match %q", id, idPattern)
}

func FullName(namespace, id string) string {
	return path.Join(namespace, "users", id)
}

type Resources struct {
	db *db.DB
}

func New(db *db.DB) *Resources {
	return &Resources{db: db}
}

func (r Resources) Get(name string) (*pb.User, error) {
	row, err := r.db.Get(name)
	if err != nil {
		return nil, err
	}
	return ToResource(row)
}

func (r Resources) Update(res *pb.User) (*pb.User, error) {
	row, err := r.db.Update(res)
	if err != nil {
		return nil, err
	}
	return ToResource(row)
}

func (r Resources) Create(res *pb.User) (*pb.User, error) {
	row, err := r.db.Create(res)
	if err != nil {
		return nil, err
	}
	return ToResource(row)
}

func (r Resources) Delete(name string) error {
	return r.db.Delete(name)
}
