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
		quiz := getQuiz()
		fmt.Fprintf(os.Stdout, "Q.%d %s\n", cnt, quiz.body)
		fmt.Print("> ")
		scanner.Scan()
		ans := scanner.Text()
		if (isRight(quiz, ans)) {
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

type Quiz struct {
  body string
  answer string
}

func getQuiz() Quiz {
  return Quiz { "2 + 4 = ?", "6" }
}

func isRight(quiz Quiz, ans string) bool {
  return ans == quiz.answer
}
