package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-go/lexer"
	"monkey-go/token"
)

const PROMPT_START = "#> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT_START)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.nextToken(); tok.Type != token.EOF; tok = l.nextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
