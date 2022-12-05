package yahtzee

import "fmt"

type Game struct {
	Players *Players
	Turn    *Turn
	MaxTurn int
}

func NewGame(ps *Players) *Game {
	maxTurn := ps.AllPlayer * NumberOfYakuType
	return &Game{
		Players: ps,
		MaxTurn: maxTurn,
	}
}

func (g *Game) NextTurn() error {
	if !g.Turn.IsCompleted() {
		return fmt.Errorf("まだターンが終了していません")
	}

	player := g.Players.Next()
	nextTurnNumber := g.Turn.Current + 1
	g.Turn = NewTurn(player, nextTurnNumber)

	return nil
}

func (g *Game) RemainTurn() int {
	return g.MaxTurn - g.Turn.Current
}
