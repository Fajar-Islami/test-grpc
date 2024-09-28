package userd

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserDomain struct {
	pg *sqlx.DB
}

func NewUserDomain(pg *sqlx.DB) UserRepo {
	return &UserDomain{pg: pg}
}

func (d *UserDomain) CreateUser(ctx context.Context, params UserEntity) (err error) {
	_, err = d.pg.NamedExecContext(ctx, `INSERT INTO users (role_id, email, password, name, last_access, created_at, updated_at) VALUES(:role_id, :email, :password, :name, :last_access, :created_at, :updated_at)`, params)
	return
}

func (d *UserDomain) GetUser(ctx context.Context) (res []ListUserEntity, err error) {
	query := `select u.id, u.role_id, r.name as role_name, u.email, u.password, u.name, u.last_access, u.created_at, u.updated_at from users u join roles r
	on u.role_id = r.id
	where u.deleted_at is null`

	err = d.pg.SelectContext(ctx, &res, query)
	return
}

func (d *UserDomain) GetUserByEmail(ctx context.Context, email string) (res UserEntity, err error) {
	query := `select id,role_id, email, password, name, last_access, created_at, updated_at from users where deleted_at is null and email = $1`

	err = d.pg.GetContext(ctx, &res, query, email)
	return
}

func (d *UserDomain) UpdateUser(ctx context.Context, params UserEntity) (err error) {
	query := `update users set name = $1,updated_at = now() where id = $2`

	_, err = d.pg.ExecContext(ctx, query, params.Name, params.ID)
	return
}

func (d *UserDomain) UpdateLastAccess(ctx context.Context, id int) (err error) {
	query := `update users set last_access = now() where id = $1`

	_, err = d.pg.ExecContext(ctx, query, id)
	return
}

func (d *UserDomain) DeleteUser(ctx context.Context, id int) (err error) {
	query := `update users set deleted_at = now() where id = $1`

	_, err = d.pg.ExecContext(ctx, query, id)
	return
}

func (d *UserDomain) GetRoleRight(ctx context.Context, roleid int) (res RoleRight, err error) {
	query := `SELECT id, role_id, route, "section", "path", r_create, r_read, r_update, r_delete, created_at, updated_at FROM role_rights WHERE role_id = $1`

	err = d.pg.GetContext(ctx, &res, query, roleid)
	return
}
