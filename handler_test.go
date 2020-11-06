package lab2

import (
	"bytes"
	"fmt"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestInputFile(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("- 1 42"),
		Output: &buff,
	}
	handler.Compute()
	c.Assert(buff.String(), Equals, "-41")
}

func (s *MySuite) TestInputErr(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("- 1 4s2"),
		Output: &buff,
	}
	err := handler.Compute()
	c.Assert(fmt.Sprint(err), Equals, "Expression contains invalid characters")
}
