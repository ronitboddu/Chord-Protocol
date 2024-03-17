package main

import (
	"Test1/models"
	"Test1/node7/listener"
	"Test1/pb"
	"sync"
)

var curr_id int32 = 8
var succ_id int32 = 11
var pred_id int32 = 7
var curr_ip_addr = "127.0.0.8"
var succ_ip_addr = "127.0.0.11"
var pred_ip_addr = "127.0.0.7"
var curr_port = "50001"
var succ_port = "50001"
var pred_port = "50001"

// var node = models.GetNode(curr_ip_addr, succ_ip_addr, pred_ip_addr, curr_port, succ_port, pred_port)
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
	//t.Register()
	var wg sync.WaitGroup
	f.AddKey("8", 8)

	wg.Add(1)
	go listener.GRPCListen(&wg, &t)
	//go t.GRPCListen(&wg)
	wg.Wait()
}
