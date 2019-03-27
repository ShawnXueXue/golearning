package main

/**
1.下载protoc:
	https://github.com/protocolbuffers/protobuf
2.安装protoc-gen-go:
	go get -u github.com/golang/protobuf/protoc-gen-go
3.编写.proto文件
4.生成.go文件:
	protoc.exe --go_out=plugins=grpc:./ arith.proto
5.安装grpc
	git clone https://github.com/grpc/grpc-go, copy至  $GOPATH/src/google.golang.org/grpc
	git clone https://github.com/google/go-genproto,  copy至  $GOPATH/src/google.golang.org/genproto
*/

import (
	"context"
	"errors"
	"fmt"
	"golearning/protorpc/pb"
	"google.golang.org/grpc"
	"net"
)

type Arith struct {
}

func (a *Arith) Multiply(ctx context.Context, req *pb.ArithRequest) (*pb.ArithResponse, error) {
	res := &pb.ArithResponse{}
	res.Pro = req.A * req.B
	return res, nil
}

func (a *Arith) Divide(ctx context.Context, req *pb.ArithRequest) (*pb.ArithResponse, error) {
	if req.GetB() == 0 {
		return nil, errors.New("divide by zero")
	}
	res := &pb.ArithResponse{}
	res.Quo = req.GetA() / req.GetB()
	res.Rem = req.GetA() % req.GetB()
	return res, nil
}

func main() {
	// 监听网络
	ln, err := net.Listen("tcp", "localhost:15535")
	if nil != err {
		fmt.Println("网络异常", err)
	}
	// 创建一个grpc句柄
	server := grpc.NewServer()
	// 注册Arith至grpc服务中
	pb.RegisterArithServiceServer(server, &Arith{})
	// 监听grpc服务
	server.Serve(ln)
}
