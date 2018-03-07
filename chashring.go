package Consistent_hashring_go

import (
	"math"
	"strconv"
	"sort"
)

const DefaultNodeVirualSpot = 40

type chashring struct {
	nodeVirualSpot int
	nodes          nodesArray
	weights        map[string]int
	useHash        hashAlgorithm
}

func NewChashring(spot int) *chashring {

	if (spot == 0) {
		spot = DefaultNodeVirualSpot
	}

	hash := &chashring{
		nodeVirualSpot: spot,
		weights:        make(map[string]int),
		useHash:        &FNV1_32_HASH{},
	}

	return hash
}

func NewChashringWithHashAlgorithm(spot int, useHash hashAlgorithm) *chashring {

	if (spot == 0) {
		spot = DefaultNodeVirualSpot
	}

	hash := &chashring{
		nodeVirualSpot: spot,
		weights:        make(map[string]int),
		useHash:        useHash,
	}

	return hash
}

func (hash *chashring) AddNodes(nodeWeight map[string]int) {

	for nodeKey, weight := range nodeWeight {
		hash.weights[nodeKey] = weight
	}

	hash.generate()
}

func (hash *chashring) AddNode(nodeKey string, weight int) {

	hash.weights[nodeKey] = weight

	hash.generate()

}

func (hash *chashring) UpdateNode(nodeKey string, weight int) {

	hash.weights[nodeKey] = weight

	hash.generate()

}

func (hash *chashring) RemoveNode(nodeKey string) {
	delete(hash.weights, nodeKey)

	hash.generate()

}

func (hash *chashring) generate() {
	var totalW int
	for _, w := range hash.weights {
		totalW += w
	}
	totalVirtualSpots := hash.nodeVirualSpot * len(hash.weights)
	hash.nodes = nodesArray{}

	for nodeKey, w := range hash.weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots)))
		for i := 1; i <= spots; i++ {
			n := node{
				key:       nodeKey,
				hashValue: hash.useHash.Hash((nodeKey + ":" + strconv.Itoa(i))),
			}
			hash.nodes = append(hash.nodes, n)
		}
	}
	hash.nodes.Sort()
}

func (hash *chashring) GetNode(str string) string {

	lenHashNodes := len(hash.nodes)

	if lenHashNodes == 0 {
		return ""
	}

	strHash := hash.useHash.Hash(str)

	index := sort.Search(len(hash.nodes), func(i int) bool {
		return hash.nodes[i].hashValue >= strHash
	});
	if (index == lenHashNodes) {
		index = 0
	}

	return hash.nodes[index].key
}
