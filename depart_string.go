package main

import (
	"fmt"
)

func main() {
	str := "this is a string"
	ox, oy := returnStr(str, 6)
	fmt.Println(ox)
	fmt.Println(oy)

	fmt.Println(string(str[3]))

	fmt.Println(str[len(str)/2:] + str[:len(str)/2])
}

func returnStr(str string, i int) (x string, y string) {
	if i >= len(str) {
		x, y = "out of the length", "error"
		return
	}
	//c := []byte(str)
	x, y = string(str[0:i]), string(str[i:])
	return

}
