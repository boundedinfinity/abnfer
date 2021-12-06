package lexer

import (
	"fmt"

	"github.com/boundedinfinity/abnfer/lexer/token_types"
	"github.com/boundedinfinity/abnfer/util"
)

func LexError(l *Lexer) LexerFn {
	l.Emit(
		token_types.SyntaxError,
		fmt.Sprintf("syntax error starting at %v", l.rest),
	)

	return nil
}

func LexStartState(l *Lexer) LexerFn {
	for {
		l.next()

		switch {
		case l.isEof() || l.rest == "":
			l.Emit(token_types.Eof, "")
			return nil
		case l.hasPrefix(util.CRLF):
			fallthrough
		case l.hasPrefixb(util.LF):
			fallthrough
		case l.hasAnyOfPrefix(util.WSP):
			l.reset()
		case l.hasPrefixb(';'):
			return LexComment
		default:
			return LexError
		}
	}
}

func LexComment(l *Lexer) LexerFn {
	l.next()
	l.reset()

	for {
		l.next()

		switch {
		case l.hasAnyOfPrefix(util.WSP, util.VCHARS):
			// ignore
		case l.hasAnyOfPrefix([]byte{util.LF}):
			fallthrough
		case l.hasPrefix(util.CRLF):
			l.Emit(token_types.Comment, l.current)
			return LexStartState
		default:
			return LexError
		}
	}
}
