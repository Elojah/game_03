package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/elojah/game_03/pkg/tile/dto"
	"github.com/elojah/game_03/pkg/ulid"
)

var tilesetSheets = map[string]dto.CreateTilesetReq{
	"RuinsWall": {
		ID: ulid.MustParse("01GHE0TD8VC0HJHAEGTWN0AF44"),
	},
}

func run(prog string, input string, output string) {
	for name, req := range tilesetSheets {
		xml, err := os.ReadFile(filepath.Join(input, req.ID.String()+".xml"))
		if err != nil {
			fmt.Println("failed to read tileset file", err)

			return
		}

		req.Set = xml

		raw, err := json.Marshal(req)
		if err != nil {
			fmt.Println("failed to marshal JSON", err)

			return
		}

		if err := os.WriteFile(
			filepath.Join(output, name+".json"),
			raw,
			0o600, //nolint: gomnd
		); err != nil {
			fmt.Println("failed to create tileset file", err)

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 3 { //nolint: gomnd
		fmt.Printf("Usage: ./%s input_directory output_directory\n", args[0])

		return
	}

	run(args[0], args[1], args[2])
}
