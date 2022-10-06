package example

import "fmt"

// 🤯🤯🤯🤯🤯

type Computer interface {
	Print()
	SetPrinter(Printer)
}
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("mac printing")
	m.printer.PrintFile("my mac is printing 😄")
}
func (m *Mac) SetPrinter(ep Epson) {
	m.printer = ep
}

type Printer interface {
	PrintFile(string)
}
type Epson struct {
}

func (epsonPrinter Epson) PrintFile(msg string) {
	fmt.Println("message from epson printer: ", msg)
}

type Linux struct {
	printer Printer
}

func (linux *Linux) Print() {
	fmt.Println("linux printing")
	linux.printer.PrintFile("slime shady")
}
func (linux *Linux) SetPrinter(p Printer) {
	linux.printer = p
}

type Kiosera struct{}

func (kiosera *Kiosera) PrintFile(msg string) {
	fmt.Println("message from kiosera printer: ", msg)
}
