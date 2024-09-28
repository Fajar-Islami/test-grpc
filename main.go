package main

import (
	"test-code/internal/infrastructrue/container"
	"test-code/internal/server"
)

func main() {
	cont := container.NewContainer()

	server.StartGRPCServer(cont)
}
