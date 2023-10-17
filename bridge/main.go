package main

var (
	hpPrinter    *Hp
	epsonPrinter *Epson

	macComputer *Mac
	winComputer *Windows
)

func Init() {
	hpPrinter = &Hp{}
	epsonPrinter = &Epson{}
	macComputer = &Mac{}
	winComputer = &Windows{}
}

func main() {
	Init() // 初始化
	service(macComputer, hpPrinter)
	service(macComputer, epsonPrinter)

	service(winComputer, hpPrinter)
	service(macComputer, epsonPrinter)
}

func service(computer Computer, printer Printer) {
	computer.SetPrinter(printer)
	computer.Print()
}
