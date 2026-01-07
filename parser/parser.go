package parser

import (
	"fmt"
	"math"
	"strconv"

	"github.com/AmKreta/golang-expression-parser/ast"
	Lexer "github.com/AmKreta/golang-expression-parser/lexer"
	Token "github.com/AmKreta/golang-expression-parser/token"
)

type Parser struct {
	lexer        *Lexer.Lexer
	currentToken *Token.Token
}

func New(lexer *Lexer.Lexer) *Parser {
	return &Parser{lexer: lexer, currentToken: lexer.GetNextToken()}
}

func (parser *Parser) eat(token *Token.Token) {
	if parser.currentToken.Type == token.Type {
		parser.currentToken = parser.lexer.GetNextToken()
	} else {
		panic(fmt.Sprintf("expected token type %s, got %s", Token.TokenNames[token.Type], Token.TokenNames[parser.currentToken.Type]))
	}
}

func (parser *Parser) parseFactor() ast.Node {
	switch parser.currentToken.Type {
	case Token.NUMBER:
		numberValue, err := strconv.ParseFloat(parser.currentToken.Literal, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse number %s", parser.currentToken.Literal))
		}
		number := ast.NewNumber(ast.ComputeSize(numberValue))
		parser.eat(parser.currentToken)
		return number
	case Token.PI:
		pi := ast.NewNumber(math.Pi)
		parser.eat(parser.currentToken)
		return pi
	case Token.PARENTHESIS_OPEN:
		parser.eat(parser.currentToken)
		expression := parser.parseExpression()
		parser.eat(parser.currentToken)
		return expression
	case Token.Square_Bracket_Open:
		parser.eat(parser.currentToken)
		expression := parser.parseExpression()
		parser.eat(parser.currentToken)
		return expression
	case Token.Curly_Bracket_Open:
		parser.eat(parser.currentToken)
		expression := parser.parseExpression()
		parser.eat(parser.currentToken)
		return expression
	default:
		panic(fmt.Sprintf("unexpected token type %s", Token.TokenNames[parser.currentToken.Type]))
	}
}

func (parser *Parser) parseUnaryExpression() ast.Node {
	var unary ast.Node

	if parser.currentToken.IsPrefixUnaryTypeToken() {
		op := parser.currentToken
		parser.eat(op)
		if parser.currentToken.IsBracketOpenTypeToken() {
			parser.eat(parser.currentToken)
			expression := parser.parseExpression()
			if parser.currentToken.IsBracketCloseTypeToken() {
				parser.eat(parser.currentToken)
			} else {
				panic(fmt.Sprintf("expected token type %s, got %s", Token.TokenNames[Token.PARENTHESIS_CLOSE], Token.TokenNames[parser.currentToken.Type]))
			}
			unary = ast.NewUnaryOperator(op.Type, expression)
		} else {
			panic(fmt.Sprintf("expected token type %s, got %s", Token.TokenNames[Token.PARENTHESIS_OPEN], Token.TokenNames[parser.currentToken.Type]))
		}
	} else {
		unary = parser.parseFactor()
		for parser.currentToken.IsSuffixUnaryTypeToken() {
			op := parser.currentToken
			parser.eat(op)
			unary = ast.NewUnaryOperator(op.Type, unary)
		}
	}

	return unary
}

func (parser *Parser) parseMultiplicativeExpression() ast.Node {
	left := parser.parseUnaryExpression()
	for parser.currentToken.IsMultiplicativeTypeToken() {
		op := parser.currentToken
		parser.eat(op)
		right := parser.parseUnaryExpression()
		left = ast.NewBinaryOperator(left, right, op.Type)
	}
	return left
}

func (parser *Parser) parseAdditiveExpression() ast.Node {
	left := parser.parseMultiplicativeExpression()
	for parser.currentToken.IsAdditiveTypeToken() {
		op := parser.currentToken
		parser.eat(op)
		right := parser.parseMultiplicativeExpression()
		left = ast.NewBinaryOperator(left, right, op.Type)
	}
	return left
}

func (parser *Parser) parseExpression() ast.Node {
	return parser.parseAdditiveExpression()
}

func (parser *Parser) Parse() ast.Node {
	return parser.parseExpression()
}

/*
  grammer:
  expression -> additive_expression
  additive_expression -> multiplicative_expression (('+' | '-') multiplicative_expression)*
  multiplicative_expression -> unary_expression (('*' | '/' | '%' | '//', '^') unary_expression)*
  unary_expression -> factor (('!', | 'sin' | 'cos' | 'tan' | 'cot' | 'sec' | 'cosec' | log)factor)*
  prefix_unary_expression -> ('sin' | 'cos' | 'tan' | 'cot' | 'sec' | 'cosec' | log)(factor | prefix_unary_expression | suffix_unary_expression)*
  suffix_unary_expression -> factor!
  factor -> number | 'pi' | '(' expression ')' | '[' expression ']' | '{' expression '}'
*/
