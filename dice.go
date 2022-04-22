package yahtzee

import (
	"math/rand"
	"sort"
)

var (
	Pips = map[int]string{
		1: "⚀",
		2: "⚁",
		3: "⚂",
		4: "⚃",
		5: "⚄",
		6: "⚅",
	}
)

type Dice struct {
	Number int
	Pip    string
}

func NewDice() *Dice {
	return &Dice{}
}

func (d *Dice) Roll(seed int64) {
	rand.Seed(seed)
	number := rand.Intn(5) + 1
	d.Number = number
	d.Pip = Pips[number]
}

type Dices []*Dice

func NewDices(num int) Dices {
	dices := Dices{}
	for i := 0; i < num; i++ {
		dices = append(dices, NewDice())
	}
	return dices
}

func (ds Dices) Roll(seed int64) {
	for _, dice := range ds {
		dice.Roll(seed)
	}
}

func (ds Dices) Open() []int {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Number < ds[j].Number
	})
	var numbers []int
	for _, dice := range ds {
		numbers = append(numbers, dice.Number)
	}
	return numbers
}

type DiceBox struct {
	// 振る前のサイコロ
	ReadyDices Dices
	// 振った後のサイコロ
	RolledDices Dices
	// 確定済みのサイコロ
	PickedDices Dices
}

func NewDiceBox(num int) DiceBox {
	return DiceBox{
		ReadyDices: NewDices(num),
	}
}
