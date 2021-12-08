package lexer

// func LexStartState(l *Lexer) LexerFn {
// 	for {
// 		l.next()

// 		switch {
// 		case l.isEof() || l.rest == "":
// 			l.Emit(token_types.Eof, "")
// 			return nil
// 		case l.hasPrefix(util.CRLF):
// 			fallthrough
// 		case l.hasPrefixb(util.LF):
// 			fallthrough
// 		case l.hasAnyOfPrefix(util.WSP):
// 			l.reset()
// 		case l.hasPrefixb(';'):
// 			return LexComment
// 		default:
// 			return LexError
// 		}
// 	}
// }
