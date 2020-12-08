package user

import (
	"fmt"
	"path"
	"regexp"

	"github.com/carefree/project/common/db"

	pb "github.com/carefree/api/project/account/user/v1"
)

// ToResource converts db.Row to proto resource.
func ToResource(r *db.Row) (*pb.User, error) {
	var res pb.User
	if err := db.Unmarshal(r.Resource, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

const idPattern = `[\w-]{4,32}`

var idRE = regexp.MustCompile(fmt.Sprintf("^%s$", idPattern))

// CheckID checks id is valid.
func CheckID(id string) error {
	if idRE.MatchString(id) {
		return nil
	}
	return fmt.Errorf("invalid BankAccount id: %q not match %q", id, idPattern)
}

// FullName returns resource full name.
func FullName(id string) string {
	return path.Join("users", id)
}

// Resources wrapper a series of database's operations.
type Resources struct {
	db *db.DB
}

// New returns a resources.
func New(db *db.DB) *Resources {
	return &Resources{db: db}
}

// Get gets a resource.
func (r Resources) Get(name string) (*pb.User, error) {
	row, err := r.db.Get(name)
	if err != nil {
		return nil, err
	}
	return ToResource(row)
}

// Update updates a resource.
func (r Resources) Update(res *pb.User) (*pb.User, error) {
	row, err := r.db.Update(res)
	if err != nil {
		return nil, err
	}
	return ToResource(row)
}

// Create creates a resource.
func (r Resources) Create(res *pb.User) (*pb.User, error) {
	row, err := r.db.Create(res)
	if err != nil {
		return nil, err
	}
	return ToResource(row)
}

// Delete deletes a resource (soft delete).
func (r Resources) Delete(name string) error {
	return r.db.Delete(name)
}

// Purge deletes a resource (hard delete).
func (r Resources) Purge(name string) error {
	return r.db.Purge(name)
}
