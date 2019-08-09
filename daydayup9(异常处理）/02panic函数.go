package main

import "fmt"

func test02(i int)  {
	var arr [3]int
	arr[i] = 999
	//panic("runtime error: index out of range")
	fmt.Println(arr)
}

func main() {

	test02(3)


	//fmt.Println("hello")
	//fmt.Println("hello")
	//fmt.Println("hello")
	//当程序遇到panic是自动终止
	//panic("hello")
	//fmt.Println("hello")
	//fmt.Println("hello")
}
