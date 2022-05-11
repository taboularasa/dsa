package problems

func NChooseK(n, k int) [][]int {
	vals := []int{}
	for i := 1; i <= n; i++ {
		vals = append(vals, i)
	}
	ret := [][]int{}
	tmp := []int{}
	nkHelper(k, vals, tmp, &ret)
	return ret
}

func nkHelper(k int, vals, tmp []int, ret *[][]int) {
	if k == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*ret = append(*ret, t)
		return
	}

	if len(vals) < 1 {
		return
	}

	t := make([]int, len(tmp))
	copy(t, tmp)
	t = append(t, vals[0])
	vals = vals[1:]

	nkHelper(k-1, vals, t, ret) // include element
	nkHelper(k, vals, tmp, ret) // exclude element
}
