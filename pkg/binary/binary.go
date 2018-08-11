package binary

type blockType uint32

// Current available block types
const (
	Constants blockType = iota
	Text
	Debug
)
