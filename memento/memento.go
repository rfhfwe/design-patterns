package main

// Memento 备忘录
type Memento struct {
	// 一些参数，包含备忘录参数...
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}
