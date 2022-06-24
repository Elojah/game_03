package tile

type Map struct {
	BackgroundColor  string
	Class            string
	CompressionLevel int
	Height           int
	HexSideLength    int
	Infinite         bool
	Layer            []Layer
	NextLayerID      int
	NextObjectID     int
	Orientation      string
	ParallaxOriginX  float64
	ParallaxOriginY  float64
	Properties       []Property
	RenderOrder      string
	StaggerAxis      string
	StaggerIndex     string
	TiledVersion     string
	TileHeight       int
	TileWidth        int
	Type             string
	Version          string
	Width            int
}

type Layer struct {
	Chunks           []Chunk
	Class            string
	Compression      string
	Data             []uint32
	DrawOrder        string
	Encoding         string
	Height           int
	ID               int
	Image            string
	Layers           []Layer
	Locked           bool
	Name             string
	Objects          []Object
	OffsetX          float64
	OffsetY          float64
	Opacity          float64
	ParallaxX        float64
	ParallaxY        float64
	Properties       []Property
	RepeatX          bool
	RepeatY          bool
	StartX           int
	StartY           int
	TintColor        string
	TransparentColor string
	Type             string
	Visible          bool
	Width            int
	X                int
	Y                int
}

type Chunk struct {
	Data   []uint32
	Height int
	Width  int
	X      int
	Y      int
}

type Object struct {
	Class      string
	Ellipse    bool
	GID        int
	Height     float64
	ID         int
	Name       string
	Point      bool
	Polygon    []Point
	Polyline   []Point
	Properties []Property
	Rotation   float64
	Template   string
	Text       Text
	Visible    bool
	Width      float64
	X          float64
	Y          float64
}

type Text struct {
	Bold       bool
	Color      string
	FontFamily string
	HAlign     string
	Italic     bool
	Kerning    bool
	PixelSize  int
	StrikeOut  bool
	Text       string
	Underline  bool
	VAlign     string
	Wrap       bool
}
