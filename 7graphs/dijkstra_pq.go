package _graphs

import (
	"container/heap"
	"math"
)

type item struct {
	node int
	dist int
	idx  int
}
type PriorityQueue []*item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	v := *pq
	return v[i].dist < v[j].dist
}

func (pq *PriorityQueue) Swap(i, j int) {
	v := *pq
	v[i], v[j] = v[j], v[i]
	v[i].idx, v[j].idx = i, j
}

func (pq *PriorityQueue) Push(val any) {
	i := val.(*item)
	i.idx = len(*pq)
	*pq = append(*pq, i)

}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	i := old[len(old)-1]
	*pq = old[:len(old)-1]
	return i
}

func dijkstraPq(source int, sink int, arr WeightedAdjacencyList) []int {
	dist := make([]int, len(arr))
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[source] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{node: source, dist: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*item)

		// Skip if we already found a better path to this node
		if curr.dist > dist[curr.node] {
			continue
		}

		// Relax all edges from current node
		for _, edge := range arr[curr.node] {
			newDist := dist[curr.node] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(pq, &item{node: edge.to, dist: newDist})
			}
		}
	}

	return dist
}
