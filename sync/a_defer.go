package main

import "fmt"

func defer_test1() {
	i := 0
	defer fmt.Println(i) // 最终打印0
	i++
}

func defer_test2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 最终打印3210
	}
}

// f returns 42
func defer_test3() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		fmt.Println(result) // 6
		result *= 7
	}()
	return 6
}
