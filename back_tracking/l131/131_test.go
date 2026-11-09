package l131

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "example one",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "example two",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			name: "no longer palindromes",
			s:    "abc",
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "even length palindrome",
			s:    "abba",
			want: [][]string{{"a", "b", "b", "a"}, {"a", "bb", "a"}, {"abba"}},
		},
		{
			name: "odd length palindrome",
			s:    "aba",
			want: [][]string{{"a", "b", "a"}, {"aba"}},
		},
		{
			name: "repeated characters",
			s:    "aaa",
			want: [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}},
		},
		{
			name: "multiple palindrome groups",
			s:    "aabb",
			want: [][]string{{"a", "a", "b", "b"}, {"aa", "b", "b"}, {"a", "a", "bb"}, {"aa", "bb"}},
		},
		{
			name: "overlapping palindromes",
			s:    "abab",
			want: [][]string{{"a", "b", "a", "b"}, {"a", "bab"}, {"aba", "b"}},
		},
	}

	// Partition order is arbitrary, but substring order within each partition
	// must be preserved. Keep multiplicities to detect duplicate partitions.
	counts := func(partitions [][]string) map[string]int {
		result := make(map[string]int)
		for _, parts := range partitions {
			key, err := json.Marshal(parts)
			if err != nil {
				t.Fatalf("encode partition: %v", err)
			}
			result[string(key)]++
		}
		return result
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.s)
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("partition(%q) = %v, want %v (in any partition order)", tt.s, got, tt.want)
			}
		})
	}
}

func TestPartitionIndependentResults(t *testing.T) {
	got := partition("aaa")
	if len(got) != 4 {
		t.Fatalf("got %d partitions, want 4", len(got))
	}

	before := make([][]string, len(got))
	for i, parts := range got {
		if len(parts) == 0 {
			t.Fatalf("partition %d is empty for a nonempty string", i)
		}
		before[i] = append([]string(nil), parts...)
	}

	for i, parts := range got {
		original := parts[0]
		parts[0] = "changed"
		for j := range got {
			if j != i && !reflect.DeepEqual(got[j], before[j]) {
				t.Errorf("mutating partition %d changed partition %d: got %v, want %v", i, j, got[j], before[j])
			}
		}
		parts[0] = original
	}
}
