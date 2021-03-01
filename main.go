package main

import (
	"fmt"
	pb "github.com/gitcloneese/testgrpc"
	"google.golang.org/grpc"
	"net"
	"context"
)

type server struct{}


func (this *server) Sayhello(ctx context.Context, in *pb.HelloReq ) (out *pb.HelloResp, err error){
	return &pb.HelloResp{Msg :"hello, " + in.Name}, nil	
	
		
}
func (this *server) Sayname(ctx context.Context, in *pb.NameReq) (out *pb.NameResp, err error){
	return &pb.NameResp{Msg : in.Name + ", 早上好! "}, nil
	
}

func main() {
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	
	//创建grpc服务
	srv := grpc.NewServer()
	pb.RegisterHelloserverServer(srv, &server{})
	
	err = srv.Serve(ln)
	
	if err != nil {
		fmt.Println("网络有问题 !", err)
	}
	
}
