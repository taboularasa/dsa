package sort

func Merge(nums []int) []int {
	mergeSortHelper(nums, 0, len(nums)-1)
	return nums
}

func mergeSortHelper(nums []int, start, end int) {
	// base case
	if start == end {
		return
	}

	// take the midpoint for a pivot
	mid := (start + end) / 2
	i := start
	j := mid + 1

	// allocate a temporary array
	aux := []int{}

	// recurse the two halves
	mergeSortHelper(nums, start, mid)
	mergeSortHelper(nums, j, end)

	// merge the two sorted halves
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			aux = append(aux, nums[i])
			i++
		} else {
			aux = append(aux, nums[j])
			j++
		}
	}

	// append leftovers
	for i <= mid {
		aux = append(aux, nums[i])
		i++
	}

	for j <= end {
		aux = append(aux, nums[j])
		j++
	}

	// copy sub array back into original
	for k, l := start, 0; k <= end; k, l = k+1, l+1 {
		nums[k] = aux[l]
	}
}
