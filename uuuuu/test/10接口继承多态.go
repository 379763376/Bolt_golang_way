package main

//使用者定义接口
type Retriver interface {
	Get(source string) string
}
//使用接口
func download(retriver Retriver)  {
	return retriver.Get("http://www.baidu.com")
}