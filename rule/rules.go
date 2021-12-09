package rule

import (
	"math"

	"github.com/boundedinfinity/abnfer/char"
)

const (
	Min = 0
	Max = math.MaxInt
)

var (
	REPETITION_SPECIFIC = []char.Pattern{
		char.DIGIT,
	}
	REPETITION_VARIABLE_MIN_MAX = []char.Pattern{
		char.DIGIT,
		char.ASTERISK,
		char.DIGIT,
	}
	REPETITION_VARIABLE_MIN_ONLY = []char.Pattern{
		char.DIGIT,
		char.ASTERISK,
	}
	REPETITION_VARIABLE_MAX_ONLY = []char.Pattern{
		char.ASTERISK,
		char.DIGIT,
	}

	Rules = map[string]*Rule{
		"ALPHA": {
			Name:     "ALPHA",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.ALPHA},
		},
		"BIT": {
			Name:     "BIT",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.BIT},
		},
		"CHAR": {
			Name:     "CHAR",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.CHAR},
		},
		"CRLF": {
			Name: "CRLF",
			Min:  0,
			Max:  math.MaxInt,
			Patterns: []char.Pattern{
				char.LF,
				char.Concat(char.CR, char.LF),
			},
		},
		"CTL": {
			Name:     "CTL",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.Range(0x00, 0x1F), char.List(0x7F)},
		},
		"DIGIT": {
			Name:     "DIGIT",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.DIGIT},
		},
		"DQUOTE": {
			Name:     "DQUOTE",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.DQUOTE},
		},
		"HEXDIG": {
			Name:     "HEXDIG",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.HEXDIG},
		},
		"HTAB": {
			Name:     "HTAB",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.HTAB},
		},
		"LF": {
			Name:     "LF",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.LF},
		},
		"LWSP": {
			Name:     "LWSP",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.SP, char.HTAB},
		},
		"OCTET": {
			Name:     "OCTET",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.OCTET},
		},
		"SP": {
			Name:     "SP",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.SP},
		},
		"VCHAR": {
			Name:     "VCHAR",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.VCHAR},
		},
		"WSP": {
			Name:     "WSP",
			Min:      0,
			Max:      math.MaxInt,
			Patterns: []char.Pattern{char.SP, char.HTAB},
		},
	}
)

func init() {

}

type Rule struct {
	Name     string
	Min      int
	Max      int
	Patterns []char.Pattern
}
