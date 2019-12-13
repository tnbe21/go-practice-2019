package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt = 0
var correctCnt = 0
const MAX_QUIZ_COUNT = 3

func main() {
	fmt.Println("クイズを開始します。")

	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		cnt += 1
		quiz := generateQuiz()
		fmt.Fprintf(os.Stdout, "Q.%d %s\n", cnt, quiz)
		fmt.Print("> ")
		scanner.Scan()
		ans := scanner.Text()
		if (isRight(ans, quiz)) {
		  correctCnt += 1
			fmt.Println("正解")
		} else {
			fmt.Println("不正解")
		}

		if cnt == MAX_QUIZ_COUNT {
			break
		}
	}
	fmt.Println("***")
	fmt.Fprintf(os.Stdout, "%d問中%d問正解\n", cnt, correctCnt)
	fmt.Println("***")
}

type quiz struct{}

func generateQuiz() string {
  return "2 + 4は？"
}

func isRight(quiz, ans string) bool {
  return true
}
