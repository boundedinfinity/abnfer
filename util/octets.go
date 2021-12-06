package util

var (
	OCTETS = []byte{}
)

func init() {
	for i := 0x00; i <= 0xFF; i++ {
		OCTETS = append(OCTETS, byte(i))
	}
}

func IsOctet(v byte) bool {
	return isIn(OCTETS, v)
}
