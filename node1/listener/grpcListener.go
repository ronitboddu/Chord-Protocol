package listener

import (
	"Test1/node1/client"
	"Test1/pb"
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChordServer struct {
	pb.UnimplementedKeyServiceServer
}

var node *pb.Node
var hashTable map[string]int32

func (c *ChordServer) Lookup(ctx context.Context, Key *pb.Key) (*pb.ResponseNode, error) {
	fmt.Println("node1 lookup call")
	if _, ok := hashTable[Key.GetKey()]; !ok {
		return client.FindKey(node.GetSuccIpAddr(), node.GetSuccPort(), Key.GetKey())
	}
	resNode := &pb.ResponseNode{}
	resNode.IpAddr = node.GetCurrIpAddr()
	resNode.Port = node.GetCurrPort()
	resNode.FoundFlag = true
	return resNode, nil
}

func GRPCListen(wg *sync.WaitGroup, n *pb.Node) {
	node = n
	hashTable = node.HashTable
	lis, err := net.Listen("tcp", node.GetCurrIpAddr()+":"+node.GetCurrPort())
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterKeyServiceServer(s, &ChordServer{})
	fmt.Printf("gRPC server started on port %s\n", node.GetCurrPort())

	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to listen for gRPC: %v", err))
	}
	defer wg.Done()
}
