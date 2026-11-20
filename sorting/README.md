# 排序、分治与快速选择

[返回总目录](../README.md)

本目录补充排序算法本身的学习指南与 Go 模板，不移动已有题解，也不表示题单均已实现。

## 如何选型

| 需求 | 思路 | 代表题 |
| --- | --- | --- |
| 获得完整顺序 | 归并排序、随机快排、堆排序 | 912 排序数组 |
| 只要第 k 大 / 小 | Quickselect：只保留目标排名所在区间 | 215 第 K 个最大元素 |
| 统计跨区间关系 | 两半有序后用指针计数，再归并 | 315 右侧小于当前的个数、493 翻转对、327 区间和的个数 |
| 链表排序 | 分割与合并，不依赖随机下标 | 148 排序链表 |
| 合并多条有序序列 | 两两分治或最小堆 | 23 合并 K 个升序链表 |

已有 [215 堆解法](../heap_priority_queue/l215/215.go)、[912 堆排序](../heap_priority_queue/l912/912.go)，新模板用于比较不同不变量。

## 模板一：稳定归并排序

```go
import "math/rand"

func MergeSort(a []int) {
    buffer := make([]int, len(a))
    var solve func(int, int)
    solve = func(left, right int) {
        if right-left <= 1 { return }
        middle := left+(right-left)/2
        solve(left, middle); solve(middle, right)
        i, j := left, middle
        for k := left; k < right; k++ {
            if j == right || i < middle && a[i] <= a[j] {
                buffer[k] = a[i]; i++
            } else {
                buffer[k] = a[j]; j++
            }
        }
        copy(a[left:right], buffer[left:right])
    }
    solve(0, len(a))
}
```

不变量：合并时两半已经有序，buffer 已写部分是两半已消费元素的有序合并。相等取左边保证稳定性。时间 O(n log n)，额外数组 O(n)、调用栈 O(log n)，原地改写输入但不是常数空间排序。

## 模板二：随机三路快速选择

```go
// k 为从 0 开始的升序排名；会改变输入顺序。
func SelectKth(a []int, k int) int {
    if k < 0 || k >= len(a) { panic("rank out of range") }
    left, right := 0, len(a)
    for {
        pivot := a[left+rand.Intn(right-left)]
        less, current, greater := left, left, right
        for current < greater {
            if a[current] < pivot {
                a[less], a[current] = a[current], a[less]
                less++; current++
            } else if a[current] > pivot {
                greater--
                a[current], a[greater] = a[greater], a[current]
            } else { current++ }
        }
        if k < less { right = less } else if k >= greater { left = greater } else { return a[k] }
    }
}
```

分区不变量：`[left,less)<pivot`、`[less,current)==pivot`、`[current,greater)` 未知、`[greater,right)>pivot`。换入未知值时不能直接 current++。相等区间一次排除，适合重复值。

期望时间 O(n)，最坏 O(n²)，迭代辅助空间 O(1)。求第 k 大（k 从 1 开始）调用 `SelectKth(a,len(a)-k)`。随机化不是最坏线性保证；库随机源适合算法分区，不用于安全用途。

## 从模板迁移

- **快排**：分区后递归处理两侧，而不是只保留一侧；随机快排期望 O(n log n)，最坏 O(n²)。递归空间也不能一概写成 O(log n)。
- **315**：归并时连同原下标排序，记录右半已移入的严格更小元素数量，相等不应误计。
- **493**：先对每个左半值推进右半指针统计 `left>2*right`，再普通归并；乘法要先提升类型。它不是普通逆序对条件。
- **327**：对前缀和归并，右半用两指针计数差值落在 `[lower,upper]` 的数量。

## 易错点与路线

统一半开区间 `[left,right)`；不要混淆排名从 0 还是 1 开始；自定义比较器必须满足严格序，不能用 `<=`；做乘法或减法比较时注意溢出。需要稳定排序、在线维护或完整顺序时，不能仅按“Quickselect 更快”选型。

练习：912 归并排序 → 215 快速选择 → 148 → 315 → 493 → 327。
