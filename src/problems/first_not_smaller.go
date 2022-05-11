package problems

func FirstNotSmaller(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	ret := left
	for left <= right {
		i := left + (right-left)/2
		if arr[i] >= target {
			ret = i
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return ret
}
