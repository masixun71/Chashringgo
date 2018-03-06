package Consistent_hashring_go

import "sort"

type node struct {
	key string
	hashValue uint32
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

