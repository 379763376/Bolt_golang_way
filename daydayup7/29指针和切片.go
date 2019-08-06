package main

import "fmt"

func main2301() {

	s := []int{1,2,3,4,5}
	//指针和切片建立联系
	p := &s
	fmt.Printf("%p\n",p)
	fmt.Printf("%p\n",s)
	//*[]int
	//var p *[]int
	fmt.Printf("%T\n",p)


	//通过指针间接操作切片元素
	//p[1] = 222
	(*p)[1] = 222
	fmt.Println(s)
	//for循环遍历
	for i:=0;i<len(s) ; i++ {
		fmt.Println((*p)[i])
	}

}
//切片指针作为函数参数
func test23(s *[]int)  {
	*s = append(*s,6,7,8,9)
}


func main()  {
	s := []int{1,2,3,4,5}
	//地址传递
	test23(&s)
	fmt.Println(s)
}