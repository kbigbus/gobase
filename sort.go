package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []int{1, 3, 4, 2, 8, 5}

	//sort.Ints(str)
	//SearchInts 根据二分法查找 所以查询结果是排序后的索引值
	fmt.Println(sort.SearchInts(str, 2))
}
