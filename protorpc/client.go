package main

import (
	"context"
	"fmt"
	"golearning/protorpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:15535", grpc.WithInsecure())
	if nil != err {
		fmt.Println("连接服务器失败", err)
	}
	defer conn.Close()
	client := pb.NewArithServiceClient(conn)
	response, e := client.Divide(context.Background(), &pb.ArithRequest{A: 5, B: 3})
	if nil != e {
		fmt.Println("divide error!", e)
	}
	fmt.Printf("quo:%d, rem:%d\n", response.Quo, response.Rem)
	response, e = client.Multiply(context.Background(), &pb.ArithRequest{A: 5, B: 3})
	if nil != e {
		fmt.Println("multiply error!", e)
	}
	fmt.Printf("pro:%d\n", response.Pro)
}
