package main

import (
	"Test1/node7/listener"
	"Test1/pb"
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var id int32 = 7
var curr_ip_addr = "127.0.0.3"
var succ_ip_addr = "127.0.0.1"
var pred_ip_addr = "127.0.0.2"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

// var node = models.GetNode(curr_ip_addr, succ_ip_addr, pred_ip_addr, curr_port, succ_port, pred_port)
var node = pb.Node{
	Id:         id,
	CurrIpAddr: curr_ip_addr,
	SuccIpAddr: succ_ip_addr,
	PredIpAddr: pred_ip_addr,
	CurrPort:   curr_port,
	SuccPort:   succ_port,
	PredPort:   pred_port,
	HashTable:  make(map[string]int32),
}

func AddKey(node *pb.Node, key string, val int32) {
	node.HashTable[key] = val
}

func RegisterNode() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, "127.0.0.254:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	random_ip, err := c.RegisterNode(ctx, &pb.NodeIp{CurrIpAddr: node.CurrIpAddr})

	if err != nil {
		panic(err)
	}
	fmt.Println("here")
	fmt.Println(random_ip.CurrIpAddr)
}

func main() {
	RegisterNode()
	var wg sync.WaitGroup
	AddKey(&node, "3", 3)

	wg.Add(1)
	go listener.GRPCListen(&wg, &node)
	wg.Wait()
}