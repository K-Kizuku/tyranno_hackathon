package repository

import (
	"context"
	"tyranno/backend/domain/model"

	"github.com/uptrace/bun"
)

type IUserRepository interface {
	GetUserByID(ctx *context.Context) (*model.User, error)
	CreateUser(ctx *context.Context) (*model.User, error)
	LoginUser(ctx *context.Context) (*model.User, error)
}
type User struct {
	repo IUserRepository
}

func (u *User) CreateUser(ctx context.Context, db *bun.DB) {

}
