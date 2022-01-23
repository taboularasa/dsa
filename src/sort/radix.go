package sort

import (
	"math"
	"strconv"
)

func Radix(a []int) []int {
	// swap the largest element to the end
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}

	// find the number of decimals of the largest element
	w := len(strconv.Itoa(a[len(a)-1]))

	// count values at each radix
	for i := 0; i < w; i++ {
		tmp := [10]int{0}
		base := int(math.Pow10(i))
		for j := 0; j < len(a); j++ {
			e := a[j]
			x := (e / base) % 10
			tmp[x]++
		}
		for j := 0; j < len(tmp)-1; j++ {
			tmp[j+1] += tmp[j]
		}
		r := make([]int, len(a))
		for j := len(a) - 1; j >= 0; j-- {
			e := a[j]
			pos := tmp[(e/base)%10]
			r[pos-1] = e
			tmp[(e/base)%10]--
		}
		a = r
	}

	return a
}
