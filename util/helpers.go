package util

import "bytes"

func isIn(xs []byte, v byte) bool {
	return bytes.Contains(xs, []byte{v})
}
