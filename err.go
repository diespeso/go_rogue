package rogue

type Err struct { CellData }

func NewErr() *Err {
	return &Err{*NewCellData(ERR)}
}

func (e *Err) OnEnter(living Living) {
	e.CellData.OnEnter(living)
	panic("Entered to a Error Cell")
}

func (e *Err) CanEnter() bool {
	return true
}