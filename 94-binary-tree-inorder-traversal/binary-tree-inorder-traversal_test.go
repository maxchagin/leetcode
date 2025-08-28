package binarytreeinordertraversal

import "testing"

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Left child only",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{2, 1},
		},
		{
			name: "Right child only",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{1, 2},
		},
		{
			name: "Example 1: [1,null,2,3]",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name: "Complete binary tree with 3 nodes",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: []int{2, 1, 3},
		},
		{
			name: "Complex tree with multiple levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Right-skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Tree with negative values",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -1,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			expected: []int{-1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := inorderTraversal(tt.root)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], result[i])
				}
			}
		})
	}
}
