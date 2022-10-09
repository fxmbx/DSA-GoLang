package main

import (
	"errors"
	"fmt"
)

const (
	TerroristDressType        = "dressT"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleton = &dressFactory{dressMap: make(map[string]dress)}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (dFactory *dressFactory) getDressByType(dressType string) (dress, error) {
	if dFactory.dressMap[dressType] != nil {
		return dFactory.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		dFactory.dressMap[dressType] = newTerroristDress()
		return dFactory.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		dFactory.dressMap[dressType] = newCounterTerroristDress()
		return dFactory.dressMap[dressType], nil
	}
	return nil, errors.New("wrong dress type")
}

type dress interface {
	getColor() string
}
type terroristDress struct {
	color string
}
type counterTerroristDress struct {
	color string
}

func (tDress *terroristDress) getColor() string {
	return tDress.color
}
func (ctDress *counterTerroristDress) getColor() string {
	return ctDress.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{
		color: "black",
	}

}
func newCounterTerroristDress() dress {
	return &counterTerroristDress{
		color: "navy blue",
	}
}
func getDressFactorySingletonInstance() *dressFactory {
	return dressFactorySingleton

}

type player struct {
	dress      dress
	playerType string
	long       int
	lat        int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingletonInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}
func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}
func (g *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}

func main() {
	game := newGame()
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	dressFactoryInstance := getDressFactorySingletonInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}

}
