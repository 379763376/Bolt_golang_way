package main

import "fmt"

type person10 struct {
	id   int
	name string
	age  int
}

func (p person10) PrintInfo1() {
	fmt.Printf("info1%p, %v\n", &p, p)
}

//建议使用
func (p *person10) PrintInfo2() {
	fmt.Printf("info2%p, %v\n", p, *p)
}

func main() {

	p := person10{1, "make", 22}
	//p.PrintInfo1()  //0xc042002400, {1 make 22}
	//p.PrintInfo2()  //0xc0420023e0, {1 make 22}

	//fmt.Println(p.PrintInfo1)
	//fmt.Println(p.PrintInfo2)
	//
	//fmt.Printf("%T\n",p.PrintInfo1)
	//fmt.Printf("%T\n",p.PrintInfo2)

	//方法值   隐式传递  隐藏的是接受者  绑定实例（对象）
	//var pfunc1 func()
	//对象相同 但是函数类型不同 不能赋值
	//函数类型相同   可以赋值
	pfunc1 := p.PrintInfo1
	pfunc1 = p.PrintInfo2
	pfunc1() // == p.PrintInfo1
	fmt.Printf("%T\n", pfunc1)

	//方法表达式  显式传参

	//pfunc1 := person10.PrintInfo1
	//pfunc1(p)
	//pfunc2 := (*person10).PrintInfo2
	//pfunc2(&p)
	//
	//
	//fmt.Printf("%T\n",pfunc1)
	//fmt.Printf("%T\n",pfunc2)
}
