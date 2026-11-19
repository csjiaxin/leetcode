package l1053

// Standalone lexicographic/factorial subtemplate and problem adapter.
// previousSingleSwap changes the rightmost reducible pivot by one swap only.
func previousSingleSwap(arr []int) []int {
	n := len(arr)

	i := n - 2
	for ; i >= 0 && arr[i] <= arr[i+1]; i-- {
	}
	if i < 0 {
		return arr
	}
	j := n - 1
	for ; j >= i && arr[j] >= arr[i]; j-- {
	}
	for ; arr[j-1] == arr[j]; j-- {

	}
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}

func prevPermOpt1(arr []int) []int { return previousSingleSwap(arr) }
