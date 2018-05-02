package rogue

type DefaultCell struct {
	CellData
}

func NewDefaultCell() *DefaultCell {
	return &DefaultCell{*NewCellData(DEFAULTCELL)}
}

func (d *DefaultCell) OnEnter(living Living) {
	d.CellData.OnEnter(living)
	panic("entered to default cell")
}
