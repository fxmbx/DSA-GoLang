package main

import (
	"fmt"
)

//an algorithm is quadratic  when processing time is proportional to number of input O(n^2)

func main() {

	var k, l int

	for k = 1; k <= 10; k++ {
		fmt.Println(" Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			var x int = l * k
			fmt.Println(x)
		}

	}
}
