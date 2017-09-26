package main

//map 键值对

import (
	"fmt"
)

func main() {
	m1 := make(map[int]string)
	m1[2] = "OK"
	m1[1] = "GOOD"

	leng := len(m1)
	fmt.Println(leng)
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println(m1)

	m2 := map[int]string{1: "on", 2: "off"}
	if v, ok := m2[2]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("get value failed")
	}
	fmt.Println(m2[1])
}
