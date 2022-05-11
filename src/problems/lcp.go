package problems

import (
	"strconv"
	"strings"
)

func LCP(input string) []string {
	results := &[]string{}
	partial := &[]string{}
	in := strings.Split(input, "")
	helper(&in, partial, results, 0)
	return *results
}

func helper(input, partial, results *[]string, idx int) {
	if idx == len(*input) {
		*results = append(*results, strings.Join(*partial, ""))
		return
	}

	if _, err := strconv.Atoi((*input)[idx]); err == nil {
		*partial = append((*partial), (*input)[idx]) // push
		helper(input, partial, results, idx+1)
		*partial = (*partial)[:len(*partial)-1] // pop
	} else {
		*partial = append(*partial, strings.ToUpper((*input)[idx])) // push
		helper(input, partial, results, idx+1)
		*partial = (*partial)[:len(*partial)-1] // pop

		*partial = append(*partial, strings.ToLower((*input)[idx])) // push
		helper(input, partial, results, idx+1)
		*partial = (*partial)[:len(*partial)-1] // pop
	}
}
