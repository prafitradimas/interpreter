package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/prafitradimas/interpreter/internal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s!", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
