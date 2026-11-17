package l1143

import "testing"

func TestLongestCommonSubsequenceEmptyPrefixes(t *testing.T) {
	for _, pair := range [][2]string{{"", ""}, {"", "abc"}, {"abc", ""}} {
		if got := longestCommonSubsequence(pair[0], pair[1]); got != 0 {
			t.Errorf("longestCommonSubsequence(%q, %q) = %d, want 0", pair[0], pair[1], got)
		}
	}
}

func TestLongestCommonSubsequenceAgainstEnumeration(t *testing.T) {
	words := []string{""}
	for length := 1; length <= 4; length++ {
		for mask := 0; mask < 1<<length; mask++ {
			word := make([]byte, length)
			for i := range word {
				word[i] = 'a' + byte((mask>>i)&1)
			}
			words = append(words, string(word))
		}
	}
	for _, first := range words {
		for _, second := range words {
			want := longestByEnumeration(first, second)
			if got := longestCommonSubsequence(first, second); got != want {
				t.Fatalf("longestCommonSubsequence(%q, %q) = %d, enumeration gives %d", first, second, got, want)
			}
		}
	}
}

// 独立枚举第一个字符串的子序列，不复用动态规划转移。
func longestByEnumeration(first, second string) int {
	best := 0
	for mask := 0; mask < 1<<len(first); mask++ {
		position, length := 0, 0
		valid := true
		for i := 0; i < len(first); i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			for position < len(second) && second[position] != first[i] {
				position++
			}
			if position == len(second) {
				valid = false
				break
			}
			position++
			length++
		}
		if valid && length > best {
			best = length
		}
	}
	return best
}
