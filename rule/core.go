package rule

import (
	"github.com/boundedinfinity/abnfer/char"
)

var (
	ALPHA = Rule{
		Name: "ALPHA",
		Min:  1,
		Max:  1,
		Pattern: AnyOf(
			char.Concat(
				char.Range(char.UPPERCASE_A, char.UPPERCASE_Z),
				char.Range(char.LOWERCASE_A, char.LOWERCASE_Z),
			),
		),
	}

	BIT = Rule{
		Name:    "BIT",
		Min:     1,
		Max:     1,
		Pattern: AnyOf(char.String("0"), char.String("1")),
	}

	CHAR = Rule{
		Name: "CHAR",
		Min:  1,
		Max:  1,
		Pattern: AnyOf(
			char.Range(char.START_OF_HEADER, char.DELETE),
		),
	}

	CR = Rule{
		Name:    "CR",
		Min:     1,
		Max:     1,
		Pattern: AnyOf(char.Single(char.CARRIAGE_RETURN)),
	}

	LF = Rule{
		Name:    "LF",
		Min:     1,
		Max:     1,
		Pattern: AnyOf(char.Single(char.LINE_FEED)),
	}

	CRLF = Rule{
		Name:    "CRLF",
		Min:     1,
		Max:     1,
		Pattern: Concat(CR, LF),
	}
)

func NewRules() Rules {
	var rm Rules

	rm.Append(ALPHA)
	rm.Append(BIT)
	rm.Append(CHAR)
	rm.Append(CRLF)

	return rm
}
