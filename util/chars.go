package util

var (
	CHARS = []byte{}
)

func IsChar(v byte) bool {
	return isIn(CHARS, v)
}

func init() {
	for i := 0x01; i <= 0x7F; i++ {
		CHARS = append(CHARS, byte(i))
	}
}
