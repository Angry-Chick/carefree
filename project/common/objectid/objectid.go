// Package objectid defines an ID format for identifying resources, and provides
// methods to manipulate it.
package objectid

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// ID is an v4 UUID
type ID uuid.UUID

// New returns an initialized ID.
func New() (ID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return ID{}, err
	}
	return ID(u), nil
}

// MustNew returns id if err is nil and panics otherwise.
func MustNew() ID {
	u, err := New()
	if err != nil {
		panic(err)
	}
	return u
}

// Base64 returns the base64 encoded uuid.
func (id ID) Base64() string {
	return base64.URLEncoding.EncodeToString([]byte(id[:]))
}

// Base64 return a base64 encoded random uuid.
func Base64() (string, error) {
	u, err := New()
	if err != nil {
		return "", err
	}
	return u.Base64(), nil
}
