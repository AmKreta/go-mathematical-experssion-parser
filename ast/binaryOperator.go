package ast

import "github.com/AmKreta/golang-expression-parser/token"

type BinaryOperator struct {
	left     Node
	right    Node
	operator token.TokenType
}

func (b *BinaryOperator) AcceptVisitor(visitor Visitor[ComputeSize]) ComputeSize {
	return visitor.VisitBinaryOperator(b)
}

func NewBinaryOperator(left Node, right Node, operator token.TokenType) *BinaryOperator {
	return &BinaryOperator{left: left, right: right, operator: operator}
}

func (b *BinaryOperator) GetLeft() Node {
	return b.left
}

func (b *BinaryOperator) GetRight() Node {
	return b.right
}

func (b *BinaryOperator) GetOperator() token.TokenType {
	return b.operator
}
