package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"test-code/internal/infrastructrue/container"
	"test-code/internal/usecase/useru"
	"test-code/internal/utils"
	"test-code/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UsersServer struct {
	// Wajib menyertakan embed struct unimplemented dari hasil generate protobuf jika tidak maka akan error
	pb.UnimplementedUsersServer

	UserUsc useru.UserUsc
}

func StartGRPCServer(cont *container.Container) {
	srv := grpc.NewServer()
	var userService UsersServer = UsersServer{UserUsc: cont.UserUsc}

	pb.RegisterUsersServer(srv, userService)

	log.Println("not listen to %d: %v", cont.Apps.HttpPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cont.Apps.HttpPort))
	if err != nil {
		log.Fatalf("could not listen to %d: %v", cont.Apps.HttpPort, err)
	}

	log.Fatal(srv.Serve(l))
}

func (s UsersServer) Login(ctx context.Context, in *pb.LoginReq) (res *pb.LoginRes, err error) {
	token, err := s.UserUsc.Login(ctx, useru.LoginReq{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	res = &pb.LoginRes{
		Status:  true,
		Message: "Successfully",
		Data: &pb.AccessToken{
			AccessToken: token,
		},
	}
	return
}

func (s UsersServer) List(ctx context.Context, in *emptypb.Empty) (res *pb.GetUserRes, err error) {
	// TODO cek role
	err = s.CheckRole(ctx, "read")
	if err != nil {
		return
	}

	// Get User
	listUser, err := s.UserUsc.GetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	var listUserPB = []*pb.UserRes{}
	for _, v := range listUser {
		data := &pb.UserRes{
			RoleId:   int32(v.RoleID),
			RoleName: v.RoleName,
			Name:     v.Name,
			Email:    v.Email,
			Id:       int32(v.ID),
			// LastAccess: v.LastAccess.Format(time.RFC3339),
		}
		if v.LastAccess != nil {
			data.LastAccess = v.LastAccess.Format(time.RFC3339)
		}
		listUserPB = append(listUserPB, data)
	}

	res = &pb.GetUserRes{
		Status:  true,
		Message: "Successfully",
		Data:    listUserPB,
	}
	return
}

func (s UsersServer) Register(ctx context.Context, in *pb.CreateUserReq) (res *pb.DefaultRes, err error) {
	err = s.CheckRole(ctx, "create")
	if err != nil {
		return
	}
	err = s.UserUsc.CreateUser(ctx, useru.UserCreateReq{
		RoleID:   int(in.RoleId),
		Email:    in.Email,
		Password: in.Password,
		Name:     in.Name,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res = &pb.DefaultRes{
		Status:  true,
		Message: "Successfully",
	}
	return
}
func (s UsersServer) Update(ctx context.Context, in *pb.UpdateUserReq) (res *pb.DefaultRes, err error) {
	err = s.CheckRole(ctx, "update")
	if err != nil {
		return
	}
	err = s.UserUsc.UpdateUser(ctx, useru.UserUpdate{
		ID:   int(in.Id),
		Name: in.Name,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res = &pb.DefaultRes{
		Status:  true,
		Message: "Successfully",
	}
	return
}

func (s UsersServer) Delete(ctx context.Context, in *pb.DeleteUserReq) (res *pb.DefaultRes, err error) {
	err = s.CheckRole(ctx, "delete")
	if err != nil {
		return
	}
	err = s.UserUsc.DeleteUser(ctx, int(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res = &pb.DefaultRes{
		Status:  true,
		Message: "Successfully",
	}
	return
}

func (s *UsersServer) CheckJWT(ctx context.Context) (data *utils.DataClaims, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = status.Error(codes.Unauthenticated, "error Unauthenticated")
		return
	}

	token := md["token"]

	fmt.Println("token", token)

	if len(token) < 1 {
		err = status.Error(codes.Unauthenticated, "error Unauthenticated")
		return
	}
	if token[0] == "" {
		err = status.Error(codes.Unauthenticated, "error Unauthenticated")
		return
	}

	extract, err := utils.CheckToken(token[0])
	if err != nil {
		err = status.Error(codes.Unauthenticated, "error Unauthenticated")
		return
	}

	return &extract.DataClaims, nil
}

func (s *UsersServer) CheckRole(ctx context.Context, method string) (err error) {
	data, err := s.CheckJWT(ctx)
	if err != nil {
		err = status.Error(codes.Unauthenticated, "error Unauthenticated")
		return
	}

	err = s.UserUsc.CheckerValidRole(
		ctx,
		useru.CheckValidRoleReq{
			RoleID: data.RoleID,
			Method: method,
		},
	)
	if err != nil {
		err = status.Error(codes.Unauthenticated, "error Unauthenticated")
		return
	}

	return
}
