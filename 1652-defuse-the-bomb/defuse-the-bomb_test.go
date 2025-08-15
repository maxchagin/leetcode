package defusethebomb

import "testing"

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name string
		code []int
		k    int
		want []int
	}{
		{
			name: "Example 1: k > 0",
			code: []int{5, 7, 1, 4},
			k:    3,
			want: []int{12, 10, 16, 13},
		},
		{
			name: "Example 2: k == 0",
			code: []int{1, 2, 3, 4},
			k:    0,
			want: []int{0, 0, 0, 0},
		},
		{
			name: "Example 3: k < 0",
			code: []int{2, 4, 9, 3},
			k:    -2,
			want: []int{12, 5, 6, 13},
		},
		{
			name: "Single element array with k > 0",
			code: []int{5},
			k:    1,
			want: []int{5},
		},
		{
			name: "Single element array with k == 0",
			code: []int{5},
			k:    0,
			want: []int{0},
		},
		{
			name: "Single element array with k < 0",
			code: []int{5},
			k:    -1,
			want: []int{5},
		},
		{
			name: "All zeros with k > 0",
			code: []int{0, 0, 0, 0},
			k:    2,
			want: []int{0, 0, 0, 0},
		},
		{
			name: "Mixed positive and negative numbers with k > 0",
			code: []int{-1, 2, -3, 4},
			k:    2,
			want: []int{-1, 1, 3, 1},
		},
		{
			name: "Mixed positive and negative numbers with k < 0",
			code: []int{-1, 2, -3, 4},
			k:    -2,
			want: []int{1, 3, 1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decrypt(tt.code, tt.k)
			if !equal(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to compare two slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
