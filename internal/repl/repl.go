package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/prafitradimas/interpreter/internal/lexer"
	"github.com/prafitradimas/interpreter/internal/token"
)

const (
	PROMPT = ">> "
)

func Start(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		lexer := lexer.New(scanner.Text())
		for _token := lexer.NextToken(); _token.Type != token.EOF; _token = lexer.NextToken() {
			fmt.Printf("%+v\n", _token)
		}
	}
}
