package main

import (
	"Test1/models"
	"Test1/node1/listener"
	"Test1/pb"
	"fmt"
	"sync"
)

type ChordServer struct {
	pb.UnimplementedKeyServiceServer
}

var id int32 = 1
var curr_ip_addr = "127.0.0.1"
var succ_ip_addr = "127.0.0.3"
var pred_ip_addr = "127.0.0.14"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

var node = pb.Node{
	Id:        id,
	CurrIp:    &pb.NodeIp{Id: id, IpAddr: curr_ip_addr, Port: curr_port},
	SuccIp:    &pb.NodeIp{Id: 3, IpAddr: succ_ip_addr, Port: succ_port},
	PredIp:    &pb.NodeIp{Id: 14, IpAddr: pred_ip_addr, Port: pred_port},
	HashTable: make(map[string]int32),
}

var t = models.Transport{Node: &node, CS: models.ChordServer{}}
var f = models.Fingers{Node: &node}

func main() {
	f.CreateFingerTable(4)
	f.PrintFingerTable(4)
	//t.RegisterNode()
	var wg sync.WaitGroup
	f.AddKey("1", 1)
	wg.Add(1)
	go listener.GRPCListen(&wg, &t)

	resNode, _ := t.FindKey("14")
	// fmt.Println(resNode.IpAddr)
	// resNode, _ = t.FindKey("7")
	// fmt.Println(resNode.IpAddr)
	// resNode, _ = t.FindKey("8")
	// fmt.Println(resNode.IpAddr)
	// resNode, _ = t.FindKey("11")
	// fmt.Println(resNode.IpAddr)
	// resNode, _ = t.FindKey("14")
	// fmt.Println(resNode.IpAddr)
	// resNode, _ := t.RPCGetSuccessor("127.0.0.7", "50001")
	fmt.Println(resNode.IpAddr, resNode.Port)
	wg.Wait()
}
