package binary

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// File represents an open ijvm binary file
type File struct {
	Blocks []*Block
	closer io.Closer
}

// FormatError represents an error in the ijvm binary format
type FormatError struct {
	offset int64
	msg    string
	value  interface{}
}

func (e *FormatError) Error() string {
	msg := e.msg
	if e.value != nil {
		msg += fmt.Sprintf(" '%v' ", e.value)
	}
	msg += fmt.Sprintf("in file at byte %#x", e.offset)
	return msg
}

// Open opens a new file on disk from a given path, and interprets it as
// an IJVM binary file.
func Open(name string) (*File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	binaryFile, err := NewFile(file)
	if err != nil {
		// we're in the error path, so having an error closing the file
		// doesn't really matter here anymore
		file.Close()
		return nil, err
	}

	binaryFile.closer = file

	return binaryFile, nil
}

// NewFile takes an io.ReaderAt and attempts to parse the data into an
// IJVM file. NewFile expects the binary to start at offset 0
func NewFile(r io.ReaderAt) (*File, error) {
	sr := io.NewSectionReader(r, 0, 1<<63-1)

	var header [4]uint8

	if _, err := r.ReadAt(header[0:], 0); err != nil {
		return nil, err
	}

	if header[0] != 0x1D || header[1] != 0xEA || header[2] != 0xDF || header[3] != 0xAD {
		return nil, &FormatError{0, "bad magic number", header}
	}

	f := new(File)

	// A well-formed binary is guaranteed to have 2 blocks, and at most
	// two extra debug blocks
	f.Blocks = make([]*Block, 0, 4)

	// skip the header
	sr.Seek(4, io.SeekStart)

	constants, err := NewBlock(sr, Constants)
	if err != nil {
		return nil, err
	}

	f.Blocks = append(f.Blocks, constants)

	text, err := NewBlock(sr, Text)
	if err != nil {
		return nil, err
	}

	f.Blocks = append(f.Blocks, text)

	debugBlock, err := NewBlock(sr, Debug)
	if err != nil {
		// if there is no debug block present at all
		// just return what we have
		if err == io.EOF {
			log.Debug("no debug block present")
			return f, nil
		}
		return nil, err
	}

	f.Blocks = append(f.Blocks, debugBlock)

	debugBlock, err = NewBlock(sr, Debug)
	if err != nil {
		if err == io.EOF {
			log.Debug("no second debug block present")
			return f, nil
		}
		return nil, err
	}

	return f, nil
}
