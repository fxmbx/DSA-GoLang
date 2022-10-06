package main

import (
	"dsa-golang/Chapter1/StructuralPatterns/bridge/example"
	"fmt"
)

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct {
}

func (dShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("drawing shape")
}

type IContour interface {
	drawCotour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (contour DrawContour) drawCotour(x [5]float32, y [5]float32) {
	fmt.Println("drawing contour")
	contour.shape.drawShape(contour.x, contour.y)
}
func (contour DrawContour) resizeByFactor(factor int) {
	fmt.Println("resizing by factor")
	contour.factor = factor
}
func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawCotour(x, y)
	contour.resizeByFactor(2)

	epSonPrinter := &example.Epson{}

	kioseraPrinter := &example.Kiosera{}

	macComputer := &example.Mac{}
	macComputer.SetPrinter(*epSonPrinter)
	macComputer.Print()

	linuxSystem := &example.Linux{}
	linuxSystem.SetPrinter(kioseraPrinter)
	linuxSystem.Print()

}
