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

func (c *ChordServer) RPCLookup(ctx context.Context, Key *pb.Key) (*pb.ResponseNode, error) {
	fmt.Println("node8 lookup call")
	if _, ok := hashTable[Key.GetKey()]; !ok {
		return client.FindKey(t.Node.SuccIp.IpAddr, t.Node.SuccIp.Port, Key.Key)
	}
	resNode := &pb.ResponseNode{}
	resNode.IpAddr = t.Node.CurrIp.IpAddr
	resNode.Port = t.Node.CurrIp.Port
	resNode.FoundFlag = true
	return resNode, nil
}

func (c *ChordServer) GetSuccessor(ctx context.Context, emp *pb.Empty) (*pb.NodeIp, error) {
	return &pb.NodeIp{Id: t.Node.SuccIp.Id, IpAddr: t.Node.SuccIp.IpAddr, Port: t.Node.SuccIp.Port}, nil
}

func (c *ChordServer) RPCClosestPrecedingFinger(ctx context.Context, id_m *pb.IdM) (*pb.NodeIp, error) {
	ft := t.Finger.FingerTable
	id, m := id_m.Id, id_m.M
	max, min := int32(0), int32(2147483647)
	max_node, min_node := &pb.NodeIp{}, &pb.NodeIp{}
	for i := m - 1; i >= 0; i-- {
		finger := models.GetFingerKey(t.Node.Id, i, m)
		if ft[finger] == nil {
			continue
		}
		finger_succ_id := ft[finger].Id
		//curr_id := t.Node.Id
		fmt.Println(finger_succ_id, max, min)
		if finger_succ_id > max {
			max = finger_succ_id
			max_node = ft[finger]
		}
		if id > finger_succ_id && (id-finger_succ_id) < min {
			min = (id - finger_succ_id)
			min_node = ft[finger]
		}
	}
	if min_node.IpAddr == "" {
		return max_node, nil
	}
	return min_node, nil
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
