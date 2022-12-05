package main

import (
	"fmt"
	"log"

	"github.com/koluku/yacht"
)

func main() {
	if err := play(); err != nil {
		log.Fatal(err)
	}
}

func play() error {
	players := yacht.NewPlayers()
	me := yacht.NewPlayer("me")
	bot := yacht.NewPlayer("bot")
	players.Add(me)
	players.Add(bot)

	game := yacht.NewGame(players)

	for !game.IsEnded() {
		turn, err := game.NextTurn()
		if err != nil {
			return err
		}
		phase := turn.Phase

		if phase.Type == yacht.PhaseTypeStandby {
			if err := phase.Next(); err != nil {
				return err
			}
		}
		if phase.Type == yacht.PhaseTypeRoll {
			phase.DiceBox.Roll()
			if err := phase.Next(); err != nil {
				return err
			}
		}
		if phase.Type == yacht.PhaseTypePick {
			phase.DiceBox.Pick([]int{0, 1, 2, 3, 4})
			if err := phase.Next(); err != nil {
				return err
			}
		}
		if phase.Type == yacht.PhaseTypeScoring {
			yaku := &yacht.Yaku{}
			yaku.Calculate(phase.DiceBox.Dices.Open())

			var currentPlayer *yacht.Player
			if turn.Current%2 == 1 {
				currentPlayer = me
			} else {
				currentPlayer = bot
			}
			switch {
			case turn.Current <= 2:
				fmt.Println("Ace")
				if err := currentPlayer.Score.FillAce(yaku.Ace); err != nil {
					return err
				}
			case turn.Current > 2 && turn.Current <= 4:
				fmt.Println("Twos")
				if err := currentPlayer.Score.FillTwos(yaku.Twos); err != nil {
					return err
				}
			case turn.Current > 4 && turn.Current <= 6:
				fmt.Println("Threes")
				if err := currentPlayer.Score.FillThrees(yaku.Threes); err != nil {
					return err
				}
			case turn.Current > 6 && turn.Current <= 8:
				fmt.Println("Fours")
				if err := currentPlayer.Score.FillFours(yaku.Fours); err != nil {
					return err
				}
			case turn.Current > 8 && turn.Current <= 10:
				fmt.Println("Fives")
				if err := currentPlayer.Score.FillFives(yaku.Fives); err != nil {
					return err
				}
			case turn.Current > 10 && turn.Current <= 12:
				fmt.Println("Sixes")
				if err := currentPlayer.Score.FillSixes(yaku.Sixes); err != nil {
					return err
				}
			case turn.Current > 12 && turn.Current <= 14:
				fmt.Println("Choice")
				if err := currentPlayer.Score.FillChoice(yaku.Choice); err != nil {
					return err
				}
			case turn.Current > 14 && turn.Current <= 16:
				fmt.Println("FourOfAKind")
				if err := currentPlayer.Score.FillFourOfAKind(yaku.FourOfAKind); err != nil {
					return err
				}
			case turn.Current > 16 && turn.Current <= 18:
				fmt.Println("FullHouse")
				if err := currentPlayer.Score.FillFullHouse(yaku.FullHouse); err != nil {
					return err
				}
			case turn.Current > 18 && turn.Current <= 20:
				fmt.Println("LittleStraight")
				if err := currentPlayer.Score.FillLittleStraight(yaku.LittleStraight); err != nil {
					return err
				}
			case turn.Current > 20 && turn.Current <= 22:
				fmt.Println("BigStraight")
				if err := currentPlayer.Score.FillBigStraight(yaku.BigStraight); err != nil {
					return err
				}
			case turn.Current > 22 && turn.Current <= 24:
				fmt.Println("Yacht")
				if err := currentPlayer.Score.FillYacht(yaku.Yacht); err != nil {
					return err
				}
			default:
				return fmt.Errorf("もう入力できないよ！")
			}

			if err := phase.Next(); err != nil {
				return err
			}
		}
	}

	fmt.Printf("%+v\n", me.Score.Point())
	fmt.Printf("%+v\n", bot.Score.Point())

	return nil
}
