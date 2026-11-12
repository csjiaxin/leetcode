# 单调栈：最近边界与贡献结算

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [Go 骨架：两种遍历方式](#code)
- [变化点与推演](#mapping)
- [复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

需要某个位置左/右侧最近的更大/更小元素，或某个值作为区间最值的作用边界。栈存下标，值通过原数组访问。

**从左向右、弹栈结算**：栈保存尚未找到右侧答案的位置；当前值满足条件时，为弹出的旧下标确定右侧答案。

**从右向左、清除候选后查询**：栈保存右侧未被支配的候选；弹掉不可能作为当前答案的值，再由栈顶给当前下标提供答案。

两种写法的相等处理不同，不能只背“严格大于就用某一个比较符号”。

<a id="code"></a>

## 2. Go 骨架：两种遍历方式

```go
package template

// 返回右侧第一个严格更大元素的下标，不存在为 -1。
func NextGreaterLeftScan(nums []int) []int {
    answer := make([]int,len(nums))
    for i := range answer { answer[i] = -1 }
    stack := make([]int,0,len(nums))
    for i,x := range nums {
        for len(stack)>0 && nums[stack[len(stack)-1]]<x {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            answer[j] = i
        }
        stack = append(stack,i) // 值保持非递增，允许相等
    }
    return answer
}

func NextGreaterRightScan(nums []int) []int {
    answer := make([]int,len(nums))
    stack := make([]int,0,len(nums))
    for i:=len(nums)-1;i>=0;i-- {
        for len(stack)>0 && nums[stack[len(stack)-1]]<=nums[i] {
            stack = stack[:len(stack)-1] // 相等不是严格更大候选
        }
        answer[i] = -1
        if len(stack)>0 { answer[i] = stack[len(stack)-1] }
        stack = append(stack,i)
    }
    return answer
}

// 非负柱高；遇到更低柱或结束时结算被弹出的柱子。
func LargestRectangle(heights []int) int64 {
    stack := []int{}
    var answer int64
    for i:=0;i<=len(heights);i++ {
        for len(stack)>0 && (i==len(heights) || heights[stack[len(stack)-1]]>heights[i]) {
            j:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            left:=-1
            if len(stack)>0 { left=stack[len(stack)-1] }
            area:=int64(heights[j])*int64(i-left-1)
            answer=max(answer,area)
        }
        if i<len(heights) { stack=append(stack,i) }
    }
    return answer
}
```

最后一个函数在 i=n 时强制清空栈，短路条件避免访问 heights[n]；不用修改输入来追加哨兵。

<a id="mapping"></a>

## 3. 变化点与推演

nums=[2,2,3]，左向右结算：

```text
i=0：栈 [0]
i=1：相等不弹，栈 [0,1]
i=2：3>2，弹 1，answer[1]=2；再弹 0，answer[0]=2
结果 [2,2,-1]
```

| 题目 | 比较 / 栈的含义 | 弹出时或查询后做什么 |
| --- | --- | --- |
| [739. 每日温度](https://leetcode.com/problems/daily-temperatures/) | 左向右遇到严格更大温度才弹 | answer[j]=i-j，不存在时为 0 |
| [496. 下一个更大元素 I](https://leetcode.com/problems/next-greater-element-i/) | 找严格更大位置 | 从下标取数值，再建立题目要求的映射 |
| [84. 柱状图中最大的矩形](https://leetcode.com/problems/largest-rectangle-in-histogram/) | 非递减柱高栈，遇低柱结算 | 面积 height[j]×(right-left-1) |
| [907. 子数组的最小值之和](https://leetcode.com/problems/sum-of-subarray-minimums/) | 为每个最小值划分左右边界 | 贡献 nums[i]×(i-L)×(R-i)，按题意取模 |

907 的一种正确归属：L 为左侧最近**严格小于**，R 为右侧最近**小于等于**。这样 [2,2] 的双元素区间归于右侧那个 2，恰计一次。反向不对称也可，但两边都严格会重复、两边都非严格会遗漏。

84 示例 [2,1,2]：末尾清栈时高度 1 覆盖三根，面积为 3。相等柱在本骨架中保留，靠较早柱最终覆盖完整宽度；不要把中间弹出的相等柱宽度当作严格小于边界。

<a id="pitfalls"></a>

## 4. 复杂度与陷阱

- 每个下标入栈一次、出栈至多一次，总时间 O(n)，额外空间 O(n)。内层 while 不等于 O(n²)。
- 存值无法直接算等待天数或宽度，应存下标。
- `>`、`>=` 的选择同时依赖遍历方向、弹栈含义及重复值归属，必须写出不变量。
- 栈空时左边界常用 -1；右侧不存在常用 n；面积宽度为 R-L-1。
- 贡献乘法先转 int64，再相乘；转型太晚不能挽救已经溢出的 int 运算。
- 栈不是滑动窗口队列，不负责自动移除过期的左端位置。窗口最大值可考虑单调队列。

已有练习：[monotonic_stack](../../monotonic_stack/)，重点 [84](../../monotonic_stack/l84/)、[907](../../monotonic_stack/l907/)。
