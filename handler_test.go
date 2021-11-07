package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestWriteResultToOutput(c *C) {
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  strings.NewReader("4 1 +"),
		Output: output,
	}
	_ = handler.Compute()
	c.Assert(output.String(), Equals, "+ 4 1")
}

func (s *HandlerSuite) TestReturnError(c *C) {
	handler := ComputeHandler{
		Input:  strings.NewReader("asdasd"),
		Output: bytes.NewBufferString(""),
	}
	err := handler.Compute()
	c.Assert(err.Error(), Equals, "invalid expression")
}
