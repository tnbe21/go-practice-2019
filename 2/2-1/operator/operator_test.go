package operator

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCreateOperator(t *testing.T) {
	t.Run("should create adder", func(t *testing.T) {
		t.Parallel()
		doCreateNormalTest(t, "+")
	})
	t.Run("should create subtractor", func(t *testing.T) {
		t.Parallel()
		doCreateNormalTest(t, "-")
	})
	t.Run("should create multiplier", func(t *testing.T) {
		t.Parallel()
		doCreateNormalTest(t, "*")
	})
	t.Run("should create divider", func(t *testing.T) {
		t.Parallel()
		doCreateNormalTest(t, "/")
	})
	t.Run("should create modulo", func(t *testing.T) {
		t.Parallel()
		doCreateNormalTest(t, "%")
	})
	t.Run("should not create operator", func(t *testing.T) {
		t.Parallel()
		_, err := CreateOperator("$")
		if err == nil {
			t.Error("we cannot get the error from incorrect arg")
		}
	})
}

func TestOperate(t *testing.T) {
	t.Run("should operate as adder", func(t *testing.T) {
		t.Parallel()
		doOperateNormalTest(t, "+", 1, 2, 3)
	})
	t.Run("should operate as subtractor", func(t *testing.T) {
		t.Parallel()
		doOperateNormalTest(t, "-", 2, 4, -2)
	})
	t.Run("should operate as multiplier", func(t *testing.T) {
		t.Parallel()
		doOperateNormalTest(t, "*", 5, 2, 10)
	})
	t.Run("should operate as divider", func(t *testing.T) {
		t.Parallel()
		doOperateNormalTest(t, "/", 9, 3, 3)
	})
	t.Run("should operate as modulo", func(t *testing.T) {
		t.Parallel()
		doOperateNormalTest(t, "%", 7, 3, 1)
	})
}

func doCreateNormalTest(t *testing.T, str string) {
	t.Helper()
	operator, err := CreateOperator(str)
	if err != nil {
		t.Error(err)
	}

	var isCorrect = false
	switch str {
	case "+":
		_, isCorrect = operator.(Operator)
	case "-":
		_, isCorrect = operator.(Operator)
	case "*":
		_, isCorrect = operator.(Operator)
	case "/":
		_, isCorrect = operator.(Operator)
	case "%":
		_, isCorrect = operator.(Operator)
	}
	if isCorrect != true {
		t.Error("we cannot get the operator impl from \"" + str + "\"")
	}
}

func doOperateNormalTest(t *testing.T, str string, n1, n2, answer int) {
	t.Helper()
	operator, err := CreateOperator(str)
	if err != nil {
		t.Error(err)
	}

	realAnswer := operator.Operate(n1, n2)
	if realAnswer != answer {
		t.Error(err)
	}
}
