package rogue

type Water struct { CellData }

func NewWater() *Water {
	return &Water{*NewCellData(WATER)}
}