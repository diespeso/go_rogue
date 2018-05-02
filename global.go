package rogue

import("fmt")

var GBoard *Board
var GState GameState = PLAYING
var GPlayer *Player

var GLivings []Living = make([]Living, 20)
var GLivingsCounter int = 0

func GUpdate() {
	GBoard.SetCell(GPlayer.GetPosition().X, GPlayer.GetPosition().Y, GPlayer)
	if GPlayer.GetIsDead() {
		//GCloseGame()
	}
	for i := 0; i < GLivingsCounter; i++ {
		if GLivings[i] != nil {
			GBoard.SetCell(GLivings[i].GetPosition().X,
				GLivings[i].GetPosition().Y, GLivings[i])
		}
	}
}

func GShowLivings() {
	fmt.Println("SHOWING LIVINGS:")
	for i := 0; i < GLivingsCounter; i++ {
		fmt.Println("--------------------")
		GLivings[i].Show()
		fmt.Println("--------------------")
	}
	
}

func GAddLiving(living Living) {
	//When adding a living, it should take the cell that
	//was on the board before it came and use it as
	//its cellUnder
	living.SetCellUnder(GBoard.GetCellFromPos(living.GetPosition()))
	if(GLivingsCounter >= len(GLivings)) {
		panic("Too much livings in GLivings")
	}
	GLivings[GLivingsCounter] = living
	GLivingsCounter++
}

func GKillLiving(living Living) {
	for i := 0; i < GLivingsCounter; i++ {
		if(GLivings[i] != nil) {
			if(GLivings[i].GetId() == living.GetId()) {
				GBoard.SetCellFromPos(GLivings[i].GetPosition(),
					GLivings[i].GetCellUnder())
				GLivings[i] = nil
			}
		}
		
	}
}

func CloseGame() {
	GState = ENDED
	if GPlayer.GetIsDead() {
		fmt.Println("The player died")
	} else {
		fmt.Println("You exxited the game, hun")
	}
}