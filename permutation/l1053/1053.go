package l1053

// Standalone lexicographic/factorial subtemplate and problem adapter.
// previousSingleSwap changes the rightmost reducible pivot by one swap only.
func previousSingleSwap(arr []int) []int {
	pivot := len(arr) - 2
	for pivot >= 0 && arr[pivot] <= arr[pivot+1] {
		pivot--
	}
	if pivot < 0 {
		return arr
	}
	successor := len(arr) - 1
	for arr[successor] >= arr[pivot] {
		successor--
	}
	for successor > pivot+1 && arr[successor] == arr[successor-1] {
		successor--
	}
	arr[pivot], arr[successor] = arr[successor], arr[pivot]
	return arr
}

func prevPermOpt1(arr []int) []int { return previousSingleSwap(arr) }
