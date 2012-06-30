package cdat

import "io"

// readInt reads 4 bytes from r starting at off, interprets them as a
// little-endian int32, and returns the result.
func readInt(r io.ReaderAt, off int64) (n int32, err error) {
	p := make([]byte, 4)
	if n, err := r.ReadAt(p, off); err != nil && n < 4 {
		return 0, err
	}
	n = int32(p[0]) | int32(p[1])<<8 | int32(p[2])<<16 | int32(p[3])<<24
	return n, nil
}

// readString reads n bytes from r starting at off, strips off garbage after
// the first null character ('\0'), and returns the resulting string.
func readString(r io.ReaderAt, off int64, n uint) (s string, err error) {
	p := make([]byte, n)
	if n, err := r.ReadAt(p, off); err != nil && err != io.EOF {
		return "", err
	} else if n < len(p) {
		p = p[0:n]
	}

	for i, v := range p {
		if v == 0 {
			return string(p[0:i]), nil
		}
	}
	return string(p), nil
}
