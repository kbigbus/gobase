package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	fmt.Println(items)
	fmt.Println(sum(items[:]...))
	for i := range items {
		items[i] *= 2
	}
	fmt.Println(items)
	fmt.Println(sum(items[:]...))

}

func sum(param ...int) (ret int) {
	for _, item := range param {
		ret += item
	}
	return
}
