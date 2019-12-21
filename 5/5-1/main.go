package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"quiz/quiz"
	"time"
)

var cnt = 0
var correctCnt = 0

const MAX_QUIZ_COUNT = 3
const TIMEOUT_PER_QUIZ = 5

func main() {
	fmt.Println("クイズを開始します。")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		cnt += 1
		quiz := quiz.GetQuiz()
		go viewQuiz(quiz.GetBody())

		scanner.Scan()
		ans := scanner.Text()
		if quiz.IsCorrect(ans) {
			correctCnt += 1
			fmt.Println("正解")
		} else {
			fmt.Println("不正解")
		}
		cnt += 1
		if cnt == MAX_QUIZ_COUNT {
			break
		}
	}
	fmt.Println("***")
	fmt.Fprintf(os.Stdout, "%d問中%d問正解\n", cnt, correctCnt)
	fmt.Println("***")
}

func viewQuiz(body string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*TIMEOUT_PER_QUIZ)
	var tickCnt = 0
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("不正解")
			cnt += 1
			return
		default:
			if tickCnt > 0 {
				fmt.Fprintf(os.Stdout, "\033[1A\r")
			}
			fmt.Fprintf(os.Stdout, "Q.%d %s %d\n> ", cnt, body, TIMEOUT_PER_QUIZ-tickCnt)
			time.Sleep(time.Second)
			tickCnt += 1
		}
	}
}
