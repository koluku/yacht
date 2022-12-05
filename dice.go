package yacht

import (
	"math/rand"
	"sort"
	"time"

	"github.com/samber/lo"
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
	ID     int
	Number int
	Emoji  string
	Status DiceStatus
}

func NewDice(id int) *Dice {
	return &Dice{
		ID:     id,
		Status: DiceStatusReady,
	}
}

func (d *Dice) Roll() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(5) + 1
	d.Number = number
	d.Emoji = DiceEmojiMap[number]
	d.Status = DiceStatusRolled
}

type Dices []*Dice

func (ds Dices) Rolls() {
	for _, dice := range ds {
		if dice.Status != DiceStatusReady {
			continue
		}
		dice.Roll()
	}
}

func (ds Dices) Picks(ids []int) {
	for _, dice := range ds {
		if lo.Contains(ids, dice.ID) {
			dice.Status = DiceStatusPicked
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
	Dices       Dices
	RolledTimes int
}

func NewDiceBox(num int) *DiceBox {
	dices := Dices{}
	for i := 0; i < num; i++ {
		dices = append(dices, NewDice(i))
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
