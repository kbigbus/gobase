package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack [LIMIT]string

func (s *Stack) Push(x string) {
	for i, v := range s {
		if v == "" {
			s[i] = x
			break
		}
	}
}

func (s *Stack) Pop() string {
	v := len(s)
	x := ""
	for i := v - 1; i >= 0; i-- {
		if s[i] != "" {
			x = s[i]
			break
		}
	}
	return x
}

func (s Stack) String() string {
	str := ""
	for i, v := range s {
		str += "[" + strconv.Itoa(i) + ":" + v + "]"
	}
	return str + "\n"
}

func main() {
	s := Stack{"i", "j"}
	fmt.Print(s)
	s.Push("x")
	fmt.Print(s)
	fmt.Print(s.Pop())
}
