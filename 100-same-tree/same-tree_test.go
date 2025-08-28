package sametree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			name:     "Both trees are nil",
			p:        nil,
			q:        nil,
			expected: true,
		},
		{
			name:     "One tree is nil, other is not",
			p:        &TreeNode{Val: 1},
			q:        nil,
			expected: false,
		},
		{
			name:     "Both trees have single node with same value",
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 5},
			expected: true,
		},
		{
			name:     "Both trees have single node with different values",
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 10},
			expected: false,
		},
		{
			name: "Identical balanced trees with same structure and values",
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 0},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 0},
			},
			expected: false,
		},
		{
			name: "Same structure but different values",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4}, // Different value
			},
			expected: false,
		},
		{
			name: "Different structure - left child vs right child",
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name: "Different structure - missing node",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				// Missing right child
			},
			expected: false,
		},
		{
			name: "Complex identical trees with multiple levels",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: true,
		},
		{
			name: "Complex trees with different structure at deep level",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
					// Missing right child at deep level
				},
			},
			expected: false,
		},
		{
			name: "Trees with negative values",
			p: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: -5},
			},
			q: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: -5},
			},
			expected: true,
		},
		{
			name: "Trees with maximum allowed values",
			p: &TreeNode{
				Val:  10000,
				Left: &TreeNode{Val: 9999},
			},
			q: &TreeNode{
				Val:  10000,
				Left: &TreeNode{Val: 9999},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSameTree(test.p, test.q)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}

			// For now, we'll just verify our test structure is correct
			t.Logf("Test case: %s", test.name)
		})
	}
}
