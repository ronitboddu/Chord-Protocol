package main

import (
	"Test1/models"
	"Test1/node3/listener"
	"Test1/pb"
	"sync"
)

var curr_id int32 = 3
var succ_id int32 = 7
var pred_id int32 = 1
var curr_ip_addr = "127.0.0.3"
var succ_ip_addr = "127.0.0.7"
var pred_ip_addr = "127.0.0.1"
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

var f = models.Fingers{Node: &node, FingerTable: make(map[int32]*pb.NodeIp)}
var t = models.Transport{Node: &node, Finger: &f}

func main() {
	//t.Register()
	GenerateFingerTable()
	var wg sync.WaitGroup
	f.AddKey("3", 3)

	wg.Add(1)
	go listener.GRPCListen(&wg, &t)
	//go t.GRPCListen(&wg)
	wg.Wait()
}

func GenerateFingerTable() {
	f.FingerTable[4] = &pb.NodeIp{Id: 7, IpAddr: "127.0.0.7", Port: "50001"}
	f.FingerTable[5] = &pb.NodeIp{Id: 7, IpAddr: "127.0.0.7", Port: "50001"}
	f.FingerTable[7] = &pb.NodeIp{Id: 7, IpAddr: "127.0.0.7", Port: "50001"}
	f.FingerTable[11] = &pb.NodeIp{Id: 3, IpAddr: "127.0.0.3", Port: "50001"}
}

// 3, 3, 8
// 7, 3, 8
