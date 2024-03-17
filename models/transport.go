package models

import (
	"Test1/pb"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChordServer struct {
	pb.UnimplementedKeyServiceServer
}

type Transport struct {
	Node   *pb.Node
	Finger *Fingers
}

func (t *Transport) Register() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, "127.0.0.254:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	random_ip, err := c.RPCRegisterNode(ctx, &pb.NodeIp{IpAddr: t.Node.CurrIp.IpAddr})

	if err != nil {
		panic(err)
	}
	fmt.Println(random_ip)
}

/*Node Listens to services*/
func (t *Transport) GRPCListen(wg *sync.WaitGroup) {
	// node = n
	// hashTable = node.HashTable
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

/*Node acts as client to forward the Findkey request to the successor node*/
func (t *Transport) FindKey(key string) (*pb.ResponseNode, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// fmt.Println(t.Node.SuccIp.IpAddr)
	conn, err := grpc.DialContext(ctx, t.Node.SuccIp.IpAddr+":"+t.Node.SuccIp.Port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// fmt.Println(t.Node.Id)
	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := c.RPCLookup(ctx, &pb.Key{Key: key, HashTable: t.Node.HashTable})

	if err != nil {
		panic(err)
	}
	return res, nil
}

/*Get successor from a node*/
func (t *Transport) GetSuccessor(node *pb.NodeIp) (*pb.NodeIp, error) {
	node_ip, node_port := node.IpAddr, node.Port
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, node_ip+":"+node_port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, _ := c.RPCGetSuccessor(ctx, &pb.Empty{})
	return &pb.NodeIp{IpAddr: res.IpAddr, Port: res.Port}, nil
}

func (t *Transport) findPredecessor(id int32) *pb.NodeIp {
	curr_node := pb.NodeIp{Id: t.Node.Id, IpAddr: t.Node.CurrIp.IpAddr, Port: t.Node.CurrIp.Port}
	succ_node := &pb.NodeIp{Id: t.Node.SuccIp.Id, IpAddr: t.Node.SuccIp.IpAddr, Port: t.Node.SuccIp.Port}
	curr_id := curr_node.Id
	succ_id := succ_node.Id
	for !((curr_id < id && id <= succ_id) || (succ_id < curr_id && (id > curr_id || id <= succ_id))) {
		curr_node = *t.Closest_preceding_finger(&curr_node, id, 4)
		succ_node, _ = t.GetSuccessor(&curr_node)
		curr_id = curr_node.Id
		succ_id = succ_node.Id
	}
	return &curr_node
}

func (t *Transport) findSuccessor(id int32) *pb.NodeIp {
	if t.Node.SuccIp.IpAddr == "" {
		return &pb.NodeIp{Id: t.Node.Id, IpAddr: t.Node.CurrIp.IpAddr, Port: t.Node.CurrIp.Port}
	}
	pred := t.findPredecessor(id)
	succ, _ := t.GetSuccessor(&pb.NodeIp{IpAddr: pred.IpAddr, Port: pred.Port})
	return succ
}

func (t *Transport) Closest_preceding_finger(node *pb.NodeIp, id int32, m int32) *pb.NodeIp {
	node_ip, node_port := node.IpAddr, node.Port
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, node_ip+":"+node_port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, _ := c.RPCClosestPrecedingFinger(ctx, &pb.IdM{Id: id, M: m})

	return res
}
