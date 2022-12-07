package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/koluku/yacht"
	"github.com/olekukonko/tablewriter"
)

func main() {
	if err := play(); err != nil {
		log.Fatal(err)
	}
}

func play() error {
	players := yacht.NewPlayers()

	fmt.Println("名前を入力してください")
	name := enterName()
	me := yacht.NewPlayer(name)
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

		fmt.Printf("%sのターンです\n", turn.Player.Name)

	LOOP:
		for {
			switch phase.Type {
			case yacht.PhaseTypeStandby:
				fmt.Printf("%sのスタンバイフェイズ\n", turn.Player.Name)
				if err := phase.Next(); err != nil {
					return err
				}

			case yacht.PhaseTypeRoll:
				fmt.Printf("%sのロールフェイズ\n", turn.Player.Name)
				phase.DiceBox.Roll()
				fmt.Println(phase.DiceBox.Dices.Rolled())
				if err := phase.Next(); err != nil {
					return err
				}

			case yacht.PhaseTypePick:
				fmt.Printf("%sのサイコロ選択フェイズ\n", turn.Player.Name)
				if turn.Current%2 == 1 {
					if phase.DiceBox.RolledTimes == 3 {
						a := phase.DiceBox.Dices.Rolled()
						fmt.Printf("3回振ったので%vをキープします\n", a)
						phase.DiceBox.Pick(a)
					} else {
						fmt.Println("半角スペース区切りでキープするサイコロの数字を選択してください")
						scanner := bufio.NewScanner(os.Stdin)
						scanner.Scan()
						inputs := strings.Split(scanner.Text(), " ")
						var a []int
						if len(inputs) == 1 && inputs[0] == "" {
							fmt.Println("キープしませんでした")
						} else {
							for _, b := range inputs {
								f, err := strconv.Atoi(b)
								if err != nil {
									return err
								}
								if f < 1 || f > 5 {
									return fmt.Errorf("1~5までの数字を入力してください")
								}
								a = append(a, f)
							}
						}
						if len(a) > 0 {
							fmt.Printf("%vをキープします\n", a)
							phase.DiceBox.Pick(a)
						}
					}
				} else {
					a := phase.DiceBox.Dices.Rolled()
					phase.DiceBox.Pick(a)
					fmt.Printf("%vをキープしました\n", a)
				}
				if err := phase.Next(); err != nil {
					return err
				}

			case yacht.PhaseTypeScoring:
				fmt.Printf("%sの得点確定フェイズ\n", turn.Player.Name)
				yaku := &yacht.Yaku{}
				yaku.Calculate(phase.DiceBox.Dices.Picked())

				var currentPlayer *yacht.Player
				if turn.Current%2 == 1 {
					currentPlayer = me
					for {
						yn, po, err := user(currentPlayer.Score, yaku)
						if err != nil {
							log.Println(err)
							continue
						}
						fmt.Printf("役名: %s, 得点: %d\n", yn, po)
						break
					}
				} else {
					currentPlayer = bot
					yn, po, err := best(currentPlayer.Score, yaku)
					if err != nil {
						return err
					}
					fmt.Printf("役名: %s, 得点: %d\n", yn, po)
				}
				if err := phase.Next(); err != nil {
					return err
				}
			default:
				break LOOP
			}
		}

		fmt.Printf("============================\n")
	}

	fmt.Printf("%sの点数: %+v\n", me.Name, me.Score.Point())
	fmt.Printf("%sの点数: %+v\n", bot.Name, bot.Score.Point())

	return nil
}

func enterName() string {
	var name string
	fmt.Scanf("%s", &name)
	if name == "" {
		return "あなた"
	}
	return name
}

func user(score *yacht.Score, yaku *yacht.Yaku) (string, int, error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Yaku", "Point"})
	yakuMap := yaku.ToMap()
	scoreFillableMap := score.FillableMap()
	for _, name := range []string{"Ace", "Twos", "Threes", "Fours", "Fives", "Sixes", "Choice", "FourOfAKind", "FullHouse", "LittleStraight", "BigStraight", "Yacht"} {
		if scoreFillableMap[name] {
			table.Append([]string{name, strconv.Itoa(yakuMap[name])})
		}
	}
	table.Render()

	var yakuName string
	fmt.Println("スコアを確定させる役名を入力してください")
	fmt.Scanf("%s", &yakuName)
	point := yakuMap[yakuName]
	if err := score.Fill(yakuName, point); err != nil {
		return "", 0, err
	}
	return yakuName, point, nil
}

func best(score *yacht.Score, yaku *yacht.Yaku) (string, int, error) {
	yakuMap := yaku.ToMap()
	var yakuSlice []map[string]any
	priority := 1
	for k, v := range yakuMap {
		yakuSlice = append(yakuSlice, map[string]any{"priority": priority, "name": k, "value": v})
		priority++
	}
	sort.Slice(yakuSlice, func(i, j int) bool { return yakuSlice[i]["value"].(int) > yakuSlice[j]["value"].(int) })
	scoreFillableMap := score.FillableMap()
	var yakuName string
	var point int
	for _, y := range yakuSlice {
		if scoreFillableMap[y["name"].(string)] {
			if err := score.Fill(y["name"].(string), y["value"].(int)); err != nil {
				return "", 0, err
			}
			yakuName = y["name"].(string)
			point = y["value"].(int)
			break
		}
	}
	return yakuName, point, nil
}
