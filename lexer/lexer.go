package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	filename     string
	line         uint
	ch           byte // current char under examination”
	// todo: use rune instead of u8
}

func New(input string, filename string) *Lexer {
	l := &Lexer{input: input, filename: filename}
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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	tok.Filename = l.filename
	tok.Line = l.line

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = newToken(token.EQ, literal, l.filename, l.line)
		} else {
			tok = newToken(token.ASSIGN, string(l.ch), l.filename, l.line)
		}
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch), l.filename, l.line)
	case '(':
		tok = newToken(token.LPAREN, string(l.ch), l.filename, l.line)
	case ')':
		tok = newToken(token.RPAREN, string(l.ch), l.filename, l.line)
	case ',':
		tok = newToken(token.COMMA, string(l.ch), l.filename, l.line)
	case '+':
		tok = newToken(token.PLUS, string(l.ch), l.filename, l.line)
	case '{':
		tok = newToken(token.LBRACE, string(l.ch), l.filename, l.line)
	case '}':
		tok = newToken(token.RBRACE, string(l.ch), l.filename, l.line)
	case '-':
		tok = newToken(token.MINUS, string(l.ch), l.filename, l.line)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = newToken(token.NOT_EQ, literal, l.filename, l.line)
		} else {
			tok = newToken(token.BANG, string(l.ch), l.filename, l.line)
		}
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch), l.filename, l.line)
	case '/':
		tok = newToken(token.SLASH, string(l.ch), l.filename, l.line)
	case '<':
		tok = newToken(token.LT, string(l.ch), l.filename, l.line)
	case '>':
		tok = newToken(token.GT, string(l.ch), l.filename, l.line)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch), l.filename, l.line)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch string, filename string, line uint) token.Token {
	return token.Token{Type: tokenType, Literal: ch, Filename: filename, Line: line}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	// TODO: add "?" as a valid identifier name
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line += 1
		}
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
