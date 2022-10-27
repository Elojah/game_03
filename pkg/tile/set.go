package tile

type Set struct {
	BackgroundColor  string          `json:"backgroundcolor,omitempty"`
	Class            string          `json:"class,omitempty"`
	Columns          int             `json:"columns,omitempty"`
	FillMode         string          `json:"fillmode,omitempty"`
	FirstGID         int             `json:"firstgid,omitempty"`
	Grid             Grid            `json:"grid,omitempty"`
	Image            string          `json:"image,omitempty"`
	ImageHeight      int             `json:"imageheight,omitempty"`
	ImageWidth       int             `json:"imagewidth,omitempty"`
	Margin           int             `json:"margin"`
	Name             string          `json:"name,omitempty"`
	ObjectAlignment  string          `json:"objectalignment,omitempty"`
	Properties       []Property      `json:"properties,omitempty"`
	Source           string          `json:"source,omitempty"`
	Spacing          int             `json:"spacing"`
	Terrains         []Terrain       `json:"terrains,omitempty"`
	TileCount        int             `json:"tilecount,omitempty"`
	TiledVersion     string          `json:"tiledversion,omitempty"`
	TileHeight       int             `json:"tileheight,omitempty"`
	TileOffset       Offset          `json:"tileoffset,omitempty"`
	TileRenderSize   string          `json:"tilerendersize,omitempty"`
	Tiles            []Tile          `json:"tiles,omitempty"`
	TileWidth        int             `json:"tilewidth,omitempty"`
	Transformations  Transformations `json:"transformations,omitempty"`
	TransparentColor string          `json:"transparentcolor,omitempty"`
	Type             string          `json:"type,omitempty"`
	Version          string          `json:"version,omitempty"`
	WangSets         []WangSet       `json:"wangsets,omitempty"`
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
	Class      string      `json:"class,omitempty"`
	Colors     []WangColor `json:"colors,omitempty"`
	Name       string      `json:"name,omitempty"`
	Properties []Property  `json:"properties,omitempty"`
	Tile       int         `json:"tile,omitempty"`
	Type       string      `json:"type,omitempty"`
	WangTiles  []WangTile  `json:"wangtiles,omitempty"`
}

type WangColor struct {
	Class       string     `json:"class,omitempty"`
	Color       string     `json:"color,omitempty"`
	Name        string     `json:"name,omitempty"`
	Probability float64    `json:"probability,omitempty"`
	Properties  []Property `json:"properties,omitempty"`
	Tile        int        `json:"tile,omitempty"`
}

type WangTile struct {
	TileID int    `json:"tileid,omitempty"`
	WangID []byte `json:"wangid,omitempty"`
}

type ObjectTemplate struct {
	Type   string `json:"type,omitempty"`
	Set    Set    `json:"set,omitempty"`
	Object Object `json:"object,omitempty"`
}
