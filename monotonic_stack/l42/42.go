package l42

// https://leetcode.com/problems/trapping-rain-water/
// Embedded subtemplate: see README.md for its invariant and tie policy.
// trappedLayers settles a basin bottom only when both bounding walls exist.
func trappedLayers(heights []int) int {
	stack := []int{}
	water := 0
	for right, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] < h {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			level := heights[left]
			if h < level {
				level = h
			}
			water += (right - left - 1) * (level - heights[bottom])
		}
		stack = append(stack, right)
	}
	return water
}

func trap(height []int) int { return trappedLayers(height) }
