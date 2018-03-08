package Consistent_hashring_go

import (
	//	"fmt"
	"testing"
)


func Benchmark_get(b *testing.B) {
	b.StopTimer()
	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 1
	nodeWeight[node2] = 1
	nodeWeight[node3] = 1
	vitualSpots := 1

	hash := NewChashring(vitualSpots)
	hash.AddNodes(nodeWeight)
	hash.RemoveNode(node3)
	hash.AddNode(node3, 1)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		hash.GetNode("wefwfwadass12321412cdscds")
	}

}

func Benchmark_crud(b *testing.B) {
	b.StopTimer()
	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 1
	nodeWeight[node2] = 1
	nodeWeight[node3] = 1
	vitualSpots := 1

	hash := NewChashring(vitualSpots)
	hash.AddNodes(nodeWeight)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		hash.RemoveNode(node3)
		hash.AddNode(node3, 1)
		hash.AddNodes(nodeWeight)
	}

}
