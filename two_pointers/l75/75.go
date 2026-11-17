package l75

// ThreeWayPartition 维护 [0,low) 小于 pivot、[low,scan) 等于 pivot、(high,n) 大于 pivot。
func ThreeWayPartition(nums []int, pivot int) {
	low, scan, high := 0, 0, len(nums)-1
	for scan <= high {
		if nums[scan] < pivot {
			nums[low], nums[scan] = nums[scan], nums[low]
			low++
			scan++
		} else if nums[scan] > pivot {
			nums[scan], nums[high] = nums[high], nums[scan]
			high--
		} else {
			scan++
		}
	}
}

// https://leetcode.com/problems/sort-colors/
// 三色划分；与右边交换后不推进 scan，换来的值还未分类。
func sortColors(nums []int) { ThreeWayPartition(nums, 1) }
