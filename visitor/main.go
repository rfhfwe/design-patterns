package main

var (
	square    *Square
	circle    *Circle
	rectangle *Rectangle

	areaCalculator    *AreaCalculator
	middleCoordinates *MiddleCoordinates
)

func Init() {
	square = &Square{side: 2}
	circle = &Circle{radius: 2}
	rectangle = &Rectangle{l: 2, b: 3}

	areaCalculator = &AreaCalculator{}
	middleCoordinates = &MiddleCoordinates{}
}

func main() {
	Init() // 初始化
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
