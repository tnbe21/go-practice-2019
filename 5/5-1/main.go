package main

import (
	"bufio"
	"fmt"
	"os"
	"quiz/quiz"
)

var cnt = 0
var correctCnt = 0

const MAX_QUIZ_COUNT = 3

func main() {
	fmt.Println("クイズを開始します。")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		cnt += 1
		quiz := quiz.GetQuiz()
		fmt.Fprintf(os.Stdout, "Q.%d %s\n", cnt, quiz.GetBody())
		fmt.Print("> ")
		scanner.Scan()
		ans := scanner.Text()
		if quiz.IsCorrect(ans) {
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
