package problems

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KthLargest(k int, initialStream, appendStream []int) []int {
	h := IntHeap(initialStream)
	heap.Init(&h)

	results := []int{}
	for _, e := range appendStream {
		heap.Push(&h, e)
		for h.Len() > k {
			heap.Pop(&h)
		}
		results = append(results, h[0])
	}
	return results
}
