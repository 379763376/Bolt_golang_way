package main

import "fmt"

func main() {
	//1.函数  定义：将代码重用的机制
	//2 函数的基本语法： func 函数名（参数列表） 返回值 {...}
	// 函数不能重名
	//3 函数的使用 run()
	fmt.Println(len("hello"))
	run()
	//4 普通参数
	test("aa") //实参  编译器的参数名提示功能
	test2(1,2)
	test3(1,"1")
}
func run()  {
	fmt.Println("run")
}
func test(a string)  { //形参
	fmt.Println(a)
}
func test2(a1,b1 int)  {
	fmt.Println(a1+b1)
}
func test3(a1 int,b1 string)  {
	fmt.Println(a1,b1)
}