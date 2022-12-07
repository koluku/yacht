package yacht

import "fmt"

type Turn struct {
	Current int
	Player  *Player
	Phase   *Phase
	DiceBox *DiceBox
}

func NewTurn(player *Player, current int) *Turn {
	return &Turn{
		Player:  player,
		Current: current,
		Phase:   NewPhase(),
	}
}

type PhaseType int

const (
	PhaseTypeStandby PhaseType = iota
	PhaseTypeRoll
	PhaseTypePick
	PhaseTypeScoring
	PhaseTypeEnd
)

type Phase struct {
	Type    PhaseType
	DiceBox *DiceBox
}

func NewPhase() *Phase {
	return &Phase{
		Type:    PhaseTypeStandby,
		DiceBox: NewDiceBox(5),
	}
}

func (p *Phase) Next() error {
	switch p.Type {
	case PhaseTypeStandby:
		p.MoveRollPhase()
		return nil
	case PhaseTypeRoll:
		if !p.CanMovePickPhase() {
			return fmt.Errorf("サイコロを振っていない場合は残すサイコロを決めることができない")
		}
		p.MovePickPhase()
		return nil
	case PhaseTypePick:
		if p.CanMoveScoringPhase() {
			p.MoveScoringPhase()
			return nil
		}
		if p.CanMoveRollPhase() {
			p.MoveRollPhase()
			return nil
		}
		return fmt.Errorf("ピック確定させて！！")
	case PhaseTypeScoring:
		p.MoveEndPhase()
	case PhaseTypeEnd:
		return fmt.Errorf("終わってるよ！！")
	}
	return nil
}

func (p *Phase) CanMoveRollPhase() bool {
	if p.DiceBox.PickedNum() == 5 {
		return false
	} else if p.DiceBox.RolledTimes == 3 {
		return false
	}
	return true
}

func (p *Phase) MoveRollPhase() {
	p.Type = PhaseTypeRoll
}

func (p *Phase) CanMovePickPhase() bool {
	return p.DiceBox.RolledNum() > 0
}

func (p *Phase) MovePickPhase() {
	p.Type = PhaseTypePick
}

func (p *Phase) CanMoveScoringPhase() bool {
	return p.DiceBox.PickedNum() == 5
}

func (p *Phase) MoveScoringPhase() {
	p.Type = PhaseTypeScoring
}

func (p *Phase) MoveEndPhase() {
	p.Type = PhaseTypeEnd
}

// func (t *Turn) NextPhase() error {
// 	var err error = nil
// 	switch p {
// 	case PhaseTypeStandby:
// 		t.MoveRollPhase()
// 	case PhaseTypeRoll:
// 		if !t.CanMovePickPhase() {
// 			err = fmt.Errorf("サイコロを振っていない場合は残すサイコロを決めることができない")
// 			break
// 		}
// 		// t.PickPhase()
// 	case PhaseTypePick:
// 		if !t.CanMoveScoringPhase() {
// 			err = fmt.Errorf("残すサイコロが5つない場合はスコアを付けることができない")
// 			break
// 		}
// 		// t.ScoringPhase()
// 	case PhaseTypeScoring:
// 		if !t.CanMoveEndPhase() {
// 			err = fmt.Errorf("役からスコアをつけなかったら終わらせることができない")
// 			break
// 		}
// 		// t.EndPhase()
// 	case PhaseTypeEnd:
// 		err = fmt.Errorf("既にターンが終わっているので次のフェイズに移行できない")
// 	}
// 	return err
// }

// func (t *Turn) MoveRollPhase() {
// 	t.Phase = PhaseTypeRoll
// }

// func (t *Turn) RollDice(seed int64) {
// 	t.DiceBox.Roll(seed)
// }

// func (t *Turn) CanMoveRollPhase() bool {
// 	return t.DiceBox.RolledTimes < 3
// }

// func (t *Turn) IsRollPhase() bool {
// 	return t.Phase == PhaseTypeRoll
// }

// func (t *Turn) MovePickPhase() {
// 	t.Phase = PhaseTypePick
// }

// func (t *Turn) PickDice(ids []int) {
// 	// t.DiceBox.Pick(ids)
// }

// func (t *Turn) ScoringPhase() {
// 	t.Phase = PhaseTypeScoring
// }

// func (t *Turn) CanMoveScoringPhase() bool {
// 	return true
// }

// func (t *Turn) MoveEndPhase() {
// 	t.Phase = PhaseTypeEnd
// }

// func (t *Turn) CanMoveEndPhase() bool {
// 	return true
// }

func (t *Turn) IsEnded() bool {
	return t.Phase.Type == PhaseTypeEnd
}
