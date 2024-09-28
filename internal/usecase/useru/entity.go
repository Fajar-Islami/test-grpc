package useru

import "time"

type UserCreateReq struct {
	ID       int    `json:"id"`
	RoleID   int    `json:"role_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type GetListUser struct {
	ID         int        `json:"id"`
	RoleID     int        `json:"role_id"`
	RoleName   string     `json:"role_name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Name       string     `json:"name"`
	LastAccess *time.Time `json:"last_access"`
}

type UserLoginReq struct {
	ID       int    `json:"id"`
	RoleID   int    `json:"role_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserUpdate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckValidRoleReq struct {
	RoleID int    `json:"roleID"`
	Method string `json:"method"`
}
