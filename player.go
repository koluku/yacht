package yahtzee

import (
	"github.com/samber/lo"
)

type Player struct {
	Name  string
	Order int
	Score *Score
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Score: NewScore(),
	}
}

type Players struct {
	Players   []*Player
	Current   int
	AllPlayer int
}

func NewPlayers() *Players {
	return &Players{
		Current: 0,
	}
}

func (ps *Players) Add(p *Player) {
	p.Order = len(ps.Players) + 1
	ps.Players = append(ps.Players, p)
	ps.AllPlayer = len(ps.Players)
}

func (ps *Players) Shuffle() {
	ps.Players = lo.Shuffle(ps.Players)
	for i, p := range ps.Players {
		p.Order = i + 1
	}
}

func (ps *Players) First() *Player {
	return ps.Players[0]
}

func (ps *Players) Next() *Player {
	next := ps.Current + 1
	if next > len(ps.Players)-1 {
		next = 0
	}
	ps.Current = next
	return ps.Players[next]
}
