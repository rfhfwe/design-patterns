package main

type IPizza interface {
	getPrice() int
}

type VeggieMania struct {
}

func (v *VeggieMania) getPrice() int {
	return 10
}

// TomatoTopping 具体装饰类
type TomatoTopping struct {
	pizza IPizza
}

// CheeseTopping 具体装饰类
type CheeseTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 100
}

func (t *CheeseTopping) getPrice() int {
	return t.pizza.getPrice() + 99
}
