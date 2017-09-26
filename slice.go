package main

//切片

import (
	"fmt"
)

func main() {
	//以数组为基准 做切割获取切片
	arr := [...]int{1, 3, 7, 4, 5, 8, 9, 6}
	s1 := arr[3:]
	fmt.Println(s1)

	//通过append 查看切片容量变化
	s2 := make([]int, 3, 10)
	fmt.Println(s2, len(s2), cap(s2))
	s2 = append(s2, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(s2, len(s2), cap(s2))
	s2 = append(s2, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Println(s2, len(s2), cap(s2))

}
