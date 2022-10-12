package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/elojah/game_03/pkg/ulid"
)

func run(prog string, param string) {
	id, err := ulid.Parse(param)
	if err != nil {
		fmt.Println("failed to parse param")

		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(id))
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s ulid\n", args[0])

		return
	}

	run(args[0], args[1])
}
