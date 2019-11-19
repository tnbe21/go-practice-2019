package operator

type subtractor struct{}

func (*subtractor) Operate(a, b int) int {
	return a - b
}
