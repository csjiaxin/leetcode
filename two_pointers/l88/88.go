package l88

// MergeBack 的写指针位于所有未读元素之后，避免覆盖 dst 的有效前缀。
func MergeBack(dst []int, m int, src []int, n int) {
	i, j, write := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && dst[i] > src[j] {
			dst[write] = dst[i]
			i--
		} else {
			dst[write] = src[j]
			j--
		}
		write--
	}
}

// https://leetcode.com/problems/merge-sorted-array/
// 从尾部归并到 nums1，第二个序列耗尽后前缀已在正确位置。
func merge(nums1 []int, m int, nums2 []int, n int) { MergeBack(nums1, m, nums2, n) }
