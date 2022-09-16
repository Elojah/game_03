package tile

import "github.com/elojah/game_03/pkg/ulid"

type Map struct {
	BackgroundColor  string     `json:"backgroundcolor"`
	Class            string     `json:"class"`
	CompressionLevel int        `json:"compressionlevel"`
	Height           int        `json:"height"`
	HexSideLength    int        `json:"hexsidelength"`
	Infinite         bool       `json:"infinite"`
	Layers           []Layer    `json:"layers"`
	NextLayerID      int        `json:"nextlayerid"`
	NextObjectID     int        `json:"nextobjectid"`
	Orientation      string     `json:"orientation"`
	ParallaxOriginX  float64    `json:"parallaxoriginx"`
	ParallaxOriginY  float64    `json:"parallaxoriginy"`
	Properties       []Property `json:"properties"`
	RenderOrder      string     `json:"renderorder"`
	StaggerAxis      string     `json:"staggeraxis"`
	StaggerIndex     string     `json:"staggerindex"`
	TiledVersion     string     `json:"tiledversion"`
	Tilesets         []Set      `json:"tilesets"`
	TileHeight       int        `json:"tileheight"`
	TileWidth        int        `json:"tilewidth"`
	Type             string     `json:"type"`
	Version          string     `json:"version"`
	Width            int        `json:"width"`
}

func NewMap() Map {
	return Map{
		BackgroundColor:  "#000000",
		CompressionLevel: -1,
		Height:           0,
		Width:            0,
		Infinite:         false,
		Layers:           nil,
		NextLayerID:      2, //nolint: gomnd
		NextObjectID:     1,
		Orientation:      "orthogonal",
		RenderOrder:      "right-down",
		TiledVersion:     "custom",
		TileHeight:       32, //nolint: gomnd
		TileWidth:        32, //nolint: gomnd
		Tilesets:         nil,
		Type:             "map",
		Version:          "1.4",

		HexSideLength:   0,
		ParallaxOriginX: 0,
		ParallaxOriginY: 0,
		Properties:      nil,
		StaggerAxis:     "",
		StaggerIndex:    "",
	}
}

type Layer struct {
	Chunks           []Chunk    `json:"chunks"`
	Class            string     `json:"class"`
	Compression      string     `json:"compression"`
	Data             string     `json:"data"`
	DrawOrder        string     `json:"draworder"`
	Encoding         string     `json:"encoding"`
	Height           int        `json:"height"`
	ID               int        `json:"id"`
	Image            string     `json:"image"`
	Layers           []Layer    `json:"layers"`
	Locked           bool       `json:"locked"`
	Name             string     `json:"name"`
	Objects          []Object   `json:"objects"`
	OffsetX          float64    `json:"offsetx"`
	OffsetY          float64    `json:"offsety"`
	Opacity          float64    `json:"opacity"`
	ParallaxX        float64    `json:"parallaxx"`
	ParallaxY        float64    `json:"parallaxy"`
	Properties       []Property `json:"properties"`
	RepeatX          bool       `json:"repeatx"`
	RepeatY          bool       `json:"repeaty"`
	StartX           int        `json:"startx"`
	StartY           int        `json:"starty"`
	TintColor        string     `json:"tintcolor"`
	TransparentColor string     `json:"transparentcolor"`
	Type             string     `json:"type"`
	Visible          bool       `json:"visible"`
	Width            int        `json:"width"`
	X                int        `json:"x"`
	Y                int        `json:"y"`
}

func NewLayer() Layer {
	return Layer{
		Compression: "",
		Data:        "",
		Encoding:    "base64",
		ID:          0,
		Name:        ulid.NewID().String(),
		Height:      0,
		Width:       0,
		X:           0,
		Y:           0,
		Opacity:     1,

		Chunks:    nil,
		Class:     "",
		DrawOrder: "",
		Image:     "",
		Layers:    nil,
		Locked:    true,
		Visible:   true,
		Type:      "tilelayer",

		Objects:          nil,
		OffsetX:          0,
		OffsetY:          0,
		ParallaxX:        0,
		ParallaxY:        0,
		Properties:       nil,
		RepeatX:          false,
		RepeatY:          false,
		StartX:           0,
		StartY:           0,
		TintColor:        "",
		TransparentColor: "",
	}
}

type Chunk struct {
	Data   []uint32 `json:"data"`
	Height int      `json:"height"`
	Width  int      `json:"width"`
	X      int      `json:"x"`
	Y      int      `json:"y"`
}

type Object struct {
	Class      string     `json:"class"`
	Ellipse    bool       `json:"ellipse"`
	GID        int        `json:"gid"`
	Height     float64    `json:"height"`
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Point      bool       `json:"point"`
	Polygon    []Point    `json:"polygon"`
	Polyline   []Point    `json:"polyline"`
	Properties []Property `json:"properties"`
	Rotation   float64    `json:"rotation"`
	Template   string     `json:"template"`
	Text       Text       `json:"text"`
	Visible    bool       `json:"visible"`
	Width      float64    `json:"width"`
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
}

type Text struct {
	Bold       bool   `json:"bold"`
	Color      string `json:"color"`
	FontFamily string `json:"fontfamily"`
	HAlign     string `json:"halign"`
	Italic     bool   `json:"italic"`
	Kerning    bool   `json:"kerning"`
	PixelSize  int    `json:"pixelsize"`
	StrikeOut  bool   `json:"strikeout"`
	Text       string `json:"text"`
	Underline  bool   `json:"underline"`
	VAlign     string `json:"valign"`
	Wrap       bool   `json:"wrap"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Property struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	PropertyType string      `json:"propertytype"`
	Value        interface{} `json:"value"`
}
