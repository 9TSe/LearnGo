package main

import (
	"fmt"
)

func a() {
	defer func() {
		/*捕获函数a内部的panic*/
		r := recover()
		fmt.Println("panic recover", r)
	}()
	panic(1)
}

func Recover() {
	defer func() {
		/*因为函数a的panic已经被函数a内部的recover捕获了
		所以main里的recover捕获不到异常，r的值是nil*/
		r := recover()
		fmt.Println("main recover", r)
	}()
	a()
	fmt.Println("main")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") //because of panic break, not run
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func complicatedRecover() {
	f()
	fmt.Println("Returned normally from f.")
}

// Foo prints and returns n.
func Foo(n int) int {
	fmt.Println(n)
	return n
}

func Fallthrough() {
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4): //not judgment
		fmt.Println("Second case")
	case Foo(2):
		fmt.Println("Third case")
	}
}

func SelectTrait() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1: // random run
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		}
	}
}
