package main

import (
	"fmt"
)

func change_slice1(slice []int) {
	//slice[0] = 1 可以修改实参
	slice = append(slice, 2) //cant
	fmt.Println("in function slice = ", slice)
}

func change_slice2(slice *[]int) {
	*slice = append(*slice, 2) //yep
}

func sliceComplement() {
	var length int = 10
	var capacity int = 20
	var slice_var1 []int = make([]int, length, capacity) // cap是切片容量，是make的可选参数
	fmt.Println(len(slice_var1), cap(slice_var1))

	var slice_var2 []int = make([]int, length)
	slice_var2 = append(slice_var2, 1, 2, 3, 4, 5)
	fmt.Println(slice_var2)

	slice_var3 := []int{}
	copy(slice_var3, slice_var3) //copy(dest, src)
	// 只从源切片srcSlice拷贝min(len(srcSlice), len(dstSlice))个元素到目标切片dstSlice里。
	// 如果dstSlice的长度是0，那一个都不会从srcSlice拷贝到dstSlice里。
	// 如果dstSlice的长度M小于srcSlice的长度N，则只会拷贝srcSlice里的前M个元素到目标切片dstSlice里

	slice := []int{} //slice != nil 它是一个长度为0的切片
	var slice2 []int //slice2 == nil 它是一个nil的切片

	change_slice1(slice)
	change_slice2(&slice2)
	fmt.Println(slice, slice2)
}

// rune == int32
func buildMap(str string, m map[rune]int) {
	/*函数内对map变量m的修改会影响main里的实参mapping*/
	for _, value := range str { // _ is index
		m[value]++
	}
}

func map_complement() {
	mapping := map[rune]int{}
	str := "abc"
	buildMap(str, mapping)

	/*
		mapping的值被buildMap修改了
	*/
	for key, value := range mapping {
		fmt.Printf("key:%v, value:%d\n", key, value)
	}
	delete(mapping, 'a') //succ
	delete(mapping, 'd') //although no effect but no error
	delete(mapping, 98)  //succ
	fmt.Println(len(mapping))

	//key必须支持==和!=比较，才能用作map的key。
	//因此切片slice，函数类型function，集合map，不能用作map的key
}

// interface1
type Felines interface {
	feet()
}

// interface2, 嵌套了interface1
type Mammal interface {
	Felines
	born()
}

// 猫实现Mammal这个interface里的所有方法
type Cat struct {
	name string
	age  int
}

func (cat Cat) feet() {
	fmt.Println("cat feet")
}

func (cat *Cat) born() {
	fmt.Println("cat born")
}

// 打印空interface的类型和具体的值
func printt(x interface{}) { //可以接受任何类型的参数
	fmt.Printf("type:%T, value:%v\n", x, x)
}

func interface_test() {
	cat := Cat{"rich", 1}
	/*Mammal有feet和born方法，2个都可以调用*/
	var a Mammal = &cat
	a.feet()
	a.born()

	var b Felines = cat
	b.feet()

	var dict interface{}
	dict = map[string]int{"a": 1}
	printt(dict) // type:map[string]int, value:map[a:1]

	var dict1 interface{}
	dict1 = make(map[string]int)
	dict1.(map[string]int)["a"] = 1
	value, ok := dict1.(map[string]int)["a"]
	fmt.Println(value, ok) // 1 true
	printt(dict1)          // type:map[string]int, value:map[a:1]

	dict2 := make(map[string]interface{})
	dict2["a"] = 1
	dict2["b"] = "c"

	// 传struct实参给空接口
	printt(cat) // type:main.Cat, value:{rich 1}

}
