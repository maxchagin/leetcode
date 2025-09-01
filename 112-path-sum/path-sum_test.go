package pathsum

import "testing"

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		target   int
		expected bool
	}{
		{
			name:     "Empty tree",
			root:     nil,
			target:   0,
			expected: false,
		},
		{
			name:     "Single node with matching target",
			root:     &TreeNode{Val: 5},
			target:   5,
			expected: true,
		},
		{
			name:     "Single node with different target",
			root:     &TreeNode{Val: 5},
			target:   10,
			expected: false,
		},
		{
			name: "Example 1: Valid path exists",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			target:   22,
			expected: true,
		},
		{
			name: "Example 2: No valid path",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			target:   5,
			expected: false,
		},
		{
			name: "Path exists in left subtree only",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			target:   6, // 1 + 2 + 3
			expected: true,
		},
		{
			name: "Path exists in right subtree only",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			target:   8, // 1 + 3 + 4
			expected: true,
		},
		{
			name: "Multiple paths, one matches",
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
				},
			},
			target:   8, // 1 + 3 + 4
			expected: true,
		},
		{
			name: "Negative values with positive target",
			root: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: -3,
				},
			},
			target:   -5, // -2 + -3
			expected: true,
		},
		{
			name: "Mixed positive and negative values",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: -2,
					},
				},
			},
			target:   13, // 10 + 5 - 2
			expected: true,
		},
		{
			name: "Zero target with non-zero path",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
			target:   0, // 1 - 1
			expected: true,
		},
		{
			name: "Zero target with non-zero path 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			target:   0,
			expected: false,
		},
		{
			name: "Deep tree with matching path",
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
			},
			target:   15, // 1 + 2 + 3 + 4 + 5
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := hasPathSum(test.root, test.target)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}

			// Placeholder for actual test execution
			t.Logf("Test case: %s", test.name)
		})
	}
}
