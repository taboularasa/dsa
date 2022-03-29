package problems

import (
	"sort"
)

func KthLargest(k int, initialStream, appendStream []int) []int {
	results := []int{}
	for _, e := range appendStream {
		initialStream = append(initialStream, e)
		sort.Slice(initialStream, func(i, j int) bool { return initialStream[i] < initialStream[j] })
		results = append(results, initialStream[len(initialStream)-k])
	}
	return results
}
