package main

import (
	"fmt"

	Lexer "github.com/AmKreta/golang-expression-parser/lexer"
	Parser "github.com/AmKreta/golang-expression-parser/parser"
)

func main() {
	input := "1 + 2!"
	lexer := Lexer.New(input)
	parser := Parser.New(lexer)
	expression := parser.Parse()
	fmt.Println(expression)

	fmt.Println(expression)
}
