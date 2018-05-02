package rogue

import(
	"fmt"
)

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}


func (p *Position) SetFromCoords(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Position) Get() *Position {
	return p
}

func (p *Position) String() string {
	return fmt.Sprintf("[%v, %v]", p.X , p.Y)
}

func (p *Position) Same(_p Position) bool {
	return p.X == _p.X && p.Y == _p.Y
}