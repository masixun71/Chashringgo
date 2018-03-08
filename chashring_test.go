package Chashringgo

import (
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

	//hash := NewChashring(vitualSpots)
	hash := NewChashringWithHashAlgorithm(vitualSpots, &FNV1_32_HASH{})


	hash.AddNodes(nodeWeight)

	if hash.GetNode("1") != node2 {
		t.Fatal("get node is error")
	}
	if hash.GetNode("2") != node1 {
		t.Fatal("get node is error")
	}
	if hash.GetNode("4") != node3 {
		t.Fatal("get node is error")
	}

	hash.RemoveNode(node3)
	hash.UpdateNode(node3, 1)
	if hash.GetNode("1") != node2 {
		t.Fatal("get node is error")
	}
	if hash.GetNode("2") != node1 {
		t.Fatal("get node is error")
	}
	if hash.GetNode("4") != node1 {
		t.Fatal("get node is error")
	}

	hash.AddNode(node3, 1)
	if hash.GetNode("1") != node2 {
		t.Fatal("get node is error")
	}
	if hash.GetNode("2") != node1 {
		t.Fatal("get node is error")
	}
	if hash.GetNode("4") != node3 {
		t.Fatal("get node is error")
	}

}
