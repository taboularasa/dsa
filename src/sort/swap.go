package sort

func Swap(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		minVal := nums[i]
		minIndex := i
		for j := i + 1; j < l; j++ {
			if nums[j] < minVal {
				minVal = nums[j]
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}
