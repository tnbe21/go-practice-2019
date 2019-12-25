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
var finished chan bool
var correctCnt = 0

const MAX_QUIZ_COUNT = 3
const TIMEOUT_PER_QUIZ = 5

func main() {
	fmt.Println("クイズを開始します。")

	for {
		doQuiz()
		if <-finished {
			finished <- false
			break
		}
	}

	fmt.Println("***")
	fmt.Fprintf(os.Stdout, "%d問中%d問正解\n", cnt, correctCnt)
	fmt.Println("***")
}

func doQuiz() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*TIMEOUT_PER_QUIZ)
	_quiz := quiz.GetQuiz()
	defer cancel()
	go scanAnswer(ctx, _quiz)
	var tickCnt = 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("不正解")
			cnt += 1
			finished <- true
			return
		default:
			if tickCnt > 0 {
				fmt.Fprintf(os.Stdout, "\033[1A\r")
			}
			fmt.Fprintf(os.Stdout, "Q.%d %s %d\n> ", cnt, _quiz.GetBody(), TIMEOUT_PER_QUIZ-tickCnt)
			time.Sleep(time.Second)
			tickCnt += 1
		}
	}
}

func scanAnswer(ctx context.Context, _quiz quiz.Quiz) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			answer := scanner.Text()
			if _quiz.IsCorrect(answer) {
				correctCnt += 1
				fmt.Println("正解")
			} else {
				fmt.Println("不正解")
			}
			cnt += 1
			finished <- true
			return
		}
	}
}
