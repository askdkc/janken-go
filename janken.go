// じゃんけんゲームを作るよ
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数値からグー、チョキ、パーを判断して返す
func handToString(hand int) string {
	switch hand {
	case 1:
		return "グー"
	case 2:
		return "チョキ"
	case 3:
		return "パー"
	}
	return ""
}

func janken() {
	// じゃんけんゲームと出力する
	fmt.Println("じゃんけんゲームだよ")
	// じゃんけんの手を入力してもらう
	fmt.Println("何を出す？  1:グー、2:チョキ、3:パー")
	// じゃんけんの手を入力してもらう
	var hand int
	fmt.Scan(&hand)

	// コンピュータの手をランダムに決める
	rand.Seed(time.Now().UnixNano())
	computer := rand.Intn(3) + 1

	// じゃんけんの結果を表示する
	fmt.Println("あなたの手:", handToString(hand))
	fmt.Println("コンピュータの手:", handToString(computer))

	// じゃんけんの結果を表示する
	if hand == computer {
		fmt.Println("あいこ")
	}
	if hand == 1 && computer == 2 {
		fmt.Println("あなたの勝ち")
	}
	if hand == 2 && computer == 3 {
		fmt.Println("あなたの勝ち")
	}
	if hand == 3 && computer == 1 {
		fmt.Println("あなたの勝ち")
	}
	if hand == 1 && computer == 3 {
		fmt.Println("あなたの負け")
	}
	if hand == 2 && computer == 1 {
		fmt.Println("あなたの負け")
	}
	if hand == 3 && computer == 2 {
		fmt.Println("あなたの負け")
	}
}

func main() {
	// じゃんけんを繰り返し実行する
	for {

		// じゃんけんを実行する
		janken()

		// じゃんけんを繰り返すかどうかを聞く
		fmt.Println("もう一度？ 0:いいえ、1:はい")
		var answer int
		fmt.Scan(&answer)
		if answer != 1 {
			break
		}
	}
}
