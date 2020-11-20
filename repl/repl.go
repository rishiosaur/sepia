package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-go/lexer"
	"monkey-go/parser"
)

const prompt string = "#> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }

		p := parser.New(l)

		program := p.ParseProgram().String()

		// if len(p.Errors()) == 0 {
		fmt.Println(program)
		// }
	}
}
