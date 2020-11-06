package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Kolbasen/lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func getInput(inputString *string) (io.Reader, error) {
	if *inputFile != "" {
		input, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		return input, nil
	}
	return strings.NewReader(*inputExpression), nil
}

func getOutput(outputString *string) (io.Writer, error) {
	if *outputFile != "" {
		output, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		return output, nil
	}
	return os.Stdout, nil
}

func main() {
	flag.Parse()

	input, err := getInput(inputFile)

	if err != nil {
		return
	}

	output, err := getOutput(outputFile)

	if err != nil {
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err = handler.Compute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
