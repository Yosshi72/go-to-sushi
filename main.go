package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/Yosshi72/sushi"
)

var (
	sushiTypes = []string{"maguro","sushi"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("寿司打ゲームを開始します！Enterキーを押してください。")
	waitForEnter()

	score := 0
	rounds := 10

	for i := 0; i < rounds; i++ {
		sushi := generateRandomSushi()
		fmt.Printf("【ラウンド%d】%sが流れてきました！\n", i+1, sushi)
		fmt.Print("ユーザー入力：")
		userInput := getUserInput()

		if strings.ToLower(userInput) == sushi {
			score++
			fmt.Println("正解！ポイントを獲得しました。")
		} else {
			fmt.Println("不正解！ポイントを失いました。")
		}

		fmt.Println()
	}

	fmt.Printf("ゲーム終了！あなたのスコアは%d/%dです。\n", score, rounds)
}

func waitForEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getUserInput() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(input)
}

func generateRandomSushi() string {
	index := rand.Intn(len(sushiTypes))
	return sushiTypes[index]
}
