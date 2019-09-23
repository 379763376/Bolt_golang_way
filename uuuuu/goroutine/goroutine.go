package main
/*
并发编程
原生支持
关键字go 建了go就会并发的执行
10个并发协程（Coroutine） 不断的打印
1.go func中的i直接使用for中的i不安全，所以需要传进去
2.main和go func是并发执行的，当main执行完后会把其他的协程杀掉了
3.使用sleep让main等待再退出
4.一般个操作系统开一百个线程没问题，但是上千个线程并发执行 就需要异步io
go 语言中就只需要加个go就可以并发执行


coroutine是轻量级“线程” 线程在任何时候都有可能被系统切换 是抢占式任务处理，系统处理完其他的还会回来继续执行完没完成的线程  需要存更多的上下文关系
非抢占式多任务处理       由协程主动交出控制权 不同于抢占式，要保存各线程的状态 只用保存资源切换的几个点
编译器/解释器/虚拟机层面的多任务    编译器级别的多任务 go有调度器调度协程 操作系统也有调度器调度协程
多个协程可能在一个或多个线程上运行  由调度器决定的


抢占式 vs 非抢占式
io操作有等待的过程，会切换 交出控制权
创建了一个10长度的数组，每个数组创建一个协程 分别对数据增加
a[i]++ 是不会主动交出控制权，所以一直在for循环 main函数也无法获取控制权
也可以手动交出runtime.GoSched()  让别人有机会执行  每个人获得的机会是不一样的  一般不这么写
 */
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	//time.Millisecond一毫秒
	//time.Minute
	time.Sleep(time.Millisecond)
}

func main2()  {
	var a[10] int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
/*
检测数据访问冲突
race condition
以下直接引用for中的i不安全 因为外层循环会改变i的值，和里面for循环读取冲突，所以需要把i传进去
go run 1.goroutine.go
go run -race goruntine.go 检测数据访问冲突
如果匿名函数使用自由变量i,go func读了i 但是main重写了i
fmt.println(a)读的时候 协程正在写
*/
/*
子程序是协程的一个特例
协程和普通函数区别
普通函数 在一个线程内 一个函数执行完把再把控制权交还main函数，再去调用下一个语句
协调 在一个或多个线程内 main和dowork之间数据和控制权可以互相通信，交换控制权

其他语言对协程的支持
c++ boost.coroutine
python 3.5 asnyc def对协程原生支持，之前是使用yield关键字实现协程

go语言的协程：
一个程序开启后会有一个调度器，负责调度协程，有些协程放到一个线程，有些可能两个或多个放到一个线程里面
调度器控制是否放一个线程里

goroutine的定义
任何函数只需要加上go就能送给调度器运行
不需要在定义时区分是否时异步函数（相当于python asnyc def 来讲的）
调度器在合适的点进行切换【切换的点并不能完全的控制】
使用-race检测数据访问的冲突

goroutine可能的切换点[只是参考 不能保证切换 不能保证在其他地方不切换]
i/o ,select channel进行树的遍历
channel
等待锁
函数调用（调度器决定是否切换）
runtime.Gosched()

查看开启一千个协程 启动了多少线程
查看实际系统上只是最多4个线程  因为4个核

*/