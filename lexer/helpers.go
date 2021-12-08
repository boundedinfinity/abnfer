package lexer

// func (l *Lexer) hasPrefixb(b byte) bool {
// 	return l.hasPrefix([]byte{b})
// }

// func (l *Lexer) hasPrefix(bs []byte) bool {
// 	return bytes.HasPrefix([]byte(l.rest), bs)
// }

// func (l *Lexer) hasAnyOfPrefix(bss ...[]byte) bool {
// 	rest := []byte(l.rest)

// 	for _, bs := range bss {
// 		for _, b := range bs {
// 			if bytes.HasPrefix(rest, []byte{b}) {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func (l *Lexer) assign() {
// 	l.current = l.input[l.start:l.position]
// 	l.rest = l.input[l.position:]
// }

// func (l *Lexer) reset() {
// 	l.start = l.position
// 	l.assign()
// }

// func (l *Lexer) isEof() bool {
// 	ln := len(l.input)
// 	return l.position > ln
// }

// func (l *Lexer) isBof() bool {
// 	return l.position < 0
// }

// func (l *Lexer) next() {
// 	l.position++
// 	l.assign()
// }

// func (l *Lexer) prev() {
// 	l.position--
// 	l.assign()
// }
