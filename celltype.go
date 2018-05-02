package rogue

//Represents the type of a cell
type CellType int

//Holds every type and its string sprite
var CellTypes map[CellType]string

type Stat struct{
	health int
	attack int
	defense int}

var LivingStats map[CellType]Stat

func init() {
	CellTypes = map[CellType]string {
		VOID: " ",
		WATER: "+",
		GROUND: "-",
		ERR: "¿",
		DEFAULTCELL: "0",
		PLAYER: "@",
		SPIKES: "^",
		DRAGON: "D",
	}

	LivingStats = map[CellType]Stat {
		PLAYER: Stat{30, 15, 10},
		DRAGON: Stat{40, 20, 30},
	}
}

const(
	VOID CellType = iota
	WATER
	GROUND
	ERR

	DEFAULTCELL
	SPIKES
	DRAGON
	PLAYER
)

//Returns the name of c when called
func(c CellType) String() string {
	switch(c) {
	case VOID:
		return "VOID"
	case WATER:
		return "WATER"
	case GROUND:
		return "GROUND"
	case ERR:
		return "ERR"
	case DEFAULTCELL:
		return "DEFAULTCELL"
	case PLAYER:
		return "PLAYER"
	case SPIKES:
		return "SPIKES"
	case DRAGON:
		return "DRAGON"
	}

	return "???"
}

//Returns the sprite representation of c
func (c CellType) GetSprite() string {
	return CellTypes[c]
}