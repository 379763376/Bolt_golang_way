package main

import "fmt"

func main() {
	//一 If选择结构
	//1.
	a1 := 700
	var b1 int
	fmt.Printf("请用户输入分数")
	fmt.Scanf("%d", &b1)
	if b1 >= a1 {
		fmt.Println("恭喜你1")
	}
	if b1 >= a1 {
		fmt.Println("恭喜你2")
	}
	if c1 := 710; c1 > a1 { //支持一个初始化语句，初始化语句和判断条件用 ； 分割
		fmt.Println("恭喜你3")
	}
	// 2 if else
	if c2 := 650; c2 > a1 { //支持一个初始化语句，初始化语句和判断条件用 ； 分割
		fmt.Println("恭喜你4")
	} else { //没有条件
		fmt.Println("需要再努力")
	}
	// 3 if elseif else
	a2 := 650
	if c2 := 650; c2 >= a1 { //支持一个初始化语句，初始化语句和判断条件用 ； 分割
		fmt.Println("恭喜你一本")
	} else if c2 >= a2 {
		fmt.Println("恭喜你二本")
	} else {
		fmt.Println("需要再努力")
	}
	/* 4 if和if_else区别
	if 都会去执行
	if_else效率高,只要一个匹配到之后的不再执行
	*/
	/* 5 if嵌套
	if condition1 {
		if condition2 {
			//todo
		}
	}
	*/
	//6，比较的时候不能  x<y<z    z<y结果未boole <z的时候就会报错


	//二 switch
	/*  1.支持多个条件的匹配
	a.不同的case之间不需要使用break
	b. fallthrough  强制往下执行一个case
	switch 变量（或表达式） {
		case val1:
			...
		case val2:
			...
		default:
			...
	}
	*/
	score := 90
	switch score {
	case 90:
		fmt.Println('C')
	case 80:
		fmt.Println("B")
	case 10, 20, 30, 40:
		fmt.Println("D")
	default:
		fmt.Println("A")

	}
	//2
	switch s1:=90;s1 {  //初始化语句   ;    条件
	case 90:
		fmt.Println('C')
	case 80:
		fmt.Println("B")
	case 10, 20, 30, 40:
		fmt.Println("D")
	default:
		fmt.Println("A")
	}
	//3
	var s2 int = 90
	switch  { //没有条件
	case s2>90://写判断语句
		fmt.Println('A')
	}
	//4
	switch s3:=90; { //只有初始化没有条件
	case s3>90:
		fmt.Println('A')
	}
	/*if和switch的选择：
		if适合区间判断，嵌套使用
		switch执行效率高，将多个满足条件的值放一起，适合做固定值的判断

		if相对switch执行效率低，分支少用if
		switch不建议使用嵌套
	 */
}
