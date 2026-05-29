package other

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	vals []int
	len  int
}

func (s *Stack) push(val int) {
	s.len++
	s.vals = append(s.vals, val)
}

func (s *Stack) pop() int {
	// check if it exists....
	curr := s.vals[s.len-1]
	s.vals = s.vals[:s.len-1]
	s.len--
	return curr
}

func evaluateRPN(expression string) (int, error) {
	chars := strings.Fields(expression)
	var stack Stack

	for _, val := range chars {
		if char, err := strconv.Atoi(val); err == nil {
			// we received and number
			stack.push(char)
		} else {
			if stack.len < 2 {
				return 0, fmt.Errorf("not enough numbers to preform operation")
			}
			// we received an arithmetic operator
			// can we have more than two numbers in operation?
			first := stack.pop()
			second := stack.pop()

			switch val {
			case "+":
				stack.push(first + second)
			case "-":
				stack.push(first - second)
			case "*":
				stack.push(first * second)
			case "/":
				stack.push(first / second)
			default:
				return 0, errors.New("unknown arithmetic operator")
			}
		}
	}
	if stack.len != 1 {
		return 0, errors.New("something is not OK")
	}

	return stack.pop(), nil
}
