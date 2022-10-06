package main

import "fmt"

type transport interface {
	navigateToDestination()
}
type Client struct {
}

func (c *Client) startNavigatoin(t transport) {
	fmt.Println("client start navigation method")
	t.navigateToDestination()
}

type Boat struct {
}

//boat is compatible with client
func (c *Boat) navigateToDestination() {
	fmt.Println("boats cut through friction with streamlined body on water")
}

type Car struct {
}

//car is not compatible with client
func (c *Car) driveToDestination() {
	fmt.Println("cars move with tyres on land")
}

//car gets an adapter to make it compatible with client
type CarAdapter struct {
	carAdapter *Car
}

func (cAdapter *CarAdapter) navigateToDestination() {
	fmt.Println("carAdapter modify car to allow client call ")
	cAdapter.carAdapter.driveToDestination()
}
