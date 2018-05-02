package rogue

type Player struct {
	LivingData
}

func NewPlayer(name string, position Position) *Player {
	stats := LivingStats[PLAYER]
	newPlayer := &Player{
		LivingData{
			*NewEntityData(PLAYER, position),
			stats.health,
			stats.attack,
			stats.defense,
			false,
			true,
		},
	}
	newPlayer.SetName(name)

	return newPlayer
}