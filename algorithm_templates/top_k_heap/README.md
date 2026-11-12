# 固定大小 K 的堆：保留最好的 K 个

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [Go 通用骨架](#code)
- [变化点与推演](#mapping)
- [复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

动态加入数据，维护最大 K 个、最小 K 个或按评分最优 K 个。堆顶应该是**当前保留候选里最差的那个**，方便新元素与淘汰门槛比较。

求最大 K 个 → 最小堆；求最小 K 个 → 最大堆。不是看到“最大”就用最大堆。

**不变量**：处理 t 个元素后，堆大小为 min(K,t)，保存已见多重集合中最大的这些元素。t>=K 时，堆顶才是第 K 大；t<K 时不存在第 K 大，应返回状态而不是误报堆顶。

这里重复值按多次出现计入排名，例如 [5,5,4] 的第2大为5，不是4；“第K个不同值”是另一种要求。

<a id="code"></a>

## 2. Go 通用骨架

```go
package template

import "container/heap"

type IntMinHeap []int
func(h IntMinHeap) Len() int { return len(h) }
func(h IntMinHeap) Less(i,j int) bool { return h[i]<h[j] }
func(h IntMinHeap) Swap(i,j int) { h[i],h[j]=h[j],h[i] }
func(h *IntMinHeap) Push(x any) { *h=append(*h,x.(int)) }
func(h *IntMinHeap) Pop() any {
    old:=*h
    x:=old[len(old)-1]
    *h=old[:len(old)-1]
    return x
}

type TopK struct { k int; values IntMinHeap }

func NewTopK(k int) *TopK {
    if k<=0 { panic("k must be positive") }
    return &TopK{k:k} // 空堆天然满足堆性质
}

// ok=false 表示已见元素不足 k 个，不存在第 k 大。
func(t *TopK) Add(x int) (value int,ok bool) {
    if len(t.values)<t.k {
        heap.Push(&t.values,x)
    } else if x>t.values[0] {
        t.values[0]=x
        heap.Fix(&t.values,0)
    }
    if len(t.values)<t.k { return 0,false }
    return t.values[0],true
}

// 返回未排序的候选副本，不暴露内部堆给调用方修改。
func(t *TopK) Values() []int {
    return append([]int(nil),t.values...)
}
```

固定过程：未满先加；满时与堆顶比较；更优才替换并恢复堆性质。变化点是元素类型和 Less / 新元素比较条件，二者必须使用同一排名规则。

<a id="mapping"></a>

## 3. 变化点与推演

K=3，依次加入 [4,13,8,2,5]：

| 加入 | 保留集合（非堆内实际顺序） | 第3大 |
| --- | --- | --- |
| 4 | {4} | 不存在 |
| 13 | {4,13} | 不存在 |
| 8 | {4,8,13} | 4 |
| 2 | {4,8,13} | 4 |
| 5 | {5,8,13} | 5 |

新值等于堆顶时可以忽略，因为保留的数值多重集合无需改变，排名答案相同；这不等于对输入去重。

| 题目 | 元素 / 比较键 | 骨架的调用方式 | 额外变化 |
| --- | --- | --- | --- |
| [703. 数据流中的第 K 大元素](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | 整数值 | 初始化流并持续 Add | 题目保证查询时满足其第K大约束；通用接口显式返回 ok |
| [215. 数组中的第 K 个最大元素](https://leetcode.com/problems/kth-largest-element-in-an-array/) | 整数值 | 遍历一次，最后取堆顶 | 要求1<=K<=n；也可使用快速选择 |
| [347. 前 K 个高频元素](https://leetcode.com/problems/top-k-frequent-elements/) | (值,出现次数)，按次数比较 | 先统计所有频次，再遍历不同值入堆 | 不是按值大小排序；通常无需输出有序 |
| [973. 最接近原点的 K 个点](https://leetcode.com/problems/k-closest-points-to-origin/) | 距离平方 | 需要保留最小K个，所以用最大堆 | 距离用 int64，避免平方溢出 |

347 的频率 map 占 O(U)，U 为不同值数量。不能直接用上述整数堆丢掉早期元素，之后才发现它出现次数最多：计数键与堆元素必须按问题正确建模；若频率在线改变，需额外维护更新或懒删除策略。

<a id="pitfalls"></a>

## 4. 复杂度与陷阱

- n 次加入总时间 O(n log(K+1))，空间 O(K)；347 另有 O(n) 统计时间和 O(U) 频率空间。
- 堆内切片不是排序结果，只保证堆顶最差；若要求排序输出，需要对副本排序或反复弹堆。
- Go 的 heap.Pop 会先把堆顶交换到末尾，自定义 Pop 只需取最后一项；不要在接口方法中再次删除堆顶。
- k<=0 要明确定义行为。本模板拒绝非法 k；“前K个返回全部现有值”和“存在第K大”是两个不同接口。
- 频率并列时若题目要求字典序，必须把次级规则也写进比较器，不能依赖 map 遍历顺序。
- 中位数通常需要双堆平衡，不是固定 K 模板直接可解；[295](https://leetcode.com/problems/find-median-from-data-stream/) 应另学双堆。

已有练习：[heap_priority_queue](../../heap_priority_queue/)、[703](../../heap_priority_queue/l703/)、[347](../../heap_priority_queue/l347/)。
