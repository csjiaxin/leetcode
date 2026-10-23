package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int64
}

type Item struct {
	Node int
	Dist int64
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Dijkstra returns the shortest distance from src to every node.
// Unreachable nodes have distance math.MaxInt64.
// All edge weights must be nonnegative.
func Dijkstra(graph [][]Edge, src int) []int64 {
	dist := make([]int64, len(graph))
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[src] = 0

	pq := &MinHeap{{Node: src, Dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)

		// A better distance was found after this item was queued.
		if cur.Dist != dist[cur.Node] {
			continue
		}

		for _, edge := range graph[cur.Node] {
			// Prevent integer overflow.
			if edge.Weight > math.MaxInt64-cur.Dist {
				continue
			}

			nextDist := cur.Dist + edge.Weight
			if nextDist < dist[edge.To] {
				dist[edge.To] = nextDist
				heap.Push(pq, Item{
					Node: edge.To,
					Dist: nextDist,
				})
			}
		}
	}

	return dist
}

func main() {
	// Directed graph. For undirected graphs, add each edge both ways.
	graph := [][]Edge{
		{{To: 1, Weight: 4}, {To: 2, Weight: 1}},
		{{To: 3, Weight: 1}},
		{{To: 1, Weight: 2}, {To: 3, Weight: 5}},
		{},
	}

	fmt.Println(Dijkstra(graph, 0)) // [0 3 1 4]
}
