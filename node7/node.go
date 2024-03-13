package main

import (
	"Test1/models"
	"Test1/node7/listener"
	"Test1/pb"
	"sync"
)

var id int32 = 7
var curr_ip_addr = "127.0.0.7"
var succ_ip_addr = "127.0.0.8"
var pred_ip_addr = "127.0.0.3"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

var node = pb.Node{
	Id:        id,
	CurrIp:    &pb.NodeIp{Id: id, IpAddr: curr_ip_addr, Port: curr_port},
	SuccIp:    &pb.NodeIp{Id: 8, IpAddr: succ_ip_addr, Port: succ_port},
	PredIp:    &pb.NodeIp{Id: 3, IpAddr: pred_ip_addr, Port: pred_port},
	HashTable: make(map[string]int32),
}

var t = models.Transport{Node: &node, CS: models.ChordServer{}}
var f = models.Fingers{Node: &node}

func main() {
	//t.Register()
	var wg sync.WaitGroup
	f.AddKey("7", 7)

	wg.Add(1)
	go listener.GRPCListen(&wg, &t)
	//go t.GRPCListen(&wg)
	wg.Wait()
}
