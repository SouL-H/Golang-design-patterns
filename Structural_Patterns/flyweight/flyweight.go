package main

import "fmt"

type dress interface {
	getColor() string
}

type dressFactory struct {
	data map[string]dress
}

func (f dressFactory) getDress(dressType string) dress {
	return f.data[dressType]
}

type terroristDress struct {
	color string
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

func (t terroristDress) getColor() string {
	return t.color
}

type counterTerroristDress struct {
	color string
}

func (c counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

const (
	TerroristPlayerType        = "t"
	CounterTerroristPlayerType = "ct"
)

var (
	dressProvider = &dressFactory{
		data: map[string]dress{
			TerroristDressType:        newTerroristDress(),
			CounterTerroristDressType: newCounterTerroristDress(),
		},
	}
)

type player struct {
	dress      dress
	playerType string
	lat, long  int
}

func newPlayer(playerType string) *player {
	var d dress
	switch playerType {
	case TerroristPlayerType:
		d = dressProvider.getDress(TerroristDressType)
	case CounterTerroristPlayerType:
		d = dressProvider.getDress(CounterTerroristDressType)
	}
	return &player{
		dress:      d,
		playerType: playerType,
	}
}

type game struct {
	terrorist        []*player
	counterTerrorist []*player
}

func newGame() *game {
	return &game{
		terrorist:        make([]*player, 0),
		counterTerrorist: make([]*player, 0),
	}
}

func (g *game) addTerrorist() {
	g.terrorist = append(g.terrorist, newPlayer(TerroristPlayerType))
}

func (g *game) addCounterTerrorist() {
	g.counterTerrorist = append(g.counterTerrorist, newPlayer(CounterTerroristPlayerType))
}

func main() {
	g := newGame()

	g.addTerrorist()
	g.addTerrorist()
	g.addTerrorist()
	g.addTerrorist()
	g.addTerrorist()

	g.addCounterTerrorist()
	g.addCounterTerrorist()
	g.addCounterTerrorist()
	g.addCounterTerrorist()
	g.addCounterTerrorist()

	for _, t := range g.terrorist {

		fmt.Printf("Pointer: %p - Value: %s\n", t.dress, t.dress)
	}

	for _, t := range g.counterTerrorist {
		fmt.Printf("Pointer: %p - Value: %s\n", t.dress, t.dress)
	}
}
