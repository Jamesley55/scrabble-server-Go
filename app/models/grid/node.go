package grid

type Word struct {
	Letters      []*Node
	IsHorizontal bool
}

type Node struct {
	Letter    *Letter
	TileType  Tile
	IsEmpty   bool
	X         int
	Y         int
	IsNewNode bool
}

type Letter struct {
	Character     string
	Value         int
	SelectionType SelectionType
}

type SelectionType string

const (
	EXCHANGE     SelectionType = "exchange"
	MANIPULATION SelectionType = "manipulation"
	NONE         SelectionType = "none"
)
