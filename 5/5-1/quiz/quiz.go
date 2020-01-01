package quiz

import (
	"math/rand"
	"time"
)

type Quiz struct {
	body   string
	answer string
}

var quizList = []Quiz{
	Quiz{"2+4は?", "6"},
	Quiz{"犬は英語で?", "dog"},
	Quiz{"日本の初代総理大臣は?", "伊藤博文"},
	Quiz{"アメリカ合衆国の現在の大統領は?", "トランプ"},
}

func GetQuiz() Quiz {
	now := time.Now()
	rand.Seed(now.UnixNano())
	choosedIdx := rand.Intn(len(quizList))
	_quiz := quizList[choosedIdx]

	removedList := append(quizList[:choosedIdx], quizList[choosedIdx+1:]...)
	quizList = make([]Quiz, len(removedList))
	copy(quizList, removedList)

	return _quiz
}

func (q *Quiz) IsCorrect(ans string) bool {
	return ans == q.answer
}

func (q *Quiz) GetBody() string {
	return q.body
}
