package main

// Standalone template plus public LeetCode adapter; README.md describes the invariant.
// joinSortedEnds removes one middle segment, not an arbitrary subsequence.
func joinSortedEnds(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	best := right
	for left := 0; left < right && (left == 0 || arr[left-1] <= arr[left]); left++ {
		for right < n && arr[left] > arr[right] {
			right++
		}
		if length := right - left - 1; length < best {
			best = length
		}
	}
	return best
}

func findLengthOfShortestSubarray(arr []int) int { return joinSortedEnds(arr) }

func main() {}
