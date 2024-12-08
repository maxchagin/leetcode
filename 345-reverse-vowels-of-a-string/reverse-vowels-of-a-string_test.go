package reversevowelsofastring

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		// Test with a string containing both uppercase and lowercase vowels
		{input: "IceCreAm", output: "AceCreIm"},

		// Test with a commonly used word
		{input: "leetcode", output: "leotcede"},

		// Test with a word containing only one vowel
		{input: "sky", output: "sky"},

		// Test with a single character being a vowel
		{input: "a", output: "a"},

		// Test with alternating upper and lower case vowels
		{input: "AeiOu", output: "uOieA"},

		// Test with repeating vowels
		{input: "Book", output: "Book"},

		// Test with no vowels present
		{input: "rhythm", output: "rhythm"},

		// Test with vowels at the beginning and end
		{input: "umbrella", output: "ambrellu"},

		// Test with all vowels and consonants
		{input: "aeioubcdf", output: "uoieabcdf"},

		// Test with very long string consisting of same characters
		{input: "aaaaa" + "bcdefghijklmnopqrstuvwxyz" + "ooooo", output: "ooooo" + "bcdufghojklmnipqrstevwxyz" + "aaaaa"},

		{input: "u'o;B,vKO\"?,.;?fGI 9;`9T`Z,;iqsn4A.:;'H3s8E9. B4TxfOiB'wvM?q'k:Q`J\"E1T7lo e7c!?glI66?JZh\"6 !C,aK! 1J?f9'`SX4Q!Q4XS`'9f?J1 !Ka,C! 6\"hZJ?66Ilg?!c7e ol7T1E\"J`Q:k'q?Mvw'BiOfxT4B .9E8s3H';:.A4nsqi;,Z`To`;o IGf?;.,?\"OKv,B;o'u", output: "u'o;B,vKO\"?,.;?fGI 9;`9T`Z,;oqsn4o.:;'H3s8i9. B4TxfAEB'wvM?q'k:Q`J\"O1T7li E7c!?glo66?JZh\"6 !C,eK! 1J?f9'`SX4Q!Q4XS`'9f?J1 !KI,C! 6\"hZJ?66alg?!c7a Il7T1e\"J`Q:k'q?Mvw'BoEfxT4B .9i8s3H';:.O4nsqE;,Z`TA`;i IGf?;.,?\"OKv,B;o'u"},
	}

	for _, test := range tests {
		result := reverseVowels(test.input)
		if result != test.output {
			t.Errorf("For input %q, expected %q, but got %q", test.input, test.output, result)
		}
	}
}
