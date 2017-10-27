package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack struct {
	index int
	list  [LIMIT]string
}

func (s *Stack) Push(x string) {
	arr := &(s.list)
	for i, v := range arr {
		if v == "" {
			arr[i] = x
			s.index = i + 1
			break
		}
	}
}

func (s *Stack) Pop() string {
	arr := &(s.list)
	v := len(arr)
	x := ""
	for i := v - 1; i >= 0; i-- {
		if arr[i] != "" {
			x = arr[i]
			s.index = i
			break
		}
	}
	return x
}

func (s Stack) String() string {
	str := ""
	for i, v := range s.list {
		str += "[" + strconv.Itoa(i) + ":" + v + "]"
	}
	return str + "\n"
}

func main() {
	s := Stack{2, [LIMIT]string{"i", "j"}}
	fmt.Print(s)
	s.Push("x")
	fmt.Print(s)
	fmt.Print(s.Pop())
}
