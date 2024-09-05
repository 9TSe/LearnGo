package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
)

// content持有服务器static目录下的所有文件
//
//go:embed static/*
var content embed.FS

func test_embed() {
	http.Handle("/", http.FileServer(http.FS(content)))
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	// 读取服务器static目录下的内容
	entries, err := content.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		resp, err := http.Get("http://localhost:8080/static/" + e.Name())
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q: %s", e.Name(), body)
	}
}

// 定义处理http请求的Handler
type Hello struct{}

// 结构体Hello实现接口类型http.Handler里的方法ServeHTTP
// 这样结构体Hello的实例就可以作为http的Handler来接收http请求，返回http响应结果
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func httpserver() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}

// 下面是http.ServerHTTP的方法签名
// type Handler interface {
//     ServeHTTP(w http.ResponseWriter, r *http.Request)
// }
