package cdat

import "io"

// A Codec encodes or decodes a slice of bytes starting at a file offset,
// returning the number of bytes encoded or decoded.  (Colobot and Ceebot
// encode their containers with an XOR encryption scheme, which is symmetric.)
type Codec func(p []byte, off int64) int

// runCodec takes a code table, a slice of bytes, and a file offset and XORs
// each byte in p with the corresponding byte from codec, returning the number
// of bytes encoded.
func runCodec(codec []byte, p []byte, off int64) int {
	var i int
	for i = 0; i < len(p); i++ {
		if int(off)+i < 4 {
			continue // first four bytes are not encoded
		}
		p[i] ^= codec[(int(off)+i)%len(codec)]
	}
	return i
}

// CeebotCodec implements the codec used in the full version of Ceebot.
func CeebotCodec(p []byte, off int64) int {
	codec := []byte{
		0x72, 0x91, 0x37, 0xdf, 0xa1, 0xcc, 0xf5, 0x67,
		0x53, 0x40, 0xd3, 0xed, 0x3a, 0xbb, 0x5e, 0x43,
		0x67, 0x9a, 0x0c, 0xed, 0x33, 0x77, 0x2f, 0xf2,
		0xe3, 0x42, 0x11, 0x5e, 0xc2,
	}
	return runCodec(codec, p, off)
}

// CeebotDemo codec implements the codec used in the demo version of Ceebot.
func CeebotDemoCodec(p []byte, off int64) int {
	codec := []byte{
		0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76,
		0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76,
		0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76,
		0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76,
		0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76,
		0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76, 0x76,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
		0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5,
	}
	return runCodec(codec, p, off)
}

// ColobotCodec implements the codec used in the full version of Colobot.
func ColobotCodec(p []byte, off int64) int {
	codec := []byte{
		0x85, 0x91, 0x73, 0xcf, 0xa2, 0xbb, 0xf4, 0x77,
		0x58, 0x39, 0x37, 0xfd, 0x2a, 0xcc, 0x5f, 0x55,
		0x96, 0x90, 0x07, 0xcd, 0x11, 0x88, 0x21,
	}
	return runCodec(codec, p, off)
}

// ColobotDemoCodec implements the codec used in the demo version of Colobot.
func ColobotDemoCodec(p []byte, off int64) int {
	codec := []byte{
		0x85, 0x91, 0x77, 0xcf, 0xa3, 0xbb, 0xf4, 0x77,
		0x58, 0x39, 0x37, 0xfd, 0x2a, 0xcc, 0x7f, 0x55,
		0x96, 0x80, 0x07, 0xcd, 0x11, 0x88, 0x21, 0x44,
		0x17, 0xee, 0xf0,
	}
	return runCodec(codec, p, off)
}

// IdentityCodec passes p through unchanged.
func IdentityCodec(p []byte, off int64) int {
	return len(p)
}

// DecodingReaderAt is an io.ReaderAt that decodes its bytes using a codec as
// they are read.
type DecodingReaderAt struct {
	codec    Codec
	readerAt io.ReaderAt
}

func NewDecodingReaderAt(w io.ReaderAt, c Codec) DecodingReaderAt {
	return DecodingReaderAt{codec: c, readerAt: w}
}

func (r DecodingReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = r.readerAt.ReadAt(p, off)
	return r.codec(p[0:n], off), err
}
