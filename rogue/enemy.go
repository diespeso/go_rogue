package rogue

import (
	"math/rand"
	
)

type Enemy interface {
	Living
	RandomMove()
}

type EnemyData struct {
	LivingData
}

func NewEnemyData(_type CellType, position Position) *EnemyData {
	stats := LivingStats[_type]
	return &EnemyData{
		LivingData{
			*NewEntityData(_type, position),
			stats.health,
			stats.attack,
			stats.defense,
			false,
			true,
		},
	}
}

func (e *EnemyData) RandomMove() {
	switch(rand.Intn(4)) {
	case 0:
		e.Move(UP)
	case 1:
		e.Move(DOWN)
	case 2:
		e.Move(LEFT)
	case 3:
		e.Move(RIGHT)
	}
}