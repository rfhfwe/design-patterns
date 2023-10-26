package main

import "log"

var (
	vendingMachine *VendingMachine
)

func Init() {
	vendingMachine = newVendingMachine(1, 10)
}

func main() {
	Init() // 初始化
	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
}
