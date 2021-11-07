package main

import (
	"flag"
	lab2 "github.com/nickname038/architecture-2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file with expression")
	outputFile      = flag.String("o", "", "Output file with result")
)

func writeError(err error) {
	_, _ = os.Stderr.WriteString(err.Error())
}

func main() {
	flag.Parse()

	var reader io.Reader = strings.NewReader(*inputExpression)
	var writer io.Writer = os.Stdout

	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			writeError(err)
			return
		}
		reader = file
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			writeError(err)
			return
		}
		writer = file
	}

	handler := lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		writeError(err)
	}
}
