package main

import "fmt"

func main(){
	/*myarr := [5]int{1,2,30,4,5}

	// 【第一种】
	// 1 表示从索引1开始，直到到索引为 2 (3-1)的元素
	mysli1 := myarr[1:3]

	// 【第二种】
	// 1 表示从索引1开始，直到到索引为 2 (3-1)的元素
	mysli2 := myarr[1:3:3]

	fmt.Println(mysli1)
	fmt.Println(mysli2)*/


	myarr := [5]int{10,20,30,40,50}
	/*fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))

	mysli1 := myarr[1:3]
	fmt.Printf("mysli1 的长度为：%d，容量为：%d\n", len(mysli1), cap(mysli1))
	fmt.Println(mysli1)

	mysli2 := myarr[1:3:3]
	fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	fmt.Println(mysli2)
	fmt.Printf("类型是: %T", mysli1)*/

	fmt.Printf("myarr[:] 的长度为：%d，容量为：%d\n", len(myarr[:]), cap(myarr[:]))
	fmt.Println(myarr[:])



	// 指针类型
	/*aint := 1     // 定义普通变量
	ptr := &aint  // 定义指针变量
	fmt.Println("普通变量存储的是：", aint)
	fmt.Println("普通变量存储的是：", *ptr)
	fmt.Println("指针变量存储的是：", &aint)
	fmt.Println("指针变量存储的是：", ptr)*/

	/*a := [3]int{89, 90, 91}
	fmt.Print(a[:])*/

	/*a := []int{3:2}
	b:=[]int{4,2}
	fmt.Println(a,b)*/

	/*sli := []int{1,2,3,4,5,6}
	g:=sli[0:3]
	a:=[]int{2:5}
	b:=[...]int{2:5}
	fmt.Printf("%d 的类型是: %T", g,g)
	fmt.Print("\n")
	fmt.Printf("%d 的类型是: %T", a,a)
	fmt.Print("\n")
	fmt.Printf("%d 的类型是: %T", b,b)
	//fmt.Print(g)
	//fmt.Print(a)*/
}
