package operator

type multiplier struct{}

func (*multiplier) Operate(a, b int) int {
	return a * b
}
