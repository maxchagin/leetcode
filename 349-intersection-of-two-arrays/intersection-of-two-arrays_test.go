package intersectionoftwoarrays

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
	}

	for i, test := range tests {
		result := intersection(test.nums1, test.nums2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test case %d failed: intersection(%v, %v) = %v; expected %v",
				i+1, test.nums1, test.nums2, result, test.expected)
		}
	}
}
