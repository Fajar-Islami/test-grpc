package useru

import (
	"context"
)

type UserUsc interface {
	CreateUser(ctx context.Context, params UserCreateReq) (err error)
	Login(ctx context.Context, params LoginReq) (token string, err error)
	GetUser(ctx context.Context) (res []GetListUser, err error)
	UpdateUser(ctx context.Context, params UserUpdate) (err error)
	DeleteUser(ctx context.Context, id int) (err error)

	CheckerValidRole(ctx context.Context, params CheckValidRoleReq) (err error)
}
