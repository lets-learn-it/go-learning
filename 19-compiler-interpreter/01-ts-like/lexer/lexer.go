package lexer

import "fmt"

func (l *Lexer) Tokenize() {
	for l.Pos < l.Len {
		l.EatWhitespaces()
		switch l.Src[l.Pos] {
		case '(':
			l.Tok = append(l.Tok, NewToken(string(l.Src[l.Pos]), OpenParen))
			l.Pos++
		case ')':
			l.Tok = append(l.Tok, NewToken(string(l.Src[l.Pos]), CloseParen))
			l.Pos++
		case '+', '-', '*', '/':
			l.Tok = append(l.Tok, NewToken(string(l.Src[l.Pos]), BinaryOperator))
			l.Pos++
		case '=':
			l.Tok = append(l.Tok, NewToken(string(l.Src[l.Pos]), Equals))
			l.Pos++
		default:
			if l.isNum() {
				start := l.Pos
				l.Pos++
				for l.Pos < l.Len && l.isNum() {
					l.Pos++
				}
				l.Tok = append(l.Tok, NewToken(string(l.Src[start:l.Pos]), Number))
			} else if l.isAlpha() {
				start := l.Pos
				l.Pos++
				for l.Pos < l.Len && l.isAlpha() || l.isNum() {
					l.Pos++
				}
				val := string(l.Src[start:l.Pos])
				// check for reserved keywords
				ttype, exists := KEYWORDS[val]
				if exists {
					l.Tok = append(l.Tok, NewToken(val, ttype))
				} else {
					l.Tok = append(l.Tok, NewToken(val, Identifier))
				}
			} else {
				fmt.Println(l)
				panic(l)
			}
		}
	}
	l.Tok = append(l.Tok, NewToken("EndOfFile", EOF))
}

func (l *Lexer) EatWhitespaces() {
	for l.Src[l.Pos] == ' ' || l.Src[l.Pos] == '\n' || l.Src[l.Pos] == '\t' || l.Src[l.Pos] == '\r' {
		l.Pos++
	}
}

func (l *Lexer) isAlpha() bool {
	if (l.Src[l.Pos] >= 'a' && l.Src[l.Pos] <= 'z') || (l.Src[l.Pos] >= 'A' && l.Src[l.Pos] <= 'Z') {
		return true
	}
	return false
}

func (l *Lexer) isNum() bool {
	if l.Src[l.Pos] >= '0' && l.Src[l.Pos] <= '9' {
		return true
	}
	return false
}

func (l *Lexer) Print() {
	for _, v := range l.Tok {
		fmt.Println(v)
	}
}
