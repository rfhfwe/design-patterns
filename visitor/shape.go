package main

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

type Circle struct {
	radius int
}

type Rectangle struct {
	l int
	b int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

func (s *Circle) accept(v Visitor) {
	v.visitForCircle(s)
}

func (s *Circle) getType() string {
	return "Circle"
}

func (s *Rectangle) accept(v Visitor) {
	v.visitForRectangle(s)
}

func (s *Rectangle) getType() string {
	return "Rectangle"
}
