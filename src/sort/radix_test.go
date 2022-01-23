package sort

import "testing"

func TestRadix(t *testing.T) {
	t.Run("it sorts all the elements", func(t *testing.T) {
		a := []int{2917, 5569, 3594, 9101, 9728, 6331, 8825, 2147, 6613, 8433, 8482, 2620, 19, 4680, 4224}
		expected := []int{19, 2147, 2620, 2917, 3594, 4224, 4680, 5569, 6331, 6613, 8433, 8482, 8825, 9101, 9728}
		a = Radix(a)
		for i := range a {
			if a[i] != expected[i] {
				t.Errorf("array not sorted: %v", a)
				t.FailNow()
			}
		}
	})
}
