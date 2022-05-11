package problems

func GetPermutations(arr []int) [][]int {
	ret := [][]int{}
	permuteHelper(arr, []int{}, &ret)
	return ret
}

func permuteHelper(arr, tmp []int, ret *[][]int) {
	if len(arr) == 0 {
		r := make([]int, len(tmp))
		copy(r, tmp)
		*ret = append(*ret, r)
		return
	}

	for i, e := range arr {
		a := make([]int, len(arr))
		copy(a, arr)
		a = append(a[:i], a[i+1:]...)
		t := append(tmp, e)
		permuteHelper(a, t, ret)
	}
}
