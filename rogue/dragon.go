package rogue

type Dragon struct {
	EnemyData
}

func NewDragon(position Position) *Dragon {
	return &Dragon{
		*NewEnemyData(DRAGON, position),
	}
}