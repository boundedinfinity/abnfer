package lexer

import (
	"fmt"

	"github.com/boundedinfinity/abnfer/tokens"
)

func LexError(l *Lexer) LexerFn {
	l.Emit(
		tokens.SyntaxError,
		fmt.Sprintf("syntax error starting at %v", l.current()),
	)

	return nil
}

func LexBegin(l *Lexer) LexerFn {
	l.Emit(tokens.FileStart, "")
	return LexLine
}

func LexLine(l *Lexer) LexerFn {
	l.print()
	for {
		switch {
		// case l.hasPrefixc(char.SPACE):
		// 	l.next()
		// case l.hasPrefixc(char.SEMICOLON):
		// 	return LexComment
		default:
			return LexError
		}
	}
}

func LexComment(l *Lexer) LexerFn {
	for {
		l.next()

		switch {
		// case l.hasPrefixc(char.LINE_FEED):
		// 	fallthrough
		// case l.hasPrefixc(char.CARRIAGE_RETURN, char.LINE_FEED):
		// 	l.Emit(tokens.Comment, l.current())
		// 	l.compact()
		// 	return LexLine

		// case l.hasPrefixc(char.SPACE):
		// 	// ignore
		default:
			return LexError
		}
	}
}
