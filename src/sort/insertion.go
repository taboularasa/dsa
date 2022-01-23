package sort

func InsertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		red := i - 1
		for red >= 0 && nums[red] > temp {
			nums[red+1] = nums[red]
			red--
		}
		nums[red+1] = temp
	}
	return nums
}
