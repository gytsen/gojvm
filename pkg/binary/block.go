package binary

import (
	"encoding/binary"
	"io"

	log "github.com/sirupsen/logrus"
)

// BlockHeader represents the header of a binary block
type BlockHeader struct {
	Origin uint32
	Size   uint32
}

// Block represents a single, arbitrary, block from the binary
type Block struct {
	Header BlockHeader
	Data   []byte
	Type   blockType
}

// NewBlock takes a io.SectionReader and reads a block section from the reader.
// The block is expected to start at offset 0.
func NewBlock(sr *io.SectionReader, t blockType) (*Block, error) {
	bh := new(BlockHeader)
	b := new(Block)

	// read the header
	if err := binary.Read(sr, binary.BigEndian, bh); err != nil {
		return nil, err
	}

	b.Header = *bh
	b.Type = t

	data := make([]byte, b.Header.Size)

	n, err := sr.Read(data)
	log.Debugf("NewBlock: read %d bytes", n)

	if err != nil {
		return nil, err
	}

	b.Data = data

	return b, nil
}
