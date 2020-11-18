package lexer

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

// New creates a new Lexer and returns a reference to it.
func New(input string) *Lexer {
	lexer := Lexer{input: input}
	lexer.consumeChar()
	return &lexer
}
