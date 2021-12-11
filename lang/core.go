package lang

import "github.com/boundedinfinity/abnfer/char"

type Pattern [][]char.Char

func (t *Pattern) CharRange(s, e char.Char) {
	for _, c := range char.Range(s, e) {
		*t = append(*t, char.Single(c))
	}
}

func (t *Pattern) CharAnyOf(cs ...char.Char) {
	*t = append(*t, cs)
}

func (t *Pattern) AppendPattern(ps ...Pattern) {
	for _, p1 := range ps {
		*t = append(*t, p1...)
	}
}

var (
	ALPHA Pattern
	BIT   Pattern
	CHAR  Pattern
	CR    Pattern
	LF    Pattern
	CRLF  Pattern
	VCHAR Pattern
	SP    Pattern
	HTAB  Pattern
	WSP   Pattern
	LWSP  Pattern
)

func init() {
	ALPHA.CharRange(char.UPPERCASE_A, char.UPPERCASE_Z)
	ALPHA.CharRange(char.LOWERCASE_A, char.LOWERCASE_Z)

	BIT.CharAnyOf(char.ZERO)
	BIT.CharAnyOf(char.ONE)

	CHAR.CharRange(char.START_OF_HEADER, char.DELETE)

	CR.CharAnyOf(char.CARRIAGE_RETURN)

	LF.CharAnyOf(char.LINE_FEED)

	CRLF.AppendPattern(LF)
	CRLF.AppendPattern(CR, LF)

	VCHAR.CharRange(char.EXCLAMATION_MARK, char.TILDE)

	SP.CharAnyOf(char.SPACE)

	HTAB.CharAnyOf(char.HORIZONTAL_TAB)

	WSP.AppendPattern(SP, HTAB)

	LWSP.AppendPattern(WSP)
	LWSP.AppendPattern(CRLF, WSP)
}
