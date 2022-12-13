package tile

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Set struct {
	ID               ulid.ID         `json:"-" xml:"-"`
	BackgroundColor  string          `json:"backgroundcolor,omitempty" xml:"backgroundcolor"`
	Class            string          `json:"class,omitempty" xml:"class"`
	Columns          int             `json:"columns,omitempty" xml:"columns,attr"`
	FillMode         string          `json:"fillmode,omitempty" xml:"fillmode"`
	FirstGID         int             `json:"firstgid,omitempty" xml:"firstgid,attr"`
	Grid             Grid            `json:"-" xml:"grid"`
	Image            string          `json:"image,omitempty" xml:"image,attr"`
	ImageHeight      int             `json:"imageheight,omitempty" xml:"imageheight,attr"`
	ImageWidth       int             `json:"imagewidth,omitempty" xml:"imagewidth,attr"`
	Margin           int             `json:"margin" xml:"margin"`
	Name             string          `json:"name,omitempty" xml:"name,attr"`
	ObjectAlignment  string          `json:"objectalignment,omitempty" xml:"objectalignment"`
	Properties       []Property      `json:"properties,omitempty" xml:"properties"`
	Source           string          `json:"source,omitempty" xml:"source"`
	Spacing          int             `json:"spacing" xml:"spacing"`
	Terrains         []Terrain       `json:"terrains,omitempty" xml:"terrains"`
	TileCount        int             `json:"tilecount,omitempty" xml:"tilecount,attr"`
	TiledVersion     string          `json:"tiledversion,omitempty" xml:"tiledversion,attr"`
	TileHeight       int             `json:"tileheight,omitempty" xml:"tileheight,attr"`
	TileOffset       Offset          `json:"-" xml:"tileoffset,attr"`
	TileRenderSize   string          `json:"tilerendersize,omitempty" xml:"tilerendersize"`
	Tiles            []Tile          `json:"tiles,omitempty" xml:"tiles"`
	TileWidth        int             `json:"tilewidth,omitempty" xml:"tilewidth,attr"`
	Transformations  Transformations `json:"-" xml:"transformations"`
	TransparentColor string          `json:"transparentcolor,omitempty" xml:"transparentcolor"`
	Type             string          `json:"type,omitempty" xml:"type,attr"`
	Version          string          `json:"version,omitempty" xml:"version,attr"`
	WangSets         []WangSet       `json:"wangsets,omitempty" xml:"wangsets>wangset"`
}

func (s Set) CleanCopy() Set {
	s.WangSets = nil

	return s
}

type FilterSet struct {
	ID ulid.ID
}

type StoreSet interface {
	InsertSet(context.Context, Set) error
	FetchSet(context.Context, FilterSet) (Set, error)
	DeleteSet(context.Context, FilterSet) error
}

type App interface {
	StoreSet
}

func NewSet() Set {
	return Set{
		BackgroundColor:  "#000000",
		Class:            "background",
		Columns:          0,
		FillMode:         "stretch",
		FirstGID:         0,
		Grid:             Grid{},
		Image:            "",
		ImageHeight:      0,
		ImageWidth:       0,
		Margin:           0,
		Name:             "",
		ObjectAlignment:  "unspecified",
		Properties:       nil,
		Source:           "",
		Spacing:          0,
		Terrains:         nil,
		TileCount:        0,
		TiledVersion:     "custom",
		TileHeight:       0,
		TileOffset:       Offset{},
		TileRenderSize:   "tile",
		Tiles:            nil,
		TileWidth:        0,
		Transformations:  Transformations{},
		TransparentColor: "#000000",
		Type:             "tileset",
		Version:          "1.4",
		WangSets:         nil,
	}
}

type Tile struct {
	Animation   []Frame    `json:"animation,omitempty"`
	Class       string     `json:"class,omitempty"`
	ID          int        `json:"id,omitempty"`
	Image       string     `json:"image,omitempty"`
	ImageHeight int        `json:"imageheight,omitempty"`
	ImageWidth  int        `json:"imagewidth,omitempty"`
	X           int        `json:"x,omitempty"`
	Y           int        `json:"y,omitempty"`
	Width       int        `json:"width,omitempty"`
	Height      int        `json:"height,omitempty"`
	ObjectGroup Layer      `json:"objectgroup,omitempty"`
	Probability float64    `json:"probability,omitempty"`
	Properties  []Property `json:"properties,omitempty"`
	Terrain     []Terrain  `json:"terrain,omitempty"`
}

type Terrain struct {
	Name       string     `json:"name,omitempty"`
	Properties []Property `json:"properties,omitempty"`
	Tile       int        `json:"tile,omitempty"`
}

type Grid struct {
	Height      int    `json:"height,omitempty"`
	Orientation string `json:"orientation,omitempty"`
	Width       int    `json:"width,omitempty"`
}

type Offset struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

type Transformations struct {
	HFlip               bool `json:"hflip,omitempty"`
	VFlip               bool `json:"vflip,omitempty"`
	Rotate              bool `json:"rotate,omitempty"`
	PreferUntransformed bool `json:"preferuntransformed,omitempty"`
}

type Frame struct {
	Duration int `json:"duration,omitempty"`
	TileID   int `json:"tileid,omitempty"`
}

type WangSet struct {
	Class      string      `json:"class,omitempty" xml:"class"`
	Colors     []WangColor `json:"colors,omitempty" xml:"wangcolor"`
	Name       string      `json:"name,omitempty" xml:"name,attr"`
	Properties []Property  `json:"properties,omitempty" xml:"properties"`
	Tile       int         `json:"tile,omitempty" xml:"tile,attr"`
	Type       WangType    `json:"type,omitempty" xml:"type,attr"`
	WangTiles  []WangTile  `json:"wangtiles,omitempty" xml:"wangtile"`
}

type WangColor struct {
	Class       string     `json:"class,omitempty"`
	Color       string     `json:"color,omitempty" xml:"color,attr"`
	Name        string     `json:"name,omitempty" xml:"name,attr"`
	Probability float64    `json:"probability,omitempty" xml:"probability,attr"`
	Properties  []Property `json:"properties,omitempty" xml:"properties"`
	Tile        int        `json:"tile,omitempty" xml:"tile,attr"`
}

type WangTile struct {
	TileID int    `json:"tileid,omitempty" xml:"tileid,attr"`
	WangID []byte `json:"wangid,omitempty" xml:"wangid,attr"`
}

type ObjectTemplate struct {
	Type   string `json:"type,omitempty"`
	Set    Set    `json:"set,omitempty"`
	Object Object `json:"object,omitempty"`
}
