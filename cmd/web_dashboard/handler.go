package main

import (
	"github.com/elojah/game_03/cmd/auth/grpc"
	"github.com/elojah/game_03/pkg/cookie"
)

type handler struct {
	cookie cookie.App

	grpc.AuthClient
}
