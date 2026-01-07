package main

import (
	"fmt"

	Interpreter "github.com/AmKreta/golang-expression-parser/interpreter"
	Lexer "github.com/AmKreta/golang-expression-parser/lexer"
	Parser "github.com/AmKreta/golang-expression-parser/parser"
)

func main() {
	input := "1 + log(1024)"
	lexer := Lexer.New(input)
	parser := Parser.New(lexer)
	interpreter := Interpreter.NewInterpreter(parser)
	result := interpreter.Interpret()
	fmt.Println(result)
}
