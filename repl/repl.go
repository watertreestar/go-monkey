package repl

import (
	"bufio"
	"fmt"
	"github.com/watertreestar/go-monkey/lexer"
	"github.com/watertreestar/go-monkey/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader,out io.Writer)  {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken();tok.Type != token.EOF;tok = l.NextToken() {
			fmt.Printf("%+v\n",tok)
		}
	}
}
