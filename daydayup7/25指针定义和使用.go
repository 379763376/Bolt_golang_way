package main

import "fmt"

func main() {

	//var i int
	//i = 100
	////fmt.Printf("%d\n",i)
	////fmt.Printf("%p\n",&i)
	////fmt.Printf("%v\n",&i)
	//
	////指针变量 定义
	//
	//var p *int  //定义一个指针变量
	//p = &i      //把变量i 地址复制指针变量p
	//
	//*p = 80  //使用指针修改值
	//fmt.Printf("i=%d,p=%p\n",i,p)
	//fmt.Printf("i=%d,p=%v",i,*p)

	//注意
	//1.默认值 nil

	//内存地址编号为0  0-255 空间  系统占用  不允许用户访问
	//空指针
	//var p *int = nil
	//fmt.Println(p)


	//野指针   指针变量指向一个位未知的空间   会报错
	//程序不允许出现野指针
	//var a int
	//var p *int
	//p = &a     //没有指向  直接操作
	//*p = 56
	//fmt.Println(p,*p)


	//new函数
	//gc垃圾回收机制  自动释放空间
	var p *int
	p = new(int)
	*p = 57
	fmt.Println(*p)

	//自动推导类型
	q := new(int)
	*q = 999
	fmt.Println(*q)

}
