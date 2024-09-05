package main

import (
	"fmt"
	"runtime"
)

func case_struct1() {
	// 基本语法
	x := 4
	if x > 10 {
		fmt.Println("x > 10")
	} else if x == 10 {
		fmt.Println("x == 10")
	} else {
		fmt.Println("x < 10")
	}

	// 在if条件前面可以加一条代码语句
	var b, c int = 1, 2
	if a := b + c; a < 42 {
		fmt.Println("a < 42")
	} else {
		fmt.Println("a >= 42")
	}

	// 在if里做类型判断
	var val interface{} = "foo"
	if str, ok := val.(string); ok { //val return 2 values, which is string and bool
		fmt.Println(str)
	}
}

func case_struct2() {
	operatingSystem := runtime.GOOS
	switch operatingSystem {
	case "darwin":
		fmt.Println("mac os")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("other os")
	}

	switch os := runtime.GOOS; os {
	case "darwin", "linux":
		fmt.Println("not window")
	default:
		fmt.Println("other os")
	}

	switch {
	case operatingSystem == "darwin":
		fmt.Println("mac os")
	default:
		fmt.Println("other os")
	}
}

func loop() {
	// Go只有for，没有while和until关键字
	i := 0
	for i := 1; i < 10; i++ { //这里的i是局部变量
		fmt.Println(i)
	}
	for i < 10 { // 相当while循环的效果
		fmt.Println(i) // begin with 0
		i++
	}
	// for i < 10  { // 如果只有一个条件，可以省略分号，也相当于while循环
	// }
	// for { // 可以忽略条件，相当于while (true)
	// }

	// continue here表示外层的for循环继续执行，继续执行时外层for循环里的i会++
	// break there表示退出外层循环，也就是退出整个循环了
here:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if i == 0 {
				continue here
			}
			fmt.Println(j)
			if j == 2 {
				break
			}
		}
	}

there:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if j == 1 {
				continue
			}
			fmt.Println(j)
			if j == 2 {
				break there
			}
		}
	}
}
