package rule

import (
	"bytes"

	"github.com/boundedinfinity/abnfer/char"
)

func equalFoldSS(a, b string) bool {
	return equalFoldSB(a, []byte(b))
}

func equalFoldSB(a string, bs []byte) bool {
	return bytes.EqualFold([]byte(a), bs)
}

func HasPrefix(s string, ps []char.Pattern) bool {
	for _, p := range ps {
		if !char.HasPrefix(s, p) {
			return false
		}
	}

	return true
}
