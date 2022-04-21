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

func (y Yaku) Calculate(dices []int) {
	_, y.Ace = Ace(dices)
	_, y.Twos = Twos(dices)
	_, y.Threes = Threes(dices)
	_, y.Fours = Fours(dices)
	_, y.Fives = Fives(dices)
	_, y.Sixes = Sixes(dices)
	_, y.Choice = Choice(dices)
	_, y.FourOfAKind = FourOfAKind(dices)
	_, y.FullHouse = FullHouse(dices)
	_, y.SmallStraight = SmallStraight(dices)
	_, y.BigStraight = BigStraight(dices)
	_, y.Yahtzee = Yahtzee(dices)
}

// 任意の組み合わせ(１の目の合計)
func Ace(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		if dice == 1 {
			value += dice
		}
	}
	return true, value
}

// 任意の組み合わせ(2の目の合計)
func Twos(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		if dice == 2 {
			value += dice
		}
	}
	return true, value
}

// 任意の組み合わせ(3の目の合計)
func Threes(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		if dice == 3 {
			value += dice
		}
	}
	return true, value
}

// 任意の組み合わせ(4の目の合計)
func Fours(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		if dice == 4 {
			value += dice
		}
	}
	return true, value
}

// 任意の組み合わせ(5の目の合計)
func Fives(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		if dice == 5 {
			value += dice
		}
	}
	return true, value
}

// 任意の組み合わせ(6の目の合計)
func Sixes(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		if dice == 6 {
			value += dice
		}
	}
	return true, value
}

// 任意の組み合わせ(すべての目の合計)
func Choice(dices []int) (bool, int) {
	var value int
	for _, dice := range dices {
		value += dice
	}
	return true, value
}

// 同じ目を４つ以上揃える(すべての目の合計)
func FourOfAKind(dices []int) (bool, int) {
	groups := lo.GroupBy(dices, func(i int) int {
		return i
	})
	if len(groups) != 2 {
		return false, 0
	}

	keys := make(map[int]bool)
	for _, group := range groups {
		keys[len(group)] = true
	}
	if keys[1] == false || keys[4] == false {
		return false, 0
	}

	var value int
	for _, dice := range dices {
		value += dice
	}
	return true, value
}

// 3つの同じ目と2つの同じ目を揃える(すべての目の合計)
func FullHouse(dices []int) (bool, int) {
	groups := lo.GroupBy(dices, func(i int) int {
		return i
	})
	if len(groups) != 2 {
		return false, 0
	}

	keys := make(map[int]bool)
	for _, group := range groups {
		keys[len(group)] = true
	}
	if keys[2] == false || keys[3] == false {
		return false, 0
	}

	var value int
	for _, dice := range dices {
		value += dice
	}
	return true, value
}

// ４つ以上の目が連続している(15点)
func SmallStraight(dices []int) (bool, int) {
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
			return true, 15
		}
	}

	return false, 0
}

// ５つの目が連続している(30点)
func BigStraight(dices []int) (bool, int) {
	sort.Ints(dices)
	previus := dices[0]
	for _, dice := range dices[1:] {
		if dice != previus+1 {
			return false, 0
		}
		previus = dice
	}
	return true, 30
}

// ５つの目がすべて同じ(50点)
func Yahtzee(dices []int) (bool, int) {
	groups := lo.GroupBy(dices, func(i int) int {
		return i
	})
	if len(groups) != 1 {
		return false, 0
	}
	return true, 50

}
