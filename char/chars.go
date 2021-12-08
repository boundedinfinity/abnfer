package char

var (
	ALPHA  = Pattern{}
	BIT    = List(0x30, 0x31) // 0,1
	CHAR   = Range(0x01, 0x7F)
	CR     = List(0x0D)
	CTL    = Pattern{}
	DIGIT  = Range(0x30, 0x39)
	DQUOTE = List(0x22)
	HEXDIG = Pattern{}
	HTAB   = List(0x09)
	LF     = List(0x0A)
	OCTET  = Range(0x00, 0xFF)
	SP     = List(0x20)
	VCHAR  = Range(0x21, 0x7E)

	MAP = map[string]Pattern{
		"ALPHA":  ALPHA,
		"BIT":    BIT,
		"CHAR":   CHAR,
		"CR":     CR,
		"CTL":    CTL,
		"DIGIT":  DIGIT,
		"DQUOTE": DQUOTE,
		"HTAB":   HTAB,
		"LF":     LF,
		"OCTET":  OCTET,
		"SP":     SP,
		"VCHAR":  VCHAR,
	}
)

func init() {
	ALPHA = append(ALPHA, Range(0x41, 0x5A)...)
	ALPHA = append(ALPHA, Range(0x61, 0x7A)...)
	CTL = append(CTL, Range(0x00, 0x1F)...)
	CTL = append(CTL, List(0x1F)...)
	HEXDIG = append(HEXDIG, DIGIT...)
	HEXDIG = append(HEXDIG, Range(0x41, 0x46)...) // A - F
}
