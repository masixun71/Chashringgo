# Consistent-hashring-go
一致性hash算法 go 实现



#### 一开始采用红黑树实现，但发现有序数组查找效率和红黑树效率一样，甚至略高，替换为有序数组，节点加入了权重，在增删节点时必须重置节点数组，红黑树插入效率也低于数据，最终替换为有序数组，只进行一次数组排序。


安装｜install
===

	go get github.com/masixun71/Chashringgo

Example
===


```
	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 1
	nodeWeight[node2] = 1
	nodeWeight[node3] = 1
	vitualSpots := 1

	//hash := NewChashring(vitualSpots)
	hash := NewChashringWithHashAlgorithm(vitualSpots, &FNV1_32_HASH{})

	hash.AddNodes(nodeWeight)
	hash.AddNode(node3, 1)
	hash.RemoveNode(node3)
	hash.UpdateNode(node3, 1)
	nodeIndex := hash.GetNode("1")
```
