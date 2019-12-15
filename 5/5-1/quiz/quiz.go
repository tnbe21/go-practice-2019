package quiz

import (
	"math/rand"
	"time"
)

type quiz struct {
	body   string
	answer string
}

var quizList = []quiz{
	quiz{"2+4は?", "6"},
	quiz{"犬は英語で?", "dog"},
	quiz{"日本の初代総理大臣は?", "伊藤博文"},
	quiz{"アメリカ合衆国の現在の大統領は?", "トランプ"},
}

func GetQuiz() quiz {
	now := time.Now()
	rand.Seed(now.UnixNano())
	choosedIdx := rand.Intn(len(quizList))
	_quiz := quizList[choosedIdx]

	removedList := append(quizList[:choosedIdx], quizList[choosedIdx+1:]...)
	quizList = make([]quiz, len(removedList))
	copy(quizList, removedList)

	return _quiz
}

func (q *quiz) IsCorrect(ans string) bool {
	return ans == q.answer
}

func (q *quiz) GetBody() string {
	return q.body
}
