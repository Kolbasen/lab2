package lab2

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {

	fileScanner := bufio.NewScanner(ch.Input)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		out, err := EvaluatePrefix(text)
		if err != nil {
			return err
		}
		ch.Output.Write([]byte(out))
	}

	return nil
}
