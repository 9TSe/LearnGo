package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"
)

type Vertex struct {
	X, Y float64
}

// 给结构体定义方法，在func关键字和方法名称之间加上结构体声明(var_name StructName)即可
// 调用方法时，会把结构体的值拷贝一份
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 如果想调用方法时改变外部结构体变量的值，方法需要使用指针接受者
// 下面的方法，每次调用add方法时就不会拷贝结构体的值
func (v *Vertex) add(n float64) {
	v.X += n
	v.Y += n
}
func structs() {
	// 结构体是一种类型，也是一系列字段的集合
	// 创建结构体变量
	var v = Vertex{1, 2}
	var z = Vertex{X: 1, Y: 2}               // 通过字段名称:值的形式来创建结构体变量
	var y = []Vertex{{1, 2}, {5, 2}, {5, 5}} // 初始化结构体切片

	// 访问结构体的字段
	v.X = 4
	// 调用结构体方法
	z.add(4)

	//匿名结构体
	point := struct {
		X, Y int
	}{1, 2}
	fmt.Println(v.Abs(), v, z, y, point)

}

func ptr() {
	p := Vertex{1, 2}
	q := &p
	r := &Vertex{1, 2}

	var ptr *Vertex = new(Vertex)
	fmt.Println(p, q, r, ptr)
}

// 接口声明
type Awesomizer interface {
	Awesomize() string
}

// 嵌套接口
type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// 结构体不会在声明的时候指定要实现某个接口
type Foo struct{}

// 相反，结构体如果实现了接口里的所有方法，那就隐式表明该结构体满足了该接口
// 可以通过接口变量来调用结构体方法
func (foo Foo) Awesomize() string {
	return "Awesome!"
}

func Interface() string {
	var foo = Foo{}
	return foo.Awesomize()
}

// 结构体嵌套
type Server struct {
	Host string
	Port int
	*log.Logger
}

func struct_embedding() {
	//server := &Server{"nitesy", 8080, log.New(os.Stdout, "Server: ", log.Lshortfile)}
	// server.Log()
	// var logger *log.Logger = server.Log()
	fmt.Println("test")
}

type Myerror interface {
	Error() string
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negative value")
	}
	return math.Sqrt(x), nil
}

func test_error() {
	val, err := sqrt(-1)
	if err != nil {
		// 处理错误
		fmt.Println(err) // negative value
		return
	}
	// 没有错误，打印结果
	fmt.Println(val)
}

func routine_test(s string) string {
	return s
}

func goroutine() {
	go routine_test("hhh")

	go func(value int) int {
		return value
	}(36)
}

func do_buffer(channelOut, channelIn chan int) {
	select {
	case channelOut <- 1:
		fmt.Println("we coud write")
	case x := <-channelIn:
		fmt.Println("we coud read", x)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func channel() {
	// 无缓冲管道中，发送操作和接收操作必须同步完成
	// m := make(chan int) // no buffer
	m := make(chan int, 2)
	m <- 1
	value := <-m
	fmt.Println(value)

	take_buf := make(chan int, 100) //buffer size : 100 int
	// for i := range take_buf {
	// 	fmt.Println(i)
	// }
	close(take_buf)

	v, ok := <-take_buf
	fmt.Println(v, ok) // recive and judge
	do_buffer(m, take_buf)
}

func print() {
	fmt.Println("Hello, 你好, नमस्ते, Привет, ᎣᏏᏲ") //基本的打印，会自动换行
	p := struct{ X, Y int }{17, 2}
	fmt.Println("My point:", p, "x coord=", p.X)       // 打印结构体和字段值
	s := fmt.Sprintln("My point:", p, "x coord=", p.X) // 打印内容到字符串变量里

	fmt.Printf("%d hex:%x bin:%b fp:%f sci:%e", 17, 17, 17, 17.0, 17.0) // C风格的格式化打印
	s2 := fmt.Sprintf("%d %f", 17, 17.0)                                // 字符串格式化

	hellomsg := `
	 "Hello" in Chinese is 你好 ('Ni Hao')
	 "Hello" in Hindi is नमस्ते ('Namaste')"
	` // 跨越多行的字符串，使用``
	fmt.Println(s, s2, hellomsg)
}

func type_select(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	default:
		fmt.Println("unknown")
	}
}

// type_select("hello")
// type_select(1)
// type_select(1.0)
