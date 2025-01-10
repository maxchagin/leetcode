package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "Example 1 - first three numbers",
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "Example 2 - first five numbers",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "Example 3 - first fifteen numbers with FizzBuzz",
			n:    15,
			want: []string{
				"1", "2", "Fizz", "4", "Buzz",
				"Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
		{
			name: "Minimum input - single number",
			n:    1,
			want: []string{"1"},
		},
		{
			name: "Two numbers - no special output",
			n:    2,
			want: []string{"1", "2"},
		},
		{
			name: "First multiple of both 3 and 5",
			n:    16,
			want: []string{
				"1", "2", "Fizz", "4", "Buzz",
				"Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz", "16",
			},
		},
		{
			name: "Section with multiple Fizz",
			n:    6,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz"},
		},
		{
			name: "Section ending with Buzz",
			n:    10,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
		},
		{
			name: "Numbers through 12 (multiple Fizz)",
			n:    12,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz"},
		},
		{
			name: "Numbers through 20 (multiple special cases)",
			n:    20,
			want: []string{
				"1", "2", "Fizz", "4", "Buzz",
				"Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
				"16", "17", "Fizz", "19", "Buzz",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
