package problems

func FindIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	i, j, k := 0, 0, 0
	ret := []int{}
	for {
		if i > len(arr1)-1 || j > len(arr2)-1 || k > len(arr3)-1 {
			break
		}

		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			ret = append(ret, arr1[i])
		}

		arr1i := arr1[i]
		arr2j := arr2[j]
		arr3k := arr3[k]

		if arr1i <= arr2j && arr1i <= arr3k {
			i++
		}
		if arr2j <= arr1i && arr2j <= arr3k {
			j++
		}
		if arr3k <= arr1i && arr3k <= arr2j {
			k++
		}
	}
	return ret
}
