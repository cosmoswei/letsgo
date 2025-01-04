package gin

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"time"
)

// 处理请求的函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 模拟一些计算工作
	time.Sleep(10 * time.Millisecond) // 模拟延迟，表现出计算的复杂性
	w.Write([]byte("Hello, World!"))
}

// 高性能的Web服务配置
func setupServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:           ":8088", // 监听端口8080
		Handler:        mux,     // 路由处理器
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	return server
}

func WebGo() {
	server := setupServer()
	// 启动服务并监听请求
	fmt.Println("Starting server on port 8088...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func fastHelloHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetBodyString("Hello, Go!")
}

func FastWeb() {
	// 使用fasthttp处理HTTP请求
	server := &fasthttp.Server{
		Handler: fastHelloHandler,
	}

	fmt.Println("Starting server on port 8088...")
	if err := server.ListenAndServe(":8088"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func PingPongWeb() {
	// 定义请求处理函数
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ping":
			// 设置响应内容
			ctx.SetContentType("text/plain; charset=utf-8")
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetBodyString("pong")
		default:
			// 处理未知路径
			ctx.Error("unsupported path", fasthttp.StatusNotFound)
		}
	}

	// 启动服务器
	port := ":8086"
	log.Printf("Server is running at http://localhost%s/ping", port)
	if err := fasthttp.ListenAndServe(port, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
