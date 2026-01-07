package ast

type Visitable[T any] interface {
	AcceptVisitor(visitor Visitor[T]) T
}

type Visitor[T any] interface {
	VisitBinaryOperator(node *BinaryOperator) T
	VisitNumber(node *Number) T
	VisitUnaryOperator(node *UnaryOperator) T
}

type Node interface {
	Visitable[ComputeSize]
}
