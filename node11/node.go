package main

import (
	"Test1/models"
	"Test1/node7/listener"
	"Test1/pb"
	"sync"
)

var id int32 = 11
var curr_ip_addr = "127.0.0.11"
var succ_ip_addr = "127.0.0.14"
var pred_ip_addr = "127.0.0.8"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

// var node = models.GetNode(curr_ip_addr, succ_ip_addr, pred_ip_addr, curr_port, succ_port, pred_port)
var node = pb.Node{
	Id:        id,
	CurrIp:    &pb.NodeIp{Id: id, IpAddr: curr_ip_addr, Port: curr_port},
	SuccIp:    &pb.NodeIp{Id: 14, IpAddr: succ_ip_addr, Port: succ_port},
	PredIp:    &pb.NodeIp{Id: 8, IpAddr: pred_ip_addr, Port: pred_port},
	HashTable: make(map[string]int32),
}

var t = models.Transport{Node: &node, CS: models.ChordServer{}}
var f = models.Fingers{Node: &node}

func main() {
	//t.Register()
	var wg sync.WaitGroup
	f.AddKey("11", 11)

	wg.Add(1)
	go listener.GRPCListen(&wg, &t)
	//go t.GRPCListen(&wg)
	wg.Wait()
}
