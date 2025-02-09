package main

import "fmt"

// type Player struct {
// 	posX float64
// 	posY float64
// }

// func (p *Player) Move(x, y float64) {
// 	p.posX += x
// 	p.posY += y
// }

// func (p *Player) Teleport(x, y float64) {
// 	p.posX = x
// 	p.posY = y
// }

type Position struct {
	X float64
	Y float64
}

func (p *Position) Move(x, y float64) {
	p.X += x
	p.Y += y
}

func (p *Position) Teleport(x, y float64) {
	p.X = x
	p.Y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{&Position{}}
}

func main() {
	p := NewPlayer()
	p.Move(1, 2)
	p.Teleport(5, 5)

	fmt.Println(p.X, p.Y)
}