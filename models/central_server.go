package main

import (
	"Test1/pb"
	"context"
	"fmt"
	"net"
	"reflect"
	"sync"

	"google.golang.org/grpc"
)

type ChordServer struct {
	pb.UnimplementedKeyServiceServer
}

var registry = make(map[string]bool)

func (c *ChordServer) RegisterNode(ctx context.Context, node_ip_addr *pb.NodeIp) (*pb.NodeIp, error) {
	fmt.Println(node_ip_addr.CurrIpAddr)
	if len(registry) == 0 {
		ip_adrr := node_ip_addr.CurrIpAddr
		registry[ip_adrr] = true
		return nil, nil
	}
	rand_list_ip_addr := reflect.ValueOf(registry).MapKeys()[0].String()
	ip_adrr := node_ip_addr.CurrIpAddr
	registry[ip_adrr] = true
	return &pb.NodeIp{CurrIpAddr: rand_list_ip_addr}, nil
}

func GRPCListen(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", "127.0.0.254:50001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterKeyServiceServer(s, &ChordServer{})
	fmt.Printf("gRPC server started on port %s\n", "50001")

	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to listen for gRPC: %v", err))
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go GRPCListen(&wg)
	wg.Wait()
}
