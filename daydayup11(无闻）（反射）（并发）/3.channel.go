package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//n,ok := <-c 可以用ok判断是否是close
	//或者以下range
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
/*
channel
用作两个routine之间数据的交换

var c chan int //定义一个channel 传入的数据是int，但是这个channel没给做出来是c == nil
所以创建一个channel需要c:make(chan int)

chanDemo中定义一个channel然后定义一个routine接收数据 不接收会datalock异常，c<- 发数据到channel n := <-c 从channel取出数据

函数是一等公民 channel也是 可以作为函数 可以作为参数 可以作为返回值

bufferChannel 添加缓冲区

channel close 发送发close，外层函数推掉 channel也就关闭了

channel的理论基础CSP communication sequential process
 */