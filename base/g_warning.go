package main

import (
	"fmt"
	"reflect"
)

func string_test() {
	var v complex64 = 3.2 + 12i  //float32
	var v2 complex128 = 3.2 + 1i //float64 these 1 cant ignore

	fmt.Println(v, v2)
	str := "abcd"
	str = "efgh"
	// str[1] = "z" // error string can't be modified
	fmt.Println(len(str), str+"ijk", reflect.TypeOf(str))
	// fmt.Println(&str[1]) // error
	fmt.Printf("%T\n", str)
}

// 全局变量允许声明后不使用，编译不会报错。
var globalVar int
var globalVar2 int = 20
var globalVar3 = 30

var (
	v1 int  = 10
	v2 bool = true
)
var (
	v5 int  // the value will be defaulted to 0
	v6 bool // the value will be defaulted to false
)
var (
	v3 = 20
	v4 = false
)

func scope_test() {
FLAG:
	globalVar = 10
	v5 = 40

	var local int
	local = 20
	fmt.Println("local var = ", local)

	var str, int_value = "10", 10
	str1, int_value1 := "20", 20
	fmt.Println(str, int_value, str1, int_value1)
	if 1 == 0 {
		goto FLAG
	}
}

func const_test() {
	const (
		class1 = 0
		class2        // class2 = 0
		class3 = iota //iota is 2, so class3 = 2
		class4        // class4 = 3
		class5 = "abc"
		class6        // class6 = "abc"
		class7 = iota // class7 is 6
	)
}

func multidimention_array() {
	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	arr2 := [2][3]int{}
	arr2[1][0] = 3

	slice1 := [][]int{}
	slice_append := []int{1, 2, 3}
	slice1 = append(slice1, slice_append) //append only use to slice
	fmt.Println(arr1, arr2, slice1)

	fmt.Println("warning: slice append")
	twoDimArray := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5}
	twoDimArray = append(twoDimArray, row1)
	twoDimArray = append(twoDimArray, row2)
	fmt.Println("twoDimArray=", twoDimArray) //[[1 2 3] [4 5]]
}

func pointer_test() {
	array := [3]int{1, 2, 3}
	var arrayPtr *[3]int = &array // C++赋值就不用加&
	fmt.Println(arrayPtr)

	var ptr_array [3]*int
	ptr_array[1] = &array[1]
	fmt.Println(*ptr_array[1])
	//结构体指针访问结构体里的成员，也是用点.，这个和C++用->不一样
}
