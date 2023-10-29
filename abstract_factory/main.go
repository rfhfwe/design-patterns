package main

import "fmt"

var (
	adidasFactory ISportsFactory
	nikeFactory   ISportsFactory
	err           error
	nikeShoe      IShoe
	nikeShirt     IShirt

	adidasShoe  IShoe
	adidasShirt IShirt
)

func Init() {
	adidasFactory, err = GetSportsFactory("adidas")
	nikeFactory, err = GetSportsFactory("nike")
	nikeShoe = nikeFactory.makeShoe()
	nikeShirt = nikeFactory.makeShirt()
	adidasShoe = adidasFactory.makeShoe()
	adidasShirt = adidasFactory.makeShirt()
}

func main() {
	Init()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
