package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEvaluatePrefix(c *C) {
	res, _ := EvaluatePrefix("+ 5 * - 4 2 3")
	c.Assert(res, Equals, "11")
}

func (s *MySuite) TestEvaluatePrefixSimple1(c *C) {
	res, _ := EvaluatePrefix("* + 5 7 - 6 3")
	c.Assert(res, Equals, "36")
}

func (s *MySuite) TestEvaluatePrefixSimple2(c *C) {
	res, _ := EvaluatePrefix("* - + 6 3 4 2")
	c.Assert(res, Equals, "10")
}

func (s *MySuite) TestEvaluatePrefixSimple3(c *C) {
	res, _ := EvaluatePrefix("* 2 - 1 + 6 ^ 4 2")
	c.Assert(res, Equals, "-42")
}

func (s *MySuite) TestEvaluatePrefixAdvanced1(c *C) {
	res, _ := EvaluatePrefix("- - + + 5 3 1 2 - + * 3 3 9 + 1 6")
	c.Assert(res, Equals, "-4")
}

func (s *MySuite) TestEvaluatePrefixAdvanced2(c *C) {
	res, _ := EvaluatePrefix("- - - - + + + + 1 2 3 4 5 6 7 8 9")
	c.Assert(res, Equals, "-15")
}

func (s *MySuite) TestEvaluatePrefixAdvanced3(c *C) {
	res, _ := EvaluatePrefix("- + * * 6 / 2 5 2 6 - * 12 2 / 9 2")
	c.Assert(res, Equals, "-8.7")
}

func (s *MySuite) TestEvaluatePrefixAdvanced4(c *C) {
	res, _ := EvaluatePrefix("- + * * 6 / 2 5 2 6 - * 12 2 / 9 ^ 2 2")
	c.Assert(res, Equals, "-10.95")
}

func (s *MySuite) TestEvaluatePrefixEmptyString(c *C) {
	res, err := EvaluatePrefix("")
	c.Assert(res, Equals, "")
	c.Assert(fmt.Sprint(err), Equals, "Expression is empty string")
}

func (s *MySuite) TestEvaluatePrefixInvalidExpression(c *C) {
	res, err := EvaluatePrefix("a*21")
	c.Assert(res, Equals, "")
	c.Assert(fmt.Sprint(err), Equals, "Expression contains invalid characters")
}

func (s *MySuite) TestEvaluatePrefixSpacesOnly(c *C) {
	res, err := EvaluatePrefix("      ")
	c.Assert(res, Equals, "")
	c.Assert(fmt.Sprint(err), Equals, "Expression is empty string")
}

func ExampleEvaluatePrefix() {
	res, _ := EvaluatePrefix("+ 5 * - 4 2 3")
	fmt.Println(res)

	// Output:
	// 11
}
