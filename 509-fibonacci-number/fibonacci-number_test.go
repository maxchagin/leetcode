package fibonaccinumber

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		n        int
		expected int
		name     string
	}{
		{
			n:        0,
			expected: 0,
			name:     "Base case F(0)",
		},
		{
			n:        1,
			expected: 1,
			name:     "Base case F(1)",
		},
		{
			n:        2,
			expected: 1,
			name:     "Fibonacci of 2",
		},
		{
			n:        3,
			expected: 2,
			name:     "Fibonacci of 3",
		},
		{
			n:        4,
			expected: 3,
			name:     "Fibonacci of 4",
		},
		{
			n:        5,
			expected: 5,
			name:     "Fibonacci of 5",
		},
		{
			n:        10,
			expected: 55,
			name:     "Fibonacci of 10",
		},
		{
			n:        15,
			expected: 610,
			name:     "Fibonacci of 15",
		},
		{
			n:        20,
			expected: 6765,
			name:     "Fibonacci of 20",
		},
		{
			n:        30,
			expected: 832040,
			name:     "Fibonacci of 30, larger value test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fib(tt.n)
			if result != tt.expected {
				t.Errorf("For n=%d, expected %d but got %d", tt.n, tt.expected, result)
			}
		})
	}
}
