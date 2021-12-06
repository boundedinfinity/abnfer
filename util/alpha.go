package util

var (
	ALPHAS = []byte{}
)

func IsAlpha(v byte) bool {
	return isIn(ALPHAS, v)
}

func init() {
	for i := 0x41; i <= 0x5A; i++ {
		ALPHAS = append(ALPHAS, byte(i))
	}
}
