package main

import (
	"operator/error"
	"operator/operator"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const USAGE_MSG = "usage: go run main.go n1 n2 operator(\"+\", \"-\", \"*\", \"/\", \"%%\")"

func main() {
	flag.Parse()
	if flag.NArg() != 3 {
		myerror.PrintError(fmt.Errorf(USAGE_MSG))
		os.Exit(1)
	}

	n1, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		myerror.PrintError(err)
		os.Exit(1)
	}

	n2, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		myerror.PrintError(err)
		os.Exit(1)
	}

	operator, err := operator.CreateOperator(flag.Arg(2))
	if err != nil {
		myerror.PrintError(err)
		os.Exit(1)
	}

	fmt.Println(operator.Operate(n1, n2))
}
