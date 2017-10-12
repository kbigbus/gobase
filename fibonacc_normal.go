package main

import (
	"fmt"
	"time"
)

func Finalen(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return Finalen(x-1) + Finalen(x-2)
	}

}

func main() {
	start := time.Now()
	for i := 0; i < 40; i++ {
		fmt.Println(Finalen(i))
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
