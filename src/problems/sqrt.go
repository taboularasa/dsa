package problems

func Sqrt(n int) int {
	left := 1
	right := n
	ret := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= n {
			ret = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ret
}
