// Package cdat implements a reader for the DAT container format used to store
// Ceebot and Colobot game data.
package cdat

import (
	"errors"
	"io"
)

// ErrBadHeader is the error used when an invalid value is read from a DAT
// header.
var ErrBadHeader = errors.New("cdat: header contains illegal values")

// ErrNameTooLong is the error used when a filename exceeds the 13 character
// limit of DAT files.
var ErrNameTooLong = errors.New("cdat: filename too long")

// ErrReadOnly is the error used when an attempt is made to add files to a
// container opened for reading.
var ErrReadOnly = errors.New("cdat: Container is read-only")

// Container represents a DAT file.
type Container struct {
	Files []File // files in this container
	c     Codec
	r     io.ReaderAt
}

// File represents a file stored in a DAT container.
type File struct {
	*io.SectionReader        // reader for this file
	Offset            int64  // offset where this file starts
	Length            int64  // length of this file
	Name              string // name of this file
}

// New creates a new Container reading from the specified io.ReaderAt.  If r
// cannot be decoded using codec, Read returns ErrBadHeader.
func New(r io.ReaderAt, codec Codec) (c *Container, err error) {
	var numFiles int32

	if n, err := readInt(r, 0); err != nil {
		return nil, err
	} else if n < 0 {
		return nil, ErrBadHeader
	} else {
		numFiles = n
	}

	c = &Container{
		Files: make([]File, 0, numFiles),
		c:     codec,
		r:     NewDecodingReaderAt(r, codec),
	}

	const size = 24 // length of file header
	for i := 0; i < int(numFiles); i++ {
		var name string
		if name, err = readString(c.r, int64(size*i+4), 16); err != nil {
			return
		} else if len(name) > 12 {
			return nil, ErrBadHeader
		}
		for _, v := range name {
			if v < 32 || v > 126 {
				return nil, ErrBadHeader
			}
		}

		var start int32
		if n, err := readInt(c.r, int64(size*i+20)); err != nil {
			return c, err
		} else if n < 4 {
			return nil, ErrBadHeader
		} else {
			start = n
		}

		var length int32
		if n, err := readInt(c.r, int64(size*i+24)); err != nil {
			return c, err
		} else if n < 0 {
			return nil, ErrBadHeader
		} else {
			length = n
		}

		c.Files = append(c.Files, File{
			SectionReader: io.NewSectionReader(c.r, int64(start), int64(length)),
			Offset:        int64(start),
			Length:        int64(length),
			Name:          name,
		})
	}

	return c, nil
}
