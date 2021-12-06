package lexer

import "github.com/boundedinfinity/abnfer/lexer/token_types"

type Lexer struct {
	Path     string
	Input    string
	current  string
	rest     string
	Ch       chan Token
	State    LexerFn
	Start    int
	Position int
}

func New(path, input string) *Lexer {
	return &Lexer{
		Path:     path,
		Input:    input,
		State:    nil,
		Ch:       make(chan Token),
		Start:    0,
		Position: 0,
	}
}

func (t *Lexer) Run() {
	for state := LexStartState; state != nil; {
		state = state(t)
	}

	t.Shutdown()
}

func (t *Lexer) Shutdown() {
	close(t.Ch)
}

func (l *Lexer) Emit(t token_types.TokenType, v string) {
	l.Ch <- Token{
		Type:  t,
		Value: v,
	}
}
