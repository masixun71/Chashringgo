package Consistent_hashring_go

import (
	"sort"
	"github.com/HuKeping/rbtree"
)

type node struct {
	key string
	hashValue uint32
}

func (n node) Less(than rbtree.Item) bool  {
	return n.hashValue < than.(node).hashValue
}


type nodesArray []node



func (n nodesArray) Len() int {
	return len(n)
}

func (n nodesArray) Less(i, j int) bool  {
	return n[i].hashValue < n[j].hashValue
}

func (n nodesArray) Swap(i, j int)  {
	n[i], n[j] = n[j], n[i]
}

func (n nodesArray) Sort()  {
	sort.Sort(n)
}

