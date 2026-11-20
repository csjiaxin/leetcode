# 树状数组 Fenwick Tree：动态前缀聚合与排名计数

[返回总目录](../../README.md) · [树目录](../readme.md) · [线段树对照](../segment/readme.md)

本目录为学习指南与参考模板，不代表所有练习已有独立题解。普通树状数组优先用于**单点增量更新＋前缀和 / 区间和**，比线段树更简洁，但通用性更弱。

## 分类与识别

| 需求 | 模型 | 代表题 |
| --- | --- | --- |
| 动态数组求和 | 单点加 delta，两个前缀相减 | 307 区域和检索——数组可修改 |
| 右侧较小元素数量 | 反向扫描、值域离散化、前缀频次 | 315 |
| 不同阈值的数对计数 | 压缩值域并二分查询阈值 | 493 翻转对 |
| 前缀和之差在范围内 | 压缩前缀和，统计排名区间频次 | 327 区间和的个数 |

## 模板一：公开 0-based，内部 1-based

```go
import "sort"

type Fenwick struct { tree []int64 }
func NewFenwick(n int) *Fenwick {
    if n < 0 { panic("negative length") }
    return &Fenwick{tree: make([]int64, n+1)}
}
func (f *Fenwick) Add(index int, delta int64) {
    if index < 0 || index >= len(f.tree)-1 { panic("index out of range") }
    for i := index+1; i < len(f.tree); i += i&-i { f.tree[i] += delta }
}
// Prefix(end) 返回 [0,end) 的和，允许 end=0 和 end=n。
func (f *Fenwick) Prefix(end int) int64 {
    if end < 0 || end >= len(f.tree) { panic("prefix out of range") }
    var sum int64
    for i := end; i > 0; i -= i&-i { sum += f.tree[i] }
    return sum
}
func (f *Fenwick) Range(left, right int) int64 {
    if left < 0 || right < left || right >= len(f.tree) { panic("range out of bounds") }
    return f.Prefix(right)-f.Prefix(left)
}
```

不变量：内部 tree[i] 保存长度为 `lowbit(i)=i&-i`、以 i 结尾的一段 1-based 数据。查询不断移除最低置位，恰好拆分出互不重叠前缀块；更新向所有覆盖该点的祖先传播。

Add、Prefix、Range 均 O(log n)，空间 O(n)。累加使用 int64，但仍应保证总和在类型范围内。Add 是增量更新，不是赋值；307 赋新值时应另外保存旧值并执行 `Add(index,new-old)`。

边界检查不能遗漏：若将非法 index=-1 转成内部 i=0，更新 `i+=i&-i` 将永不推进。

## 模板二：315 的离散化频次计数

```go
func CountSmaller(nums []int) []int {
    sorted := append([]int(nil), nums...)
    sort.Ints(sorted)
    unique := sorted[:0]
    for _, x := range sorted {
        if len(unique) == 0 || unique[len(unique)-1] != x { unique = append(unique, x) }
    }
    counts := NewFenwick(len(unique))
    result := make([]int, len(nums))
    for i := len(nums)-1; i >= 0; i-- {
        rank := sort.SearchInts(unique, nums[i])
        result[i] = int(counts.Prefix(rank))
        counts.Add(rank, 1)
    }
    return result
}
```

扫描 i 前，树内只保存右侧已处理元素的频次。`Prefix(rank)` 不包含自己这个排名，因此只统计严格更小值；相等值不会被计入。重复值共享排名，但每次出现都加 1。时间 O(n log n)，空间 O(n)，不修改输入。

## 与其他结构的边界

- 静态区间和：普通前缀和即可，查询 O(1)，无需树状数组。
- 区间最小值 / 最大值、任意结合运算：普通前缀和模板不能直接替换运算符，优先考虑线段树。
- 区间加、单点查：差分后可用一棵树；区间加、区间和：需推导两棵树公式，不是直接调用 Add 更新两个端点就完成区间和。
- 通过前缀频次寻找第 k 个元素：需要频次非负，使前缀单调；可进阶学习二进制提升。
- 493 查询阈值不一定在压缩数组中，要用二分边界定位；乘 2 前先提升类型。

检查：空树、单点、Prefix(0)、Prefix(n)、重复值、负增量、赋值与增量区别。练习：307 → 315 → 493 → 327；与[归并计数](../../sorting/README.md)对照。
