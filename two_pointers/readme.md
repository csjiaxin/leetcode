# 双指针（Two Pointers）

## 本页可复用模板：双指针与确定后继

先复制以下 helper，再复制任一题适配器。实际题解在每个文件内嵌同名 helper，可独立提交，不导入仓库本地算法包。不同移动/贪心依据使用不同命名模板，不把“用了两个变量”当作同一不变量。

```go
import (
	"sort"
	"strconv"
)

func Compact(nums []int, keep func(out []int, value int) bool) int {
	write := 0
	for _, value := range nums {
		if keep(nums[:write], value) {
			nums[write] = value
			write++
		}
	}
	return write
}

func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

func Distance(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

func Vowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func Alnum(c byte) bool { return c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }

func Lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func PalindromeRange(s string, l, r int) bool {
	answer := true
	Opposite(l, r, func(l, r int) (int, int) {
		if s[l] != s[r] {
			answer = false
			return r, r
		}
		return l + 1, r - 1
	})
	return answer
}

func AdvanceTwo(n, m int, step func(i, j int) (int, int)) (int, int) {
	i, j := 0, 0
	for i < n && j < m {
		i, j = step(i, j)
	}
	return i, j
}

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

func PreviousVisible(s string, i int) int {
	skipped := 0
	for i >= 0 {
		if s[i] == '#' {
			skipped++
			i--
		} else if skipped > 0 {
			skipped--
			i--
		} else {
			break
		}
	}
	return i
}

func ForRuns[T comparable](values []T, visit func(start, end int)) {
	for start := 0; start < len(values); {
		end := start + 1
		for end < len(values) && values[end] == values[start] {
			end++
		}
		visit(start, end)
		start = end
	}
}

func Partition(nums []int, keepLeft func(int) bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if keepLeft(nums[left]) {
			left++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		right--
	}
	return left
}

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

func ChooseEnds(n int, chooseLow func(i int) bool) []int {
	answer := make([]int, n+1)
	low, high := 0, n
	for i := 0; i < n; i++ {
		if chooseLow(i) {
			answer[i] = low
			low++
		} else {
			answer[i] = high
			high--
		}
	}
	answer[n] = low
	return answer
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Floyd[S comparable](start S, next func(S) (S, bool)) (S, bool) {
	slow, fast := start, start
	for {
		var ok bool
		fast, ok = next(fast)
		if !ok {
			var zero S
			return zero, false
		}
		fast, ok = next(fast)
		if !ok {
			var zero S
			return zero, false
		}
		slow, ok = next(slow)
		if !ok {
			var zero S
			return zero, false
		}
		if slow == fast {
			return slow, true
		}
	}
}

func NodeStep(node *ListNode) (*ListNode, bool) {
	if node == nil || node.Next == nil {
		return nil, false
	}
	return node.Next, true
}

func CycleEntry[S comparable](start S, next func(S) (S, bool)) (S, bool) {
	meeting, ok := Floyd(start, next)
	if !ok {
		var zero S
		return zero, false
	}
	for start != meeting {
		start, _ = next(start)
		meeting, _ = next(meeting)
	}
	return start, true
}

func Middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode
	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}
	return previous
}

func NthPredecessor(dummy *ListNode, n int) *ListNode {
	if n <= 0 {
		return nil
	}
	runner := dummy
	for step := 0; step < n; step++ {
		if runner.Next == nil {
			return nil
		}
		runner = runner.Next
	}
	previous := dummy
	for runner.Next != nil {
		runner = runner.Next
		previous = previous.Next
	}
	return previous
}

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
```

| 题目 | 采用的局部模板 | 实际实现 |
| --- | --- | --- |
| 26 | `Compact` | [实际 Go 文件](l26/26.go) |
| 80 | `Compact` | [实际 Go 文件](l80/80.go) |
| 27 | `Compact` | [实际 Go 文件](l27/27.go) |
| 283 | `Compact` | [实际 Go 文件](l283/283.go) |
| 167 | `Opposite` | [实际 Go 文件](l167/167.go) |
| 15 | `Opposite` | [实际 Go 文件](l15/15.go) |
| 16 | `Opposite + Distance` | [实际 Go 文件](l16/16.go) |
| 18 | `Opposite` | [实际 Go 文件](l18/18.go) |
| 611 | `Opposite` | [实际 Go 文件](l611/611.go) |
| 11 | `Opposite` | [实际 Go 文件](l11/11.go) |
| 42 | `Opposite` | [实际 Go 文件](l42/42.go) |
| 344 | `Opposite` | [实际 Go 文件](l344/344.go) |
| 345 | `Opposite + Vowel` | [实际 Go 文件](l345/345.go) |
| 125 | `Opposite + Alnum + Lower` | [实际 Go 文件](l125/125.go) |
| 680 | `Opposite + PalindromeRange` | [实际 Go 文件](l680/680.go) |
| 977 | `Opposite` | [实际 Go 文件](l977/977.go) |
| 392 | `AdvanceTwo` | [实际 Go 文件](l392/392.go) |
| 986 | `AdvanceTwo` | [实际 Go 文件](l986/986.go) |
| 88 | `MergeBack` | [实际 Go 文件](l88/88.go) |
| 844 | `PreviousVisible` | [实际 Go 文件](l844/844.go) |
| 443 | `ForRuns` | [实际 Go 文件](l443/443.go) |
| 905 | `Partition` | [实际 Go 文件](l905/905.go) |
| 75 | `ThreeWayPartition` | [实际 Go 文件](l75/75.go) |
| 942 | `ChooseEnds` | [实际 Go 文件](l942/942.go) |
| 141 | `Floyd + NodeStep` | [实际 Go 文件](l141/141.go) |
| 142 | `Floyd + CycleEntry + NodeStep` | [实际 Go 文件](l142/142.go) |
| 202 | `Floyd` | [实际 Go 文件](l202/202.go) |
| 287 | `Floyd + CycleEntry` | [实际 Go 文件](l287/287.go) |
| 876 | `Middle` | [实际 Go 文件](l876/876.go) |
| 234 | `Middle + ReverseList` | [实际 Go 文件](l234/234.go) |
| 19 | `NthPredecessor` | [实际 Go 文件](l19/19.go) |
| 1574：拼接有序前后缀 | `joinSortedEnds` | [实际 Go 文件](l1574/legacy/1574.go) |

为使整页 Go 示例可合并编译，重名适配器在文档中加题号/DP/Greedy 后缀；实际 Go 文件仍保留官方函数签名。链表节点等支撑类型仅在 helper 区声明一次。

双指针的核心不是“定义两个下标”，而是利用顺序、单调性或结构关系，让指针协同移动，避免重复枚举。

题单入口：[LeetCode Two Pointers](https://leetcode.com/problem-list/two-pointers/)。本文按经典模型整理代表题，不是该动态题单的完整枚举。

## 目录

- [1. 核心思想与模型速查](#core)
- [2. 相向双指针](#opposite)
- [3. 同向读写指针](#read-write)
- [4. 双序列归并与匹配](#merge)
- [5. 快慢指针与滑动窗口](#fast-slow)
- [6. 15. 三数之和：模板与图解](#leetcode-15)
- [7. 11. 盛最多水的容器：模板与图解](#leetcode-11)
- [8. 42. 接雨水：模板与图解](#leetcode-42)
- [9. 142. 环形链表 II：模板与图解](#leetcode-142)
- [10. 分区指针、易错点与练习路线](#practice)
- [11. 原参考链接（保留）](#references)

<a id="core"></a>

## 1. 核心思想与模型速查

做题前先回答三个问题：

1. 每个指针代表什么，维护的区间是什么？
2. 什么条件下移动哪个指针？
3. **为什么移动后不会漏解？**

| 模型 | 指针含义 | 常见信号 | 代表题 |
|---|---|---|---|
| 相向双指针 | 左右边界向中间移动 | 有序数对、回文、两端高度 | 167、15、11、42、125 |
| 同向读写指针 | 一个扫描，一个维护有效结果 | 原地删除、去重、移动元素 | 26、27、80、283 |
| 双序列指针 | 分别扫描两个序列 | 合并、比较、匹配 | 88、392、844、986 |
| 快慢 / 间距指针 | 不同速度或固定距离 | 中点、环、倒数位置 | 876、141、142、19 |
| 滑动窗口 | 连续区间的左右边界 | 子数组、子串、区间约束 | 3、209、713、76 |
| 分区与多指针 | 各类别的边界 | 原地分类、重排 | 75、905 |

分类可以重叠，重点是识别不变量。两个指针各自最多单向移动 n 次时，总移动次数是 O(n)，不是 O(n²)；排序、外层枚举等成本需要另外计算。

<a id="opposite"></a>

## 2. 相向双指针

### 2.1 有序数对：根据大小排除候选

| 题目 | 核心思路 | 时间（包含排序） |
|---|---|---|
| [167. 两数之和 II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | 和小移左，和大移右 | O(n) |
| [15. 三数之和](https://leetcode.com/problems/3sum/) | 排序，固定一个数，其余做两数之和 | O(n²) |
| [16. 最接近的三数之和](https://leetcode.com/problems/3sum-closest/) | 固定一个数，维护最小差值 | O(n²) |
| [18. 四数之和](https://leetcode.com/problems/4sum/) | 固定两个数，其余相向查找 | O(n³) |
| [611. 有效三角形的个数](https://leetcode.com/problems/valid-triangle-number/) | 固定最大边，一次统计一段合法左端点 | O(n²) |

167 的排除依据（数组升序）：

- `nums[l]+nums[r] < target`：当前 l 与任何不超过 r 的右端点搭配都不够大，排除 l。
- `nums[l]+nums[r] > target`：当前 r 与任何不小于 l 的左端点搭配都太大，排除 r。

611：排序后固定最大边 `nums[k]`，令 `l=0, r=k-1`。若 `nums[l]+nums[r]>nums[k]`，则 l 到 r-1 都可作为第一条边，一次贡献 `r-l`，然后 `r--`；否则 `l++`。

### 2.2 回文、反转与两端选择

| 题目 | 核心思路 |
|---|---|
| [344. 反转字符串](https://leetcode.com/problems/reverse-string/) | 两端交换，向中间收缩 |
| [345. 反转字符串中的元音字母](https://leetcode.com/problems/reverse-vowels-of-a-string/) | 两端各自寻找元音，再交换 |
| [125. 验证回文串](https://leetcode.com/problems/valid-palindrome/) | 跳过非字母数字字符，忽略大小写后比较 |
| [680. 验证回文串 II](https://leetcode.com/problems/valid-palindrome-ii/) | 首次失配时检查跳过左字符或右字符 |
| [977. 有序数组的平方](https://leetcode.com/problems/squares-of-a-sorted-array/) | 比较两端绝对值，从结果末尾填入最大平方 |

680 不能随便删除一边：需要验证 `[l+1,r]` 或 `[l,r-1]` 是否为回文。最多额外做两次线性检查，总时间仍为 O(n)。

<a id="read-write"></a>

## 3. 同向读写指针

**读指针扫描输入，写指针维护有效前缀。** 不变量：`nums[:write]` 是已处理部分的正确结果。

| 题目 | 保留条件 / 操作 |
|---|---|
| [27. 移除元素](https://leetcode.com/problems/remove-element/) | 保留不等于 val 的元素 |
| [26. 删除有序数组中的重复项](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | 与最后保留的元素不同才写入 |
| [80. 删除有序数组中的重复项 II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/) | 每个值最多保留两次 |
| [283. 移动零](https://leetcode.com/problems/move-zeroes/) | 稳定写入非零元素，再把尾部补零 |
| [443. 压缩字符串](https://leetcode.com/problems/string-compression/) | 读取连续字符组，写入字符与计数 |

### 27：基本模板

```go
func removeElement(nums []int, val int) int {
	return Compact(nums, func(_ []int, value int) bool { return value != val })
}
```

返回有效长度，尾部内容不重要。时间 O(n)，额外空间 O(1)。

### 26 / 80：统一为最多保留 k 次

```go
func keepAtMostK(nums []int, k int) int {
	if k <= 0 {
		return 0
	}
	return Compact(nums, func(out []int, value int) bool { return len(out) < k || out[len(out)-k] != value })
}
```

`k=1` 对应 26，`k=2` 对应 80。比较的是已保留结果中的 `nums[write-k]`，不是原数组中读指针之前第 k 个元素。写入位置不会超过当前读取位置，因此不会破坏尚未读取的数据。

<a id="merge"></a>

## 4. 双序列归并与匹配

| 题目 | 核心思路 | 时间 |
|---|---|---|
| [88. 合并两个有序数组](https://leetcode.com/problems/merge-sorted-array/) | 从后往前填，避免覆盖未读元素 | O(m+n) |
| [392. 判断子序列](https://leetcode.com/problems/is-subsequence/) | 长串持续扫描，匹配成功才推进短串 | O(m+n) |
| [844. 比较含退格的字符串](https://leetcode.com/problems/backspace-string-compare/) | 从后向前扫描，各自跳过被删除字符 | O(m+n) |
| [986. 区间列表的交集](https://leetcode.com/problems/interval-list-intersections/) | 计算交集，推进结束更早的区间 | O(m+n) |

986：当前两个闭区间的交集为 `[max(a.start,b.start), min(a.end,b.end)]`，左端点不大于右端点时存在交集。之后推进结束更早的区间，因为它不可能再与另一序列后面的区间相交。前提是每组区间各自有序且内部不相交。

<a id="fast-slow"></a>

## 5. 快慢指针与滑动窗口

### 5.1 快慢 / 固定间距指针

| 题目 | 指针设计 | 关键点 |
|---|---|---|
| [876. 链表的中间结点](https://leetcode.com/problems/middle-of-the-linked-list/) | slow 一步，fast 两步 | 常见模板在偶数长度时返回第二个中点 |
| [141. 环形链表](https://leetcode.com/problems/linked-list-cycle/) | slow 一步，fast 两步 | 环内一定相遇 |
| [142. 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii/) | 相遇后，一个回头部，双方每次一步 | 再次相遇即环入口 |
| [19. 删除链表的倒数第 N 个结点](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | 哑节点配合固定间距 | 最终让慢指针停在待删节点前驱 |
| [234. 回文链表](https://leetcode.com/problems/palindrome-linked-list/) | 找中点、反转后半段、比较 | 注意奇偶长度，需要时恢复链表 |
| [202. 快乐数](https://leetcode.com/problems/happy-number/) | 将数位平方和视为后继函数 | 判断到达 1 或进入其他环 |
| [287. 寻找重复数](https://leetcode.com/problems/find-the-duplicate-number/) | 下标到数值形成函数图 | 依赖题目数值范围，转为找环入口 |

快慢指针不限于链表：存在确定后继函数 `x → f(x)` 时，也可以研究轨迹是否入环。

### 5.2 滑动窗口

滑动窗口维护连续区间，是双指针的重要分支。关键前提是状态可增量维护，并且指针移动具有可利用的单调性。

| 模型 | 代表题 | 收缩 / 统计方式 |
|---|---|---|
| 最长合法区间 | [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/)、[424](https://leetcode.com/problems/longest-repeating-character-replacement/)、[1004](https://leetcode.com/problems/max-consecutive-ones-iii/) | 不合法时收缩 |
| 最短满足区间 | [209](https://leetcode.com/problems/minimum-size-subarray-sum/)、[76](https://leetcode.com/problems/minimum-window-substring/) | 满足时更新答案并继续收缩 |
| 合法区间计数 | [713](https://leetcode.com/problems/subarray-product-less-than-k/) | 调整到合法后累计 `right-left+1` |
| 恰好 K 个不同值 | [992](https://leetcode.com/problems/subarrays-with-k-different-integers/) | `atMost(K)-atMost(K-1)` |
| 定长窗口 | [438](https://leetcode.com/problems/find-all-anagrams-in-a-string/)、[567](https://leetcode.com/problems/permutation-in-string/) | 保持长度并维护频次 |

209 的常规窗口依赖正数条件；允许负数后不能直接套用。713 的常规乘积窗口依赖正整数条件，并需处理 `k<=1`。详见[滑动窗口总结](../sliding_window/README.md)。

<a id="leetcode-15"></a>

## 6. [15. 三数之和](https://leetcode.com/problems/3sum/)：模板与图解

### 思路与移动依据

先排序，固定 `nums[i]`，在 `[i+1,n-1]` 中寻找和为 `-nums[i]` 的两个数。

- 三数之和小于 0：当前 l 即使配上当前最大的 r 也不够大，`l++`。
- 三数之和大于 0：当前 r 即使配上当前最小的 l 也太大，`r--`。
- 等于 0：记录答案，移动两端并跳过重复值。
- `nums[i]>0` 时后面全为正数，可以停止。

### Go 模板

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	answer := [][]int{}
	for i := 0; i+2 < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		Opposite(i+1, len(nums)-1, func(l, r int) (int, int) {
			sum := int64(nums[i]) + int64(nums[l]) + int64(nums[r])
			if sum < 0 {
				return l + 1, r
			}
			if sum > 0 {
				return l, r - 1
			}
			answer = append(answer, []int{nums[i], nums[l], nums[r]})
			l++
			r--
			for l < r && nums[l] == nums[l-1] {
				l++
			}
			for l < r && nums[r] == nums[r+1] {
				r--
			}
			return l, r
		})
	}
	return answer
}
```

时间 O(n²)；双指针扫描自身额外空间 O(1)，排序栈空间随实现而定，返回结果另计。`int64` 足够覆盖本题数值约束，不代表任意三个 int64 都不会溢出。

### 逐步图解

```text
输入：[-1, 0, 1, 2, -1, -4]
排序：[-4,-1,-1, 0, 1, 2]
下标： 0  1  2  3  4  5
```

| 固定 i | l | r | 三数之和 | 操作 |
|---|---|---|---|---|
| 0 (-4) | 1 (-1) | 5 (2) | -3 | l++ |
| 0 (-4) | 2 (-1) | 5 (2) | -3 | l++ |
| 0 (-4) | 3 (0) | 5 (2) | -2 | l++ |
| 0 (-4) | 4 (1) | 5 (2) | -1 | l++，本轮结束 |
| 1 (-1) | 2 (-1) | 5 (2) | 0 | 记录 [-1,-1,2]，两端内移 |
| 1 (-1) | 3 (0) | 4 (1) | 0 | 记录 [-1,0,1]，两端内移 |
| 2 (-1) | — | — | — | 与前一个固定值相同，跳过 |
| 3 (0) | 4 (1) | 5 (2) | 3 | r--，本轮结束 |

去重示例：`[-2,0,0,2,2]` 固定 -2，第一次找到 `[-2,0,2]` 后，l/r 内移并跳过重复值，不能再输出同一个值三元组。

**边界：** 长度不足 3 返回空结果；`[0,0,0,0]` 只输出一个 `[0,0,0]`；三个元素必须来自不同下标。

<a id="leetcode-11"></a>

## 7. [11. 盛最多水的容器](https://leetcode.com/problems/container-with-most-water/)：模板与图解

### 思路与证明：为什么移动短板？

`area(l,r) = min(height[l],height[r]) × (r-l)`。

假设 `height[l] <= height[r]`，固定 l，把右端点换成任意 `r'<r`：

```text
area(l,r') <= height[l] × (r'-l)
           <= height[l] × (r-l)
           =  area(l,r)
```

因此当前 l 不可能再参与更优答案，可以 `l++`。右边更矮时对称。两端等高时移动任意一边即可。

**移动短板只是保留可能变好的机会，不保证下一步面积一定增大。**

### Go 模板

```go
func maxArea(height []int) int {
	answer := 0
	Opposite(0, len(height)-1, func(l, r int) (int, int) {
		answer = max(answer, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			return l + 1, r
		}
		return l, r - 1
	})
	return answer
}
```

时间 O(n)，额外空间 O(1)。本题约束下 int 可容纳面积；推广到更大数值范围时应先转换再用 int64 乘法。

### 逐步图解

```text
下标：  0 1 2 3 4 5 6 7 8
高度： [1,8,6,2,5,4,8,3,7]
        L               R   → 面积 1×8=8，移动左侧短板
          L             R   → 面积 7×7=49，移动右侧短板
          L           R     → 面积 3×6=18，继续移动右侧
```

| l | r | 短板高度 | 宽度 | 面积 | 下一步 |
|---|---|---|---|---|---|
| 0 | 8 | 1 | 8 | 8 | l++ |
| 1 | 8 | 7 | 7 | 49 | r-- |
| 1 | 7 | 3 | 6 | 18 | r-- |
| 1 | 6 | 8 | 5 | 40 | 等高，模板选择 l++ |
| 2 | 6 | 6 | 4 | 24 | l++ |
| 3 | 6 | 2 | 3 | 6 | l++ |
| 4 | 6 | 5 | 2 | 10 | l++ |
| 5 | 6 | 4 | 1 | 4 | l++，结束 |

最大面积为 49。中间柱子不会减少容器容量；本题不是接雨水。

**边界：** 全零答案为 0；两个元素只计算一次；模板对空数组 / 单元素返回 0。

<a id="leetcode-42"></a>

## 8. [42. 接雨水](https://leetcode.com/problems/trapping-rain-water/)：模板与图解

### 从公式到双指针

位置 i 的水量为 `min(左侧最高柱,右侧最高柱)-height[i]`，这里两侧最高值都包含当前位置，水量不会为负数。

维护 `[0,l]` 的最高柱 leftMax 与 `[r,n-1]` 的最高柱 rightMax。每轮先用当前位置更新它们：

- `leftMax <= rightMax`：右边已经存在不低于 leftMax 的柱子，l 位置水位确定为 leftMax，结算后 `l++`。
- 否则：左边已经存在不低于 rightMax 的柱子，r 位置水位确定为 rightMax，结算后 `r--`。

中间未处理区域可能出现更高柱子，但不会改变较低侧当前被结算位置的水位。**不是猜另一侧最终最高值，而是已有足够高的挡板。**

### Go 模板

```go
func trap(height []int) int {
	leftMax, rightMax, answer := 0, 0, 0
	Opposite(0, len(height)-1, func(l, r int) (int, int) {
		if height[l] <= height[r] {
			leftMax = max(leftMax, height[l])
			answer += leftMax - height[l]
			return l + 1, r
		}
		rightMax = max(rightMax, height[r])
		answer += rightMax - height[r]
		return l, r - 1
	})
	return answer
}
```

时间 O(n)，额外空间 O(1)。`l<=r` 表示每个位置结算一次；空数组不会进入循环。不要把其他变体的比较对象、最高值更新时机和循环条件随意拼接。

### 逐步图解

选择能同时展示左右结算的例子 `[4,2,0,3,2,5,1]`：

```text
#：柱子，~：积水

高度 5 | . . . . . # .
高度 4 | # ~ ~ ~ ~ # .
高度 3 | # ~ ~ # ~ # .
高度 2 | # # ~ # # # .
高度 1 | # # ~ # # # #
下标   | 0 1 2 3 4 5 6
水量   | 0 2 4 1 2 0 0  → 总水量 9
```

表中最高值均为本轮更新之后的值：

| l | r | leftMax | rightMax | 结算位置 | 新增水量 | 累计 |
|---|---|---|---|---|---|---|
| 0 | 6 | 4 | 1 | 右侧 6 | 1-1=0 | 0 |
| 0 | 5 | 4 | 5 | 左侧 0 | 4-4=0 | 0 |
| 1 | 5 | 4 | 5 | 左侧 1 | 4-2=2 | 2 |
| 2 | 5 | 4 | 5 | 左侧 2 | 4-0=4 | 6 |
| 3 | 5 | 4 | 5 | 左侧 3 | 4-3=1 | 7 |
| 4 | 5 | 4 | 5 | 左侧 4 | 4-2=2 | 9 |
| 5 | 5 | 5 | 5 | 左侧 5 | 5-5=0 | 9 |

第二轮开始，右侧已经发现高度 5 的柱子，因此左边高度上限为 4 的各位置都可以安全结算。

### 与 11 的区别

| 对比 | 11 盛最多水的容器 | 42 接雨水 |
|---|---|---|
| 目标 | 选一对边界使面积最大 | 累加每个位置的积水 |
| 状态 | 两端高度与宽度 | 两侧最高柱 |
| 移动依据 | 排除不可能更优的端点 | 结算水位已经确定的位置 |
| 中间柱子 | 不影响这对边界的容量计算 | 占据空间，必须扣除柱高 |

**边界：** 单调数组、全相等数组均不积水；少于三根柱子答案为 0；不要只看左右紧邻柱子，要看两侧最高柱。

<a id="leetcode-142"></a>

## 9. [142. 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii/)：模板与图解

### 两个阶段

1. 判环：slow 每次一步，fast 每次两步；fast 或 fast.Next 为空说明无环。
2. 找入口：相遇后，让一个指针回到 head，另一个留在相遇点；两者每次各走一步，再次相遇的位置就是入口。

### Go 模板

```go
func detectCycle(head *ListNode) *ListNode { entry, _ := CycleEntry(head, NodeStep); return entry }
```

时间 O(n)，额外空间 O(1)。判断节点是否相同必须比较指针，不能比较 Val；不同节点的值可能相同。题目中的 pos 只是描述尾节点连接位置，不是函数入参。

### 逐步图解

```text
A → B → C → D → E
        ↑       ↓
        └───────┘

入口 C；头部到入口距离 a=2；环长 c=3。
```

阶段一（初始均在 A，先移动再判断）：

| 轮次 | slow（一步） | fast（两步） | 结果 |
|---|---|---|---|
| 初始 | A | A | 初始相同不能直接判有环 |
| 1 | B | C | 未相遇 |
| 2 | C | E | 未相遇 |
| 3 | D | D | 相遇，但 D 不是入口 |

阶段二：entry 回到 A，slow 留在 D：

| 轮次 | entry | slow | 结果 |
|---|---|---|---|
| 初始 | A | D | 每次各走一步 |
| 1 | B | E | 未相遇 |
| 2 | C | C | 找到入口 C |

### 为什么回到头部后会在入口相遇？

设头部到入口距离为 a，入口沿环到相遇点距离为 b（`0<=b<c`），环长为 c。slow 到相遇时可能已经绕环 k 圈：

```text
slow 距离 s = a + b + k*c
fast 距离   = 2*s
距离差 s    = m*c       （m >= 1）
因此 a+b ≡ 0 (mod c)
     a   ≡ -b (mod c)
```

从相遇点沿环走 a 步，最终位置是 `(b+a) mod c = 0`，恰好回到入口；从 head 走 a 步也到入口。

若 a>0，head 出发的指针在前 a 步尚未进入环，而另一个指针一直在环内，因此不可能提前相遇；若 a=0，两者在第二阶段开始时就处于入口。

**边界：** 空链表、单节点无环返回 nil；单节点自环返回自身；入口可能为 head；不能把第一次相遇点直接当作入口。

<a id="practice"></a>

## 10. 分区指针、易错点与练习路线

### 10.1 75：荷兰国旗分区

[75. 颜色分类](https://leetcode.com/problems/sort-colors/) 用三个指针维护四个区间：

```text
[0,low)     全是 0
[low,i)     全是 1
[i,high]    尚未处理
(high,n)    全是 2
```

```go
func sortColors(nums []int) { ThreeWayPartition(nums, 1) }
```

时间 O(n)，额外空间 O(1)。前提是元素只有 0、1、2。**交换到右边之后不执行 i++**，因为换回来的元素尚未处理。

### 10.2 常见错误

1. **没有证明就移动指针**：移动应对应“排除不可能的候选”或“完成某个位置的处理”。
2. **忽略排序影响**：排序会改变原始下标，并产生时间与栈空间成本；需要原下标时应额外保存信息。
3. **混淆循环边界**：不同元素配对通常 `l<r`；逐点结算或处理未知区间时可能 `l<=r`，应由不变量决定。
4. **去重不完整**：15 / 18 的固定值与匹配成功后的左右值都要考虑去重。
5. **原地覆盖未读数据**：同向写入通常保证 write 不超过 read；88 则从后往前填。
6. **Go 字符串字节与字符混淆**：`s[i]` 返回字节；ASCII 题可直接处理，按 Unicode 码点处理时考虑 `[]rune`。
7. **看到连续区间就套窗口**：先检查约束是否支持单调收缩，尤其注意负数。
8. **整数溢出**：先转为 int64 再求和 / 乘法，不是算完后再转换；仍需结合数值范围确认安全。
9. **链表比较节点值**：环检测比较节点身份（指针），不是值。

### 10.3 推荐刷题顺序

| 阶段 | 题目 | 练习目标 |
|---|---|---|
| 基础移动 | 344 → 125 → 27 → 26 → 283 | 边界、扫描、有效前缀 |
| 相向排除 | 167 → 15 → 16 → 18 → 611 | 单调性、去重、批量计数 |
| 贪心证明 | 11 → 42 | 区分排除候选与确定水位 |
| 双序列协作 | 88 → 392 → 844 → 986 | 选择移动哪一侧 |
| 链表与函数图 | 876 → 141 → 142 → 19 → 234 → 287 | 速度差、固定间距、环入口 |
| 综合应用 | 80 → 75 → 3 → 209 → 713 → 76 → 992 | 不变量与窗口条件 |

**一句话记忆：相向指针排除候选，读写指针维护结果，归并指针推进序列，快慢指针利用距离，滑动窗口维护区间。**

<a id="references"></a>

## 11. 原参考链接（保留）

- 综合刷题经验：https://leetcode.com/discuss/study-guide/1688903/Solved-all-two-pointers-problems-in-100-days.

RainWater：

- https://leetcode.com/problems/trapping-rain-water/submissions/
- https://leetcode.com/problems/container-with-most-water/
- https://leetcode.com/problems/trapping-rain-water/discuss/17357/Sharing-my-simple-c%2B%2B-code%3A-O(n)-time-O(1)-space
- https://leetcode.com/problems/trapping-rain-water/discuss/153992/Java-O(n)-time-and-O(1)-space-(with-explanations).

原有实现：[11.go](l11/11.go)、[42.go](l42/variants/from_rainwater/42.go)。本文模板与原实现可能采用不同的等高处理或循环边界，不影响各自的不变量；原实现与测试保留在合并后的题目目录。

## 目录归并：按移动规则选择模板

以下是当前目录入口；上面的历史讲义继续保留。实现和测试成对迁移，不把未完成练习替换成另一份答案。

| 不变量 | 当前题目入口 | 指针如何移动 |
| --- | --- | --- |
| 有序区间中排除不可能的数对 | [167](l167)、[15](l15)、[16](l16)、[18](l18)、[611](l611) | 和偏小移左端、偏大移右端；外层枚举另计复杂度 |
| 两端已经匹配或交换完成 | [125](l125)、[344](l344)、[345](l345)、[680](l680)、[977](l977) | 从两侧向中间推进，允许题目规定的跳过 |
| `nums[:write]` 是正确的有效前缀 | [26](l26)、[27](l27)、[80](l80)、[283](l283)、[443](l443) | read 扫描，满足保留条件才推进 write |
| 两个序列各自已扫描前缀不再需要 | [88](l88)、[392](l392)、[844](l844)、[986](l986) | 由排序、匹配或区间结束点决定移动方向 |
| 速度差或固定间距不变 | [141](l141)、[142](l142)、[19](l19)、[234](l234)、[876](l876)、[202](l202)、[287](l287) | 不强行套入相向指针循环 |
| 已处理端点的容器/水位可确定 | [11](l11)、[42](l42) | 11 排除短板；42 结算水位，证明不同 |
| 未使用数始终位于 `[low, high]` | [942](l942) | I 取最小值、D 取最大值，构造任意合法排列 |
| 连续窗口状态可增量更新 | [滑动窗口汇总](../sliding_window/README.md) | 3、76、209、424、438、567、713、992、1004 已实际迁入该目录 |

读写模板的移动规则可以复用，但保留条件必须由题目决定：

```text
write := 0
for read, value := range nums {
    // 此时 write <= read；nums[:write] 是已处理部分的正确结果。
    _ = read
    if keep(value) {
        nums[write] = value
        write++
    }
}
// 有效结果为 nums[:write]；keep 是题目提供的保留条件。
```

`rainwater/l11` 已移至 [l11](l11)；原可运行的 42 演示完整保存在 [l42/variants/from_rainwater](l42/variants/from_rainwater)，主目录 [l42](l42) 的独立练习和原测试现也使用同一 Opposite 模板。接雨水的单调栈版本仍在 [monotonic_stack/l42](../monotonic_stack/l42)，因为它按凹槽层而非两侧水位结算，不合并成一个强制共用的循环。

历史重复练习放在目标题目的 `variants/from_two_pointers`，仍是独立 Go 包；各来源现均按相同子模板独立实现，并运行原测试，不通过隐藏练习文件制造通过。路径映射和已知未完成包见 [reorganization.json](reorganization.json)。


## 其余题目的模板适配器

以下代码调用本页已经声明的 helper；复制到对应题解时使用实际 Go 文件中的官方函数名。

### 26

```go
func removeDuplicates26(nums []int) int {
	return Compact(nums, func(out []int, value int) bool { return len(out) < 1 || value != out[len(out)-1] })
}
```

### 80

```go
func removeDuplicates80(nums []int) int {
	return Compact(nums, func(out []int, value int) bool { return len(out) < 2 || value != out[len(out)-2] })
}
```

### 283

```go
func moveZeroes(nums []int) {
	length := Compact(nums, func(_ []int, value int) bool { return value != 0 })
	for i := length; i < len(nums); i++ {
		nums[i] = 0
	}
}
```

### 167

```go
func twoSum(numbers []int, target int) []int {
	result := []int{}
	Opposite(0, len(numbers)-1, func(l, r int) (int, int) {
		sum := int64(numbers[l]) + int64(numbers[r])
		if sum < int64(target) {
			return l + 1, r
		}
		if sum > int64(target) {
			return l, r - 1
		}
		result = []int{l + 1, r + 1}
		return r, r
	})
	return result
}
```

### 16

```go
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	desired := int64(target)
	best := int64(nums[0]) + int64(nums[1]) + int64(nums[2])
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		Opposite(i+1, len(nums)-1, func(l, r int) (int, int) {
			sum := int64(nums[i]) + int64(nums[l]) + int64(nums[r])
			if Distance(sum, desired) < Distance(best, desired) {
				best = sum
			}
			if sum < desired {
				return l + 1, r
			}
			if sum > desired {
				return l, r - 1
			}
			return r, r
		})
		if best == desired {
			break
		}
	}
	return int(best)
}
```

### 18

```go
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	answer := [][]int{}
	for i := 0; i+3 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j+2 < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			Opposite(j+1, len(nums)-1, func(l, r int) (int, int) {
				sum := int64(nums[i]) + int64(nums[j]) + int64(nums[l]) + int64(nums[r])
				if sum < int64(target) {
					return l + 1, r
				}
				if sum > int64(target) {
					return l, r - 1
				}
				answer = append(answer, []int{nums[i], nums[j], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				return l, r
			})
		}
	}
	return answer
}
```

### 611

```go
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	answer := 0
	for largest := 2; largest < len(nums); largest++ {
		Opposite(0, largest-1, func(l, r int) (int, int) {
			if int64(nums[l])+int64(nums[r]) > int64(nums[largest]) {
				answer += r - l
				return l, r - 1
			}
			return l + 1, r
		})
	}
	return answer
}
```

### 344

```go
func reverseString(s []byte) {
	Opposite(0, len(s)-1, func(l, r int) (int, int) { s[l], s[r] = s[r], s[l]; return l + 1, r - 1 })
}
```

### 345

```go
func reverseVowels(s string) string {
	data := []byte(s)
	Opposite(0, len(data)-1, func(l, r int) (int, int) {
		if !Vowel(data[l]) {
			return l + 1, r
		}
		if !Vowel(data[r]) {
			return l, r - 1
		}
		data[l], data[r] = data[r], data[l]
		return l + 1, r - 1
	})
	return string(data)
}
```

### 125

```go
func isPalindrome125(s string) bool {
	answer := true
	Opposite(0, len(s)-1, func(l, r int) (int, int) {
		if !Alnum(s[l]) {
			return l + 1, r
		}
		if !Alnum(s[r]) {
			return l, r - 1
		}
		if Lower(s[l]) != Lower(s[r]) {
			answer = false
			return r, r
		}
		return l + 1, r - 1
	})
	return answer
}
```

### 680

```go
func validPalindrome(s string) bool {
	answer := true
	Opposite(0, len(s)-1, func(l, r int) (int, int) {
		if s[l] != s[r] {
			answer = PalindromeRange(s, l+1, r) || PalindromeRange(s, l, r-1)
			return r, r
		}
		return l + 1, r - 1
	})
	return answer
}
```

### 977

```go
func sortedSquares(nums []int) []int {
	answer := make([]int, len(nums))
	write := len(nums) - 1
	l, _ := Opposite(0, len(nums)-1, func(l, r int) (int, int) {
		a, b := nums[l]*nums[l], nums[r]*nums[r]
		if a > b {
			answer[write] = a
			write--
			return l + 1, r
		}
		answer[write] = b
		write--
		return l, r - 1
	})
	if len(nums) > 0 {
		answer[0] = nums[l] * nums[l]
	}
	return answer
}
```

### 392

```go
func isSubsequence(s string, t string) bool {
	matched, _ := AdvanceTwo(len(s), len(t), func(i, j int) (int, int) {
		if s[i] == t[j] {
			return i + 1, j + 1
		}
		return i, j + 1
	})
	return matched == len(s)
}
```

### 986

```go
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	answer := [][]int{}
	AdvanceTwo(len(firstList), len(secondList), func(i, j int) (int, int) {
		left, right := max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])
		if left <= right {
			answer = append(answer, []int{left, right})
		}
		if firstList[i][1] < secondList[j][1] {
			return i + 1, j
		}
		if firstList[i][1] > secondList[j][1] {
			return i, j + 1
		}
		return i + 1, j + 1
	})
	return answer
}
```

### 88

```go
func merge(nums1 []int, m int, nums2 []int, n int) { MergeBack(nums1, m, nums2, n) }
```

### 844

```go
func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		i, j = PreviousVisible(s, i), PreviousVisible(t, j)
		if i < 0 || j < 0 {
			return i == j
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}
```

### 443

```go
func compress(chars []byte) int {
	write := 0
	ForRuns(chars, func(start, end int) {
		value := chars[start]
		chars[write] = value
		write++
		if count := end - start; count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[write] = byte(digit)
				write++
			}
		}
	})
	return write
}
```

### 905

```go
func sortArrayByParity(nums []int) []int {
	Partition(nums, func(value int) bool { return value%2 == 0 })
	return nums
}
```

### 942

```go
func diStringMatch(s string) []int {
	return ChooseEnds(len(s), func(i int) bool { return s[i] == 'I' })
}
```

### 141

```go
func hasCycle(head *ListNode) bool { _, found := Floyd(head, NodeStep); return found }
```

### 202

```go
func isHappy(n int) bool {
	next := func(value int) (int, bool) {
		sum := 0
		for value > 0 {
			digit := value % 10
			sum += digit * digit
			value /= 10
		}
		return sum, true
	}
	meeting, _ := Floyd(n, next)
	return meeting == 1
}
```

### 287

```go
func findDuplicate(nums []int) int {
	entry, _ := CycleEntry(0, func(i int) (int, bool) { return nums[i], true })
	return entry
}
```

### 876

```go
func middleNode(head *ListNode) *ListNode { return Middle(head) }
```

### 234

```go
func isPalindrome234(head *ListNode) bool {
	second := ReverseList(Middle(head))
	left, right := head, second
	answer := true
	for right != nil {
		if left.Val != right.Val {
			answer = false
			break
		}
		left = left.Next
		right = right.Next
	}
	ReverseList(second)
	return answer
}
```

### 19

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	previous := NthPredecessor(dummy, n)
	if previous != nil {
		previous.Next = previous.Next.Next
	}
	return dummy.Next
}
```

### 1574：拼接有序前后缀

```go
func findLengthOfShortestSubarray(arr []int) int { return joinSortedEnds(arr) }
```
