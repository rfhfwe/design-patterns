package main

import "fmt"

var (
	caretaker  *Caretaker  // 负责人
	originator *Originator // 原发器
)

func Init() {
	caretaker = &Caretaker{
		mementoArray: make([]*Memento, 0),
	}
	originator = &Originator{
		state: "A",
	}
}

func main() {

	Init() // 初始化

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento()) // 向负责人保存备忘录

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento()) // 向负责人保存备忘录

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento()) // 向负责人保存备忘录

	originator.restoreMemento(caretaker.getMemento(1)) // 恢复到先前index=1时
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0)) // 恢复到先前index=0时
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
