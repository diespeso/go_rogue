package rogue

import(
	"fmt"
)

//Base interface for all Cells and derivates
type Cell interface {
	GetType() CellType
	SetType(_type CellType)
	GetSprite() string
	SetSprite(sprite string)
	OnEnter(living Living)
	CanEnter(living Living) bool
}

//Implements Cell
type CellData struct {
	_type CellType
	sprite string
}

func NewCellData(_type CellType) *CellData {
	return &CellData{_type, _type.GetSprite()}
}

func (c *CellData) GetType() CellType {
	return c._type
}

func (c *CellData) SetType(_type CellType) {
	*c = *NewCellData(_type)
}

func (c *CellData) GetSprite() string {
	return c.sprite
}

func (c *CellData) SetSprite(sprite string) {
	c.sprite = sprite
}

func (c *CellData) OnEnter(living Living) {
	fmt.Printf("%s entered to %s\n", living.GetName(), c._type)
}

func (c *CellData) CanEnter(living Living) bool {
	return false
}

func (c CellData) String() string {
	return fmt.Sprintf("[%s, %v]", c._type, c.GetSprite())
}