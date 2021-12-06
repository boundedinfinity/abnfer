package util

var (
	HEXDIGS = []byte{}
)

func IsHexdig(v byte) bool {
	return isIn(HEXDIGS, v)
}

func init() {

	for i := 0x30; i <= 0x39; i++ {
		HEXDIGS = append(HEXDIGS, DIGITS...)
		HEXDIGS = append(HEXDIGS, 'A', 'B', 'C', 'D', 'E', 'F')
	}
}
