package yahtzee

import (
	"fmt"

	"github.com/samber/lo"
)

type Score struct {
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

	IsFilledAce           bool
	IsFilledTwos          bool
	IsFilledThrees        bool
	IsFilledFours         bool
	IsFilledFives         bool
	IsFilledSixes         bool
	IsFilledChoice        bool
	IsFilledFourOfAKind   bool
	IsFilledFullHouse     bool
	IsFilledSmallStraight bool
	IsFilledBigStraight   bool
	IsFilledYahtzee       bool
}

func NewScore() *Score {
	return &Score{
		Ace:           0,
		Twos:          0,
		Threes:        0,
		Fours:         0,
		Fives:         0,
		Sixes:         0,
		Choice:        0,
		FourOfAKind:   0,
		FullHouse:     0,
		SmallStraight: 0,
		BigStraight:   0,
		Yahtzee:       0,

		IsFilledAce:           false,
		IsFilledTwos:          false,
		IsFilledThrees:        false,
		IsFilledFours:         false,
		IsFilledFives:         false,
		IsFilledSixes:         false,
		IsFilledChoice:        false,
		IsFilledFourOfAKind:   false,
		IsFilledFullHouse:     false,
		IsFilledSmallStraight: false,
		IsFilledBigStraight:   false,
		IsFilledYahtzee:       false,
	}
}

func (s *Score) FillAce(point int) error {
	if s.IsFilledAce {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Ace = point
	s.IsFilledAce = true
	return nil
}

func (s *Score) FillTwos(point int) error {
	if s.IsFilledTwos {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Twos = point
	s.IsFilledTwos = true
	return nil
}

func (s *Score) FillThrees(point int) error {
	if s.IsFilledThrees {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Threes = point
	s.IsFilledThrees = true
	return nil
}

func (s *Score) FillFours(point int) error {
	if s.IsFilledFours {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Fours = point
	s.IsFilledFours = true
	return nil
}

func (s *Score) FillFives(point int) error {
	if s.IsFilledFives {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Fives = point
	s.IsFilledFives = true
	return nil
}

func (s *Score) FillSixes(point int) error {
	if s.IsFilledSixes {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Sixes = point
	s.IsFilledSixes = true
	return nil
}

func (s *Score) FillChoice(point int) error {
	if s.IsFilledChoice {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Choice = point
	s.IsFilledChoice = true
	return nil
}

func (s *Score) FillFourOfAKind(point int) error {
	if s.IsFilledFourOfAKind {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.FourOfAKind = point
	s.IsFilledFourOfAKind = true
	return nil
}

func (s *Score) FillFullHouse(point int) error {
	if s.IsFilledFullHouse {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.FullHouse = point
	s.IsFilledFullHouse = true
	return nil
}

func (s *Score) FillSmallStraight(point int) error {
	if s.IsFilledSmallStraight {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.SmallStraight = point
	s.IsFilledSmallStraight = true
	return nil
}

func (s *Score) FillBigStraight(point int) error {
	if s.IsFilledBigStraight {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.BigStraight = point
	s.IsFilledBigStraight = true
	return nil
}

func (s *Score) FillYahtzee(point int) error {
	if s.IsFilledYahtzee {
		return fmt.Errorf("既に埋めてあるよ")
	}
	s.Yahtzee = point
	s.IsFilledYahtzee = true
	return nil
}

func (s *Score) Point() int {
	return lo.Sum(
		[]int{
			s.Ace,
			s.Twos,
			s.Threes,
			s.Fours,
			s.Fives,
			s.Sixes,
			s.Choice,
			s.FourOfAKind,
			s.FullHouse,
			s.SmallStraight,
			s.BigStraight,
			s.Yahtzee,
		},
	)
}

func (s *Score) IsFilledAll() bool {
	return lo.None(
		[]bool{
			s.IsFilledAce,
			s.IsFilledTwos,
			s.IsFilledThrees,
			s.IsFilledFours,
			s.IsFilledFives,
			s.IsFilledSixes,
			s.IsFilledChoice,
			s.IsFilledFourOfAKind,
			s.IsFilledFullHouse,
			s.IsFilledSmallStraight,
			s.IsFilledBigStraight,
			s.IsFilledYahtzee,
		},
		[]bool{false},
	)
}
