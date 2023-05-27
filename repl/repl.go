package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ekrako/interperter-go/lexer"
	"github.com/ekrako/interperter-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text() // line is a string
		l := lexer.New(line)   // l is a pointer to a lexer.Lexer struct
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok) // tok is a token.Token struct
		}
	}
}
