package sort

func MergeSort(nums []int) []int {
	mergeSortHelper(nums, 0, len(nums)-1)
	return nums
}

func mergeSortHelper(nums []int, start, end int) {
	// leaf worker
	if start == end {
		return
	}
	// internal node worker
	mid := (start + end) / 2
	i := start
	j := mid + 1
	aux := []int{}

	mergeSortHelper(nums, start, mid)
	mergeSortHelper(nums, j, end)

	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			aux = append(aux, nums[i])
			i++
		} else {
			aux = append(aux, nums[j])
			j++
		}
	}
	// gather phase
	for i <= mid {
		aux = append(aux, nums[i])
		i++
	}

	for j <= end {
		aux = append(aux, nums[j])
		j++
	}

	for k, l := start, 0; k <= end; k, l = k+1, l+1 {
		nums[k] = aux[l]
	}
}
