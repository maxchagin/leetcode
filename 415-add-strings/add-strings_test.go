package addstrings

import "testing"

func TestAddStrings(t *testing.T) {
	tests := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{
			name: "Example 1 - different length numbers",
			num1: "11",
			num2: "123",
			want: "134",
		},
		{
			name: "Example 2 - three digit plus two digit",
			num1: "456",
			num2: "77",
			want: "533",
		},
		{
			name: "Example 3 - both zeros",
			num1: "0",
			num2: "0",
			want: "0",
		},
		{
			name: "Addition with carry",
			num1: "999",
			num2: "1",
			want: "1000",
		},
		{
			name: "Large numbers",
			num1: "1234567890",
			num2: "9876543210",
			want: "11111111100",
		},
		{
			name: "Single digit numbers",
			num1: "5",
			num2: "7",
			want: "12",
		},
		{
			name: "One empty and one digit",
			num1: "9",
			num2: "",
			want: "9",
		},
		{
			name: "Numbers with multiple carries",
			num1: "999999",
			num2: "999999",
			want: "1999998",
		},
		{
			name: "Add zero to number",
			num1: "500",
			num2: "0",
			want: "500",
		},
		{
			name: "Different length with multiple carries",
			num1: "98765",
			num2: "1234",
			want: "99999",
		},
		{
			name: "Big",
			num1: "3876620623801494171",
			num2: "6529364523802684779",
			want: "10405985147604178950",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.num1, tt.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
