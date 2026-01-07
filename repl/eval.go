package repl

import (
	"fmt"
	"os"

	Ast "github.com/AmKreta/golang-expression-parser/ast"
	Interpreter "github.com/AmKreta/golang-expression-parser/interpreter"
	Lexer "github.com/AmKreta/golang-expression-parser/lexer"
	Parser "github.com/AmKreta/golang-expression-parser/parser"
)

func eval(input string) Ast.ComputeSize {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()
	lexer := Lexer.New(input)
	parser := Parser.New(lexer)
	interpreter := Interpreter.NewInterpreter(parser)
	result := interpreter.Interpret()
	return result
}
