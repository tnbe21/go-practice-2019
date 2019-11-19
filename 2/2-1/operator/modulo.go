package operator

type modulo struct{}

func (*modulo) Operate(a, b int) int {
	return a % b
}
