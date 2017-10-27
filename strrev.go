package main

import (
	"fmt"
)

func main() {
	str := "google"
	fmt.Println(reverseString(str))
	fmt.Println(strrev(str))
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//只支持英文
func strrev(str string) string {
	clen := len(str)
	rc := make([]byte, clen) //英文转化
	//rc := make([]int32, clen) //中文转化
	// for i := clen - 1; i >= 0; i-- {
	// 	rc[i] = str[clen-i-1]
	// }

	for i := range str {
		rc[i] = str[clen-i-1]
	}
	return string(rc)
}
