package main

import (
	"convertor/convertor"
	"convertor/error"
	"flag"
	"fmt"
	"os"
)

const USAGE_MSG = "usage: go run main.go pngPath (\"jpg\", \"gif\")"

func main() {
	src := flag.String("src", "", "src png filePath")
	dst := flag.String("dst", "", "converted file type(extension): jpg or gif")
	flag.Parse()
	if *src == "" || *dst == "" {
		myerror.PrintError(fmt.Errorf(USAGE_MSG))
		os.Exit(1)
	}

	if err := convertor.Convert(*src, *dst); err != nil {
		myerror.PrintError(err)
		os.Exit(1)
	}
}
