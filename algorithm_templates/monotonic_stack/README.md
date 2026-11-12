# 单调栈：最近边界与贡献结算

[返回总目录](../README.md)

## 目录

- [1. 核心思路](#core)
- [2. 739：每日温度](#lc739)
- [3. 496：下一个更大元素](#lc496)
- [4. 84：柱状图最大矩形](#lc84)
- [5. 907：子数组最小值贡献](#lc907)
- [6. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

栈保存尚未确定答案的下标。遍历当前元素，遇到能确定栈顶边界的值，就弹栈结算，然后入栈当前下标。

```text
遍历i → while栈顶满足弹出条件：弹出j并结算 → i入栈
```

每个下标最多入栈、出栈一次，所以总时间O(n)。比较符号取决于“严格更大/更小”和重复值归属，不能脱离题目机械背符号。以下代码块互相独立，均不依赖其他题实现。

<a id="lc739"></a>

## 2. [739. 每日温度](https://leetcode.com/problems/daily-temperatures/)

[练习](l739/739.go) · [测试](l739/739_test.go)

**题意**：对每一天求还需等几天才有更高温度，不存在为0。

**栈含义**：等待更高温度的下标，温度非递增。当前温度严格更大，才为旧下标确定答案。

```go
package template

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := []int{}
	for i, x := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < x {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[j] = i - j
		}
		stack = append(stack, i)
	}
	return answer
}
```

**例子**：[30,30,40]，i=2同时解决前两天下标，答案[2,1,0]。相等不是更暖，不能弹。

**复杂度**：时间O(n)，空间O(n)，包括输出和栈。

<a id="lc496"></a>

## 3. [496. 下一个更大元素 I](https://leetcode.com/problems/next-greater-element-i/)

[练习](l496/496.go) · [测试](l496/496_test.go)

**题意**：nums1是nums2的子集，两数组各自无重复。为nums1每个值查找它在nums2右侧的第一个更大值。

**变化**：仍弹出等待右侧更大值的下标，但结果存为“旧值→下一个更大值”的map，最后按nums1顺序输出。

```go
package template

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := make(map[int]int)
	stack := []int{}
	for i, x := range nums2 {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < x {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[nums2[j]] = x
		}
		stack = append(stack, i)
	}
	answer := make([]int, len(nums1))
	for i, x := range nums1 {
		if value, ok := next[x]; ok {
			answer[i] = value
		} else {
			answer[i] = -1
		}
	}
	return answer
}
```

**例子**：nums1=[4,1,2]、nums2=[1,3,4,2]，答案[-1,3,-1]。不是寻找右侧最大值。

**复杂度**：时间O(n+m)，空间O(n+m)，n、m为两数组长度。

<a id="lc84"></a>

## 4. [84. 柱状图中最大的矩形](https://leetcode.com/problems/largest-rectangle-in-histogram/)

[练习](l84/84.go) · [测试](l84/84_test.go)

**题意**：每根柱宽1，求能组成的最大矩形面积。

**栈含义**：柱高非递减，遇更矮柱时结算旧柱j。弹出后新栈顶left限制左侧扩展，宽为i-left-1。相等柱保留，较早的柱最终覆盖完整宽度。

```go
package template

func largestRectangleArea(heights []int) int {
	stack := []int{}
	var answer int64
	for i := 0; i <= len(heights); i++ {
		for len(stack) > 0 && (i == len(heights) || heights[stack[len(stack)-1]] > heights[i]) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			answer = max(answer, int64(heights[j])*int64(i-left-1))
		}
		if i < len(heights) {
			stack = append(stack, i)
		}
	}
	return int(answer)
}
```

**例子**：[2,1,2]，末尾清栈时高度1覆盖宽3，答案3。i=n强制结算，短路条件避免访问heights[n]。

**复杂度**：时间O(n)，空间O(n)。相等保留时left可为等高柱，不能误称每次都是“严格更小左边界”。

<a id="lc907"></a>

## 5. [907. 子数组的最小值之和](https://leetcode.com/problems/sum-of-subarray-minimums/)

[练习](l907/907.go) · [测试](l907/907_test.go)

**题意**：所有非空子数组的最小值求和，对10^9+7取模。

**贡献法**：为i找左侧最近严格小于L、右侧最近小于等于R。以i为指定最小值的区间数是(i-L)×(R-i)，贡献再乘arr[i]。

```go
package template

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i, x := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= x {
			stack = stack[:len(stack)-1]
		}
		left[i] = -1
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		right[i] = n
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	const mod int64 = 1000000007
	var answer int64
	for i, x := range arr {
		contribution := int64(x) * int64(i-left[i]) * int64(right[i]-i)
		answer = (answer + contribution) % mod
	}
	return int(answer)
}
```

**例子**：[2,2]，left=[-1,-1]、right=[1,2]，贡献2和4，总和6。双元素区间归给右侧的2，防止重复。

**复杂度**：时间O(n)，空间O(n)。先转int64再乘；本题约束下乘积可表示。

<a id="pitfalls"></a>

## 6. 对比与易错点

| 题目 | 结算内容 | 相等值处理 |
| --- | --- | --- |
| 739 | 等待天数i-j | 相等不弹 |
| 496 | 下一个更大值 | 题目值唯一 |
| 84 | 高×宽 | 本实现相等保留，由早柱覆盖完整范围 |
| 907 | 最小值×左右选择数 | 左严格、右非严格，统一归属 |

- 栈存下标才能算距离、宽度；不要混淆当前i和被弹出的j。
- 左不存在用-1，右不存在用n，但宽度公式与贡献公式不同。
- 907两侧都严格会重复计数，两侧都非严格会遗漏。
- 遍历方向和弹栈含义改变时，比较符号也要重新推导。
- 窗口最大值的过期处理需要单调队列，不是这个栈骨架。

已有分类：[monotonic_stack](../../monotonic_stack/)。
