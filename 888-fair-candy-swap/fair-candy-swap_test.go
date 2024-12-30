package faircandyswap

import (
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	tests := []struct {
		name       string
		aliceSizes []int
		bobSizes   []int
		expected   []int
	}{
		{
			// Test case 1: Simple case with equal length arrays
			name:       "Equal length arrays",
			aliceSizes: []int{1, 1},
			bobSizes:   []int{2, 2},
			expected:   []int{1, 2},
		},
		{
			// Test case 2: Different length arrays
			name:       "Different length arrays",
			aliceSizes: []int{2},
			bobSizes:   []int{1, 3},
			expected:   []int{2, 3},
		},
		{
			// Test case 3: Large difference in total candies
			name:       "Large difference in totals",
			aliceSizes: []int{1, 2, 5},
			bobSizes:   []int{2, 4},
			expected:   []int{5, 4},
		},
		{
			// Test case 4: Multiple possible solutions (testing one valid solution)
			name:       "Multiple possible solutions",
			aliceSizes: []int{1, 1, 2},
			bobSizes:   []int{2, 2, 2},
			expected:   []int{1, 2},
		},
		{
			// Test case 5: Large numbers
			name:       "Large numbers",
			aliceSizes: []int{35, 17, 4, 24, 10},
			bobSizes:   []int{63, 21},
			expected:   []int{24, 21},
		},
		{
			// Test case 6: Small difference between totals
			name:       "Small difference between totals",
			aliceSizes: []int{1, 2, 3},
			bobSizes:   []int{2, 2, 2},
			expected:   []int{3, 2},
		},
		{
			// Test case 8: Sequential numbers
			name:       "Sequential numbers",
			aliceSizes: []int{1, 2, 3, 4},
			bobSizes:   []int{2, 3, 4, 5},
			expected:   []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fairCandySwap(tt.aliceSizes, tt.bobSizes)
			// Check if the result is valid
			if !isValidSwap(tt.aliceSizes, tt.bobSizes, result) {
				t.Errorf("fairCandySwap(%v, %v) = %v, want %v",
					tt.aliceSizes, tt.bobSizes, result, tt.expected)
			}
		})
	}
}

// Helper function to verify if the swap results in equal total candies
func isValidSwap(alice, bob, swap []int) bool {
	if len(swap) != 2 {
		return false
	}

	// Calculate initial sums
	aliceSum, bobSum := sum(alice), sum(bob)

	// Find the swapped values in arrays
	foundAlice := false
	foundBob := false
	for _, v := range alice {
		if v == swap[0] {
			foundAlice = true
			break
		}
	}
	for _, v := range bob {
		if v == swap[1] {
			foundBob = true
			break
		}
	}

	if !foundAlice || !foundBob {
		return false
	}

	// Calculate new sums after swap
	newAliceSum := aliceSum - swap[0] + swap[1]
	newBobSum := bobSum - swap[1] + swap[0]

	return newAliceSum == newBobSum
}

// Helper function to calculate sum of array
func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
