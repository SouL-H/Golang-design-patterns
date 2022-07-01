package main

import "fmt"

type Notebook interface {
	SetupPrinter(p Printer)
	PrintDoc(doc string)
}

type Printer interface {
	Print(doc string)
}

type Mac struct {
	Printer Printer
}

func (m *Mac) SetupPrinter(p Printer) {
	m.Printer = p
}

func (m *Mac) PrintDoc(doc string) {
	fmt.Printf("Print some doc by Mac: %s\n", doc)
	m.Printer.Print(doc)
}

type Windows struct {
	Printer Printer
}

func (w *Windows) SetupPrinter(p Printer) {
	w.Printer = p
}

func (w *Windows) PrintDoc(doc string) {
	fmt.Printf("Print some doc by Windows: %s\n", doc)
	w.Printer.Print(doc)
}

type Epson struct{}

func (e *Epson) Print(doc string) {
	fmt.Printf("Epson is ready to print some doc: %s\n", doc)
}

type HP struct{}

func (H *HP) Print(doc string) {
	fmt.Printf("HP is ready to print some doc: %s\n", doc)
}

func main() {
	e := &Epson{}
	h := &HP{}

	m := &Mac{}
	m.SetupPrinter(e)
	m.PrintDoc("1")
	m.SetupPrinter(h)
	m.PrintDoc("2")

	w := &Windows{}
	w.SetupPrinter(e)
	w.PrintDoc("3")
	w.SetupPrinter(h)
	w.PrintDoc("4")
}
