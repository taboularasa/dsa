package problems

import "sort"

func OnlineMedian(stream []int) []int {
	tmp := []int{}
	ret := []int{}

	for _, e := range stream {
		tmp = append(tmp, e)
		sort.Ints(tmp)
		l := len(tmp)
		if l == 1 {
			ret = append(ret, e)
			continue
		}
		median := 0
		if l%2 == 0 {
			median = (tmp[l/2-1] + tmp[l/2]) / 2
		} else {
			median = tmp[l/2]
		}
		ret = append(ret, median)
	}
	return ret
}
