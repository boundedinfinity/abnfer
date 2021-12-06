package util

var (
	VCHARS = []byte{}
)

func IsVchar(v byte) bool {
	return isIn(VCHARS, v)
}

func init() {
	for i := 0x21; i <= 0x7E; i++ {
		VCHARS = append(VCHARS, byte(i))
	}
}
