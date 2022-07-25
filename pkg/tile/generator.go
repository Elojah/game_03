package tile

// Params required for map generation.
type Params struct {
	Density    uint64
	SizeMin    uint64
	SizeMax    uint64
	Distortion uint64
	PathMin    uint64
	PathMax    uint64
}

// Generator generates a new map.
type Generator struct {
	Params       Params
	Intermediate [][]int
}
