package yacht

import (
	"math/rand"
	"sort"
	"time"
)

type DiceStatus int

const (
	DiceStatusReady DiceStatus = iota
	DiceStatusRolled
	DiceStatusPicked
)

var (
	DiceEmojiMap = map[int]string{
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
	Emoji  string
	Status DiceStatus
}

func NewDice() *Dice {
	return &Dice{
		Status: DiceStatusReady,
	}
}

func (d *Dice) Roll() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6) + 1
	d.Number = number
	d.Emoji = DiceEmojiMap[number]
	d.Status = DiceStatusRolled
}

type Dices []*Dice

func (ds Dices) Rolls() {
	for _, dice := range ds {
		if dice.Status == DiceStatusReady {
			dice.Roll()
		}
	}
}

func (ds Dices) Picks(numbers []int) {
	for _, number := range numbers {
		for _, dice := range ds {
			if dice.Status != DiceStatusRolled {
				continue
			}
			if dice.Number == number {
				dice.Status = DiceStatusPicked
			}
		}
	}
}

// Pick以外のサイコロをReadyに戻す
func (ds Dices) Clear() {
	for _, dice := range ds {
		if dice.Status == DiceStatusPicked {
			continue
		}
		dice.Status = DiceStatusReady
	}
}

func (ds Dices) Readied() []int {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Number < ds[j].Number
	})
	var numbers []int
	for _, dice := range ds {
		if dice.Status != DiceStatusReady {
			continue
		}
		numbers = append(numbers, dice.Number)
	}
	return numbers
}

func (ds Dices) Rolled() []int {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Number < ds[j].Number
	})
	var numbers []int
	for _, dice := range ds {
		if dice.Status != DiceStatusRolled {
			continue
		}
		numbers = append(numbers, dice.Number)
	}
	return numbers
}

func (ds Dices) Picked() []int {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Number < ds[j].Number
	})
	var numbers []int
	for _, dice := range ds {
		if dice.Status != DiceStatusPicked {
			continue
		}
		numbers = append(numbers, dice.Number)
	}
	return numbers
}

type DiceBox struct {
	Dices       Dices
	RolledTimes int
}

func NewDiceBox(num int) *DiceBox {
	dices := Dices{}
	for i := 0; i < num; i++ {
		dices = append(dices, NewDice())
	}
	return &DiceBox{
		Dices: dices,
	}
}

func (db DiceBox) ReadyNum() int {
	var num int
	for _, dice := range db.Dices {
		if dice.Status == DiceStatusReady {
			num++
		}
	}
	return num
}

func (db DiceBox) RolledNum() int {
	var num int
	for _, dice := range db.Dices {
		if dice.Status == DiceStatusRolled {
			num++
		}
	}
	return num
}

func (db DiceBox) PickedNum() int {
	var num int
	for _, dice := range db.Dices {
		if dice.Status == DiceStatusPicked {
			num++
		}
	}
	return num
}

func (db *DiceBox) Roll() {
	db.Dices.Clear()
	db.Dices.Rolls()
	db.RolledTimes++
}

func (db *DiceBox) Pick(ids []int) {
	db.Dices.Picks(ids)
}
