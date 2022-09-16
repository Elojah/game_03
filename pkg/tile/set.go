package tile

type Set struct {
	BackgroundColor  string          `json:"backgroundcolor"`
	Class            string          `json:"class"`
	Columns          int             `json:"columns"`
	FillMode         string          `json:"fillmode"`
	FirstGID         int             `json:"firstgid"`
	Grid             Grid            `json:"grid"`
	Image            string          `json:"image"`
	ImageHeight      int             `json:"imageheight"`
	ImageWidth       int             `json:"imagewidth"`
	Margin           int             `json:"margin"`
	Name             string          `json:"name"`
	ObjectAlignment  string          `json:"objectalignment"`
	Properties       []Property      `json:"properties"`
	Source           string          `json:"source"`
	Spacing          int             `json:"spacing"`
	Terrains         []Terrain       `json:"terrains"`
	TileCount        int             `json:"tilecount"`
	TiledVersion     string          `json:"tiledversion"`
	TileHeight       int             `json:"tileheight"`
	TileOffset       Offset          `json:"tileoffset"`
	TileRenderSize   string          `json:"tilerendersize"`
	Tiles            []T             `json:"tiles"`
	TileWidth        int             `json:"tilewidth"`
	Transformations  Transformations `json:"transformations"`
	TransparentColor string          `json:"transparentcolor"`
	Type             string          `json:"type"`
	Version          string          `json:"version"`
	WangSets         []WangSet       `json:"wangsets"`
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

type Grid struct {
	Height      int    `json:"height"`
	Orientation string `json:"orientation"`
	Width       int    `json:"width"`
}

type Offset struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Transformations struct {
	HFlip               bool `json:"hflip"`
	VFlip               bool `json:"vflip"`
	Rotate              bool `json:"rotate"`
	PreferUntransformed bool `json:"preferuntransformed"`
}

type T struct {
	Animation   []Frame    `json:"animation"`
	Class       string     `json:"class"`
	ID          int        `json:"id"`
	Image       string     `json:"image"`
	ImageHeight int        `json:"imageheight"`
	ImageWidth  int        `json:"imagewidth"`
	X           int        `json:"x"`
	Y           int        `json:"y"`
	Width       int        `json:"width"`
	Height      int        `json:"height"`
	ObjectGroup Layer      `json:"objectgroup"`
	Probability float64    `json:"probability"`
	Properties  []Property `json:"properties"`
	Terrain     []Terrain  `json:"terrain"`
}

type Frame struct {
	Duration int `json:"duration"`
	TileID   int `json:"tileid"`
}

type Terrain struct {
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
	Tile       int        `json:"tile"`
}

type WangSet struct {
	Class      string      `json:"class"`
	Colors     []WangColor `json:"colors"`
	Name       string      `json:"name"`
	Properties []Property  `json:"properties"`
	Tile       int         `json:"tile"`
	Type       string      `json:"type"`
	WangTiles  []WangTile  `json:"wangtiles"`
}

type WangColor struct {
	Class       string     `json:"class"`
	Color       string     `json:"color"`
	Name        string     `json:"name"`
	Probability float64    `json:"probability"`
	Properties  []Property `json:"properties"`
	Tile        int        `json:"tile"`
}

type WangTile struct {
	TileID int    `json:"tileid"`
	WangID []byte `json:"wangid"`
}

type ObjectTemplate struct {
	Type   string `json:"type"`
	Set    Set    `json:"set"`
	Object Object `json:"object"`
}
