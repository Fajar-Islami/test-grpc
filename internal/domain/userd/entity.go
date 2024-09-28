package userd

import "time"

type UserEntity struct {
	ID         int        `db:"id"`
	RoleID     int        `db:"role_id"`
	Email      string     `db:"email"`
	Password   string     `db:"password"`
	Name       string     `db:"name"`
	LastAccess *time.Time `db:"last_access"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}

type ListUserEntity struct {
	UserEntity
	Role_name string `db:"role_name"`
}

type RoleRight struct {
	ID         int        `db:"id"`
	RoleID     int        `db:"role_id"`
	Route      string     `db:"route"`
	Section    string     `db:"section"`
	Path       string     `db:"path"`
	R_create   int        `db:"r_create"`
	R_read     int        `db:"r_read"`
	R_update   int        `db:"r_update"`
	R_delete   int        `db:"r_delete"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}

// SELECT id, role_id, route, "section", "path", r_create, r_read, r_update, r_delete, created_at, updated_at FROM public.role_rights WHERE id IN (2);
