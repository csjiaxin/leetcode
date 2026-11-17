package main

// Standalone template plus public LeetCode adapter; README.md describes the invariant.
// monotonicBoundaries settles right boundaries on pop and reads left on push.
// pop defines the strict/non-strict tie ownership for this specific problem.
func monotonicBoundaries(nums []int, pop func(previous, current int) bool) ([]int, []int) {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i, x := range nums {
		right[i] = n
		for len(stack) > 0 && pop(nums[stack[len(stack)-1]], x) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[j] = i
		}
		left[i] = -1
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return left, right
}

func totalStrength(strength []int) int {
	const mod int64 = 1_000_000_007
	n := len(strength)
	prefix := make([]int64, n+1)
	prefixOfPrefix := make([]int64, n+2)
	for i, value := range strength {
		prefix[i+1] = (prefix[i] + int64(value)) % mod
	}
	for i, value := range prefix {
		prefixOfPrefix[i+1] = (prefixOfPrefix[i] + value) % mod
	}
	left, right := monotonicBoundaries(strength, func(previous, current int) bool { return previous >= current })
	var result int64
	for i, value := range strength {
		sumRight := (prefixOfPrefix[right[i]+1] - prefixOfPrefix[i+1] + mod) % mod
		sumLeft := (prefixOfPrefix[i+1] - prefixOfPrefix[left[i]+1] + mod) % mod
		weighted := (int64(i-left[i])*sumRight - int64(right[i]-i)*sumLeft) % mod
		if weighted < 0 {
			weighted += mod
		}
		result = (result + int64(value)*weighted) % mod
	}
	return int(result)
}

func main() {}
