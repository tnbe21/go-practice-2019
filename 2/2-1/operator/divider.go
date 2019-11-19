package operator

type divider struct{}

func (*divider) Operate(a, b int) int {
	return a / b
}
