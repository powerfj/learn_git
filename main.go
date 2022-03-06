/*
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回200
*/

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func f1(w http.ResponseWriter, r *http.Request) {

	//http.request 是结构体类型，包括map[string][]string类型的Header头域
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Connection": {"keep-alive"},
	//	}
	//由于map键值对里面的value是切片类型，通过strings.Join转换成string类型；
	//将request里面的Header写给Response的Header；
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ""))
	}
	//读取环境变量并添加到response Header;
	ver := os.Getenv("TERM_PROGRAM_VERSION")
	w.Header().Add("VERSION", ver)
	//打印客户端地址和返回码
	log.Println("客户端地址是", r.RemoteAddr)
	log.Println("http 返回码是", http.StatusOK)
}

func main() {

	//给/assignment/httpServer注册handler函数
	http.HandleFunc("/assignment/httpServer/", f1)

	//定义监听端口；
	//The handler is typically nil, in which case the DefaultServeMux is used.
	err := http.ListenAndServe("127.0.0.1:7070", nil)
	if err != nil {
		log.Fatal(err)
	}

}
