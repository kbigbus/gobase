package main

//struct

import (
	"fmt"
)

//底层结构
type Person struct {
	Name string
}

type User struct {
	Person
	Age int
}

//值拷贝 方法
func (u User) SetName(name string) {
	u.Name = name
	return
}

//地址拷贝 方法
func (u *User) SetNamePointer(name string) {
	u.Name = name
	return
}

func main() {
	//初始化嵌套结构体
	u := User{Person{Name: "joe"}, Age: 5}
	//值拷贝设置name  无效
	u.SetName("ketty")
	fmt.Println(u)
	//地址拷贝设置name 有效
	u.SetNamePointer("john")
	fmt.Println(u)
	//设置嵌套name 有效
	u.Person.Name = "lucy"
	fmt.Println(u)
}
