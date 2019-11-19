package operator

type adder struct{}

func (*adder) Operate(a, b int) int {
	return a + b
}
