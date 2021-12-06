package util

const (
	CR   = 0x0D
	LF   = 0x0A
	SP   = 0x20
	HTAB = 0x09
)

var (
	CRLF = []byte{CR, LF}
	WSP  = []byte{SP, HTAB}
	LWSP = []byte{}
)

func init() {
	LWSP = append(LWSP, WSP...)
	LWSP = append(LWSP, CRLF...)
}

// func IsCr(v byte) bool {
// 	return v == CR
// }

// func IsCrlf(v byte) bool {
// 	return isIn(CRLF, v)
// }

// func IsLf(v byte) bool {
// 	return v == LF
// }

// func IsLwsp(v byte) bool {
// 	return isIn(LWSP, v)
// }

// func IsSp(v byte) bool {
// 	return v == SP
// }

// func IsWsp(v byte) bool {
// 	return isIn(WSP, v)
// }

// func IsHtab(v byte) bool {
// 	return v == HTAB
// }
