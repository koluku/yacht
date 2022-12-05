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

func (g *Game) NextTurn() (*Turn, error) {
	// ゲームが始まっていない場合はTurnを新規作成する
	if g.Turn == nil {
		player := g.Players.First()
		turn := NewTurn(player, 1)
		g.Turn = turn
		return turn, nil
	}

	if !g.Turn.IsEnded() {
		return nil, fmt.Errorf("まだターンが終了していません")
	}

	player := g.Players.Next()
	nextTurnNumber := g.Turn.Current + 1
	turn := NewTurn(player, nextTurnNumber)
	g.Turn = turn
	return turn, nil
}

func (g *Game) RemainTurn() int {
	return g.MaxTurn - g.Turn.Current
}

func (g *Game) IsEnded() bool {
	if g.Turn == nil {
		return false
	}
	if g.RemainTurn() > 0 {
		return false
	}
	return true
}
