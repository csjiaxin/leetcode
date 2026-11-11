# 优先级队列题目

优先级队列适用于：候选集合不断变化，每次需要取出当前最小、最大或优先级最高的元素。堆只保证堆顶最优，**底层切片并非整体有序**。

## 题型与推荐练习

### 1. 堆基础与模拟

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [1046. 最后一块石头的重量](https://leetcode.com/problems/last-stone-weight/) | 简单 | 最大堆，每次弹出最大的两块石头，放回差值 |
| [2558. 从数量最多的堆取走礼物](https://leetcode.com/problems/take-gifts-from-the-richest-pile/) | 简单 | 反复更新当前最大值 |
| [1962. 移除石子使总数最小](https://leetcode.com/problems/remove-stones-to-minimize-the-total/) | 中等 | 最大堆，每次移除最大堆的一半石子 |
| [1753. 移除石子的最大得分](https://leetcode.com/problems/maximum-score-from-removing-stones/) | 中等 | 最大堆模拟取两堆，也可推导数学解法；本目录已有 `1753` |
| [912. 排序数组](https://leetcode.com/problems/sort-an-array/) | 中等 | 练习建堆与堆排序；本目录已有 `912` |

注意：1046 固定取最大的两块；1049 可以自由选择，目标通常转为背包 DP，不能直接套用相同贪心。

### 2. Top K：固定大小的候选集合

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [215. 数组中的第 K 个最大元素](https://leetcode.com/problems/kth-largest-element-in-an-array/) | 中等 | 保留 k 个最大元素，用大小为 k 的最小堆 |
| [703. 数据流中的第 K 大元素](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | 简单 | 每次加入后只保留最大的 k 个，堆顶就是第 k 大 |
| [347. 前 K 个高频元素](https://leetcode.com/problems/top-k-frequent-elements/) | 中等 | 先计数，再按频率维护 k 个候选；也可用桶排序 |
| [692. 前 K 个高频单词](https://leetcode.com/problems/top-k-frequent-words/) | 中等 | 频率相同还要比较字典序，注意堆顶代表“最该淘汰”的候选 |
| [973. 最接近原点的 K 个点](https://leetcode.com/problems/k-closest-points-to-origin/) | 中等 | 保留 k 个最近点，用最大堆淘汰最远点；比较距离平方即可 |

口诀：**保留最大的 k 个，用最小堆；保留最小的 k 个，用最大堆。** 堆顶是候选集合中最差、最先被淘汰的元素。典型时间 O(n log k)，堆空间 O(k)；频率题还需要计数表。

### 3. 多路归并与有序候选扩展

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [23. 合并 K 个升序链表](https://leetcode.com/problems/merge-k-sorted-lists/) | 困难 | 最小堆维护每条链表当前节点，弹出后加入该链表的下一个节点 |
| [378. 有序矩阵中第 K 小的元素](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/) | 中等 | 把每一行看成有序序列，最小堆归并 |
| [373. 查找和最小的 K 对数字](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/) | 中等 | 固定第一个数组的下标形成多条有序候选链，弹出后扩展下一对 |
| [632. 最小区间](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/) | 困难 | 堆维护各列表当前最小元素，额外维护当前最大值 |
| [264. 丑数 II](https://leetcode.com/problems/ugly-number-ii/) | 中等 | 最小堆生成候选并去重；也有三指针 DP 解法 |

关键：不要一次性生成所有组合。只保留每条有序序列的“前沿”，需要时再扩展，并根据生成方式避免重复。

### 4. 贪心：动态选择当前最优候选

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [871. 最低加油次数](https://leetcode.com/problems/minimum-number-of-refueling-stops/) | 困难 | 将已到达站点的油量加入最大堆，不足时选最多的补给 |
| [1642. 可以到达的最远建筑](https://leetcode.com/problems/furthest-building-you-can-reach/) | 中等 | 最小堆保留使用梯子的最大高度差，其余消耗砖块 |
| [630. 课程表 III](https://leetcode.com/problems/course-schedule-iii/) | 困难 | 按截止日期排序，超期时从最大堆移除耗时最长课程 |
| [2530. 执行 K 次操作后的最大分数](https://leetcode.com/problems/maximal-score-after-applying-k-operations/) | 中等 | 最大堆取分，更新为向上取整的三分之一后放回 |
| [857. 雇佣 K 名工人的最低成本](https://leetcode.com/problems/minimum-cost-to-hire-k-workers/) | 困难 | 按工资质量比排序，最大堆维护 k 个工人的最小质量总和 |

堆只是高效选择工具，**贪心正确性需要额外证明**。例如 871 只能选择已到达的站点，而不能直接选择全局油量最多的站点。

### 5. 调度、资源分配与事件模拟

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [2402. 会议室 III](https://leetcode.com/problems/meeting-rooms-iii/) | 困难 | 空闲房间按编号排序，忙碌房间按结束时间、编号排序，使用两个最小堆 |
| [1834. 单线程 CPU](https://leetcode.com/problems/single-threaded-cpu/) | 中等 | 按入队时间扫描，堆按处理时间、原下标选择任务 |
| [1882. 使用服务器处理任务](https://leetcode.com/problems/process-tasks-using-servers/) | 中等 | 空闲服务器按权重、编号，忙碌服务器按完成时间排序 |
| [621. 任务调度器](https://leetcode.com/problems/task-scheduler/) | 中等 | 最大堆选择剩余最多的任务，配合冷却队列；也有计数公式解法 |
| [767. 重构字符串](https://leetcode.com/problems/reorganize-string/) | 中等 | 最大堆优先使用高频字符，避免连续选同一字符 |

常见细节：先释放所有在当前时刻已结束的资源；没有可用任务或资源时，时间直接跳到下一事件；并列优先级要按题目规则处理。

### 6. 双堆、懒删除与动态统计

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [295. 数据流的中位数](https://leetcode.com/problems/find-median-from-data-stream/) | 困难 | 最大堆维护较小一半，最小堆维护较大一半，保持数量平衡 |
| [480. 滑动窗口中位数](https://leetcode.com/problems/sliding-window-median/) | 困难 | 双堆 + 懒删除，并区分逻辑元素数与切片长度 |
| [239. 滑动窗口最大值](https://leetcode.com/problems/sliding-window-maximum/) | 困难 | 最大堆存值和下标，弹出过期堆顶；单调队列可进一步优化到 O(n) |

懒删除只在过期元素到达堆顶时删除。被埋在内部的过期元素仍占空间：例如 239 的简单懒删除堆最坏可能增长到 O(n)，不是自动维持 O(k)。

### 7. 图上的优先扩展

| 题目 | 难度 | 核心思路 |
| --- | --- | --- |
| [743. 网络延迟时间](https://leetcode.com/problems/network-delay-time/) | 中等 | 非负边权 Dijkstra，最小堆弹出当前最短距离 |
| [1631. 最小体力消耗路径](https://leetcode.com/problems/path-with-minimum-effort/) | 中等 | Dijkstra 的路径代价改为沿途边权最大值 |
| [1514. 概率最大的路径](https://leetcode.com/problems/path-with-maximum-probability/) | 中等 | 最大堆按当前成功概率扩展 |
| [407. 接雨水 II](https://leetcode.com/problems/trapping-rain-water-ii/) | 困难 | 从边界最小高度开始向内扩展，维护水位边界 |

Dijkstra 常将更新后的状态重新入堆，弹出时跳过旧状态，不必支持原地 decrease-key。普通 BFS 适用于等权边；堆不能让标准 Dijkstra 正确处理负权边。

## Go：`container/heap`

标准库通过 `heap.Interface` 使用堆。它嵌入 `sort.Interface` 的三个方法，再要求实现 `Push`、`Pop`。标准库接口使用 `any`，并非泛型接口。

```go
import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}

// 在函数内部使用：
// h := &MinHeap{5, 2, 8}
// heap.Init(h)
// heap.Push(h, 1)
// if h.Len() > 0 {
//     smallest := heap.Pop(h).(int)
//     _ = smallest
// }
```

最大堆将 `Less` 改为 `h[i] > h[j]`，同时使用对应的接收者类型。

| 操作 | 用法 | 复杂度 |
| --- | --- | --- |
| 建堆 | `heap.Init(h)` | O(n) |
| 插入 | `heap.Push(h, x)` | O(log n) |
| 弹出堆顶 | `heap.Pop(h)` | O(log n) |
| 查看堆顶 | `(*h)[0]`，先检查非空 | O(1) |
| 修改后修复 | `heap.Fix(h, index)` | O(log n) |
| 删除指定元素 | `heap.Remove(h, index)` | O(log n) |

### 常见坑与现有示例

- 使用 `heap.Push` / `heap.Pop`，不要直接调用自定义的 `Push` / `Pop`：标准库会先调整堆，再调用这些底层方法。自定义 `Pop` 应删除**切片末尾**。
- `Less` 必须是严格比较，不能使用 `<=`。结构体有并列规则时，逐字段比较。
- 修改已有元素的优先级后，需要调用 `heap.Fix`；如果元素保存下标，`Swap` 必须同步更新下标。
- 本目录的 [pq.go](pq.go) 是结构体最大堆示例，包含 `priority`、`index`、`update` 和 `heap.Fix`。`index` 只有在正确维护且元素仍在堆中时才有效。
- 对指针切片，`Pop` 时将旧末尾设为 `nil`，避免底层数组继续持有对象引用；`pq.go` 已包含这种处理。
- 时间和累计值可能很大，例如会议室调度的累计结束时间，应按题目范围考虑 `int64`；避免通过可能溢出的减法比较优先级。

## 本仓库实现与测试

上述 32 道题的实现与测试位于对应的 `l<题号>/` 目录。题目包复用 [internal/pqueue/queue.go](internal/pqueue/queue.go) 中基于 `container/heap` 的泛型优先级队列；912 使用原地堆排序。

```sh
go test ./heap_priority_queue/internal/... ./heap_priority_queue/l... -count=1 -cover
go vet ./heap_priority_queue/internal/... ./heap_priority_queue/l...
```

部分题目的测试包含固定随机种子的朴素算法对照、浮点误差判断或合法答案验证，不依赖堆内部的元素顺序。调度题使用 `int64` 处理累计时间。

这些实现用于本仓库学习与测试；复制到 LeetCode 时，需要将内部优先队列一并改为同文件的堆实现，并移除仓库内导入路径。`pq.go` 展示了不依赖仓库辅助包的标准接口写法。

## 建议练习顺序

`1046 → 215 → 703 → 347 → 23 → 373 → 295 → 871 → 1642 → 1834 → 630 → 480 → 743 → 407`

先掌握最小堆/最大堆，再学固定大小 Top K、多路归并，最后练贪心、双堆、调度和图搜索。各题也可能有比堆更合适的替代解法。

## 参考资料

- [Go 官方 container/heap 文档](https://pkg.go.dev/container/heap)
- [LeetCode 堆（优先队列）题目](https://leetcode.com/tag/heap-priority-queue/)
- [Master Heap by Solving 23 Questions in 4 Patterns](https://leetcode.com/discuss/general-discussion/1127238/master-heap-by-solving-23-questions-in-4-patterns-category)
