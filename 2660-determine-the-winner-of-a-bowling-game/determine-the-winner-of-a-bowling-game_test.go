package determinethewinnerofabowlinggame

import "testing"

func TestIsWinner(t *testing.T) {
	tests := []struct {
		name    string
		player1 []int
		player2 []int
		want    int
	}{
		{
			name:    "Example 1 - Player 1 wins with double points",
			player1: []int{5, 10, 3, 2},
			player2: []int{6, 5, 7, 3},
			want:    1,
		},
		{
			name:    "Example 2 - Player 2 wins with consecutive strikes",
			player1: []int{3, 5, 7, 6},
			player2: []int{8, 10, 10, 2},
			want:    2,
		},
		{
			name:    "Example 3 - Draw game with short arrays",
			player1: []int{2, 3},
			player2: []int{4, 1},
			want:    0,
		},
		{
			name:    "Example 4 - Player 2 wins with multiple strikes",
			player1: []int{1, 1, 1, 10, 10, 10, 10},
			player2: []int{10, 10, 10, 10, 1, 1, 1},
			want:    2,
		},
		{
			name:    "Single turn - Player 1 wins",
			player1: []int{10},
			player2: []int{5},
			want:    1,
		},
		{
			name:    "All zeros - Draw game",
			player1: []int{0, 0, 0},
			player2: []int{0, 0, 0},
			want:    0,
		},
		{
			name:    "Alternating strikes - Player 2 wins",
			player1: []int{10, 0, 10, 0},
			player2: []int{0, 10, 0, 10},
			want:    0,
		},
		{
			name:    "Equal scores with different patterns",
			player1: []int{5, 5, 5, 5},
			player2: []int{10, 0, 10, 0},
			want:    2,
		},
		{
			name:    "Long sequence with strikes",
			player1: []int{10, 10, 10, 5, 5, 5},
			player2: []int{5, 5, 5, 10, 10, 10},
			want:    1,
		},
		{
			name:    "Complex case with mixed scores",
			player1: []int{8, 10, 9, 2, 10, 3},
			player2: []int{10, 2, 8, 10, 3, 9},
			want:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWinner(tt.player1, tt.player2); got != tt.want {
				t.Errorf("isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
