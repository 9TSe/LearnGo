package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex //default is unlocked

func printHello() {
	mutex.Lock()
	fmt.Println("Hello")
	mutex.Unlock()
}

func test() {
	fmt.Println("test once") // x1
}

func print(once *sync.Once) { //must be pointer
	once.Do(test)
}

func once() {
	var wg sync.WaitGroup
	var once sync.Once
	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			once.Do(printHello)
			print(&once)
		}()
	}
	wg.Wait()
	// just do once
	fmt.Println("func once end")
}

type Singleton struct {
	member int
}

var instance *Singleton
var once2 sync.Once

func getInstance() *Singleton {
	once2.Do(func() { //only once
		fmt.Println("init instance")
		instance = &Singleton{}
		instance.member = 10
	})
	fmt.Println("once: ", instance.member)
	return instance
}

func singletonOnce() {
	var wg2 sync.WaitGroup
	wg2.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg2.Done()
			instance = getInstance()
		}()
	}
	wg2.Wait()
	fmt.Println("func singletonOnce end")
}

//  如果多个goroutine执行, 都去调用once.Do(f)，只有某次的函数f调用返回了，所有Do方法调用才会返回
//否则Do方法会一直阻塞等待。如果在f里继续调用同一个once变量的Do方法，就会死锁了，因为Do在等待f返回，f又在等待Do返回。
//  如果once.Do(f)方法调用的函数f发生了panic，那Do也会认为函数f已经return了。

type Counter struct {
	num     int
	rwMutex sync.RWMutex
}

func (c *Counter) getNum() int {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	return c.num
}

func (c *Counter) addNum() {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.num++
}

func mutexRW() {
	var wg sync.WaitGroup
	counter := &Counter{}
	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			counter.getNum()
			counter.addNum()
		}()
	}
	wg.Wait()
	fmt.Println("counterNum = ", counter.num)
}

func syncCond() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	size := 10
	wg.Add(size + 1)

	for i := 0; i < size; i++ {
		i := i //变量遮蔽
		//如果没有 i := i 这样的操作，所有的 goroutine 会共享同一个 i 变量
		//导致它们在执行时打印出的 i 可能都是循环结束后的值。
		go func() {
			defer wg.Done()
			/*调用Wait方法时，要对L加锁*/
			cond.L.Lock()
			fmt.Printf("%d ready\n", i)
			/*Wait实际上是会先解锁cond.L，再阻塞当前goroutine
			  这样其它goroutine调用上面的cond.L.Lock()才能加锁成功，才能进一步执行到Wait方法，
			  等待被Broadcast或者signal唤醒。
			  Wait被Broadcast或者Signal唤醒的时候，会再次对cond.L加锁，加锁后Wait才会return
			*/
			cond.Wait()
			fmt.Printf("%d done\n", i)
			cond.L.Unlock()
		}()
	}

	/*这里sleep 2秒，确保目标goroutine都处于Wait阻塞状态
	  如果调用Broadcast之前，目标goroutine不是处于Wait状态，会死锁
	*/
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		cond.Broadcast()
	}()
	wg.Wait()
}

func statistics_str() {
	m := sync.Map{}
	str := "abcdefg"
	for _, value := range str {
		fmt.Println(value)
		temp, ok := m.Load(value)
		if !ok {
			m.Store(value, 1)
		} else {
			m.Store(value, temp.(int)+1) // .(int) 类型断言
		}
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println("statistics_str key ,value :", key, value)
		return true
	})
}

var m1 sync.Map

func addMap(key int) {
	m1.Store(key, 1)
}

func routineWriteMap() {
	var wgx sync.WaitGroup
	size := 2
	wgx.Add(size)
	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wgx.Done()
			addMap(i)
		}()
	}
	wgx.Wait()

	m1.Range(func(key, value interface{}) bool {
		fmt.Println("toutinewritemap, key,value :", key, value)
		return true
	})
}
