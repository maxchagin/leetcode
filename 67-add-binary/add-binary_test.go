package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "Example 1 - basic addition with carry",
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			name: "Example 2 - longer numbers",
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
		{
			name: "Both zeros",
			a:    "0",
			b:    "0",
			want: "0",
		},
		{
			name: "Different lengths with carry",
			a:    "1111",
			b:    "1",
			want: "10000",
		},
		{
			name: "Single digit with carry",
			a:    "1",
			b:    "1",
			want: "10",
		},
		{
			name: "One empty string and one digit",
			a:    "",
			b:    "1",
			want: "1",
		},
		{
			name: "Multiple carries",
			a:    "1111",
			b:    "1111",
			want: "11110",
		},
		{
			name: "Long number with no carry",
			a:    "1010",
			b:    "0101",
			want: "1111",
		},
		{
			name: "One zero and one binary number",
			a:    "0",
			b:    "1111",
			want: "1111",
		},
		{
			name: "Complex case with multiple carries",
			a:    "10101010",
			b:    "11001100",
			want: "101110110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
