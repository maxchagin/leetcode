package maximumdepthofbinarytree

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: 3,
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
			expected: 3,
		},
		{
			name: "Balanced tree with equal depth",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 2,
		},
		{
			name: "Tree with deeper left subtree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 3,
		},
		{
			name: "Tree with deeper right subtree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			expected: 4,
		},
		{
			name: "Complex tree with multiple levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			expected: 4,
		},
		{
			name: "Tree with negative values",
			root: &TreeNode{
				Val: -10,
				Left: &TreeNode{
					Val: -5,
				},
				Right: &TreeNode{
					Val: -15,
					Left: &TreeNode{
						Val: -20,
					},
				},
			},
			expected: 3,
		},
		{
			name: "Tree with zero values",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 3,
		},
		{
			name: "Tree with maximum constraint values",
			root: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: -100,
					Right: &TreeNode{
						Val: 100,
					},
				},
				Right: &TreeNode{
					Val: -100,
				},
			},
			expected: 3,
		},
		{
			name: "Tree where left and right subtrees have same depth",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDepth(tt.root)
			if result != tt.expected {
				t.Errorf("maxDepth() = %v, want %v", result, tt.expected)
			}
		})
	}
}
