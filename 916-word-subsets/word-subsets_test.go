package wordsubsets

import (
	"reflect"
	"sort"
	"testing"
)

func TestWordSubsets(t *testing.T) {
	tests := []struct {
		name   string
		words1 []string
		words2 []string
		want   []string
	}{
		{
			name:   "Example 1 - basic case with vowels",
			words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
			words2: []string{"e", "o"},
			want:   []string{"facebook", "google", "leetcode"},
		},
		{
			name:   "Example 2 - basic case with consonant and vowel",
			words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
			words2: []string{"l", "e"},
			want:   []string{"apple", "google", "leetcode"},
		},
		{
			name:   "Empty words2 array - all words1 should be included",
			words1: []string{"cat", "dog", "rat"},
			words2: []string{},
			want:   []string{"cat", "dog", "rat"},
		},
		{
			name:   "Single character subset",
			words1: []string{"aaa", "bbb", "ccc"},
			words2: []string{"a"},
			want:   []string{"aaa"},
		},
		{
			name:   "Multiple character subset with repetition",
			words1: []string{"bella", "label", "roem"},
			words2: []string{"ll"},
			want:   []string{"bella", "label"},
		},
		{
			name:   "Multiple constraints test",
			words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
			words2: []string{"e", "oo"},
			want:   []string{"facebook", "google"},
		},
		{
			name:   "All words are universal",
			words1: []string{"aaa", "aaaa", "aaaaa"},
			words2: []string{"a"},
			want:   []string{"aaa", "aaaa", "aaaaa"},
		},
		{
			name:   "Complex character frequency test",
			words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
			words2: []string{"lo", "eo"},
			want:   []string{"google", "leetcode"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordSubsets(tt.words1, tt.words2)
			// Sort both slices to ensure consistent comparison
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
