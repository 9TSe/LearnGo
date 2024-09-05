package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func error_closure() { //多个闭包使用一个变量, 导致实际与预期不符, 非顺序打印(1,2,3,4,5...)
	var wg sync.WaitGroup
	/* wg跟踪10个goroutine */
	size := 10
	wg.Add(size)
	/* 开启10个goroutine并发执行 */
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	/* Wait会一直阻塞，直到wg的计数器为0*/
	wg.Wait()
	fmt.Println("end")
}

func mendOne() {
	var wg sync.WaitGroup
	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i) // 传递i的初始值
	}
	wg.Wait()
	fmt.Println("end")
}

func mendTwo() {
	var wg sync.WaitGroup
	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		/*定义一个新的变量*/
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	/* Wait会一直阻塞，直到wg的计数器为0*/
	wg.Wait()
	fmt.Println("end")
}

type Cat struct {
	name string
	age  int
}

func fetchChannel(ch chan Cat) {
	value := <-ch //wait ch <- a
	fmt.Printf("type: %T, value: %v\n", value, value)
}

func channel_goroutine1() {
	ch := make(chan Cat) //无缓冲区通道
	a := Cat{"yingduan", 1}
	// 启动一个goroutine，用于从ch这个通道里获取数据
	go fetchChannel(ch)
	// 往cha这个通道里发送数据
	ch <- a
	// main这个goroutine在这里等待2秒
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}

func addData(ch chan int) {
	/*
		每3秒往通道ch里发送一次数据
	*/
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(3 * time.Second)
	}
	// 数据发送完毕，关闭通道
	close(ch)
}

func channel_goroutine2() {
	ch := make(chan int, 10)
	// 开启一个goroutine，用于往通道ch里发送数据
	go addData(ch)

	/*
		for循环取完channel里的值后，因为通道close了，再次获取会拿到对应数据类型的零值
		如果通道不close，for循环取完数据后就会阻塞报错
	*/
	for {
		value, ok := <-ch
		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("finish")
			break
		}
	}
	// 0 1 2 3 4 5 6 7 8 9 finish
}

func write(ch chan<- int) {
	/*
		参数ch是只写channel，不能从channel读数据，否则编译报错
		receive from send-only type chan<- int
	*/
	ch <- 10
}

func read(ch <-chan int) {
	/*
		参数ch是只读channel，不能往channel里写数据，否则编译报错
		send to receive-only type <-chan int
	*/
	fmt.Println(<-ch)
}

func channelReadWriteOnly() {
	ch := make(chan int)
	go write(ch)
	go read(ch)

	// 等待3秒，保证write和read这2个goroutine都可以执行完成
	time.Sleep(3 * time.Second)

	//close后，如果channel还有值，接收方可以一直从channel里获取值，直到channel里的值都已经取完。
	//close后，如果channel里没有值了，接收方继续从channel里取值，会得到channel里存的数据类型对应的默认零值
}

// sync.WaitGroup 作为函数参数时需要传递指针
