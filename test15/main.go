package main

import "fmt"

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name string
	age int
	gender string
	position string
	company
}

func main()  {
	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company: myCom,
	}

	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}