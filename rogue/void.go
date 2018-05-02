package rogue

type Void struct{ CellData }

func NewVoid() *Void {
	return &Void{*NewCellData(VOID)}
}

func (v *Void) OnEnter(living Living) {
	v.CellData.OnEnter(living)
	panic("Shouldn't be able to enter to a Void Cell")
}

func (v *Void) CanEnter(living Living) bool {
	return false;
}