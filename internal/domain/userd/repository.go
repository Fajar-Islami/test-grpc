package userd

import (
	"context"
)

type UserRepo interface {
	CreateUser(ctx context.Context, params UserEntity) (err error)
	GetUser(ctx context.Context) (res []ListUserEntity, err error)
	GetUserByEmail(ctx context.Context, email string) (res UserEntity, err error)
	UpdateUser(ctx context.Context, params UserEntity) (err error)
	UpdateLastAccess(ctx context.Context, id int) (err error)
	DeleteUser(ctx context.Context, id int) (err error)

	GetRoleRight(ctx context.Context, roleid int) (res RoleRight, err error)
}
