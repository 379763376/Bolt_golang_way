package main

import "fmt"

func main() {
	sum := 0
	//1.匿名函数 这个函数使用的sum是返回值定义的变量 而不是外部函数变量
	x,y,z := func(a,b int)(i,j,sum int){
		i= a
		j =b
		sum = a+b
		return
	}(1,3)
	fmt.Println(x,y,z,sum)
	//2.匿名函数 这个函数使用的sum是返回值定义的变量 而不是外部函数变量
	func(a,b int){
		sum = a+b
		return
	}(1,3)
	fmt.Println(sum)
	//3匿名函数定义方式
	//3.1
	f := func() int{
		sum ++
		return 1
	}
	f()
	//3.2
	fmt.Println(func(){
		sum++
	})
	//3.3
	type anonymityFuncType func() int
	var anonymityFunc anonymityFuncType
	anonymityFunc = f
	a := anonymityFunc()
	fmt.Println(a)

	//4.递归
	b := add100(100)
	fmt.Println(b)
}
func add100(num int)  int{
	if num == 1{
		return 1
	}
	return num+add100(num-1)
}
