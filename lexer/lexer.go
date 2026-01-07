package lexer

import (
	"fmt"

	"github.com/AmKreta/golang-expression-parser/token"
)

type Lexer struct {
	input       string
	position    int
	inputLength int
}

func (lexer *Lexer) jump() {
	lexer.position++
}

func (lexer *Lexer) jumpN(n int) {
	lexer.position += n
}

func (lexer *Lexer) peek() string {
	return string(lexer.input[lexer.position+1])
}

func (lexer *Lexer) peekN(n int) string {
	return string(lexer.input[lexer.position : lexer.position+n])
}

func (lexer *Lexer) getCurrentChar() string {
	if lexer.position >= lexer.inputLength {
		return ""
	}
	return string(lexer.input[lexer.position])
}

func (lexer *Lexer) isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

func (lexer *Lexer) isLetter(char string) bool {
	return (char >= "a" && char <= "z") || (char >= "A" && char <= "Z")
}

func (lexer *Lexer) isSpace(char string) bool {
	return char == " "
}

func (lexer *Lexer) readNumber() string {
	num := ""
	currentchar := lexer.getCurrentChar()
	for lexer.isDigit(currentchar) {
		num += currentchar
		lexer.jump()
		if lexer.position >= lexer.inputLength {
			break
		}
		currentchar = lexer.getCurrentChar()
	}
	return num
}

func (lexer *Lexer) readIdentifier() string {
	identifier := ""
	currentChar := lexer.getCurrentChar()
	for lexer.isLetter(currentChar) {
		identifier += currentChar
		lexer.jump()
		if lexer.position >= lexer.inputLength {
			break
		}
		currentChar = lexer.getCurrentChar()
	}
	return identifier
}

func (lexer *Lexer) skipSpace() {
	currentchar := lexer.getCurrentChar()
	for lexer.isSpace(currentchar) {
		lexer.jump()
		if lexer.position >= lexer.inputLength {
			break
		}
		currentchar = lexer.getCurrentChar()
	}
}

func (lexer *Lexer) GetNextToken() *token.Token {
	if lexer.position >= lexer.inputLength {
		return token.NewToken(token.EOF)
	}
	lexer.skipSpace()
	currentchar := lexer.getCurrentChar()

	switch currentchar {
	case token.TokenTypeValMap[token.PLUS]:
		lexer.jump()
		return token.NewToken(token.PLUS)
	case token.TokenTypeValMap[token.MINUS]:
		lexer.jump()
		return token.NewToken(token.MINUS)
	case token.TokenTypeValMap[token.MULTIPLY]:
		lexer.jump()
		return token.NewToken(token.MULTIPLY)
	case token.TokenTypeValMap[token.DIVIDE]:
		if lexer.peek() == "/" {
			lexer.jumpN(2)
			return token.NewToken(token.INTEGER_DIVIDE)
		}
		lexer.jump()
		return token.NewToken(token.DIVIDE)
	case token.TokenTypeValMap[token.MODULO]:
		lexer.jump()
		return token.NewToken(token.MODULO)
	case token.TokenTypeValMap[token.POWER]:
		lexer.jump()
		return token.NewToken(token.POWER)
	case token.TokenTypeValMap[token.FACTORIAL]:
		lexer.jump()
		return token.NewToken(token.FACTORIAL)
	case token.TokenTypeValMap[token.PARENTHESIS_OPEN]:
		lexer.jump()
		return token.NewToken(token.PARENTHESIS_OPEN)
	case token.TokenTypeValMap[token.PARENTHESIS_CLOSE]:
		lexer.jump()
		return token.NewToken(token.PARENTHESIS_CLOSE)
	case token.TokenTypeValMap[token.Square_Bracket_Open]:
		lexer.jump()
		return token.NewToken(token.Square_Bracket_Open)
	case token.TokenTypeValMap[token.Square_Bracket_Close]:
		lexer.jump()
		return token.NewToken(token.Square_Bracket_Close)
	case token.TokenTypeValMap[token.Curly_Bracket_Open]:
		lexer.jump()
		return token.NewToken(token.Curly_Bracket_Open)
	case token.TokenTypeValMap[token.Curly_Bracket_Close]:
		lexer.jump()
		return token.NewToken(token.Curly_Bracket_Close)
	}

	isDigit := lexer.isDigit(currentchar)
	if isDigit {
		num := lexer.readNumber()
		return token.NewTokenWithLiteral(token.NUMBER, num)
	}

	isLetter := lexer.isLetter(currentchar)
	if isLetter {
		identifier := lexer.readIdentifier()
		switch identifier {
		case token.TokenTypeValMap[token.LOGARITHM]:
			return token.NewToken(token.LOGARITHM)
		case token.TokenTypeValMap[token.SINE]:
			return token.NewToken(token.SINE)
		case token.TokenTypeValMap[token.COSINE]:
			return token.NewToken(token.COSINE)
		case token.TokenTypeValMap[token.TANGENT]:
			return token.NewToken(token.TANGENT)
		case token.TokenTypeValMap[token.COTANGENT]:
			return token.NewToken(token.COTANGENT)
		case token.TokenTypeValMap[token.SECANT]:
			return token.NewToken(token.SECANT)
		case token.TokenTypeValMap[token.COSECANT]:
			return token.NewToken(token.COSECANT)
		case token.TokenTypeValMap[token.PI]:
			return token.NewToken(token.PI)
		}
	}

	panic(fmt.Sprintf("illegal character: %s at position %d", currentchar, lexer.position))
}
