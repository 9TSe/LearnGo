package main

import (
	"fmt"
	//. "math" 省略math.前缀
	//_ "math"  自动调用init函数
	ma "math" //用别名访问
	"sync"
	"sync/atomic"
	"time"
)

func math_test() {
	fmt.Println(ma.Sqrt(32)) // .
}

func Swap() {
	var newValue int32 = 200
	var oldValue int32 = 100

	old := atomic.SwapInt32(&oldValue, newValue)
	fmt.Println("old value:", old, "new value:", oldValue) // 100
}

func CompareAndSwap() {
	var dst int32 = 100
	oldValue := atomic.LoadInt32(&dst)
	var newValue int32 = 200
	swapped := atomic.CompareAndSwapInt32(&dst, oldValue, newValue)
	fmt.Println("swapped:", swapped, "dst:", dst) // true 200
	fmt.Printf("%v\n", swapped)                   // true
}

var wg sync.WaitGroup

func Add() {
	var sum int32 = 0
	size := 100
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(i int32) {
			defer wg.Done()
			atomic.AddInt32(&sum, i) //not sum += 1 导致结果紊乱
		}(int32(i))
	}
	wg.Wait()
	fmt.Println("sum:", sum) // 4950
}

func Store() {
	var oldValue int32 = 100
	var num int32 = 50
	atomic.StoreInt32(&oldValue, num)
	result := atomic.LoadInt32(&oldValue)
	fmt.Println("result:", result) // 50
}

func loadConfig() map[string]string {
	// 从数据库或者文件系统中读取配置信息，然后以map的形式存放在内存里
	return make(map[string]string)
}

func requests() chan int {
	// 将从外界中接收到的请求放入到channel里
	return make(chan int)
}

func Value() {
	// config变量用来存放该服务的配置信息
	var config atomic.Value
	// 初始化时从别的地方加载配置文件，并存到config变量里
	config.Store(loadConfig())
	go func() {
		// 每10秒钟定时拉取最新的配置信息，并且更新到config变量里
		for {
			time.Sleep(10 * time.Second)
			// 对应于赋值操作 config = loadConfig()
			config.Store(loadConfig())
		}
	}()
	// 创建协程，每个工作协程都会根据它所读取到的最新的配置信息来处理请求
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				// 对应于取值操作 c := config
				// 由于Load()返回的是一个interface{}类型，所以我们要先强制转换一下
				c := config.Load().(map[string]string)
				// 这里是根据配置信息处理请求的逻辑...
				_, _ = r, c
			}
		}()
	}
}
