package main

import "fmt"

type game interface {
	Start()
	Fight()
	End()
}

type gamer struct {
	g game
}

func (gmr *gamer) Play() {
	gmr.g.Start()
	gmr.g.Fight()
	gmr.g.End()
}

type chessGame struct {
	players []string
}

func (c *chessGame) Start() {
	fmt.Println("start game")
}

func (c *chessGame) Fight() {
	fmt.Println("fight with each other")
}

func (c *chessGame) End() {
	fmt.Println("end game")
}

func main() {
	cg := chessGame{players: []string{"Chris", "Hanna"}}
	g := gamer{g: &cg}
	g.Play()
}
