package random

import "github.com/google/uuid"

func NewUUIDString() string {
	return uuid.NewString()
}
