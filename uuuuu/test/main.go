package main

import (
	"Bolt_golang_way/uuuuu/test/mock"
	"Bolt_golang_way/uuuuu/test/real"
	"fmt"
	"time"
)

type Retriver interface {
	Get(url string) string
}//使用者定义了一个接口Get
type Poster interface {
	Post(url string,
		form map[string]string) string
}//接口二定义 Post

func download(r Retriver) string{
	return r.Get("http://www.baidu.com")
} //使用者使用接口
func post(poster Poster)  string {
	return poster.Post("http://www.baidu.com",
		map[string]string{
			"name":"bolt",
			"age":"23",
		})
}//接口二使用

//接口一 接口二 组合
type RetriverPoster interface {
	Retriver
	Poster
}
func session(s RetriverPoster) string {
	//s.Get()
	s.Post("http://www.baidu.com", map[string]string{
		"contents":"接口调用时修改初始实现时的值",
	})
	return s.Get("http://www.baidu.com")
}
func main() {
	var r Retriver //定义了一个接口对象，要当参数传递需要实现
	r = mock.Retriver1{"this is my interface test1"} //实现者实现了接口
	fmt.Println(download(r))//使用者使用
	fmt.Println(download(mock.Retriver1{"this is two"}))
	r = &mock.Retriver1{"this is my interface test1"} //接收者自带指针


	r = &real.Retriver2{} //实现者实现了接口,是一个指针接收者
	//fmt.Println(download(r)) //使用者使用
	r = &real.Retriver2{
		"Mozilla/5。0",
		time.Minute,
	}
	inspect(r)//判断类型，方法一：switch case
	//判断类型，方法二：type assertion r.(类型)取得类型
	realRetriver := r.(*real.Retriver2)
	fmt.Println(realRetriver.UserAgent)
	//panic: interface conversion: main.Retriver is *real.Retriver2, not mock.Retriver1
	//r.(类型)  是把接口变量肚子里的值强制类型转换，转换成功就是ok,否则err
	if mockRetriver,ok := r.(mock.Retriver1);ok{
		fmt.Println(mockRetriver.Contents)
	}else{
		fmt.Println("not a mock retriver")
	}

	//不能直接使用r 因为r只有接口的能力，只有get,没有post
	retriver := mock.Retriver1{
		"接口实现中定义的content"} //结构体创建了 具有了get post方法
	fmt.Println("try a session")
	//fmt.Println(session(retriver))
	//执行结构体对象的post(是值传递，方法内部修改了拷贝的结构体，不会改变结构体对象本身)
	//执行get打印输出
	fmt.Println(session(&retriver))
}

func inspect(r Retriver) {
	fmt.Println("%T %v\n",r,r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case mock.Retriver1:
		fmt.Println("contents:", v.Contents)
	case *real.Retriver2:
		fmt.Println("UserAgent", v.UserAgent)
	}
}
