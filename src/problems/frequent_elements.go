package problems

import (
	"container/heap"
)

type IntCount struct {
	val   int
	count int
}

type IntCountHeap []IntCount

func (h IntCountHeap) Len() int           { return len(h) }
func (h IntCountHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h IntCountHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntCountHeap) Push(x any) {
	*h = append(*h, x.(IntCount))
}
func (h *IntCountHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FrequentElements(arr []int, k int) []int {
	c := map[int]int{}
	for _, e := range arr {
		c[e] += 1
	}

	ich := &IntCountHeap{}
	for k, v := range c {
		ic := IntCount{
			val:   k,
			count: v,
		}
		*ich = append(*ich, ic)
	}
	heap.Init(ich)

	ret := []int{}
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(ich).(IntCount).val)
	}

	return ret
}
