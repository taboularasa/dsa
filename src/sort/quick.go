package sort

import (
	"math/rand"
	"time"

	"github.com/taboularasa/dsa/src/partition"
)

func QuickSort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, start, end int) {
	if start >= end {
		return
	}
	rand.Seed(time.Now().UnixNano())
	pivot := rand.Intn(end-start) + start
	nums[start], nums[pivot] = nums[pivot], nums[start]
	smaller := partition.Lomuto(nums, start)
	quickSortHelper(nums, start, smaller-1)
	quickSortHelper(nums, smaller+1, end)
}
