package rogue_test

import(
	"testing"
	"fmt"
	."rogue"
)

func AssertEquals(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		panic(fmt.Sprintf("%s != %s", a, b))
	}
}

func TestBoard(t *testing.T) {

}

func TestGlobal(t *testing.T) {
	GBoard = NewBoard(10)
	GBoard.Fill(NewGround())
	GBoard.Show()

	GPlayer = NewPlayer("new", Position{1, 6})
	GAddLiving(GPlayer)
	GUpdate()
	GBoard.Show()

	GPlayer.Move(UP)
	GBoard.Show()

	GPlayer.Move(UP)
	GBoard.Show()

	fmt.Println("ENDED GLOBAL TEST")

}

func TestEnemy(t *testing.T) {
	newEnemy := NewEnemyData(DRAGON, Position{1, 5})
	fmt.Println(newEnemy)
	newEnemy.Show()
	GAddLiving(newEnemy)
	GUpdate()
	GBoard.Show()
	fmt.Println("ENDED ENEMY TEST")
}

func TestDragon(t *testing.T) {
	newDragon := NewDragon(Position{5, 5})
	GAddLiving(newDragon)
	GUpdate()
	GBoard.Show()
}

func TestPlayer(t *testing.T) {
	newPlayer := NewPlayer("volta", Position{3, 2})
	GAddLiving(newPlayer)
	GUpdate()
	GBoard.Show()
	GShowLivings()
}