package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(10)
	intList.PushBack(550)
	intList.PushBack(40)
	fmt.Println(*intList.Back())
	fmt.Println(*intList.Front())

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))

	}
}
