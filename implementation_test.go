package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToPrefixTwoOperands(c *C) {
	res, err := PostfixToPrefix("4 1 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ 4 1")
}

func (s *MySuite) TestPostfixToPrefixThreeOperands(c *C) {
	res, err := PostfixToPrefix("3 4 ^ 10 /")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "/ ^ 3 4 10")
}

func (s *MySuite) TestPostfixToPrefixFourOperands(c *C) {
	res, err := PostfixToPrefix("5 6 + 8 9 - *")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "* + 5 6 - 8 9")
}

func (s *MySuite) TestPostfixToPrefixEightOperands(c *C) {
	res, err := PostfixToPrefix("1 2 / 3 * 4 - 5 6 / 7 8 + / +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ - * / 1 2 3 4 / / 5 6 + 7 8")
}

func (s *MySuite) TestPostfixToPrefixTenOperands(c *C) {
	res, err := PostfixToPrefix("1 2 3 4 * 5 6 - / - 7 + 8 9 * 10 - / +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ 1 / + - 2 / * 3 4 - 5 6 7 - * 8 9 10")
}

func (s *MySuite) TestPostfixToPrefixEmptyString(c *C) {
	_, err := PostfixToPrefix("")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func (s *MySuite) TestPostfixToPrefixInvalidStringTwoParametrs(c *C) {
	_, err := PostfixToPrefix("1 +")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func (s *MySuite) TestPostfixToPrefixInvalidOperators(c *C) {
	_, err := PostfixToPrefix("1 5 .")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func (s *MySuite) TestPostfixToPrefixInvalidExprecion(c *C) {
	_, err := PostfixToPrefix("+ 1 / + - 2 / * 3 4 - 5 6 7 - * 8 9")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func (s *MySuite) TestPostfixToPrefixInvalidOperands(c *C) {
	_, err := PostfixToPrefix("+ 1 / + - p / * x 4 - 5 6 7 - * 8 9")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func ExamplePostfixToPrefix() {
	res1, _ := PostfixToPrefix("1 2 3 4 5 - * / -")
	fmt.Println(res1)

	// Output:
	// - 1 / 2 * 3 - 4 5
}
