package intersectionoftwoarraysii

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}

	for i, test := range tests {
		result := intersect(test.nums1, test.nums2)
		sort.Ints(result)
		sort.Ints(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test case %d failed: intersect(%v, %v) = %v; expected %v",
				i+1, test.nums1, test.nums2, result, test.expected)
		}
	}
}
