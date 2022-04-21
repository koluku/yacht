package yahtzee

import (
	"sort"

	"github.com/samber/lo"
)

const (
	NumberOfYakuType = 12
)

type Yaku struct {
	Ace           int
	Twos          int
	Threes        int
	Fours         int
	Fives         int
	Sixes         int
	Choice        int
	FourOfAKind   int
	FullHouse     int
	SmallStraight int
	BigStraight   int
	Yahtzee       int
}

func (y *Yaku) Calculate(dices []int) {
	y.Ace = Ace(dices)
	y.Twos = Twos(dices)
	y.Threes = Threes(dices)
	y.Fours = Fours(dices)
	y.Fives = Fives(dices)
	y.Sixes = Sixes(dices)
	y.Choice = Choice(dices)
	y.FourOfAKind = FourOfAKind(dices)
	y.FullHouse = FullHouse(dices)
	y.SmallStraight = SmallStraight(dices)
	y.BigStraight = BigStraight(dices)
	y.Yahtzee = Yahtzee(dices)
}

// 任意の組み合わせ(１の目の合計)
func Ace(dices []int) int {
	var value int
	for _, dice := range dices {
		if dice == 1 {
			value += dice
		}
	}
	return value
}

// 任意の組み合わせ(2の目の合計)
func Twos(dices []int) int {
	var value int
	for _, dice := range dices {
		if dice == 2 {
			value += dice
		}
	}
	return value
}

// 任意の組み合わせ(3の目の合計)
func Threes(dices []int) int {
	var value int
	for _, dice := range dices {
		if dice == 3 {
			value += dice
		}
	}
	return value
}

// 任意の組み合わせ(4の目の合計)
func Fours(dices []int) int {
	var value int
	for _, dice := range dices {
		if dice == 4 {
			value += dice
		}
	}
	return value
}

// 任意の組み合わせ(5の目の合計)
func Fives(dices []int) int {
	var value int
	for _, dice := range dices {
		if dice == 5 {
			value += dice
		}
	}
	return value
}

// 任意の組み合わせ(6の目の合計)
func Sixes(dices []int) int {
	var value int
	for _, dice := range dices {
		if dice == 6 {
			value += dice
		}
	}
	return value
}

// 任意の組み合わせ(すべての目の合計)
func Choice(dices []int) int {
	var value int
	for _, dice := range dices {
		value += dice
	}
	return value
}

// 同じ目を４つ以上揃える(すべての目の合計)
func FourOfAKind(dices []int) int {
	groups := lo.GroupBy(dices, func(i int) int {
		return i
	})
	if len(groups) != 2 {
		return 0
	}

	keys := make(map[int]bool)
	for _, group := range groups {
		keys[len(group)] = true
	}
	if !keys[1] || !keys[4] {
		return 0
	}

	var value int
	for _, dice := range dices {
		value += dice
	}
	return value
}

// 3つの同じ目と2つの同じ目を揃える(すべての目の合計)
func FullHouse(dices []int) int {
	groups := lo.GroupBy(dices, func(i int) int {
		return i
	})
	if len(groups) != 2 {
		return 0
	}

	keys := make(map[int]bool)
	for _, group := range groups {
		keys[len(group)] = true
	}
	if !keys[2] || !keys[3] {
		return 0
	}

	var value int
	for _, dice := range dices {
		value += dice
	}
	return value
}

// ４つ以上の目が連続している(15点)
func SmallStraight(dices []int) int {
	sort.Ints(dices)
	stack := [][]int{dices[0:4], dices[1:5]}
	var ok bool
	for _, sdices := range stack {
		ok = true
		previus := sdices[0]
		for _, dice := range sdices[1:] {
			if dice != previus+1 {
				ok = false
				break
			}
			previus = dice
		}
		if ok {
			return 15
		}
	}

	return 0
}

// ５つの目が連続している(30点)
func BigStraight(dices []int) int {
	sort.Ints(dices)
	previus := dices[0]
	for _, dice := range dices[1:] {
		if dice != previus+1 {
			return 0
		}
		previus = dice
	}
	return 30
}

// ５つの目がすべて同じ(50点)
func Yahtzee(dices []int) int {
	groups := lo.GroupBy(dices, func(i int) int {
		return i
	})
	if len(groups) != 1 {
		return 0
	}
	return 50

}
