package main

import (
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestNormal(t *testing.T) {
	t.Run("should do main as adder", func(t *testing.T) {
		t.Parallel()
		doMain(t, 1, 2, "+")
	})
	t.Run("should do main as subtractor", func(t *testing.T) {
		t.Parallel()
		doMain(t, 1, 2, "-")
	})
	t.Run("should do main as multiplier", func(t *testing.T) {
		t.Parallel()
		doMain(t, 1, 2, "*")
	})
	t.Run("should do main as divider", func(t *testing.T) {
		t.Parallel()
		doMain(t, 1, 2, "/")
	})
	t.Run("should do main as modulo", func(t *testing.T) {
		t.Parallel()
		doMain(t, 1, 2, "%")
	})
}

func doMain(t *testing.T, n1, n2 int, operator string) {
	t.Helper()
	os.Args = []string{"[[cmd]]", strconv.Itoa(n1), strconv.Itoa(n2), operator}
	main()
}
