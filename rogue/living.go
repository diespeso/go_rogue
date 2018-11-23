package rogue

import("fmt")

type Living interface {
	Entity

	Show()

	GetHealth() int
	SetHealth(health int)

	GetAttack() int 
	SetAttack(attack int)

	GetDefense() int
	SetDefense(defense int)

	GetIsDead() bool
	SetIsDead(dead bool)

	Damage(attack int) bool
	Update()
	Kill()

	GetCanMove() bool
	SetCanMove(canMove bool)
	Move(direction Direction)
}

type LivingData struct {
	EntityData

	health int 
	attack int 
	defense int 
	isDead bool

	canMove bool
}

func NewLivingData(_type CellType, position Position) *LivingData {
	return &LivingData{*NewEntityData(_type, position),
		30,
		10,
		5,
		false,
		true,
	}
}

func (l *LivingData) Show() {
	fmt.Printf("%v", l)
	fmt.Printf("health: %v\nattack: %v\ndefense: %v\n",
		l.health, l.attack, l.defense)
}

func (l *LivingData) GetHealth() int {
	return l.health
}

func (l *LivingData) SetHealth(health int) {
	l.health = health
}

func (l *LivingData) GetAttack() int {
	return l.attack
}

func (l *LivingData) SetAttack(attack int) {
	l.attack = attack
}

func (l *LivingData) GetDefense() int {
	return l.defense
}

func (l *LivingData) SetDefense(defense int) {
	l.defense = defense
}

func (l *LivingData) GetIsDead() bool {
	return l.isDead
}

func (l *LivingData) SetIsDead(isDead bool) {
	l.isDead = isDead
}

func (l *LivingData) Damage(attack int) bool {
	attack -= l.defense
	if attack >= 0 {
		l.health -= attack
	} else {
		l.health -= 0
	}

	l.Update()
	//debug
	fmt.Println(l.health)
	return l.health != 0
}

func (l *LivingData) Update() {
	if l.health == 0 {
		l.isDead = true
		l.Kill()
		GUpdate()
	}
}

func (l *LivingData) Kill() {
	GKillLiving(l)
}

func (l *LivingData) GetCanMove() bool {
	return true
}

func (l *LivingData) SetCanMove(canMove bool) {
	l.canMove = canMove
}

//First checks de future position called "nextPosition".
func (l *LivingData) Move(direction Direction) {
	nextPosition := l.GetPosition()
	l.SetDirection(direction)
	switch(direction) {
	case UP:
		if l.position.Y > 0 {
			nextPosition.Y--
		}
	case DOWN:
		if l.position.Y != GBoard.GetSize() - 1 {
			nextPosition.Y++
		}
	case LEFT:
		if l.position.X != 0 {
			nextPosition.X--
		}
	case RIGHT:
		if l.position.X != GBoard.GetSize() - 1 {
			nextPosition.X++
		}
	}

	//When the Living moves, the cell that was under that Entity is colocated on the Board, so the 
	//"renderer" (which doesn't really exists) can render it instead of the Entity that was on top
	//on said Entity.
	if GBoard.GetCellFromPos(nextPosition).CanEnter(l) && l.GetCanMove() {
		GBoard.SetCellFromPos(l.GetPosition(), l.GetCellUnder())
		l.SetCellUnder(GBoard.GetCellFromPos(nextPosition))
		l.SetPosition(nextPosition)
		l.GetCellUnder().OnEnter(l) //triggers the cell under
	}

	GUpdate()
}
