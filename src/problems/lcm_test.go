package problems

import (
	"fmt"
	"testing"
)

func TestLCM(t *testing.T) {
	input := []string{"a", "1", "b", "2"}
	res := LCM(input)
	fmt.Println(res)
	t.FailNow()
}
