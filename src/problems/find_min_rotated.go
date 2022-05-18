package problems

func FindMinRotated(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	left := 0
	right := len(arr) - 1
	ret := -1
	pivot := arr[len(arr)-1]
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < pivot {
			ret = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ret
}
