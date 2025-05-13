package token

import (
	"unicode"
	"unicode/utf8"
)

type Tokenizer struct {
	input        string
	position     int
	readPosition int
	ch           rune
	line         int
	column       int
}

func New(input string) *Tokenizer {
	t := &Tokenizer{input: input, line: 1, column: 0}
	t.readChar()
	return t
}

func (t *Tokenizer) readChar() {
	if t.readPosition >= len(t.input) {
		t.ch = 0
	} else {
		t.ch = rune(t.input[t.readPosition])
	}

	t.position = t.readPosition
	t.readPosition += utf8.RuneLen(t.ch)

	if t.ch == '\n' {
		t.line++
		t.column = 0
	} else {
		t.column++
	}
}

func (t *Tokenizer) NextToken() Token {
	var tok Token

	t.skipWhitespace()

	switch t.ch {
	case ':':
		if t.peekChar() == '=' {
			t.readChar()
			tok = Token{Type: ASSIGN, Literal: ":=", Line: t.line, Column: t.column - 1}
		} else {
			tok = newToken(ILLEGAL, t.ch, t.line, t.column)
		}
	case ';':
		tok = newToken(SEMICOLON, t.ch, t.line, t.column)
	case '(':
		tok = newToken(LPAREN, t.ch, t.line, t.column)
	case ')':
		tok = newToken(RPAREN, t.ch, t.line, t.column)
	case ',':
		tok = newToken(COMMA, t.ch, t.line, t.column)
	case '.':
		tok = newToken(DOT, t.ch, t.line, t.column)
	case '+':
		tok = newToken(PLUS, t.ch, t.line, t.column)
	case '-':
		tok = newToken(MINUS, t.ch, t.line, t.column)
	case '!':
		if t.peekChar() == '=' {
			t.readChar()
			tok = Token{Type: NEQ, Literal: "!=", Line: t.line, Column: t.column - 1}
		} else {
			tok = newToken(BANG, t.ch, t.line, t.column)
		}
	case '&':
		if t.peekChar() == '&' {
			t.readChar()
			tok = Token{Type: AND, Literal: "&&", Line: t.line, Column: t.column - 1}
		}
	case '|':
		if t.peekChar() == '|' {
			t.readChar()
			tok = Token{Type: OR, Literal: "||", Line: t.line, Column: t.column - 1}
		}
	case '*':
		tok = newToken(ASTERISK, t.ch, t.line, t.column)
	case '/':
		tok = newToken(SLASH, t.ch, t.line, t.column)
	case '<':
		tok = newToken(LT, t.ch, t.line, t.column)
	case '>':
		tok = newToken(GT, t.ch, t.line, t.column)
	case '=':
		tok = newToken(EQ, t.ch, t.line, t.column)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if unicode.IsLetter(t.ch) {
			tok.Line = t.line
			tok.Column = t.column
			tok.Literal = t.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if unicode.IsDigit(t.ch) {
			tok.Line = t.line
			tok.Column = t.column
			tok.Literal = t.readNumber()
			tok.Type = NUMBER
			return tok
		} else {
			tok = newToken(ILLEGAL, t.ch, t.line, t.column)
		}
	}

	t.readChar()
	return tok
}

func (t *Tokenizer) skipWhitespace() {
	for {
		switch t.ch {
		case ' ', '\t', '\n', '\r':
			t.readChar()
		case '/':
			if t.peekChar() == '/' {
				// 跳过单行注释
				for t.ch != '\n' && t.ch != 0 {
					t.readChar()
				}
			} else {
				return
			}
		default:
			return
		}
	}
}

func (t *Tokenizer) readIdentifier() string {
	position := t.position
	for unicode.IsLetter(t.ch) || unicode.IsDigit(t.ch) || t.ch == '_' {
		t.readChar()
	}
	return t.input[position:t.position]
}

func (t *Tokenizer) readNumber() string {
	position := t.position
	for unicode.IsDigit(t.ch) {
		t.readChar()
	}
	return t.input[position:t.position]
}

func (t *Tokenizer) peekChar() rune {
	if t.readPosition >= len(t.input) {
		return 0
	}
	return rune(t.input[t.readPosition])
}

func newToken(tokenType TokenType, ch rune, line, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    line,
		Column:  column,
	}
}
