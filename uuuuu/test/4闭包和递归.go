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


}