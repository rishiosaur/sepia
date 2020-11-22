package main

import (
	"fmt"
	"os"
	"os/user"
	"sepia/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Sepia programming language.!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
