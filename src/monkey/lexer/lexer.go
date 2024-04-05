package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NE, Literal: "!="}
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if l.isLetter(l.ch) {
			literal := l.readIdentifier()
			typ := token.LookupTokenType(literal)
			tok = token.Token{Type: typ, Literal: literal}
			return tok
		} else if l.isDigit(l.ch) {
			literal := l.readNumber()
			tok = token.Token{Type: token.INT, Literal: literal}
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: ""}
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
