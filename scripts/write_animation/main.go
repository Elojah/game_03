package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	"github.com/elojah/game_03/pkg/ulid"
)

var (
	npcSheets = map[string]string{
		"Altar": "01GTXTWQQ0GEK0SZZ83FXPDB3E",
	}

	npcAnimations = map[string]dto.CreateAnimationReq{
		"main": {
			Animation: entity.Animation{
				Name:         "main",
				Start:        1,
				End:          39,
				Rate:         4,
				FrameWidth:   224,
				FrameHeight:  288,
				FrameStart:   0,
				FrameEnd:     38,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
	}

	charSheets = map[string]string{
		"BlueNinja":   "01GF16HSRMZ5HD82BSPTZ8YAAY",
		"BlueSamurai": "01GF16HSRMMR8ME1FM50A5BE8E",
		"Boy":         "01GF16HSRMMYB38HS5X955K2HJ",
		"Cavegirl":    "01GF16HSRMSYJBMM38F8QMDZ41",
		"Cavegirl2":   "01GF16HSRM9QPEGNQ43HTE9RBP",
		"Caveman":     "01GF16HSRMA04AFK11TC5FHS5J",
		"Caveman2":    "01GF16HSRMTCHG0SEKPKWTRXQH",
		"Child":       "01GF16HSRMPP1NRHV2MFTWGCCA",
		"DarkNinja":   "01GF16HSRMX7RJRB5QD9H28M9K",
		"EggBoy":      "01GF16HSRMDY74YT3TJSY9YWQE",
		"EggGirl":     "01GF16HSRMNRV1N8J1Q1X0QR1E",
		"Eskimo":      "01GF16HSRM7PH0DYGM4W6NNE3E",
		"EskimoNinja": "01GF16HSRMFZDHC84HH7WTF1ES",
		"GoldKnight":  "01GF16HSRM83CP64RB0YKCR6AH",
		"GrayNinja":   "01GF16HSRMF8X4227SMB4F9HA6",
		"Greenman":    "01GF16HSRMKY2Y398SNJ1XXB7Z",
		"GreenNinja":  "01GF16HSRMJWCBT1RNWKFXGDH8",
		"Inspector":   "01GF16HSRM2977NCTE7WN2CGPZ",
		"Knight":      "01GF16HSRMYC30729BW2DSTAAB",
		"Lion":        "01GF16HSRMJ16YDZXW5K6P8DD3",
		"MaskedNinja": "01GF16HSRMX2N3MVESA3S42V4Q",
		"MaskFrog":    "01GF16HSRMY486XKF297RF0VH0",
		"Master":      "01GF16HSRMWKJ0Q56XMM8XFGX8",
		"Monk":        "01GF16HSRMS4AZYZC1FG39FW61",
		"Monk2":       "01GF16HSRMAW36696WXAEGJ85R",
		"Noble":       "01GF16HSRMW6SJYHMEQJE5VCP0",
		"OldMan":      "01GF16HSRMPANM50203S1W2C6Y",
		"OldMan2":     "01GF16HSRM1TJ7ZN1CS86HQ8AV",
		"OldMan3":     "01GF16HSRMJ6HXQYYNFHBP5VTF",
		"OldWoman":    "01GF16HSRM02YP79VAWSJ676RK",
		"Princess":    "01GF16HSRMSATBGJ24AN42C7ZE",
		"RedNinja":    "01GF16HSRMZEYNNSFMTJ30Q9ND",
		"RedNinja2":   "01GF16HSRMWCFBPC1VH3Z4W2HR",
		"RedSamurai":  "01GF16HSRMAPHDQT6NY5NXEYTH",
		"Samurai":     "01GF16HSRMHKVM3C5TQZZGNETZ",
		"Skeleton":    "01GF16HSRMGYMVF4HYRQRFQQ23",
		"Villager":    "01GF16HSRMF9HPPEWQ3HP2X1PE",
		"Villager2":   "01GF16HSRMEEF4NN3Y6MMHX2AF",
		"Villager3":   "01GF16HSRM45BTZQYH9519PQFZ",
		"Villager4":   "01GF16HSRMDF7V4CX148EJCFMJ",
		"Woman":       "01GF16HSRMAF049BW351816QPW",
	}

	charAnimations = map[string]dto.CreateAnimationReq{
		"idle_down": {
			Animation: entity.Animation{
				Name:         "idle_down",
				Start:        0,
				End:          1,
				Rate:         1,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"idle_left": {
			Animation: entity.Animation{
				Name:         "idle_left",
				Start:        2,
				End:          3,
				Rate:         1,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"idle_right": {
			Animation: entity.Animation{
				Name:         "idle_right",
				Start:        3,
				End:          4,
				Rate:         1,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"idle_up": {
			Animation: entity.Animation{
				Name:         "idle_up",
				Start:        1,
				End:          2,
				Rate:         1,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"walk_down": {
			Animation: entity.Animation{
				Name:         "walk_down",
				Sequence:     []int64{0, 4, 8, 12},
				Rate:         4,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"walk_left": {
			Animation: entity.Animation{
				Name:         "walk_left",
				Sequence:     []int64{2, 6, 10, 14},
				Rate:         4,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"walk_right": {
			Animation: entity.Animation{
				Name:         "walk_right",
				Sequence:     []int64{3, 7, 11, 15},
				Rate:         4,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
		"walk_up": {
			Animation: entity.Animation{
				Name:         "walk_up",
				Sequence:     []int64{1, 5, 9, 13},
				Rate:         4,
				FrameWidth:   16,
				FrameHeight:  16,
				FrameStart:   0,
				FrameEnd:     27,
				FrameMargin:  0,
				FrameSpacing: 0,
			},
		},
	}
)

func encodeULID(s string) string {
	id, err := ulid.Parse(s)
	if err != nil {
		fmt.Println("failed to parse id", err)

		return ""
	}

	return base64.StdEncoding.EncodeToString(id)
}

func generateAnimations(path string, sheets map[string]string, anims map[string]dto.CreateAnimationReq) {
	for name, id := range sheets {
		for _, ba := range anims {
			ba := ba
			ba.EntityTemplate = name
			ba.SheetID = ulid.MustParse(id)

			raw, err := json.Marshal(ba)
			if err != nil {
				fmt.Println("failed to marshal JSON", err)

				return
			}

			if err := os.WriteFile(
				filepath.Join(path, name, ba.Name+".json"),
				raw,
				0o600, //nolint: gomnd
			); err != nil {
				fmt.Println("failed to create animation file", err)

				return
			}
		}
	}

}

func run(prog string, param string) {
	generateAnimations(param, charSheets, charAnimations)
	generateAnimations(param, npcSheets, npcAnimations)
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s directory\n", args[0])

		return
	}

	run(args[0], args[1])
}
