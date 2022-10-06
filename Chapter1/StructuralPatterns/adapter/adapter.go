package main

import "fmt"

type IProcess interface {
	process()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee Convert method")
}

type Adapter struct {
	adaptee Adaptee
}

func (adapt Adapter) process() {
	fmt.Println("Adapter Process method")
	adapt.adaptee.convert()
}
func main() {
	var processor IProcess = Adapter{}
	processor.process()

}
