package symmetrictree

import "testing"

// TestIsSymmetric tests the isSymmetric function with various test cases
func TestIsSymmetric(t *testing.T) {
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
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "Symmetric tree with two levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: true,
		},
		{
			name: "Asymmetric tree with two levels - different values",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
		{
			name: "Symmetric tree with three levels - Example 1 from problem",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: true,
		},
		{
			name: "Asymmetric tree with three levels - Example 2 from problem",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: false,
		},
		{
			name: "Symmetric tree with single child nodes",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: true,
		},
		{
			name: "Complex symmetric tree with multiple levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: true,
		},
		{
			name: "Leetcode 1xx",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			expected: true,
		},
		{
			name: "Complex asymmetric tree with same values but different structure",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: false,
		},
		{
			name: "Tree with negative values - symmetric",
			root: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: -2,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			expected: true,
		},
		{
			name: "Tree with negative values - asymmetric",
			root: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: -2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name:     "Large symmetric tree (maximum constraint)",
			root:     createLargeSymmetricTree(),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSymmetric(tt.root)
			if result != tt.expected {
				t.Errorf("isSymmetric() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to create a large symmetric tree for testing
func createLargeSymmetricTree() *TreeNode {
	// Create a symmetric tree with multiple levels
	// This simulates a tree near the constraint limit (1000 nodes)
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
}
