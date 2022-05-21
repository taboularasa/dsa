package problems

func PeakOfMountain(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= arr[mid-1] && arr[mid] >= arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return -1
}
