package ast

import "github.com/AmKreta/golang-expression-parser/token"

type UnaryOperator struct {
	operator token.TokenType
	operand  Node
}

func (u *UnaryOperator) AcceptVisitor(visitor Visitor[ComputeSize]) ComputeSize {
	return visitor.VisitUnaryOperator(u)
}

func NewUnaryOperator(operator token.TokenType, operand Node) *UnaryOperator {
	return &UnaryOperator{operator: operator, operand: operand}
}

func (u *UnaryOperator) GetOperator() token.TokenType {
	return u.operator
}

func (u *UnaryOperator) GetOperand() Node {
	return u.operand
}
