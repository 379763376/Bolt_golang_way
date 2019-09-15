package main

//1.定义别名
type Queue []int

func (q *Queue) Push(v int)  {
	*q = append(*q,v)
}
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//2.使用组合
/*
元素类型为要扩展的类型指针
如果直接传数据需要把数据拷贝，所以直接传地址

定义自己的结构体方法
对象 != nil
对象.属性 != null
 */
//type myTreeNode struct {
//	node *BoltTree
//}
//
//func main() {
//	var boltTree BoltTree
//	boltTree = BoltTree{Value:2}
//	myTree := myTreeNode{&boltTree}
//	fmt.Println(myTree)
//}
