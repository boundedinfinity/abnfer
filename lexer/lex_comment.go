package lexer

// func LexComment(l *Lexer) LexerFn {
// 	l.next()
// 	l.reset()

// 	for {
// 		l.next()

// 		switch {
// 		case l.hasAnyOfPrefix(util.WSP, util.VCHARS):
// 			// ignore
// 		case l.hasAnyOfPrefix([]byte{util.LF}):
// 			fallthrough
// 		case l.hasPrefix(util.CRLF):
// 			l.Emit(token_types.Comment, l.current)
// 			return LexStartState
// 		default:
// 			return LexError
// 		}
// 	}
// }
