package main

// Caretaker 负责人
type Caretaker struct {
	mementoArray []*Memento
}

// addMemento 添加备忘录
func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

// getMemento 根据index获取备忘录
func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}
