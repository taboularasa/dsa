package problems

func MinMaxWeight(weights []int, d int) int {
	maxWeight := -1
	for i := 0; i < len(weights); i++ {
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}
	minPtr := maxWeight
	maxPtr := 0
	for i := 0; i < len(weights); i++ {
		maxPtr += weights[i]
	}
	ret := maxPtr
	for minPtr <= maxPtr {
		mid := minPtr + (maxPtr-minPtr)/2
		if feasible(weights, mid, d) {
			ret = mid
			maxPtr = mid - 1
		} else {
			minPtr = mid + 1
		}
	}

	return ret
}

func feasible(weights []int, maxWeight int, d int) bool {
	days := 1
	capacity := maxWeight
	n := len(weights)
	i := 0
	for i < n {
		if weights[i] <= capacity {
			capacity -= weights[i]
			i += 1
		} else {
			days += 1
			capacity = maxWeight
		}
	}
	return days <= d
}
