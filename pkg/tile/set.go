package tile

type Set struct {
	BackgroundColor  string
	Class            string
	Columns          int
	FillMode         string
	FirstGID         int
	Grid             Grid
	Image            string
	ImageHeight      int
	ImageWidth       int
	Margin           int
	Name             string
	ObjectAlignment  string
	Properties       []Property
	Source           string
	Spacing          int
	Terrains         []Terrain
	TileCount        int
	TiledVersion     string
	TileHeight       int
	TileOffset       Offset
	TileRenderSize   string
	Tiles            []T
	TileWidth        int
	Transformations  Transformations
	TransparentColor string
	Type             string
	Version          string
	WangSets         []WangSet
}

type Grid struct {
	Height      int
	Orientation string
	Width       int
}

type Offset struct {
	X int
	Y int
}

type Transformations struct {
	HFlip               bool
	VFlip               bool
	Rotate              bool
	PreferUntransformed bool
}

type T struct {
	Animation   []Frame
	Class       string
	ID          int
	Image       string
	ImageHeight int
	ImageWidth  int
	X           int
	Y           int
	Width       int
	Height      int
	ObjectGroup Layer
	Probability float64
	Properties  []Property
	Terrain     []Terrain
}

type Frame struct {
	Duration int
	TileID   int
}

type Terrain struct {
	Name       string
	Properties []Property
	Tile       int
}

type WangSet struct {
	Class      string
	Colors     []WangColor
	Name       string
	Properties []Property
	Tile       int
	Type       string
	WangTiles  []WangTile
}

type WangColor struct {
	Class       string
	Color       string
	Name        string
	Probability float64
	Properties  []Property
	Tile        int
}

type WangTile struct {
	TileID int
	WangID []byte
}

type ObjectTemplate struct {
	Type   string
	Set    Set
	Object Object
}

type Property struct {
	Name         string
	Type         string
	PropertyType string
	Value        interface{}
}

type Point struct {
	X float64
	Y float64
}
