package licensekeyformatting

import "testing"

func TestLicenseKeyFormatting(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{
			name: "Example 1 - Basic case with mixed case",
			s:    "5F3Z-2e-9-w",
			k:    4,
			want: "5F3Z-2E9W",
		},
		{
			name: "Example 2 - Groups of two",
			s:    "2-5g-3-J",
			k:    2,
			want: "2-5G-3J",
		},
		{
			name: "Single character input",
			s:    "a",
			k:    1,
			want: "A",
		},
		{
			name: "All lowercase to uppercase",
			s:    "abc-def-ghi",
			k:    3,
			want: "ABC-DEF-GHI",
		},
		{
			name: "Only numbers with dashes",
			s:    "2-4-6-8-0",
			k:    2,
			want: "24-68-0",
		},
		{
			name: "Long string with mixed characters",
			s:    "r4-p2-si-j9-q5-f0",
			k:    5,
			want: "R4-P2SIJ-9Q5F0",
		},
		{
			name: "String with consecutive dashes",
			s:    "aa--bb--cc",
			k:    2,
			want: "AA-BB-CC",
		},
		{
			name: "Single group larger than k",
			s:    "abcdefgh",
			k:    3,
			want: "AB-CDE-FGH",
		},
		{
			name: "All characters are dashes",
			s:    "----",
			k:    2,
			want: "",
		},
		{
			name: "Mixed alphanumeric with k=1",
			s:    "a1-b2-c3",
			k:    1,
			want: "A-1-B-2-C-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.s, tt.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
