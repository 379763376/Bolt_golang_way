package main

import "fmt"
//1. 一个包里不能有重复名称的函数
//func run()  {
//	fmt.Println("run")
//}

func func1(args ...int)  {
	//不能将不定参的名称传给另一个不定参，如果要传值需要获取指定个数的数据
	//func2(args) 错误写法
	func2(args[0],args[1],args[2])
	func2(args[0:len(args)]...)//[:3] [2:3]
	func2(args[:]...)
	func2(args...)
}
func func2(args ...int)  {
	fmt.Println(args)
}
func main() {
	//2.不定长参数的调用
	func1(1,2,3,4,5,6)
	//3.函数的嵌套使用 在一个函数中调用另一个函数
	//4.函数嵌套的执行过程
	t1(10,20)
	//5. 函数返回值
	var result = func3()
	var result2  =  func4(1,5)
	fmt.Println(result,result2)
	//6. 函数返回多个值
	a,b,c := func5(1,5)
	fmt.Println(a,b,c)
	//6.匿名变量  丢弃接收的数据
	_,_,sum := func5(1,5)
	fmt.Println(sum)
}
func t1(x1,x2 int){
	t2(x1,x2)
}
func t2(x1,x2 int)  {
	sum := x1+x2
	fmt.Println(sum)
}
func func3() int {  //返回整型数据  通过return返回数据
	var a int = 5
	var b int =1
	var sum = a+ b
	return sum
}
func func4(a,b int) (sum int) {  //返回整型数据  通过return返回数据
	sum = a+ b
	return //sum
}
func func5(a,b int) (a1,b1,sum int) {  //返回多个值
	sum = a+ b
	a1=a
	b1=b
	return //sum
}