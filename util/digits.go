package util

var (
	DIGITS = []byte{}
)

func init() {
	for i := 0x30; i <= 0x39; i++ {
		DIGITS = append(DIGITS, byte(i))
	}
}

func IsDigit(v byte) bool {
	return isIn(DIGITS, v)
}
