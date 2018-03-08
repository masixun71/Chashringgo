package Chashringgo

import (
	//	"fmt"
	"testing"
)

const (
	node1 = "192.168.1.1"
	node2 = "192.168.1.2"
	node3 = "192.168.1.3"
	node4 = "192.168.1.4"
	node5 = "192.168.1.5"
	node6 = "192.168.1.6"
	node7 = "192.168.1.7"
	node8 = "192.168.1.8"
	node9 = "192.168.1.9"
	node10 = "192.168.1.10"
	node11 = "192.168.1.11"
	node12 = "192.168.1.12"
	node13 = "192.168.1.13"
	node14 = "192.168.1.14"
	node15 = "192.168.1.15"
	node16 = "192.168.1.16"
	node17 = "192.168.1.17"
	node18 = "192.168.1.18"
)

func TestHash(t *testing.T) {

	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 1
	nodeWeight[node2] = 1
	nodeWeight[node3] = 1
	vitualSpots := 1

	hash := NewChashring(vitualSpots)
	hash.AddNodes(nodeWeight)

	//var ss []rbtree.Item
	//hash.nodes.Ascend(hash.nodes.Min(), func(i rbtree.Item) bool {
	//	ss = append(ss, i)
	//	return true
	//})


	//node1Count := 0
	//node2Count := 0
	//node3Count := 0
	//for i :=1;i <=10; i++ {
	//	tKey := hash.GetNode(strconv.Itoa(i))
	//	t.Log(tKey)
	//	if tKey == node1 {
	//		node1Count += 1
	//	}
	//	if tKey == node2 {
	//		node2Count += 1
	//
	//	}
	//	if tKey == node3 {
	//		node3Count += 1
	//
	//	}
	//}
	//t.Logf("nodes node1:%v, node2:%v, node3:%v", node1Count, node2Count, node3Count)
	//t.Log(hash)
	if hash.GetNode("1") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("2"))
	}
	if hash.GetNode("4") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("3"))
	}

	hash.RemoveNode(node3)
	if hash.GetNode("1") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("2"))
	}
	if hash.GetNode("4") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("3"))
	}

	hash.AddNode(node3, 1)
	if hash.GetNode("1") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("2"))
	}
	if hash.GetNode("4") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("3"))
	}

}
