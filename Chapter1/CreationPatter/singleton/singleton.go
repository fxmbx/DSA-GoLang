package main

import (
	"fmt"
	"sync"
)

/*
"Singleton is a creational pattern that allows ensuring only one instance from a “struct” type along all the execution time."
*/

type singletonCon struct {
	url  string
	port string
}

var lock = &sync.Mutex{}

var singletonConInstance *singletonCon

func getInstance() *singletonCon {
	if singletonConInstance == nil {

		fmt.Println("creating new instance")
		lock.Lock()
		defer lock.Unlock()
		singletonConInstance = &singletonCon{
			url:  "url",
			port: "port",
		}
	} else {
		fmt.Println("instance already created")
	}
	return singletonConInstance
}

func main() {
	for i := 0; i < 5; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
