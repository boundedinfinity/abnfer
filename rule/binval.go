package rule

import (
	"github.com/boundedinfinity/abnfer/char"
)

type NumVal struct {
	Min     int
	Max     int
	Pattern []char.Pattern
}

func ParseNumVal(s string) (NumVal, error) {
	var numVal NumVal
	// bs := []byte(s)
	// bl := len(bs)

	switch {
	case HasPrefix(s, REPETITION_VARIABLE_MIN_MAX):
	case HasPrefix(s, REPETITION_VARIABLE_MIN_ONLY):
	case HasPrefix(s, REPETITION_VARIABLE_MAX_ONLY):
	case HasPrefix(s, REPETITION_SPECIFIC):
	}

	return numVal, nil
}
