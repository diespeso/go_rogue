package rogue

type Ground struct{ CellData }

func NewGround() *Ground {
	return &Ground{*NewCellData(GROUND)}
}

func (g *Ground) CanEnter(living Living) bool {
	return true
}