package main

import (
	"dsa-golang/Chapter1/StructuralPatterns/proxy/example"
	"fmt"
)

type IRealObject interface {
	performAction()
}

type RealObject struct {
}

func (realObj *RealObject) performAction() {
	fmt.Println("Real object performing action ")
}

type Proxy struct {
	realObject *RealObject
}

func (virtualProxy *Proxy) performAction() {
	if virtualProxy.realObject == nil {
		fmt.Println("proxy creating real object")
		virtualProxy.realObject = &RealObject{}
	}
	virtualProxy.realObject.performAction()
}

func main() {
	var object Proxy = Proxy{}
	object.performAction()

	fmt.Println()
	//car proxy
	carProxy := example.NewCarProxy(19)
	carProxy.Drive()
}
