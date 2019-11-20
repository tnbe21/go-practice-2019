package operator

import (
	"fmt"
)

type Operator interface {
	Operate(a, b int) int
}

func CreateOperator(str string) (Operator, error) {
	switch str {
	case "+":
		return &adder{}, nil
	case "-":
		return &subtractor{}, nil
	case "*":
		return &multiplier{}, nil
	case "/":
		return &divider{}, nil
	case "%":
		return &modulo{}, nil
	default:
		return nil, fmt.Errorf("usage: go run main.go n1 n2 operator(\"+\", \"-\", \"*\", \"/\", \"%%\")")
	}
}
