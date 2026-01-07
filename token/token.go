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

func (t *Token) IsAdditiveTypeToken() bool {
	return t.Type == PLUS || t.Type == MINUS
}

func (t *Token) IsMultiplicativeTypeToken() bool {
	return t.Type == MULTIPLY || t.Type == DIVIDE || t.Type == MODULO || t.Type == INTEGER_DIVIDE || t.Type == POWER
}

func (t *Token) IsUnaryTypeToken() bool {
	return t.Type == LOGARITHM || t.Type == FACTORIAL || t.Type == SINE || t.Type == COSINE || t.Type == TANGENT || t.Type == COTANGENT || t.Type == SECANT || t.Type == COSECANT
}

func (t *Token) isFactorTypeToken() bool {
	return t.Type == NUMBER || t.Type == PI || t.Type == PARENTHESIS_OPEN || t.Type == Square_Bracket_Open || t.Type == Curly_Bracket_Open
}
