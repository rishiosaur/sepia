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
	"strings"
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
		if err != nil || !strings.HasSuffix(file, ".sp") {
			fmt.Println("START ERROR")
		}
		data := string(_data)

		l := lexer.New(data)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			return
		}

		evaluator.Eval(program, machine)
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, err := io.WriteString(out, "❌ PARSE ERROR: "+msg+"\n")
		check(err)
	}
}
