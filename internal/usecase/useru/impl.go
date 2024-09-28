package useru

import (
	"context"
	"fmt"
	"log"
	"time"

	"test-code/internal/domain/userd"
	"test-code/internal/utils"
)

type UseUsecase struct {
	userRepo userd.UserRepo
}

func NewUseUsecase(userRepo userd.UserRepo) UserUsc {
	return &UseUsecase{userRepo: userRepo}
}

func (u *UseUsecase) CreateUser(ctx context.Context, params UserCreateReq) (err error) {
	tn := time.Now()
	err = u.userRepo.CreateUser(ctx, userd.UserEntity{
		// ID:         params.ID,
		RoleID:     params.RoleID,
		Email:      params.Email,
		Password:   params.Password,
		Name:       params.Name,
		LastAccess: nil,
		Created_at: &tn,
		Updated_at: &tn,
	})
	if err != nil {
		log.Println("error at CreateUser : ", err.Error())
	}

	return
}

func (u *UseUsecase) Login(ctx context.Context, params LoginReq) (token string, err error) {
	resUser, err := u.userRepo.GetUserByEmail(ctx, params.Email)
	if err != nil {
		log.Println("error at Login : ", err.Error())
		return
	}

	if resUser.Password != params.Password {
		err = fmt.Errorf("invalid password")
		log.Println("error at Login : ", err.Error())
		return
	}

	createToken := utils.NewToken(utils.DataClaims{
		ID:         resUser.ID,
		RoleID:     resUser.RoleID,
		Email:      resUser.Email,
		Password:   resUser.Password,
		Name:       resUser.Name,
		LastAccess: resUser.LastAccess,
	})

	token, err = createToken.Create()
	if err != nil {
		log.Println("error at Login : ", err.Error())
		return
	}
	return
}

func (u *UseUsecase) GetUser(ctx context.Context) (res []GetListUser, err error) {
	resUsers, err := u.userRepo.GetUser(ctx)
	if err != nil {
		log.Println("error at GetUser : ", err.Error())
		return
	}

	if len(resUsers) > 0 {
		for _, v := range resUsers {
			res = append(res, GetListUser{
				ID:         v.ID,
				RoleID:     v.RoleID,
				RoleName:   v.Role_name,
				Email:      v.Email,
				Password:   v.Password,
				Name:       v.Name,
				LastAccess: v.LastAccess,
			})
		}
	}
	return
}

func (u *UseUsecase) UpdateUser(ctx context.Context, params UserUpdate) (err error) {
	err = u.userRepo.UpdateUser(ctx, userd.UserEntity{ID: params.ID, Name: params.Name})
	if err != nil {
		log.Println("error at UpdateUser : ", err.Error())
		return
	}

	return
}

func (u *UseUsecase) DeleteUser(ctx context.Context, id int) (err error) {
	err = u.userRepo.DeleteUser(ctx, id)
	if err != nil {
		log.Println("error at DeleteUser : ", err.Error())
		return
	}

	return
}

func (u *UseUsecase) CheckerValidRole(ctx context.Context, params CheckValidRoleReq) (err error) {
	resRoleRight, err := u.userRepo.GetRoleRight(ctx, params.RoleID)
	if err != nil {
		log.Println("error at CheckerValidRole : ", err.Error())
		return
	}

	var isValid bool = false

	switch params.Method {
	case "create":
		if resRoleRight.R_create == 1 {
			isValid = true
		}
	case "read":
		if resRoleRight.R_read == 1 {
			isValid = true
		}
	case "update":
		if resRoleRight.R_update == 1 {
			isValid = true
		}
	case "delete":
		if resRoleRight.R_delete == 1 {
			isValid = true
		}
	}

	if !isValid {
		err = fmt.Errorf("invalid role")
	}

	return
}
