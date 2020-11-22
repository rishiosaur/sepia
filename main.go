package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"sepia/evaluator"
	"sepia/lexer"
	"sepia/objects"
	"sepia/parser"
	"sepia/repl"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) == 1 {
		user, err := user.Current()
		check(err)

		fmt.Printf("Hello %s! Welcome to the Sepia programming language.\n", user.Username)

		repl.Start(os.Stdin, os.Stdout)
	} else {
		machine := objects.NewMachine()
		file := os.Args[1]

		_data, err := ioutil.ReadFile(file)
		check(err)
		data := string(_data)

		l := lexer.New(data)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, machine)

		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "❌ PARSE ERROR: "+msg+"\n")
	}
}
