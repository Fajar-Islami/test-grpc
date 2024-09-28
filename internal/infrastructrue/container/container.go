package container

import (
	"fmt"
	"test-code/internal/helper"
	"test-code/internal/infrastructrue/postgre"
	"test-code/internal/utils"

	"github.com/joho/godotenv"

	"test-code/internal/domain/userd"
	"test-code/internal/usecase/useru"
)

type (
	Container struct {
		Apps    Apps
		UserUsc useru.UserUsc
	}

	Apps struct {
		Name           string `env:"apps_appName"`
		Host           string `env:"apps_host"`
		Version        string `env:"apps_version"`
		SwaggerAddress string `env:"apps_swagger_address"`
		HttpPort       int    `env:"apps_httpport"`
		SecretJwt      string `env:"apps_secretJwt"`
		CtxTimeout     int    `env:"apps_timeout"`
	}
)

func NewContainer() *Container {
	err := godotenv.Load(fmt.Sprintf("%s/%s", helper.ProjectRootPath, ".env"))
	if err != nil {
		panic(err)
	}

	var appsConf = Apps{
		Name:           utils.EnvString("apps_appName"),
		Host:           utils.EnvString("apps_host"),
		Version:        utils.EnvString("apps_version"),
		SwaggerAddress: utils.EnvString("apps_swagger_address"),
		HttpPort:       utils.EnvInt("apps_httpport"),
		SecretJwt:      utils.EnvString("apps_secretJwt"),
		CtxTimeout:     utils.EnvInt("apps_timeout"),
	}

	utils.InitJWT(utils.EnvString("apps_secretJwt"))

	postgre, err := postgre.Init()
	if err != nil {
		panic(err)
	}

	userRepo := userd.NewUserDomain(postgre)
	userUsc := useru.NewUseUsecase(userRepo)

	cont := Container{
		Apps:    appsConf,
		UserUsc: userUsc,
	}

	return &cont
}
