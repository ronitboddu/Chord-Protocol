package models

import (
	"Test1/pb"
	"fmt"
	"math"
)

type Fingers struct {
	Node        *pb.Node
	fingerTable map[int32]*pb.NodeIp
}

func (f *Fingers) AddKey(key string, val int32) {
	f.Node.HashTable[key] = val
}

func (f *Fingers) CreateFingerTable(m int) {
	f.fingerTable = make(map[int32]*pb.NodeIp)
	for i := 0; i < m; i++ {
		id := f.Node.Id
		key := GetFingerKey(id, i, m)
		f.fingerTable[key] = nil
	}
}

func (f *Fingers) PrintFingerTable(m int) {
	for i := 0; i < m; i++ {
		id := f.Node.Id
		key := GetFingerKey(id, i, m)
		if f.fingerTable[key] == nil {
			fmt.Println(key, "nil")
		} else {
			fmt.Println(key, f.fingerTable[key].IpAddr)
		}
	}
}

func GetFingerKey(id int32, i int, m int) int32 {
	return (id + int32(math.Pow(float64(2), float64(i)))) % int32(math.Pow(2.0, float64(m)))
}
