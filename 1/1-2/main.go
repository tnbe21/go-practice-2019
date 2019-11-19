package main

import (
	"flag"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() != 3 {
		log.Fatalln("usage: go run main.go n1 n2 operator(\"+\", \"-\", \"*\", \"/\", \"%\")")
	}

	n1, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	n2, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalln(err)
	}

	operator := flag.Arg(2)
	switch operator {
	case "+":
		println(n1 + n2)
	case "-":
		println(n1 - n2)
	case "*":
		println(n1 * n2)
	case "/":
		println(n1 / n2)
	case "%":
		println(n1 % n2)
	default:
		log.Fatalln("usage: go run main.go n1 n2 operator(\"+\", \"-\", \"*\", \"/\", \"%\")")
	}
}
