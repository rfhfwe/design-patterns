package main

type TerroristDress struct {
	color string
}

type CounterTerroristDress struct {
	color string
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func (t *CounterTerroristDress) getColor() string {
	return t.color
}
