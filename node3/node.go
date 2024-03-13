package main

import (
	"Test1/models"
	"Test1/node3/listener"
	"Test1/pb"
	"sync"
)

var id int32 = 3
var curr_ip_addr = "127.0.0.3"
var succ_ip_addr = "127.0.0.7"
var pred_ip_addr = "127.0.0.1"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

var node = pb.Node{
	Id:        id,
	CurrIp:    &pb.NodeIp{Id: id, IpAddr: curr_ip_addr, Port: curr_port},
	SuccIp:    &pb.NodeIp{Id: 7, IpAddr: succ_ip_addr, Port: succ_port},
	PredIp:    &pb.NodeIp{Id: 1, IpAddr: pred_ip_addr, Port: pred_port},
	HashTable: make(map[string]int32),
}

var t = models.Transport{Node: &node, CS: models.ChordServer{}}
var f = models.Fingers{Node: &node}

func main() {
	//t.Register()
	var wg sync.WaitGroup
	f.AddKey("3", 3)

	wg.Add(1)
	go listener.GRPCListen(&wg, &t)
	//go t.GRPCListen(&wg)
	wg.Wait()
}
