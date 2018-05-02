package rogue

type Direction int

const(
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (d Direction) String() string {
	switch(d) {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	}
	return "???"
}