package main

import "fmt"

// 声明一个 Profile 的结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 重点在于这个星号: *
func (person *Profile) increase_age() {
	person.age += 1
}

func main() {
	myself := Profile{name: "小明", age: 24, gender: "male"}
	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.increase_age()
	fmt.Printf("当前年龄：%d", myself.age)
}