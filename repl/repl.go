package repl

import (
	"bufio"
	"fmt"
	"io"
	"sepia/evaluator"
	"sepia/lexer"
	"sepia/objects"
	"sepia/parser"
)

const prompt string = "§ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	machine := objects.NewMachine()
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, machine)

		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect()+"\n")
			check(err)
		}

	}
}
func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, err := io.WriteString(out, "❌ PARSE ERROR: "+msg+"\n")
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
