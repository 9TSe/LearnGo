package main

import "fmt"

func declare() {
	var a, b int = 10, 12
	const c = "not clear"
	const h = 14
	i := 15
	const const_value int = 16
	fmt.Println(a, b, c, h, i, const_value)
	const (
		d = iota
		e = iota
		f
		g = 1 << iota
	)
	fmt.Println(d, e, f, g)
}

func function() {
	var str, x = function1("hhh", 3)
	var y, z = function2(3, 72)
	fmt.Println(str, x, y, z)
}

func function1(param1 string, param2 int) (string, int) {
	return param1, param2
	//return "function", 3
}

func function2(param1, param2 int) (identifier1 string, identifier2 int) {
	identifier1 = string(param2)
	identifier2 = param1
	fmt.Println("prama2 2 string is", identifier1)
	return
	// return identifier1, identifier2
}

func lambda() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}

func closure1() func() int { //该函数返回一个返回值类型为int的匿名函数
	this_value := 114
	foo := func() int { return this_value }
	return foo
}

// func another_closure() func() int {
// 	this_value = 300 //error, unioned declaration of this_value
// 	return foo
// }

func closure2() (func() int, int) {
	value := 514
	inner := func() int {
		value += 1
		return value
	}
	inner()
	fmt.Println(value)
	return inner, value
}

func args() {
	fmt.Println(args_function(1, 2, 3, 4, 5))
	fmt.Println(args_function(1, 2, 3))
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(args_function(nums...))
}

func args_function(args ...int) int {
	total := 1
	for _, v := range args {
		total += v
	}
	return total
}
