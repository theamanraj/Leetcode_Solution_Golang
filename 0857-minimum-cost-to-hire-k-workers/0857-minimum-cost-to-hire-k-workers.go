package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	works := make([][2]int, len(quality))
	for i := 0; i < len(quality); i++ {
		works[i] = [2]int{wage[i], quality[i]}
	}
	w := Works(works)
	sort.Sort(w)
	pq := &IntHeap{}
	heap.Init(pq)
	res, sumq := math.MaxFloat64, 0
	for i, work := range w {
		heap.Push(pq, work[1])
		sumq += work[1]
		if pq.Len() > K {
			sumq -= heap.Pop(pq).(int)
		}
		if pq.Len() == K && res > float64(sumq)*w.Ratio(i) {
			res = float64(sumq) * w.Ratio(i)
		}
	}
	return res
}

type Works [][2]int

func (w Works) Len() int {
	return len(w)
}

func (w Works) Less(i, j int) bool {
	return w.Ratio(i) < w.Ratio(j)
}

func (w Works) Swap(i, j int) {
	w[i][0], w[j][0] = w[j][0], w[i][0]
	w[i][1], w[j][1] = w[j][1], w[i][1]
}
func (w Works) Ratio(i int) float64 {
	return float64(w[i][0]) / float64(w[i][1])
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}