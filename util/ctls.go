package util

const (
	NUL    = 0x00
	EOF    = 0x00
	DQUOTE = 0x22
	EQUAL  = '='
	DASH   = '-'
)

var (
	CTLS = []byte{}
)

func init() {
	for i := 0x00; i <= 0x1F; i++ {
		CTLS = append(CTLS, byte(i))
	}
}

func IsEqual(v byte) bool {
	return v == EQUAL
}

func IsCtl(v byte) bool {
	return isIn(CTLS, v)
}

func IsDquote(v byte) bool {
	return v == DQUOTE
}
