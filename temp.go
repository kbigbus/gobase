package main

//interface  接口

import (
	"fmt"
)

//定义基础接口
type Base interface {
	SetName()
	GetName() string
}

type Person struct {
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
	return
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	//var p Person
	p := Person{"bigbus"}
	fmt.Println()
	// switch v:=p.(type){
	// case "Base":
	// 	fmt.Println("person")
	// default:
	// 	fmt.Println("unknown")
	// }
}
