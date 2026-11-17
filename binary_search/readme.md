# 二分查找：统一边界与特殊结构子模板

[返回模板总目录](../algorithm_templates/README.md)

FirstTrue 的循环不变量是：答案始终在尚未排除的闭区间与右侧返回哨兵中；谓词必须先 false 后 true。查找元素、重复边界、答案二分都只改变谓词与答案转换。

81 含重复值、162/1901 寻找任意峰值采用独立的二分排除内核；它们不能把不单调的坡度伪装成 FirstTrue 的全局单调谓词。2024 对长度做二分，固定长度检查分别统计两种字符；2498 的可行性利用有序石子的隔点跳跃约束。

## 复制约定

每题 `.go` 文件已内嵌所需模板，无需导入仓库的算法包。文档将固定模板只列一次，逐题展示适配器；每个适配器与它列出的模板组成一个独立代码组。不同题的同名 Constructor、类型或入口请分别使用。将列出的标准库 import 放在文件顶部，再放模板与适配器；复制提交时优先直接使用对应题目的完整 `.go`。

本组使用 Go 1.21+ 的 min/max 和 Go 1.18+ 的泛型。平台已提供的 ListNode 等类型沿用平台定义。

## 固定模板

### FirstTrue

<!-- template: FirstTrue -->
```go
func FirstTrue(n int, predicate func(int) bool) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if predicate(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
```

## 题目索引

| 题目 | 模板 | 本地代码与测试 |
| --- | --- | --- |
| [33](https://leetcode.com/problems/search-in-rotated-sorted-array/) | FirstTrue | [代码](l33/33.go) · [测试](l33/33_test.go) |
| [34](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | FirstTrue | [代码](l34/34.go) · [测试](l34/34_test.go) |
| [81](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/) | 特殊结构二分 | [代码](l81/81.go) · [测试](l81/81_test.go) |
| [153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | FirstTrue | [代码](l153/153.go) · [测试](l153/153_test.go) |
| [162](https://leetcode.com/problems/find-peak-element/) | 特殊结构二分 | [代码](l162/162.go) · [测试](l162/162_test.go) |
| [410](https://leetcode.com/problems/split-array-largest-sum/) | FirstTrue | [代码](l410/410.go) · [测试](l410/410_test.go) |
| [704](https://leetcode.com/problems/binary-search/) | FirstTrue | [代码](l704/704.go) · [测试](l704/704_test.go) |
| [729](https://leetcode.com/problems/my-calendar-i/) | FirstTrue | [代码](l729/729.go) · [测试](l729/729_test.go) |
| [875](https://leetcode.com/problems/koko-eating-bananas/) | FirstTrue | [代码](l875/875.go) · [测试](l875/875_test.go) |
| [888](https://leetcode.com/problems/fair-candy-swap/) | FirstTrue | [代码](l888/888.go) · [测试](l888/888_test.go) |
| [1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) | FirstTrue | [代码](l1011/1011.go) · [测试](l1011/1011_test.go) |
| [1170](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/) | FirstTrue | [代码](l1170/1170.go) · [测试](l1170/1170_test.go) |
| [1608](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/) | FirstTrue | [代码](l1608/1608.go) · [测试](l1608/1608_test.go) |
| [1760](https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/) | FirstTrue | [代码](l1760/1760.go) · [测试](l1760/1760_test.go) |
| [1855](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/) | FirstTrue | [代码](l1855/1855.go) · [测试](l1855/1855_test.go) |
| [1901](https://leetcode.com/problems/find-a-peak-element-ii/) | 特殊结构二分 | [代码](l1901/1901.go) · [测试](l1901/1901_test.go) |
| [2024](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/) | FirstTrue | [代码](l2024/2024.go) · [测试](l2024/2024_test.go) |
| [2089](https://leetcode.com/problems/find-target-indices-after-sorting-array/) | FirstTrue | [代码](l2089/2089.go) · [测试](l2089/2089_test.go) |
| [2498](https://leetcode.com/problems/frog-jump-ii/) | FirstTrue | [代码](l2498/2498.go) · [测试](l2498/2498_test.go) |

## 逐题适配器

### [33](https://leetcode.com/problems/search-in-rotated-sorted-array/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l33/33.go; templates: FirstTrue; imports: none -->
```go
// 33：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums)
	pivot := FirstTrue(n, func(i int) bool { return nums[i] <= nums[n-1] })
	left, right := 0, pivot
	if target <= nums[n-1] {
		left, right = pivot, n
	}
	index := left + FirstTrue(right-left, func(i int) bool { return nums[left+i] >= target })
	if index < right && nums[index] == target {
		return index
	}
	return -1
}
```

### [34](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l34/34.go; templates: FirstTrue; imports: none -->
```go
// 34：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func searchRange(nums []int, target int) []int {
	left := FirstTrue(len(nums), func(i int) bool { return nums[i] >= target })
	right := FirstTrue(len(nums), func(i int) bool { return nums[i] > target })
	if left == right {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
}
```

### [81](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)

需要模板：本题的独立排除内核。

<!-- solution: binary_search/l81/81.go; templates: none; imports: none -->
```go
// 81：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
```

### [153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l153/153.go; templates: FirstTrue; imports: none -->
```go
// 153：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func findMin(nums []int) int {
	last := nums[len(nums)-1]
	index := FirstTrue(len(nums), func(i int) bool { return nums[i] <= last })
	return nums[index]
}
```

### [162](https://leetcode.com/problems/find-peak-element/)

需要模板：本题的独立排除内核。

<!-- solution: binary_search/l162/162.go; templates: none; imports: none -->
```go
// 162：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
```

### [410](https://leetcode.com/problems/split-array-largest-sum/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l410/410.go; templates: FirstTrue; imports: none -->
```go
// 410：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func splitArray(nums []int, k int) int {
	low, high := 0, 0
	for _, value := range nums {
		if value > low {
			low = value
		}
		high += value
	}
	return low + FirstTrue(high-low+1, func(offset int) bool {
		limit := low + offset
		parts, total := 1, 0
		for _, value := range nums {
			if total+value > limit {
				parts++
				total = 0
			}
			total += value
		}
		return parts <= k
	})
}
```

### [704](https://leetcode.com/problems/binary-search/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l704/704.go; templates: FirstTrue; imports: none -->
```go
// 704：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func search(nums []int, target int) int {
	index := FirstTrue(len(nums), func(i int) bool { return nums[i] >= target })
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}
```

### [729](https://leetcode.com/problems/my-calendar-i/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l729/729.go; templates: FirstTrue; imports: none -->
```go
// 729：复用本文件的边界模板；特殊结构按对应排除不变量处理。
type interval struct{ start, end int }
type MyCalendar struct{ bookings []interval }

func Constructor() MyCalendar { return MyCalendar{} }
func (c *MyCalendar) Book(start, end int) bool {
	if start >= end {
		return false
	}
	position := FirstTrue(len(c.bookings), func(i int) bool { return c.bookings[i].start >= start })
	if position > 0 && c.bookings[position-1].end > start {
		return false
	}
	if position < len(c.bookings) && c.bookings[position].start < end {
		return false
	}
	c.bookings = append(c.bookings, interval{})
	copy(c.bookings[position+1:], c.bookings[position:])
	c.bookings[position] = interval{start, end}
	return true
}
```

### [875](https://leetcode.com/problems/koko-eating-bananas/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l875/875.go; templates: FirstTrue; imports: none -->
```go
// 875：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func minEatingSpeed(piles []int, h int) int {
	highest := 0
	for _, pile := range piles {
		if pile > highest {
			highest = pile
		}
	}
	return 1 + FirstTrue(highest, func(offset int) bool {
		speed := offset + 1
		hours := int64(0)
		for _, pile := range piles {
			hours += int64((pile-1)/speed + 1)
			if hours > int64(h) {
				return false
			}
		}
		return true
	})
}
```

### [888](https://leetcode.com/problems/fair-candy-swap/)

需要模板：`FirstTrue`。

标准库：`sort`。

<!-- solution: binary_search/l888/888.go; templates: FirstTrue; imports: sort -->
```go
// 888：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func fairCandySwap(aliceSizes, bobSizes []int) []int {
	sumA, sumB := 0, 0
	for _, value := range aliceSizes {
		sumA += value
	}
	for _, value := range bobSizes {
		sumB += value
	}
	if (sumB-sumA)%2 != 0 {
		return []int{}
	}
	values := append([]int(nil), bobSizes...)
	sort.Ints(values)
	for _, a := range aliceSizes {
		b := a + (sumB-sumA)/2
		index := FirstTrue(len(values), func(i int) bool { return values[i] >= b })
		if index < len(values) && values[index] == b {
			return []int{a, b}
		}
	}
	return []int{}
}
```

### [1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l1011/1011.go; templates: FirstTrue; imports: none -->
```go
// 1011：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func shipWithinDays(weights []int, days int) int {
	low, high := 0, 0
	for _, weight := range weights {
		if weight > low {
			low = weight
		}
		high += weight
	}
	return low + FirstTrue(high-low+1, func(offset int) bool {
		capacity := low + offset
		needed, load := 1, 0
		for _, weight := range weights {
			if load+weight > capacity {
				needed++
				load = 0
			}
			load += weight
		}
		return needed <= days
	})
}
```

### [1170](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/)

需要模板：`FirstTrue`。

标准库：`sort`。

<!-- solution: binary_search/l1170/1170.go; templates: FirstTrue; imports: sort -->
```go
// 1170：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func frequency(s string) int {
	if len(s) == 0 {
		return 0
	}
	smallest, count := s[0], 0
	for i := range s {
		if s[i] < smallest {
			smallest, count = s[i], 1
		} else if s[i] == smallest {
			count++
		}
	}
	return count
}
func numSmallerByFrequency(queries, words []string) []int {
	values := make([]int, len(words))
	for i, word := range words {
		values[i] = frequency(word)
	}
	sort.Ints(values)
	answer := make([]int, len(queries))
	for i, query := range queries {
		value := frequency(query)
		answer[i] = len(values) - FirstTrue(len(values), func(j int) bool { return values[j] > value })
	}
	return answer
}
```

### [1608](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/)

需要模板：`FirstTrue`。

标准库：`sort`。

<!-- solution: binary_search/l1608/1608.go; templates: FirstTrue; imports: sort -->
```go
// 1608：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func specialArray(nums []int) int {
	values := append([]int(nil), nums...)
	sort.Ints(values)
	count := func(x int) int {
		return len(values) - FirstTrue(len(values), func(i int) bool { return values[i] >= x })
	}
	x := FirstTrue(len(values)+1, func(x int) bool { return count(x) <= x })
	if count(x) == x {
		return x
	}
	return -1
}
```

### [1760](https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l1760/1760.go; templates: FirstTrue; imports: none -->
```go
// 1760：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func minimumSize(nums []int, maxOperations int) int {
	highest := 0
	for _, value := range nums {
		if value > highest {
			highest = value
		}
	}
	return 1 + FirstTrue(highest, func(offset int) bool {
		size := offset + 1
		operations := int64(0)
		for _, value := range nums {
			operations += int64((value - 1) / size)
			if operations > int64(maxOperations) {
				return false
			}
		}
		return true
	})
}
```

### [1855](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l1855/1855.go; templates: FirstTrue; imports: none -->
```go
// 1855：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func maxDistance(nums1 []int, nums2 []int) int {
	answer := 0
	for i, value := range nums1 {
		if i >= len(nums2) {
			break
		}
		end := i + FirstTrue(len(nums2)-i, func(offset int) bool { return nums2[i+offset] < value })
		if end-1-i > answer {
			answer = end - 1 - i
		}
	}
	return answer
}
```

### [1901](https://leetcode.com/problems/find-a-peak-element-ii/)

需要模板：本题的独立排除内核。

<!-- solution: binary_search/l1901/1901.go; templates: none; imports: none -->
```go
// 1901：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func findPeakGrid(mat [][]int) []int {
	left, right := 0, len(mat[0])-1
	for left <= right {
		column := left + (right-left)/2
		row := 0
		for r := 1; r < len(mat); r++ {
			if mat[r][column] > mat[row][column] {
				row = r
			}
		}
		if column > 0 && mat[row][column-1] > mat[row][column] {
			right = column - 1
		} else if column+1 < len(mat[0]) && mat[row][column+1] > mat[row][column] {
			left = column + 1
		} else {
			return []int{row, column}
		}
	}
	return nil
}
```

### [2024](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l2024/2024.go; templates: FirstTrue; imports: none -->
```go
// 2024：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func maxConsecutiveAnswers(answerKey string, k int) int {
	feasible := func(length int) bool {
		if length == 0 {
			return true
		}
		ts := 0
		for right := range answerKey {
			if answerKey[right] == 'T' {
				ts++
			}
			if right >= length && answerKey[right-length] == 'T' {
				ts--
			}
			if right+1 >= length && min(ts, length-ts) <= k {
				return true
			}
		}
		return false
	}
	return FirstTrue(len(answerKey)+1, func(length int) bool { return !feasible(length) }) - 1
}
```

### [2089](https://leetcode.com/problems/find-target-indices-after-sorting-array/)

需要模板：`FirstTrue`。

标准库：`sort`。

<!-- solution: binary_search/l2089/2089.go; templates: FirstTrue; imports: sort -->
```go
// 2089：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func targetIndices(nums []int, target int) []int {
	values := append([]int(nil), nums...)
	sort.Ints(values)
	left := FirstTrue(len(values), func(i int) bool { return values[i] >= target })
	right := FirstTrue(len(values), func(i int) bool { return values[i] > target })
	answer := make([]int, 0, right-left)
	for i := left; i < right; i++ {
		answer = append(answer, i)
	}
	return answer
}
```

### [2498](https://leetcode.com/problems/frog-jump-ii/)

需要模板：`FirstTrue`。

<!-- solution: binary_search/l2498/2498.go; templates: FirstTrue; imports: none -->
```go
// 2498：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func maxJump(stones []int) int {
	n := len(stones)
	if n < 2 {
		return 0
	}
	return FirstTrue(stones[n-1]-stones[0]+1, func(limit int) bool {
		if stones[1]-stones[0] > limit || stones[n-1]-stones[n-2] > limit {
			return false
		}
		for i := 2; i < n; i++ {
			if stones[i]-stones[i-2] > limit {
				return false
			}
		}
		return true
	})
}
```

## 验证与原有推导

```sh
go test ./binary_search/...
```

题解文件与文档使用相同声明和适配器；修改模板后需要同时验证每题包以及由文档重新组装的代码。原有推导、详细例子和扩展题单保存在 [整理前笔记](legacy/README.before-inline-templates.md)。
