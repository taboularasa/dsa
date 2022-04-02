package problems

import "sort"

func OnlineMedian(stream []int) []int {
	tmp := []int{}
	ret := []int{}

	for _, e := range stream {
		if len(tmp) == 0 {
			tmp = append(tmp, e)
			ret = append(ret, e)
			continue
		}

		i := sort.SearchInts(tmp, e)
		if i < len(tmp) {
			tmp = append(tmp[:i+1], tmp[i:]...)
			tmp[i] = e
		} else {
			tmp = append(tmp, e)
		}

		median := 0

		switch l := len(tmp); {
		case l == 2:
			median = (tmp[0] + tmp[1]) / 2
		case l%2 == 0:
			median = (tmp[l/2-1] + tmp[l/2]) / 2
		default:
			median = tmp[l/2]
		}

		ret = append(ret, median)
	}

	return ret
}
