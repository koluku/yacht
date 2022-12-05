package yahtzee

import (
	"fmt"
)

type Turn struct {
	Current int
	Player  *Player
	Phase   PhaseType
	DiceBox *DiceBox
}

func NewTurn(player *Player, current int) *Turn {
	return &Turn{
		Player:  player,
		Current: current,
		Phase:   PhaseTypeBegin,
		DiceBox: NewDiceBox(5),
	}
}

type PhaseType int

const (
	PhaseTypeBegin PhaseType = iota
	PhaseTypeRoll
	PhaseTypePick
	PhaseTypeScoring
	PhaseTypeEnd
)

type Phase struct {
	Turn    int
	Player  *Player
	Type    PhaseType
	DiceBox *DiceBox
}

func (t *Turn) NextPhase() error {
	var err error = nil
	switch t.Phase {
	case PhaseTypeBegin:
		t.MoveRollPhase()
	case PhaseTypeRoll:
		if !t.CanMovePickPhase() {
			err = fmt.Errorf("サイコロを振っていない場合は残すサイコロを決めることができない")
			break
		}
		// t.PickPhase()
	case PhaseTypePick:
		if !t.CanMoveScoringPhase() {
			err = fmt.Errorf("残すサイコロが5つない場合はスコアを付けることができない")
			break
		}
		// t.ScoringPhase()
	case PhaseTypeScoring:
		if !t.CanMoveEndPhase() {
			err = fmt.Errorf("役からスコアをつけなかったら終わらせることができない")
			break
		}
		// t.EndPhase()
	case PhaseTypeEnd:
		err = fmt.Errorf("既にターンが終わっているので次のフェイズに移行できない")
	}
	return err
}

func (t *Turn) MoveRollPhase() {
	t.Phase = PhaseTypeRoll
}

func (t *Turn) RollDice(seed int64) {
	t.DiceBox.Roll(seed)
}

func (t *Turn) CanMoveRollPhase() bool {
	return t.DiceBox.RolledTimes < 3
}

func (t *Turn) IsRollPhase() bool {
	return t.Phase == PhaseTypeRoll
}

func (t *Turn) MovePickPhase() {
	t.Phase = PhaseTypePick
}

func (t *Turn) PickDice(ids []int) {
	// t.DiceBox.Pick(ids)
}

func (t *Turn) CanMovePickPhase() bool {
	return t.DiceBox.ReadyNum() == 0
}

func (t *Turn) ScoringPhase() {
	t.Phase = PhaseTypeScoring
}

func (t *Turn) CanMoveScoringPhase() bool {
	return true
}

func (t *Turn) MoveEndPhase() {
	t.Phase = PhaseTypeEnd
}

func (t *Turn) CanMoveEndPhase() bool {
	return true
}

func (t *Turn) IsCompleted() bool {
	return t.Phase == PhaseTypeEnd
}
