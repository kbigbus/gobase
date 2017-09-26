package main

//冒泡排序

import (
	"fmt"
)

func main() {
	arr := [...]int{1, 3, 7, 4, 5, 8, 9, 6}
	arrlen := len(arr)
	tmp := 0
	for i := 0; i < arrlen; i++ {
		for j := i + 1; j < arrlen; j++ {
			if arr[i] < arr[j] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	fmt.Println(arr)

}
