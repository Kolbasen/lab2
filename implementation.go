package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/apaxa-go/eval"
)

func isOperand(symbol string) bool {
	return symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/"
}

func getExpression(arr []string, operandIndex int) string {
	operand := arr[operandIndex]
	val1 := arr[operandIndex+1]
	val2 := arr[operandIndex+2]

	return "float64(" + val1 + operand + val2 + ")"
}

func removeElements(arr []string, startIndex int, count int) []string {
	copy(arr[startIndex:], arr[startIndex+count:])
	arr = arr[:len(arr)-count]
	return arr
}

func stringToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// EvaluatePrefix performs a mathematical calculation of input string in prefix form.
// It returns the result in string format
func EvaluatePrefix(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("Expression is empty string")
	}

	s := strings.Split(input, " ")

	for {
		for i := len(s) - 1; i >= 0; i-- {
			if isOperand(s[i]) {
				stringExpression := getExpression(s, i)
				expr, err := eval.ParseString(stringExpression, "")
				if err != nil {
					return "", fmt.Errorf("Cannot parse expression")
				}

				c, err := expr.EvalToInterface(nil)
				if err != nil {
					return "", fmt.Errorf("Cannot evaluate expression")
				}

				fl := c.(float64)
				s[i] = fmt.Sprint(fl)
				s = removeElements(s, i+1, 2)
				break
			} else if s[i] == "^" {
				val1, err := stringToFloat(s[i+1])
				val2, err := stringToFloat(s[i+2])
				if err != nil {
					return "", fmt.Errorf("Cannot parse expression")
				}

				res := math.Pow(val1, val2)
				s[i] = fmt.Sprint(res)
				s = removeElements(s, i+1, 2)
				break
			} else {
				_, err := stringToFloat(s[i])
				if err != nil {
					return "", fmt.Errorf("Expression contains invalid characters")
				}
			}
		}
		if len(s) == 1 {
			result := s[0]
			return result, nil
		}
	}
}
