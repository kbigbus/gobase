package main

import (
	"fmt"
	"time"
)

var arr1 = [50]int{0, 1}

func Fibonacc_arr(i int) int {
	if i <= 1 {
		return arr1[i]
	} else {
		arr1[i] = arr1[i-1] + arr1[i-2]
		return arr1[i]
	}
}
func main() {
	var output int
	start := time.Now()
	for i := 0; i < 50; i++ {
		output = Fibonacc_arr(i)
		fmt.Println(output)
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
