# Top K堆与双堆：按题目独立实现

[返回总目录](../README.md)

## 目录

- [1. 核心思路](#core)
- [2. 215：数组第K大](#lc215)
- [3. 703：数据流第K大](#lc703)
- [4. 347：前K个高频元素](#lc347)
- [5. 973：最近K个点](#lc973)
- [6. 295：数据流中位数（双堆对照）](#lc295)
- [7. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

保留最好的K个候选，堆顶是其中最差者，方便淘汰：最大K个用最小堆，最小K个用最大堆。

```text
加入新候选 → 大小超过K则弹出最差 → 剩余候选始终是已见最好的K个
```

重复值按出现次数计排名；堆内切片不是有序结果。295需要两个堆平衡，不能直接套固定K；放在本篇用于对照。

以下每题包含自己的container/heap辅助类型，代码块相互独立。Less决定堆方向，接口Pop只移除末尾，由heap.Pop负责堆顶交换。

<a id="lc215"></a>

## 2. [215. 数组中的第K个最大元素](https://leetcode.com/problems/kth-largest-element-in-an-array/)

[练习](l215/215.go) · [测试](l215/215_test.go)

**题意**：返回多重集合意义的第k大，不是第k个不同值，题设1<=k<=n。

**状态**：保留最大k个，最小堆顶即第k大。

```go
package template

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
```

**例子**：[5,5,4]、k=2，答案5，不去重。

**复杂度**：时间O(n log(k+1))，空间O(k)。也可快速选择，但不是本模板重点。

<a id="lc703"></a>

## 3. [703. 数据流中的第K大元素](https://leetcode.com/problems/kth-largest-element-in-a-stream/)

[练习](l703/703.go) · [测试](l703/703_test.go)

**题意**：初始化已有流，每次Add后返回当前第k大；题设保证调用返回时元素足够。

**变化**：将215中的堆作为持久状态，构造时依次Add初始数据；不足k时内部堆顶只是现有最小值，构造不使用这个返回值。

```go
package template

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

type KthLargest struct {
	k      int
	values minHeap
}

func Constructor(k int, nums []int) KthLargest {
	result := KthLargest{k: k}
	for _, x := range nums {
		result.Add(x)
	}
	return result
}

func (s *KthLargest) Add(val int) int {
	heap.Push(&s.values, val)
	if len(s.values) > s.k {
		heap.Pop(&s.values)
	}
	return s.values[0]
}
```

**例子**：k=3、初始[4,5,8,2]，依次Add(3)、Add(5)，返回4、5。

**复杂度**：构造O(n log(k+1))，每次Add O(log(k+1))，空间O(k)。通用接口若允许查询不足k个的流，应另返回ok，不把堆顶误称第k大。

<a id="lc347"></a>

## 4. [347. 前K个高频元素](https://leetcode.com/problems/top-k-frequent-elements/)

[练习](l347/347.go) · [测试](l347/347_test.go)

**题意**：返回出现次数最多的k个不同值，顺序任意。

**变化**：先统计所有频次，堆元素为(value,count)，比较count而不是value。不能在频次尚未统计完成时直接丢弃值。

```go
package template

import "container/heap"

type entry struct{ value, count int }
type minHeap []entry

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, x := range nums {
		freq[x]++
	}
	h := &minHeap{}
	for x, count := range freq {
		heap.Push(h, entry{x, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	answer := make([]int, h.Len())
	for i, e := range *h {
		answer[i] = e.value
	}
	return answer
}
```

**例子**：[100,1,1,1,2,2]、k=1，答案[1]，不是[100]。

**复杂度**：U个不同值，平均时间O(n+U log(k+1))，空间O(U+k)。题设答案集合唯一，组内输出顺序不限。

<a id="lc973"></a>

## 5. [973. 最接近原点的K个点](https://leetcode.com/problems/k-closest-points-to-origin/)

[练习](l973/973.go) · [测试](l973/973_test.go)

**题意**：保留距离原点最近的k个点，输出顺序任意。

**变化**：保留最小k个距离，所以用最大堆淘汰最远者；平方距离即可，无需开方。

```go
package template

import "container/heap"

type point struct {
	x, y     int
	distance int64
}
type maxHeap []point

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(point)) }
func (h *maxHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func kClosest(points [][]int, k int) [][]int {
	h := &maxHeap{}
	for _, p := range points {
		x, y := int64(p[0]), int64(p[1])
		heap.Push(h, point{p[0], p[1], x*x + y*y})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	answer := make([][]int, h.Len())
	for i, p := range *h {
		answer[i] = []int{p.x, p.y}
	}
	return answer
}
```

**例子**：[[1,3],[-2,2]]、k=1，平方距离10与8，返回[[-2,2]]。重复坐标仍是多个输入点，不能用集合去重。

**复杂度**：时间O(n log(k+1))，空间O(k)。平方前转int64，覆盖本题坐标范围。

<a id="lc295"></a>

## 6. [295. 数据流的中位数](https://leetcode.com/problems/find-median-from-data-stream/)

[练习](l295/295.go) · [测试](l295/295_test.go)

**题意**：动态加入数字，查询中位数，保证查询时至少一个数字。

**双堆不变量**：low为较小一半的最大堆，high为较大一半的最小堆；low的数量等于high或多1，且low中所有值<=high中所有值。

```go
package template

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

type MedianFinder struct {
	low  maxHeap
	high minHeap
}

func Constructor() MedianFinder { return MedianFinder{} }

func (m *MedianFinder) AddNum(num int) {
	if len(m.low) == 0 || num <= m.low[0] {
		heap.Push(&m.low, num)
	} else {
		heap.Push(&m.high, num)
	}
	if len(m.low) > len(m.high)+1 {
		heap.Push(&m.high, heap.Pop(&m.low).(int))
	}
	if len(m.high) > len(m.low) {
		heap.Push(&m.low, heap.Pop(&m.high).(int))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if len(m.low) > len(m.high) {
		return float64(m.low[0])
	}
	return (float64(m.low[0]) + float64(m.high[0])) / 2
}
```

**例子**：加入[5,15,1,3]，中位数依次5、10、5、4。偶数时先转float64再加，避免整数相加溢出；精度覆盖本题范围。

**复杂度**：Add O(log(n+1))，查询O(1)，空间O(n)。双堆保存全部数据，不是固定K空间。

<a id="pitfalls"></a>

## 7. 对比与易错点

| 题目 | 排名键 | 堆方向 | 保留数量 |
| --- | --- | --- | --- |
| 215 / 703 | 值 | 最小堆淘汰小值 | K |
| 347 | 频次 | 最小堆淘汰低频 | K，另有频率map |
| 973 | 距离平方 | 最大堆淘汰远点 | K |
| 295 | 值 | 较小半最大堆＋较大半最小堆 | 全部 |

- 空堆天然满足堆性质；若用非空切片直接初始化，需要heap.Init。
- 自定义Pop只取切片末尾，不要在其中再次执行删除堆顶。
- 输出顺序有要求时另外排序副本；堆内部切片并非有序。
- 同频率若要求字典序，要补次级比较；不能依赖map遍历顺序。
- 固定K模板不自动支持删除流中旧元素，也不能直接求任意分位数。

已有分类：[heap_priority_queue](../../heap_priority_queue/)。
