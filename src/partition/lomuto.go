package partition

func Lomuto(nums []int, start int) int {
	smaller := start
	for bigger := start + 1; bigger < len(nums); bigger++ {
		if nums[bigger] < nums[start] {
			smaller++
			nums[smaller], nums[bigger] = nums[bigger], nums[smaller]
		}
	}
	nums[start], nums[smaller] = nums[smaller], nums[start]
	return smaller
}
