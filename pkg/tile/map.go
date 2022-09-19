package tile

import "github.com/elojah/game_03/pkg/ulid"

type Map struct {
	BackgroundColor  string     `json:"backgroundcolor,omitempty"`
	Class            string     `json:"class,omitempty"`
	CompressionLevel int        `json:"compressionlevel,omitempty"`
	Height           int        `json:"height,omitempty"`
	HexSideLength    int        `json:"hexsidelength,omitempty"`
	Infinite         bool       `json:"infinite,omitempty"`
	Layers           []Layer    `json:"layers,omitempty"`
	NextLayerID      int        `json:"nextlayerid,omitempty"`
	NextObjectID     int        `json:"nextobjectid,omitempty"`
	Orientation      string     `json:"orientation,omitempty"`
	ParallaxOriginX  float64    `json:"parallaxoriginx,omitempty"`
	ParallaxOriginY  float64    `json:"parallaxoriginy,omitempty"`
	Properties       []Property `json:"properties,omitempty"`
	RenderOrder      string     `json:"renderorder,omitempty"`
	StaggerAxis      string     `json:"staggeraxis,omitempty"`
	StaggerIndex     string     `json:"staggerindex,omitempty"`
	TiledVersion     string     `json:"tiledversion,omitempty"`
	Tilesets         []Set      `json:"tilesets,omitempty"`
	TileHeight       int        `json:"tileheight,omitempty"`
	TileWidth        int        `json:"tilewidth,omitempty"`
	Type             string     `json:"type,omitempty"`
	Version          string     `json:"version,omitempty"`
	Width            int        `json:"width,omitempty"`
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
	Chunks           []Chunk    `json:"chunks,omitempty"`
	Class            string     `json:"class,omitempty"`
	Compression      string     `json:"compression,omitempty"`
	Data             string     `json:"data,omitempty"`
	DrawOrder        string     `json:"draworder,omitempty"`
	Encoding         string     `json:"encoding,omitempty"`
	Height           int        `json:"height,omitempty"`
	ID               int        `json:"id,omitempty"`
	Image            string     `json:"image,omitempty"`
	Layers           []Layer    `json:"layers,omitempty"`
	Locked           bool       `json:"locked,omitempty"`
	Name             string     `json:"name,omitempty"`
	Objects          []Object   `json:"objects,omitempty"`
	OffsetX          float64    `json:"offsetx,omitempty"`
	OffsetY          float64    `json:"offsety,omitempty"`
	Opacity          float64    `json:"opacity,omitempty"`
	ParallaxX        float64    `json:"parallaxx,omitempty"`
	ParallaxY        float64    `json:"parallaxy,omitempty"`
	Properties       []Property `json:"properties,omitempty"`
	RepeatX          bool       `json:"repeatx,omitempty"`
	RepeatY          bool       `json:"repeaty,omitempty"`
	StartX           int        `json:"startx,omitempty"`
	StartY           int        `json:"starty,omitempty"`
	TintColor        string     `json:"tintcolor,omitempty"`
	TransparentColor string     `json:"transparentcolor,omitempty"`
	Type             string     `json:"type,omitempty"`
	Visible          bool       `json:"visible,omitempty"`
	Width            int        `json:"width,omitempty"`
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
	Data   []uint32 `json:"data,omitempty"`
	Height int      `json:"height,omitempty"`
	Width  int      `json:"width,omitempty"`
	X      int      `json:"x,omitempty"`
	Y      int      `json:"y,omitempty"`
}

type Object struct {
	Class      string     `json:"class,omitempty"`
	Ellipse    bool       `json:"ellipse,omitempty"`
	GID        int        `json:"gid,omitempty"`
	Height     float64    `json:"height,omitempty"`
	ID         int        `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Point      bool       `json:"point,omitempty"`
	Polygon    []Point    `json:"polygon,omitempty"`
	Polyline   []Point    `json:"polyline,omitempty"`
	Properties []Property `json:"properties,omitempty"`
	Rotation   float64    `json:"rotation,omitempty"`
	Template   string     `json:"template,omitempty"`
	Text       Text       `json:"text,omitempty"`
	Visible    bool       `json:"visible,omitempty"`
	Width      float64    `json:"width,omitempty"`
	X          float64    `json:"x,omitempty"`
	Y          float64    `json:"y,omitempty"`
}

type Text struct {
	Bold       bool   `json:"bold,omitempty"`
	Color      string `json:"color,omitempty"`
	FontFamily string `json:"fontfamily,omitempty"`
	HAlign     string `json:"halign,omitempty"`
	Italic     bool   `json:"italic,omitempty"`
	Kerning    bool   `json:"kerning,omitempty"`
	PixelSize  int    `json:"pixelsize,omitempty"`
	StrikeOut  bool   `json:"strikeout,omitempty"`
	Text       string `json:"text,omitempty"`
	Underline  bool   `json:"underline,omitempty"`
	VAlign     string `json:"valign,omitempty"`
	Wrap       bool   `json:"wrap,omitempty"`
}

type Point struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Property struct {
	Name         string      `json:"name,omitempty"`
	Type         string      `json:"type,omitempty"`
	PropertyType string      `json:"propertytype,omitempty"`
	Value        interface{} `json:"value,omitempty"`
}
