package lexer

import "github.com/ytakaya/nasparse/token"

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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '<':
		tokenType, ident := l.readTag()
		tok.Type = tokenType
		tok.Literal = ident
	case '(':
		if l.peekChar() == '#' {
			tok.Type = token.NaspToken
			tok.Literal = l.readNasp()
		} else {
			tok.Type = token.TextToken
			tok.Literal = l.readText()
		}
	default:
		tok.Type = token.TextToken
		tok.Literal = l.readText()
	}

	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
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

func (l *Lexer) readTag() (token.TokenType, string) {
	var t token.TokenType
	if l.peekChar() == '/' {
		t = token.EndTagToken
	}

	prevChar := l.ch
	position := l.position
	for l.ch != '>' {
		prevChar = l.ch
		l.readChar()
	}
	l.readChar()

	if prevChar == '/' {
		t = token.SelfClosingTagToken
	}
	if t == token.ErrorToken {
		t = token.StartTagToken
	}
	return t, l.input[position:l.position]
}

func (l *Lexer) readText() string {
	position := l.position
	for !isWhiteSpace(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNasp() string {
	position := l.position
	for l.ch != ')' {
		l.readChar()
	}
	l.readChar()
	return l.input[position:l.position]
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
