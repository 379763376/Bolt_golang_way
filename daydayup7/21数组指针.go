package main

import "fmt"

func main2101() {
	arr := [3]int{1, 2, 3}
	//定义指针  指向数组  数组指针
	var p *[3]int
	//指针和数组建立关系
	p = &arr
	fmt.Println(*p)

	//自动推导类型创建数组指针
	//p := &arr
	//通过指针操作数组
	//(*数组指针变量)[下标] = 值
	(*p)[0] = 111

	p[1] = 222
	fmt.Println(*p)

}

func main2102() {
	arr := [5]int{1, 2, 3, 4, 5}
	//指针变量和要存储数据类型要相同
	p1 := &arr
	p2 := &arr[0]
	//p1和p2在内存中指向相同的地址
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}

func swaps(p *[5]int) {  //指针数组
	(*p)[0] = 111
}

func main() {
	a := [5]int{1,2,3,4,5}
	swaps(&a)   //传地址

	fmt.Println(a)

}
