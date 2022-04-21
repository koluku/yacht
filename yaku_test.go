package yahtzee

import (
	"testing"
)

func TestAce(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  1,
		},
		{
			name:         "役が成立するが0点",
			dices:        []int{2, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Ace(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestTwos(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  2,
		},
		{
			name:         "役が成立するが0点",
			dices:        []int{1, 3, 3, 4, 5},
			expectVerify: true,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Twos(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestThrees(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  3,
		},
		{
			name:         "役が成立するが0点",
			dices:        []int{1, 2, 4, 4, 5},
			expectVerify: true,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Threes(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestFours(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  4,
		},
		{
			name:         "役が成立するが0点",
			dices:        []int{1, 2, 3, 5, 5},
			expectVerify: true,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Fours(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestFives(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  5,
		},
		{
			name:         "役が成立するが0点",
			dices:        []int{1, 2, 3, 4, 6},
			expectVerify: true,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Fives(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestSixes(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{2, 3, 4, 5, 6},
			expectVerify: true,
			expectValue:  6,
		},
		{
			name:         "役が成立するが0点",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Sixes(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestChoice(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  15,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Choice(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestFourOfAKind(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 4, 4, 4, 4},
			expectVerify: true,
			expectValue:  17,
		},
		{
			name:         "2:3で役が成立しない",
			dices:        []int{2, 2, 3, 3, 3},
			expectVerify: false,
			expectValue:  0,
		},
		{
			name:         "1:2:2で役が成立しない",
			dices:        []int{1, 2, 2, 3, 3},
			expectVerify: false,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := FourOfAKind(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestFullHouse(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{2, 2, 3, 3, 3},
			expectVerify: true,
			expectValue:  13,
		},
		{
			name:         "1:2:2で役が成立しない",
			dices:        []int{1, 2, 2, 3, 3},
			expectVerify: false,
			expectValue:  0,
		},
		{
			name:         "1:4で役が成立しない",
			dices:        []int{1, 4, 4, 4, 4},
			expectVerify: false,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := FullHouse(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestSmallStraight(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "前に4連続で役が成立する",
			dices:        []int{1, 2, 3, 4, 6},
			expectVerify: true,
			expectValue:  15,
		},
		{
			name:         "後ろに4連続で役が成立する",
			dices:        []int{2, 3, 4, 5, 6},
			expectVerify: true,
			expectValue:  15,
		},
		{
			name:         "5連続で役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  15,
		},
		{
			name:         "途中で連続が途切れて役が成立しない",
			dices:        []int{1, 2, 3, 5, 6},
			expectVerify: false,
			expectValue:  0,
		},
		{
			name:         "不揃いで役が成立しない",
			dices:        []int{1, 4, 4, 5, 5},
			expectVerify: false,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := SmallStraight(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestBigStraight(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: true,
			expectValue:  30,
		},
		{
			name:         "役が成立する",
			dices:        []int{2, 3, 4, 5, 6},
			expectVerify: true,
			expectValue:  30,
		},
		{
			name:         "4連続で役が成立しない",
			dices:        []int{1, 2, 3, 4, 6},
			expectVerify: false,
			expectValue:  0,
		},
		{
			name:         "途中で連続が途切れて役が成立しない",
			dices:        []int{1, 2, 3, 5, 6},
			expectVerify: false,
			expectValue:  0,
		},
		{
			name:         "不揃いで役が成立しない",
			dices:        []int{1, 4, 4, 5, 5},
			expectVerify: false,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := BigStraight(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}

func TestYahtzee(t *testing.T) {
	cases := []struct {
		name         string
		dices        []int
		expectVerify bool
		expectValue  int
	}{
		{
			name:         "役が成立する",
			dices:        []int{6, 6, 6, 6, 6},
			expectVerify: true,
			expectValue:  50,
		},
		{
			name:         "役が成立しない",
			dices:        []int{1, 2, 3, 4, 5},
			expectVerify: false,
			expectValue:  0,
		},
		{
			name:         "1:4で役が成立しない",
			dices:        []int{1, 4, 4, 4, 4},
			expectVerify: false,
			expectValue:  0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			verify, value := Yahtzee(tt.dices)
			if verify != tt.expectVerify {
				t.Errorf("verify test failed, expect: %v, verify: %v", tt.expectVerify, verify)
			}
			if value != tt.expectValue {
				t.Errorf("value test failed, expect: %v, value: %v", tt.expectValue, value)
			}
		})
	}
}
