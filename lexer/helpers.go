package lexer

import (
	"fmt"

	"github.com/boundedinfinity/abnfer/lang"
)

func (l *Lexer) hasPrefix(p lang.Pattern) bool {
	return lang.HasPrefix(l.rest(), p)
}

func (l *Lexer) current() string {
	return l.input[l.start:l.position]
}

func (l *Lexer) rest() string {
	return l.input[l.position:]
}

func (l *Lexer) print() {
	fmt.Printf("current:[%v], rest: [%v]\n", l.current(), l.rest())
}

func (l *Lexer) compact() {
	l.start = l.position
	l.print()
}

func (l *Lexer) next() {
	l.position++
	l.print()
}

func (l *Lexer) prev() {
	l.position--
	l.print()
}
