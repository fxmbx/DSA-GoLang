package main

import "fmt"

type StudentTuple struct {
	regNo, name any
}

func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square: ", square, "cube: ", cube)

	student1 := &StudentTuple{171766, "Olaore oluwafunmibi olumuyiwa"}
	fmt.Printf("stuent name: %s, student reg number: %d", student1.name, student1.regNo)
}

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}
