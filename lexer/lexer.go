package lexer

import (
	"strings"

	"github.com/boundedinfinity/abnfer/char"
	"github.com/boundedinfinity/abnfer/tokens"
)

type Event struct {
	Type  tokens.Token
	Value string
}

type LexerFn func(*Lexer) LexerFn

type Lexer struct {
	Path     string
	input    string
	current  string
	Events   chan Event
	state    LexerFn
	start    int
	position int
}

func New(path, input string) *Lexer {
	return &Lexer{
		Path:     path,
		input:    input,
		state:    nil,
		Events:   make(chan Event),
		start:    0,
		position: 0,
	}
}

func (t *Lexer) Run() {
	for state := LexBegin; state != nil; {
		state = state(t)
	}

	t.Shutdown()
}

func (t *Lexer) Shutdown() {
	close(t.Events)
}

func (l *Lexer) Emit(t tokens.Token, v string) {
	l.Events <- Event{
		Type:  t,
		Value: v,
	}
}

func (l *Lexer) hasPrefix(s string) bool {
	return strings.HasPrefix(l.input[l.position:], s)
}

func (l *Lexer) hasPrefixc(cs ...char.Char) bool {
	var bs []byte

	for _, c := range cs {
		bs = append(bs, byte(c))
	}

	return l.hasPrefix(string(bs))
}

func (l *Lexer) reset() {
	l.start = l.position
	l.current = ""
}

func (l *Lexer) next() {
	l.position++
	l.current = l.input[l.start:l.position]
}

func (l *Lexer) prev() {
	l.position--
	l.current = l.input[l.start:l.position]
}
