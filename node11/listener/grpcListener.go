package listener

import (
	"Test1/models"
	"Test1/node7/client"
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

var t *models.Transport
var hashTable map[string]int32

func (c *ChordServer) Lookup(ctx context.Context, Key *pb.Key) (*pb.ResponseNode, error) {
	fmt.Println("node1 lookup call")
	if _, ok := hashTable[Key.GetKey()]; !ok {
		return client.FindKey(t.Node.SuccIp.IpAddr, t.Node.SuccIp.Port, Key.Key)
	}
	resNode := &pb.ResponseNode{}
	resNode.IpAddr = t.Node.CurrIp.IpAddr
	resNode.Port = t.Node.CurrIp.Port
	resNode.FoundFlag = true
	return resNode, nil
}
func (c *ChordServer) RPCGetSuccessor(ctx context.Context, emp *pb.Empty) (*pb.NodeIp, error) {
	return &pb.NodeIp{IpAddr: t.Node.SuccIp.IpAddr, Port: t.Node.SuccIp.Port}, nil
}

func GRPCListen(wg *sync.WaitGroup, transport *models.Transport) {
	t = transport
	hashTable = t.Node.HashTable
	lis, err := net.Listen("tcp", t.Node.CurrIp.IpAddr+":"+t.Node.CurrIp.Port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterKeyServiceServer(s, &ChordServer{})
	fmt.Printf("gRPC server started on port %s\n", t.Node.CurrIp.Port)

	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to listen for gRPC: %v", err))
	}
	defer wg.Done()
}
