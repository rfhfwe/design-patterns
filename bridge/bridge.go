package main

import "fmt"

// Printer 实施类接口
type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

type Windows struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac.")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows.")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Epson 具体实施类
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
