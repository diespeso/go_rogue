package rogue

import("fmt")

type Board struct {
	size int
	matrix [][]Cell
}

func NewBoard(size int) *Board {
	newBoard := &Board{size, make([][]Cell, size, size)}
	newBoard.initialize()
	return newBoard
}

func (b *Board) Show() {
	for i := 0; i < b.GetSize(); i++ {
		for j := 0; j < b.GetSize(); j++ {
			fmt.Printf("%v ", b.GetMatrix()[j][i].GetSprite())
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) GetMatrix() [][]Cell {
	return b.matrix
}

func (b *Board) GetCell(x, y int) Cell {
	return b.matrix[x][y]
}

func (b *Board) GetCellFromPos(position Position) Cell {
	return b.matrix[position.X][position.Y]
}

func (b *Board) SetCell(x, y int, cell Cell) {
	b.matrix[x][y] = cell
}

func (b *Board) SetCellFromPos(position Position, cell Cell) {
	b.matrix[position.X][position.Y] = cell
}

func (b *Board) initialize() *Board {
	for i := 0; i < b.GetSize(); i++ {
		b.GetMatrix()[i] = make([]Cell, b.GetSize())
	}
	return b
}

func (b *Board) Fill(cell Cell) {
	for i := 0; i < b.GetSize(); i++ {
		for j := 0; j < b.GetSize(); j++ {
			b.GetMatrix()[i][j] = cell
		}
	}
}