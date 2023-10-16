package main

import "fmt"

func main() {
	pizza := &VeggieMania{}
	// 进行装饰
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	// 再进行一层装饰
	pizzaWithChessAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Println(pizzaWithChessAndTomato.getPrice())
}
