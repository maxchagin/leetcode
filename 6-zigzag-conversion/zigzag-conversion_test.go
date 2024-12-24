package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		// Test case 1: Standard example with 3 rows
		{
			name:     "Standard 3 rows",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		// Test case 2: Standard example with 4 rows
		{
			name:     "Standard 4 rows",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		// Test case 3: Single row (no zigzag)
		{
			name:     "Single row",
			s:        "ABCDEF",
			numRows:  1,
			expected: "ABCDEF",
		},
		// Test case 4: Two rows
		{
			name:     "Two rows",
			s:        "ABCDEF",
			numRows:  2,
			expected: "ACEBDF",
		},
		// Test case 5: Single character
		{
			name:     "Single character",
			s:        "A",
			numRows:  1,
			expected: "A",
		},
		// Test case 6: Number of rows larger than string length
		{
			name:     "Rows more than length",
			s:        "ABC",
			numRows:  4,
			expected: "ABC",
		},
		// Test case 7: Empty string
		{
			name:     "Empty string",
			s:        "",
			numRows:  3,
			expected: "",
		},
		// Test case 8: String with special characters
		{
			name:     "Special characters",
			s:        "A.B,C",
			numRows:  2,
			expected: "ABC.,",
		},
		// Test case 9: Long string with small rows
		{
			name:     "Long string small rows",
			s:        "ABCDEFGHIJKLMNOP",
			numRows:  2,
			expected: "ACEGIKMOBDFHJLNP",
		},
		// Test case 10: Mixed case letters
		{
			name:     "Mixed case letters",
			s:        "aBcDeFgH",
			numRows:  3,
			expected: "aeBDFHcg",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := convert(test.s, test.numRows); result != test.expected {
				t.Errorf("convert(%q, %d) = %q, want %q",
					test.s, test.numRows, result, test.expected)
			}
		})
	}
}
