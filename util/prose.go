package util

var (
	PROSE = []byte{}
)

func IsProse(v byte) bool {
	return isIn(PROSE, v)
}

func init() {
	for i := 0x20; i <= 0x7E; i++ {
		if i != '<' && i != '>' {
			PROSE = append(PROSE, byte(i))
		}
	}
}
