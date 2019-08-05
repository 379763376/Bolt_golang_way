package main

import "fmt"

//结构体嵌套结构体
//父类
type Person struct {
	id   int
	name string
	age  int
}
//子类
type Student struct {
	Person      //匿名字段

	score int
}


func main0101() {
	//对象创建
	//顺序初始化
	var s1 Student = Student{Person{101,"小明",18},98}
	fmt.Println(s1)

	//var p1 Person = Person{102,"小亮",28}
	//fmt.Println(p1)
	//自动推导类型
	s2 := Student{Person{102,"小亮",28},87}
	fmt.Println(s2)

	//指定初始化成员 没有初始化的部分 使用默认值
	s3 := Student{score:97}
	fmt.Println(s3)

	s4 := Student{Person{name:"小红"},100}
	fmt.Println(s4)


}

func main()  {
	var s1 Student = Student{Person{101,"小明",18},98}
	s1.score = 89
	s1.Person.age = 20
	s1.age = 30

	s1.Person = Person{110,"小兰",20}
	fmt.Println(s1)
}