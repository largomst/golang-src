package main

import (
	"antlr-calc/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// lexer
	is := antlr.NewInputStream("1+2*3")
	lexer := parser.NewExprLexer(is)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())

	}

}
