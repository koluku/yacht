package yahtzee

import (
	"testing"
)

func TestAce(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 1,
		},
		{
			name:   "役が成立するが0点",
			dices:  []int{2, 2, 3, 4, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Ace(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestTwos(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 2,
		},
		{
			name:   "役が成立するが0点",
			dices:  []int{1, 3, 3, 4, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Twos(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestThrees(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 3,
		},
		{
			name:   "役が成立するが0点",
			dices:  []int{1, 2, 4, 4, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Threes(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestFours(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 4,
		},
		{
			name:   "役が成立するが0点",
			dices:  []int{1, 2, 3, 5, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Fours(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestFives(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 5,
		},
		{
			name:   "役が成立するが0点",
			dices:  []int{1, 2, 3, 4, 6},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Fives(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestSixes(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{2, 3, 4, 5, 6},
			expect: 6,
		},
		{
			name:   "役が成立するが0点",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Sixes(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestChoice(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 15,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Choice(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestFourOfAKind(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 4, 4, 4, 4},
			expect: 17,
		},
		{
			name:   "2:3で役が成立しない",
			dices:  []int{2, 2, 3, 3, 3},
			expect: 0,
		},
		{
			name:   "1:2:2で役が成立しない",
			dices:  []int{1, 2, 2, 3, 3},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := FourOfAKind(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestFullHouse(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{2, 2, 3, 3, 3},
			expect: 13,
		},
		{
			name:   "1:2:2で役が成立しない",
			dices:  []int{1, 2, 2, 3, 3},
			expect: 0,
		},
		{
			name:   "1:4で役が成立しない",
			dices:  []int{1, 4, 4, 4, 4},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := FullHouse(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestSmallStraight(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "前に4連続で役が成立する",
			dices:  []int{1, 2, 3, 4, 6},
			expect: 15,
		},
		{
			name:   "後ろに4連続で役が成立する",
			dices:  []int{2, 3, 4, 5, 6},
			expect: 15,
		},
		{
			name:   "5連続で役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 15,
		},
		{
			name:   "途中で連続が途切れて役が成立しない",
			dices:  []int{1, 2, 3, 5, 6},
			expect: 0,
		},
		{
			name:   "不揃いで役が成立しない",
			dices:  []int{1, 4, 4, 5, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := SmallStraight(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestBigStraight(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 30,
		},
		{
			name:   "役が成立する",
			dices:  []int{2, 3, 4, 5, 6},
			expect: 30,
		},
		{
			name:   "4連続で役が成立しない",
			dices:  []int{1, 2, 3, 4, 6},
			expect: 0,
		},
		{
			name:   "途中で連続が途切れて役が成立しない",
			dices:  []int{1, 2, 3, 5, 6},
			expect: 0,
		},
		{
			name:   "不揃いで役が成立しない",
			dices:  []int{1, 4, 4, 5, 5},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := BigStraight(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}

func TestYahtzee(t *testing.T) {
	cases := []struct {
		name   string
		dices  []int
		expect int
	}{
		{
			name:   "役が成立する",
			dices:  []int{6, 6, 6, 6, 6},
			expect: 50,
		},
		{
			name:   "役が成立しない",
			dices:  []int{1, 2, 3, 4, 5},
			expect: 0,
		},
		{
			name:   "1:4で役が成立しない",
			dices:  []int{1, 4, 4, 4, 4},
			expect: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			value := Yahtzee(tt.dices)
			if value != tt.expect {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expect, value)
			}
		})
	}
}
