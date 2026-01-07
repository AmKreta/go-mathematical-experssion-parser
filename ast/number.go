package ast

type Number struct {
	value ComputeSize
}

func (num *Number) AcceptVisitor(visitor Visitor[ComputeSize]) ComputeSize {
	return visitor.VisitNumber(num)
}

func NewNumber(value ComputeSize) *Number {
	return &Number{value: ComputeSize(value)}
}
