package main

import (
	"fmt"
	"time"
)

func Finalen() func() int {
	var x, y = 0, 1
	return func() int {
		var temp = x
		x, y = y, x+y
		return temp
	}
}

func main() {
	start := time.Now()
	var f = Finalen()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

//  temp  x    y
//  0    1    1

//   1   1    2

//   1   2    3

//   2   3    5

//   3   5    8

//
