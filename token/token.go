package token

import "fmt"

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType) *Token {
	val, ok := TokenTypeValMap[tokenType]
	if !ok {
		panic(fmt.Sprintf("token type %d not found", tokenType))
	}
	return &Token{Type: tokenType, Literal: val}
}

func NewTokenWithLiteral(tokenType TokenType, literal string) *Token {
	_, ok := TokenTypeValMap[tokenType]
	if !ok {
		panic(fmt.Sprintf("token type %d not found", tokenType))
	}
	return &Token{Type: tokenType, Literal: literal}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Literal: %s}", TokenNames[t.Type], t.Literal)
}
