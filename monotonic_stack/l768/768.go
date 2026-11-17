package l768

// https://leetcode.com/problems/max-chunks-to-make-sorted-ii/
// Embedded subtemplate: see README.md for its invariant and tie policy.
func sortedChunks(arr []int) int {
	maxima := []int{}
	for _, value := range arr {
		if len(maxima) == 0 || maxima[len(maxima)-1] <= value {
			maxima = append(maxima, value)
			continue
		}
		mergedMax := maxima[len(maxima)-1]
		for len(maxima) > 0 && maxima[len(maxima)-1] > value {
			maxima = maxima[:len(maxima)-1]
		}
		maxima = append(maxima, mergedMax)
	}
	return len(maxima)
}

func maxChunksToSorted(arr []int) int { return sortedChunks(arr) }
