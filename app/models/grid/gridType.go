package grid

type Direction int

const (
	direction_Horizontal Direction = iota
	direction_Vertical   Direction = iota
)

type Tile int

const (
	BASIC Tile = iota
	DOUBLE_LETTER
	TRIPLE_LETTER
	DOUBLE_WORD
	TRIPLE_WORD
	STAR
	ROW
	COL
	EMPTY
)
