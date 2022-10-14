package main

import (
	"github.com/elojah/game_03/cmd/auth/grpc"
)

type handler struct {
	grpc.AuthClient
}
