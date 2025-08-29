package balancedbinarytree

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "Balanced tree with two levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			name: "Example 1: Balanced tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: true,
		},
		{
			name: "Example 2: Unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name: "Left-skewed unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Right-skewed unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Tree with height difference of 1 (balanced)",
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
			expected: true,
		},
		{
			name: "Tree with height difference of 2 (unbalanced)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
		{
			name: "Complex balanced tree",
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
			expected: true,
		},
		{
			name: "Complex balanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
						},
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
				},
			},
			expected: true,
		},
		{
			name: "Tree with negative values (balanced)",
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
			expected: true,
		},
		{
			name: "Tree with zero values (unbalanced)",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: false,
		},
		{
			name: "Tree with boundary values",
			root: &TreeNode{
				Val: 10000,
				Left: &TreeNode{
					Val: -10000,
					Left: &TreeNode{
						Val: 5000,
					},
					Right: &TreeNode{
						Val: -5000,
					},
				},
				Right: &TreeNode{
					Val: 20000,
				},
			},
			expected: true,
		},
		{
			name: "Large height difference (unbalanced)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalanced(tt.root)
			if result != tt.expected {
				t.Errorf("isBalanced() = %v, want %v", result, tt.expected)
			}
		})
	}
}
