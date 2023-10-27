package main

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

var _ IGun = (*Gun)(nil)

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun{
			name:  "AK 47 gun",
			power: 4,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}
