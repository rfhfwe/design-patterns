package main

// Originator 原发器
type Originator struct {
	// 一些参数，包含非备忘录参数+备忘录参数...
	state string
}

// createMemento 创建一个备忘录
func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

// restoreMemento 存储备忘录的状态
func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}
