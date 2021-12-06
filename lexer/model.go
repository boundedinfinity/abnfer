package lexer

import (
	"github.com/boundedinfinity/abnfer/lexer/token_types"
)

//go:generate enumer -standalone=true -package=token_types -name=TokenType -items-from=token-type.txt

type LexerFn func(*Lexer) LexerFn

type Token struct {
	Type  token_types.TokenType
	Value string
}
