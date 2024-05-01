package uuid

import "github.com/google/uuid"

type V4 string

// New Return a new uuid string version 4
func New() string {
	return uuid.NewString()
}

// IsValid return a valid uuid
func IsValid(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func (v V4) String() string {
	return string(v)
}
