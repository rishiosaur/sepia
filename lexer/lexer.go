package lexer

import (
	"monkey-go/token"
	"monkey-go/util"
)

// Lexer is a structure that lexes a given input.
type Lexer struct {
	input           string
	position        int
	readingPosition int
	currentChar     byte
}

func (lexer *Lexer) consumeChar() {
	if lexer.readingPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readingPosition]
	}

	lexer.position = lexer.readingPosition
	lexer.readingPosition++

}

func (lexer *Lexer) nextToken() token.Token {
	var t token.Token

	lexer.skipWhitespace()

	switch lexer.currentChar {
	case '=':
		t = newToken(token.ASSIGN, lexer.currentChar)

	case ';':
		t = newToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		t = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		t = newToken(token.RPAREN, lexer.currentChar)
	case ',':
		t = newToken(token.COMMA, lexer.currentChar)
	case '+':
		t = newToken(token.PLUS, lexer.currentChar)
	case '{':

		t = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		t = newToken(token.RBRACE, lexer.currentChar)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if util.IsLetter(lexer.currentChar) {
			t.Literal = lexer.consumeIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if util.IsDigit(lexer.currentChar) {
			t.Literal = lexer.consumeInteger()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, lexer.currentChar)
		}
	}
	lexer.consumeChar()

	return t
}

func (lexer *Lexer) skipWhitespace() {
	for util.IsWhitespace(lexer.currentChar) {
		lexer.consumeChar()
	}
}

func (lexer *Lexer) consumeInteger() string {
	position := lexer.position
	for util.IsDigit(lexer.currentChar) {
		lexer.consumeChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) consumeIdentifier() string {
	position := lexer.position
	for util.IsLetter(lexer.currentChar) {
		lexer.consumeChar()
	}

	return lexer.input[position:lexer.position]
}

func newToken(tokenType token.Type, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

// New creates a new Lexer and returns a reference to it.
func New(input string) *Lexer {
	lexer := Lexer{input: input}
	lexer.consumeChar()
	return &lexer
}
