package Consistent_hashring_go

import (
	//	"fmt"
	"testing"
	"fmt"
)

const (
	node1 = "192.168.1.1"
	node2 = "192.168.1.2"
	node3 = "192.168.1.3"
)

func getNodesCount(nodes nodesArray) (int, int, int) {
	node1Count := 0
	node2Count := 0
	node3Count := 0

	for _, node := range nodes {
		if node.key == node1 {
			node1Count += 1
		}
		if node.key == node2 {
			node2Count += 1

		}
		if node.key == node3 {
			node3Count += 1

		}
	}
	return node1Count, node2Count, node3Count
}

func TestHash(t *testing.T) {

	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 1
	nodeWeight[node2] = 1
	nodeWeight[node3] = 1
	vitualSpots := 1

	hash := NewChashring(vitualSpots)
	hash.AddNodes(nodeWeight)
	c1, c2, c3 := getNodesCount(hash.nodes)
	fmt.Printf("len of nodes is %v after AddNodes node1:%v, node2:%v, node3:%v\n", len(hash.nodes), c1, c2, c3)

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
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after AddNodes node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

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
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after RemoveNode node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

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
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after AddNode node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

}
