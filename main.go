package main

import (
	"fmt"

	Lexer "github.com/AmKreta/golang-expression-parser/lexer"
	Token "github.com/AmKreta/golang-expression-parser/token"
)

func main() {
	input := "1 + 2 * 3 /4  * sin(45)"
	lexer := Lexer.New(input)

	for {
		token := lexer.GetNextToken()
		if token.Type == Token.EOF {
			break
		}
		fmt.Println(token)
	}
}
