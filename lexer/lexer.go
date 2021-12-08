package lexer

// type Lexer struct {
// 	Path     string
// 	input    string
// 	current  string
// 	rest     string
// 	Ch       chan Token
// 	state    LexerFn
// 	start    int
// 	position int
// }

// func New(path, input string) *Lexer {
// 	return &Lexer{
// 		Path:     path,
// 		input:    input,
// 		state:    nil,
// 		Ch:       make(chan Token),
// 		start:    0,
// 		position: 0,
// 	}
// }

// func (t *Lexer) Run() {
// 	for state := LexStartState; state != nil; {
// 		state = state(t)
// 	}

// 	t.Shutdown()
// }

// func (t *Lexer) Shutdown() {
// 	close(t.Ch)
// }

// func (l *Lexer) Emit(t token_types.TokenType, v string) {
// 	l.Ch <- Token{
// 		Type:  t,
// 		Value: v,
// 	}
// }
