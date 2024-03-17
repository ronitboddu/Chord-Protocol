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

var curr_id int32 = 1
var succ_id int32 = 3
var pred_id int32 = 14
var curr_ip_addr = "127.0.0.1"
var succ_ip_addr = "127.0.0.3"
var pred_ip_addr = "127.0.0.14"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

var node = pb.Node{
	Id:        curr_id,
	CurrIp:    &pb.NodeIp{Id: curr_id, IpAddr: curr_ip_addr, Port: curr_port},
	SuccIp:    &pb.NodeIp{Id: succ_id, IpAddr: succ_ip_addr, Port: succ_port},
	PredIp:    &pb.NodeIp{Id: pred_id, IpAddr: pred_ip_addr, Port: pred_port},
	HashTable: make(map[string]int32),
}

var f = models.Fingers{Node: &node}
var t = models.Transport{Node: &node, Finger: &f}

func main() {
	//f.CreateFingerTable(4)
	//f.PrintFingerTable(4)
	//t.RegisterNode()
	var wg sync.WaitGroup
	f.AddKey("1", 1)
	wg.Add(1)
	go listener.GRPCListen(&wg, &t)
	resNode := t.Closest_preceding_finger(&pb.NodeIp{Id: 3, IpAddr: "127.0.0.3", Port: "50001"}, 1, 4)
	//resNode, _ := t.FindKey("8")
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
