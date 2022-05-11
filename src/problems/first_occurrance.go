package problems

func FirstOccurance(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	ret := -1
	for left <= right {
		i := left + (right-left)/2
		if arr[i] == target {
			ret = i
			right = i - 1
		} else if arr[i] > target {
			right = i - 1
		} else if arr[i] < target {
			left = i + 1
		}
	}
	return ret
}
