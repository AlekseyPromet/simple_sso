package models

import (
	"context"

	"github.com/google/uuid"
)

type IStorage interface {
	Logining(context.Context, LoginPass) error
	Register(context.Context, User) (uuid.UUID, error)
}
