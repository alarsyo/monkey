package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/alarsyo/monkey/lexer"
	"github.com/alarsyo/monkey/token"
)

// PROMPT displayed to user in REPL
const PROMPT = ">> "

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
