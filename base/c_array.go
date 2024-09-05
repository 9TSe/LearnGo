package main

import (
	"fmt"
	"time"
)

func array() {
	var arr [10]int
	arr[3] = 4
	i := arr[3]

	var arr1 = [3]int{1, 2, 3}
	arr2 := [2]int{1, 2}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := [...]int{1: 3, 2: 4}
	fmt.Println(i, arr1, arr2, arr3, arr4)
}

func slice() {
	var arr1 []int                            // 声明切片，和数组类型声明类似，不需要指定长度
	var arr2 = []int{1, 2, 3, 4}              // 声明和初始化切片
	arr3 := []int{1, 2, 3, 4}                 // 简写
	chars := []string{0: "a", 2: "c", 1: "b"} // ["a", "b", "c"] 前面的数字是索引

	//var b = arr[lo:hi]	// 通过下标索引从已有的数组或切片创建新切片，下标前闭后开，取值从lo到hi-1
	var b1 = arr2[1:4]         // 取切片a的下标索引从1到3的值赋值给新切片b
	var b2 = arr3[:3]          // :前面没有值表示起始索引是0，等同于a[0:3]
	var b3 = arr2[3:]          // :后面没有值表示结束索引是len(a)，等同于a[3:len(a)]
	arr1 = append(arr1, 17, 3) // 往切片里添加新元素(17, 3), 如果切片底层数组容量不足，会创建新的底层数组
	c := append(b1, b2...)     // 把切片a和b的值拼接起来，组成新切片

	// 使用make来创建切片
	arr4 := make([]byte, 5, 5) // make的第2个参数是切片长度，第3个参数是切片容量
	arr5 := make([]byte, 5)    // 第3个切片容量参数可选，即可以不传值

	// 根据数组来创建切片
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // 切片s指向了数组x的内存空间，改变切片s的值，也会影响数组x的值, 类似与引用

	fmt.Println(arr1, arr2, arr3, chars, b1, b2, b3, c, arr4, arr5, s, x)

	for o, p := range arr2 {
		fmt.Println(o, p) // o is index, p is value
	}
	for _, p := range arr2 {
		fmt.Println(p) // p is value
	}
	for o := range arr2 {
		fmt.Println(o) // o is index
	}

	for range time.Tick(time.Second * 2) {
		fmt.Println("hello") //2second run
	}
}

type Vertex3 struct {
	url  string
	name string
}

func set() {
	m := make(map[string]int)
	m["one"] = 1
	delete(m, "one")

	value, ok := m["one"]

	var n = map[string]Vertex3{
		"google": {"www.google.com", "google"},
		"baidu":  {"www.baidu.com", "baidu"},
	}
	fmt.Println(value, ok, n)

	for key, value := range m {
		fmt.Println(key, value)
	}
}
