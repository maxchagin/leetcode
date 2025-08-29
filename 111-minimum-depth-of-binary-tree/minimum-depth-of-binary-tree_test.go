package minimumdepthofbinarytree

import "testing"

func TestMinDepth(t *testing.T) {
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
			name:     "Single node tree (leaf)",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "Tree with only left child",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: 2,
		},
		{
			name: "Tree with only right child",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: 2,
		},
		{
			name: "Balanced tree with both children",
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
			name: "Tree where left subtree has shorter path",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 2,
		},
		{
			name: "Tree where right subtree has shorter path",
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
			expected: 2,
		},
		{
			name: "Right-skewed tree (Example 2 from problem)",
			root: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expected: 5,
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expected: 5,
		},
		{
			name: "Complex tree with multiple leaf paths",
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
							Val: 8,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: 3, // Shortest path: 1 -> 3 -> 6
		},
		{
			name: "Tree with node having one null child",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 3, // Shortest path: 1 -> 2 -> 4
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
			expected: 2, // Shortest path: -10 -> -5
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
			expected: 2, // Shortest path: 0 -> 0 (right child)
		},
		{
			name: "Tree with boundary values",
			root: &TreeNode{
				Val: 1000,
				Left: &TreeNode{
					Val: -1000,
				},
				Right: &TreeNode{
					Val: 500,
					Left: &TreeNode{
						Val: 250,
					},
				},
			},
			expected: 2, // Shortest path: 1000 -> -1000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDepth(tt.root)
			if result != tt.expected {
				t.Errorf("minDepth() = %v, want %v", result, tt.expected)
			}
		})
	}
}
