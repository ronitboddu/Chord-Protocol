package models

import (
	"Test1/pb"
	"fmt"
	"math"
)

type Fingers struct {
	Node        *pb.Node
	FingerTable map[int32]*pb.NodeIp
}

func (f *Fingers) AddKey(key string, val int32) {
	f.Node.HashTable[key] = val
}

func (f *Fingers) CreateFingerTable(m int) {
	f.FingerTable = make(map[int32]*pb.NodeIp)
	for i := 0; i < m; i++ {
		id := f.Node.Id
		key := GetFingerKey(id, int32(i), int32(m))
		f.FingerTable[key] = nil
	}
}

func (f *Fingers) PrintFingerTable(m int) {
	for i := 0; i < m; i++ {
		id := f.Node.Id
		key := GetFingerKey(id, int32(i), int32(m))
		if f.FingerTable[key] == nil {
			fmt.Println(key, "nil")
		} else {
			fmt.Println(key, f.FingerTable[key].IpAddr)
		}
	}
}

func GetFingerKey(id int32, i int32, m int32) int32 {
	return (id + int32(math.Pow(float64(2), float64(i)))) % int32(math.Pow(2.0, float64(m)))
}

func Between(key int32, a int32, b int32) bool {
	if a > b {
		return a < key || b >= key
	} else if b > a {
		return a < key && b >= key
	} else if a == b {
		return a != key
	}
	return false
}

// 3, 3, 8
// 7, 3, 8
// 3, 3, 5
