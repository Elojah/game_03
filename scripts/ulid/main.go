package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/elojah/game_03/pkg/ulid"
)

func decodeULID(s string) string {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("failed to parse b64", err)

		return ""
	}

	id := ulid.ID(b)

	return id.String()
}

func run(prog string, s string) {
	fmt.Println(decodeULID(s))
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s ulid\n", args[0])

		return
	}

	run(args[0], args[1])
}
