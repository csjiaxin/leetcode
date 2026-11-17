// Package l1092 solves LeetCode 1092: Shortest Common Supersequence.
// https://leetcode.com/problems/shortest-common-supersequence/
package l1092

import prefixdp "leetcode/dp/strings"

// shortestCommonSupersequence returns a shortest string containing both inputs
// as subsequences. dp[i][j] is its minimum length for str1[:i] and str2[:j].
// The full length table takes O(m*n) time and space. Backtracking appends each
// output byte once, followed by an O(m+n) reversal. Strings contain ASCII letters.
func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := prefixdp.Table[int](m, n,
		func(i, j int) int {
			return i + j // On a boundary, the other prefix is the answer.
		},
		func(i, j, up, left, diagonal int) int {
			if str1[i-1] == str2[j-1] {
				return diagonal + 1 // Share the matching final character.
			}
			if up < left {
				return up + 1
			}
			return left + 1
		},
	)

	// Walk from the final prefixes to the empty prefixes. The chosen characters
	// are discovered right to left, so collect them in a reverse buffer.
	reversed := make([]byte, 0, dp[m][n])
	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			reversed = append(reversed, str1[i-1])
			i--
			j--
		} else if dp[i-1][j] < dp[i][j-1] {
			reversed = append(reversed, str1[i-1])
			i--
		} else {
			// Either side is valid on a tie; choose str2 consistently.
			reversed = append(reversed, str2[j-1])
			j--
		}
	}
	for i > 0 {
		i--
		reversed = append(reversed, str1[i])
	}
	for j > 0 {
		j--
		reversed = append(reversed, str2[j])
	}
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed[left], reversed[right] = reversed[right], reversed[left]
	}
	return string(reversed)
}
