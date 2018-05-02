package rogue

import(
	"strconv"
	"fmt"
)

type Entity interface {
	Cell

	GetName() string
	SetName(name string)

	GetId() int
	SetId(id int)

	GetDirection() Direction
	SetDirection(direction Direction)

	GetPosition() Position
	SetPosition(position Position)

	GetCellUnder() Cell
	SetCellUnder(cellUnder Cell)

	GetLivingInFront() Living

}

type EntityData struct {
	CellData
	id int
	name string
	direction Direction
	position Position

	cellUnder Cell
}

var counter int = 0
func aumentCounter() {
	counter++
}


func NewEntityData(_type CellType, position Position) *EntityData {
	defer aumentCounter()
	return &EntityData{
		*NewCellData(_type),
		counter,
		_type.String() + strconv.Itoa(counter),
		UP,
		position,
		NewDefaultCell(),
	}

}

func (e *EntityData) String() string {
	return fmt.Sprintf(
		"cellData: %v\nid:%d\nname: %s\ndirection: %v\nposition: %v\ncellUnder: %v\n",
		e.CellData, e.id, e.name, e.direction, e.position, e.cellUnder)
}

func (e *EntityData) GetName() string {
	return e.name
}

func (e *EntityData) SetName(name string) {
	e.name = name
}

func (e *EntityData) GetId() int {
	return e.id
}

func (e *EntityData) SetId(id int) {
	e.id = id
}

func (e *EntityData) GetDirection() Direction {
	return e.direction
}

func (e *EntityData) SetDirection(direction Direction) {
	e.direction = direction
}

func (e *EntityData) GetPosition() Position {
	return e.position
}

func (e *EntityData) SetPosition(position Position) {
	e.position = position
}

func (e *EntityData) GetCellUnder() Cell {
	return e.cellUnder
}

func (e *EntityData) SetCellUnder(cellUnder Cell) {
	e.cellUnder = cellUnder
}


func (e *EntityData) GetLivingInFront() Living {
	inFrontPosition := NewPosition(-1, -1)
	switch(e.direction) {
	case UP:
		inFrontPosition = NewPosition(e.position.X, e.position.Y - 1)

		if inFrontPosition.Y < 0 {
			return nil
		}
	case DOWN:
		inFrontPosition = NewPosition(e.position.X, e.position.Y + 1)

		if inFrontPosition.Y >= GBoard.GetSize() {
			return nil
		}
	case LEFT:
		inFrontPosition = NewPosition(e.position.X - 1, e.position.Y)

		if inFrontPosition.X < 0 {
			return nil
		}
	case RIGHT:
		inFrontPosition = NewPosition(e.position.X + 1, e.position.Y)

		if inFrontPosition.X >= GBoard.GetSize() {
			return nil
		}
	}

	fmt.Printf("Viewing: %v\n", inFrontPosition)

	for _, living := range GLivings {
		if living != nil {
			if living.GetPosition() == inFrontPosition {
				return living
			}
		}
	}

	return nil
}