package util

var (
	BITS = []byte{}
)

func IsBit(v byte) bool {
	return isIn(BITS, v)
}

func init() {
	BITS = append(BITS, '0', '1')
}
