package allpossiblefullbinarytrees

import (
	"reflect"
	"testing"
)

func serialize(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	result := []interface{}{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	// Trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}

// Converts list of trees into list of serialized slices
func serializeForest(forest []*TreeNode) [][]interface{} {
	var res [][]interface{}
	for _, tree := range forest {
		res = append(res, serialize(tree))
	}
	return res
}

func TestAllPossibleFBT(t *testing.T) {
	tests := []struct {
		n        int
		expected [][]interface{}
	}{
		// // Test 1: smallest valid input
		// {1, [][]interface{}{{0}}},
		// // Test 2: only root and two children
		// {3, [][]interface{}{{0, 0, 0}}},
		// Test 3: n=5 small tree cases
		{5, [][]interface{}{
			{0, 0, 0, nil, nil, 0, 0},
			{0, 0, 0, 0, 0},
		}},
		// // Test 4: n=7 bigger case (multiple trees)
		// {7, [][]interface{}{
		// 	{0, 0, 0, nil, nil, 0, 0, nil, nil, 0, 0},
		// 	{0, 0, 0, nil, nil, 0, 0, 0, 0},
		// 	{0, 0, 0, 0, 0, 0, 0},
		// 	{0, 0, 0, 0, 0, nil, nil, nil, nil, 0, 0},
		// 	{0, 0, 0, 0, 0, nil, nil, 0, 0},
		// }},
		// // Test 5: even number of nodes should return empty
		// {2, [][]interface{}{}},
		// // Test 6: n=9, valid but more trees than before
		// // Not enumerating all, just checking non-empty
		// {9, nil},
		// // Test 7: n=11
		// {11, nil},
		// // Test 8: n=13
		// {13, nil},
		// // Test 9: n=15
		// {15, nil},
		// // Test 10: largest constraint (20) is even, should return empty
		// {20, [][]interface{}{}},
	}

	for idx, tt := range tests {
		result := allPossibleFBT(tt.n)
		serialized := serializeForest(result)

		if tt.expected != nil {
			if !reflect.DeepEqual(serialized, tt.expected) {
				t.Errorf("Test %d failed for n=%d: expected %v, got %v",
					idx+1, tt.n, tt.expected, serialized)
			}
		} else {
			// For larger n we only check non-emptiness when odd
			if tt.n%2 == 1 && len(serialized) == 0 {
				t.Errorf("Test %d failed for n=%d: expected non-empty result, got empty",
					idx+1, tt.n)
			}
			if tt.n%2 == 0 && len(serialized) != 0 {
				t.Errorf("Test %d failed for n=%d: expected empty result, got %v",
					idx+1, tt.n, serialized)
			}
		}
	}
}
