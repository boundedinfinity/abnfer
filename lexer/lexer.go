package lexer

import (
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
