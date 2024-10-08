// server.go
package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"net"
)

// 自定义 JSON 编解码器
type jsonCodec struct{}

func (jc *jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jc *jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (jc *jsonCodec) Name() string {
	return "json"
}

// 定义请求和响应结构
type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}

// 实现 gRPC 服务
type myService struct{}

func (s *myService) SayHello(ctx context.Context, req *helloRequest) (*helloResponse, error) {
	return &helloResponse{Message: "Hello " + req.Name}, nil
}

func GrpcRun() {
	// 注册自定义的 Codec
	encoding.RegisterCodec(&jsonCodec{})

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	// 注册服务
	// pb.RegisterMyServiceServer(s, &myService{})   // 如果使用 protobuf，这里要用 pb 自动生成的代码
	go s.Serve(lis)

	fmt.Println("Server is running on port 50051...")
	select {} // 阻止主线程退出
}
