package main

var (
	shirtItem Subject

	observerFirst  *Customer
	observerSecond *Customer
)

func Init() {
	shirtItem = newItem("Nick Shirt")

	observerFirst = &Customer{id: "abc@gmail.com"}
	observerSecond = &Customer{id: "xyz@gmail.com"}
}

func main() {
	Init() // 初始化
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.notifyAll()
}
