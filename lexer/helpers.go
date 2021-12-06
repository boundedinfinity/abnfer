package lexer

import (
	"bytes"
)

func (l *Lexer) hasPrefixb(b byte) bool {
	return l.hasPrefix([]byte{b})
}

func (l *Lexer) hasPrefix(bs []byte) bool {
	return bytes.HasPrefix([]byte(l.rest), bs)
}

func (l *Lexer) hasAnyOfPrefix(bss ...[]byte) bool {
	rest := []byte(l.rest)

	for _, bs := range bss {
		for _, b := range bs {
			if bytes.HasPrefix(rest, []byte{b}) {
				return true
			}
		}
	}

	return false
}

func (l *Lexer) assign() {
	l.current = l.Input[l.Start:l.Position]
	l.rest = l.Input[l.Position:]
}

func (l *Lexer) reset() {
	l.Start = l.Position
	l.assign()
}

func (l *Lexer) isEof() bool {
	ln := len(l.Input)
	return l.Position > ln
}

func (l *Lexer) isBof() bool {
	return l.Position < 0
}

func (l *Lexer) next() {
	l.Position++
	l.assign()
}

func (l *Lexer) prev() {
	l.Position--
	l.assign()
}
