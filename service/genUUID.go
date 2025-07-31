package service

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	gen := uuid.New()
	id := gen.String()
	return id
}
