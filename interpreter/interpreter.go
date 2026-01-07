package interpreter

import (
	"fmt"
	"math"

	"github.com/AmKreta/golang-expression-parser/ast"
	"github.com/AmKreta/golang-expression-parser/parser"
	"github.com/AmKreta/golang-expression-parser/token"
	"github.com/AmKreta/golang-expression-parser/util"
)

type Interpreter struct {
	expression ast.Node
}

func NewInterpreter(parser *parser.Parser) *Interpreter {
	expression := parser.Parse()
	return &Interpreter{expression: expression}
}

func (i *Interpreter) VisitBinaryOperator(node *ast.BinaryOperator) ast.ComputeSize {
	left := node.GetLeft().AcceptVisitor(i)
	right := node.GetRight().AcceptVisitor(i)
	operator := node.GetOperator()
	switch operator {
	case token.PLUS:
		return left + right
	case token.MINUS:
		return left - right
	case token.MULTIPLY:
		return left * right
	case token.DIVIDE:
		return left / right
	case token.MODULO:
		int_a := int64(left)
		int_b := int64(right)
		return ast.ComputeSize(int_a % int_b)
	case token.INTEGER_DIVIDE:
		return left / right
	case token.POWER:
		float_a := float64(left)
		float_b := float64(right)
		return ast.ComputeSize(math.Pow(float_a, float_b))
	default:
		panic(fmt.Sprintf("unexpected operator %s", token.TokenNames[operator]))
	}
}

func (i *Interpreter) VisitUnaryOperator(node *ast.UnaryOperator) ast.ComputeSize {
	operand := node.GetOperand().AcceptVisitor(i)
	operator := node.GetOperator()
	switch operator {
	case token.FACTORIAL:
		return ast.ComputeSize(util.Factorial(int64(operand)))
	case token.SINE:
		return ast.ComputeSize(math.Sin(float64(operand)))
	case token.COSINE:
		return ast.ComputeSize(math.Cos(float64(operand)))
	case token.TANGENT:
		return ast.ComputeSize(math.Tan(float64(operand)))
	case token.COTANGENT:
		return ast.ComputeSize(math.Cos(float64(operand)) / math.Sin(float64(operand)))
	case token.SECANT:
		return ast.ComputeSize(1 / math.Cos(float64(operand)))
	case token.COSECANT:
		return ast.ComputeSize(1 / math.Sin(float64(operand)))
	case token.LOGARITHM:
		return ast.ComputeSize(math.Log(float64(operand)))
	default:
		panic(fmt.Sprintf("unexpected operator %s", token.TokenNames[operator]))
	}
}

func (i *Interpreter) VisitNumber(node *ast.Number) ast.ComputeSize {
	return node.GetValue()
}

func (i *Interpreter) Interpret() ast.ComputeSize {
	return i.expression.AcceptVisitor(i)

}
