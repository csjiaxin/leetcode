# 优先队列：统一 Queue 与候选状态适配

[返回模板总目录](../algorithm_templates/README.md)

Queue 的 less(a,b) 表示 a 的优先级更高，Peek/Pop 的前提是堆非空。固定容量 Top K、双堆平衡、有序来源归并、事件调度及图松弛共享候选容器，各自保留合法性和更新顺序。

Top K 入堆后超额弹出最差候选；调度先释放结束事件；双堆维护大小及顺序；最短路弹出后丢弃过期状态。912 是原地堆排序，单列 siftDown 子模板；不能把 container/heap 的堆序和排序数组的顺序混淆。

## 复制约定

每题 `.go` 文件已内嵌所需模板，无需导入仓库的算法包。文档将固定模板只列一次，逐题展示适配器；每个适配器与它列出的模板组成一个独立代码组。不同题的同名 Constructor、类型或入口请分别使用。将列出的标准库 import 放在文件顶部，再放模板与适配器；复制提交时优先直接使用对应题目的完整 `.go`。

本组使用 Go 1.21+ 的 min/max 和 Go 1.18+ 的泛型。平台已提供的 ListNode 等类型沿用平台定义。

## 固定模板

### Queue

<!-- template: Queue -->
```go
type heapData[T any] struct {
	items []T
	less  func(T, T) bool
}

func (d heapData[T]) Len() int           { return len(d.items) }
func (d heapData[T]) Less(i, j int) bool { return d.less(d.items[i], d.items[j]) }
func (d heapData[T]) Swap(i, j int)      { d.items[i], d.items[j] = d.items[j], d.items[i] }
func (d *heapData[T]) Push(x any)        { d.items = append(d.items, x.(T)) }
func (d *heapData[T]) Pop() any {
	n := len(d.items) - 1
	x := d.items[n]
	var zero T
	d.items[n] = zero
	d.items = d.items[:n]
	return x
}

// Queue must not be copied after use. less(a,b) means a has higher priority.
// Peek and Pop require a nonempty queue.
type Queue[T any] struct{ d heapData[T] }

// NewQueue copies the initial values and builds the heap in O(n).
func NewQueue[T any](less func(T, T) bool, values ...T) *Queue[T] {
	q := &Queue[T]{d: heapData[T]{items: append([]T(nil), values...), less: less}}
	heap.Init(&q.d)
	return q
}
func (q *Queue[T]) Len() int { return q.d.Len() }
func (q *Queue[T]) Peek() T  { return q.d.items[0] }
func (q *Queue[T]) Push(x T) { heap.Push(&q.d, x) }
func (q *Queue[T]) Pop() T   { return heap.Pop(&q.d).(T) }
```

### HeapSort

<!-- template: HeapSort -->
```go
func siftDown(a []int, root, end int) {
	for root < end/2 {
		child := 2*root + 1
		if child+1 < end && a[child+1] > a[child] {
			child++
		}
		if a[root] >= a[child] {
			return
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}
```

## 题目索引

| 题目 | 模板 | 本地代码与测试 |
| --- | --- | --- |
| [23](https://leetcode.com/problems/merge-k-sorted-lists/) | Queue | [代码](l23/23.go) · [测试](l23/23_test.go) |
| [215](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Queue | [代码](l215/215.go) · [测试](l215/215_test.go) |
| [239](https://leetcode.com/problems/sliding-window-maximum/) | Queue | [代码](l239/239.go) · [测试](l239/239_test.go) |
| [264](https://leetcode.com/problems/ugly-number-ii/) | Queue | [代码](l264/264.go) · [测试](l264/264_test.go) |
| [295](https://leetcode.com/problems/find-median-from-data-stream/) | Queue | [代码](l295/295.go) · [测试](l295/295_test.go) |
| [347](https://leetcode.com/problems/top-k-frequent-elements/) | Queue | [代码](l347/347.go) · [测试](l347/347_test.go) |
| [373](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/) | Queue | [代码](l373/373.go) · [测试](l373/373_test.go) |
| [378](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/) | Queue | [代码](l378/378.go) · [测试](l378/378_test.go) |
| [407](https://leetcode.com/problems/trapping-rain-water-ii/) | Queue | [代码](l407/407.go) · [测试](l407/407_test.go) |
| [480](https://leetcode.com/problems/sliding-window-median/) | Queue | [代码](l480/480.go) · [测试](l480/480_test.go) |
| [621](https://leetcode.com/problems/task-scheduler/) | Queue | [代码](l621/621.go) · [测试](l621/621_test.go) |
| [630](https://leetcode.com/problems/course-schedule-iii/) | Queue | [代码](l630/630.go) · [测试](l630/630_test.go) |
| [632](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/) | Queue | [代码](l632/632.go) · [测试](l632/632_test.go) |
| [692](https://leetcode.com/problems/top-k-frequent-words/) | Queue | [代码](l692/692.go) · [测试](l692/692_test.go) |
| [703](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | Queue | [代码](l703/703.go) · [测试](l703/703_test.go) |
| [743](https://leetcode.com/problems/network-delay-time/) | Queue | [代码](l743/743.go) · [测试](l743/743_test.go) |
| [767](https://leetcode.com/problems/reorganize-string/) | Queue | [代码](l767/767.go) · [测试](l767/767_test.go) |
| [857](https://leetcode.com/problems/minimum-cost-to-hire-k-workers/) | Queue | [代码](l857/857.go) · [测试](l857/857_test.go) |
| [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) | Queue | [代码](l871/871.go) · [测试](l871/871_test.go) |
| [912](https://leetcode.com/problems/sort-an-array/) | HeapSort | [代码](l912/912.go) · [测试](l912/912_test.go) |
| [973](https://leetcode.com/problems/k-closest-points-to-origin/) | Queue | [代码](l973/973.go) · [测试](l973/973_test.go) |
| [1046](https://leetcode.com/problems/last-stone-weight/) | Queue | [代码](l1046/1046.go) · [测试](l1046/1046_test.go) |
| [1514](https://leetcode.com/problems/path-with-maximum-probability/) | Queue | [代码](l1514/1514.go) · [测试](l1514/1514_test.go) |
| [1631](https://leetcode.com/problems/path-with-minimum-effort/) | Queue | [代码](l1631/1631.go) · [测试](l1631/1631_test.go) |
| [1642](https://leetcode.com/problems/furthest-building-you-can-reach/) | Queue | [代码](l1642/1642.go) · [测试](l1642/1642_test.go) |
| [1753](https://leetcode.com/problems/maximum-score-from-removing-stones/) | Queue | [代码](l1753/1753.go) · [测试](l1753/1753_test.go) |
| [1834](https://leetcode.com/problems/single-threaded-cpu/) | Queue | [代码](l1834/1834.go) · [测试](l1834/1834_test.go) |
| [1882](https://leetcode.com/problems/process-tasks-using-servers/) | Queue | [代码](l1882/1882.go) · [测试](l1882/1882_test.go) |
| [1962](https://leetcode.com/problems/remove-stones-to-minimize-the-total/) | Queue | [代码](l1962/1962.go) · [测试](l1962/1962_test.go) |
| [2402](https://leetcode.com/problems/meeting-rooms-iii/) | Queue | [代码](l2402/2402.go) · [测试](l2402/2402_test.go) |
| [2530](https://leetcode.com/problems/maximal-score-after-applying-k-operations/) | Queue | [代码](l2530/2530.go) · [测试](l2530/2530_test.go) |
| [2558](https://leetcode.com/problems/take-gifts-from-the-richest-pile/) | Queue | [代码](l2558/2558.go) · [测试](l2558/2558_test.go) |

## 逐题适配器

### [23](https://leetcode.com/problems/merge-k-sorted-lists/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l23/23.go; templates: Queue; imports: container/heap -->
```go
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists reuses nodes; the heap holds one frontier per nonempty list.
// O(N log k) time and O(k) extra space for N total nodes.
func mergeKLists(lists []*ListNode) *ListNode {
	q := NewQueue(func(a, b *ListNode) bool { return a.Val < b.Val })
	for _, node := range lists {
		if node != nil {
			q.Push(node)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for q.Len() > 0 {
		node := q.Pop()
		if node.Next != nil {
			q.Push(node.Next)
		}
		tail.Next = node
		tail = node
	}
	if tail != dummy {
		tail.Next = nil
	}
	return dummy.Next
}
```

### [215](https://leetcode.com/problems/kth-largest-element-in-an-array/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l215/215.go; templates: Queue; imports: container/heap -->
```go
// findKthLargest keeps the largest k values in a minimum heap.
// O(n log k) time, O(k) space. Requires 1 <= k <= len(nums).
func findKthLargest(nums []int, k int) int {
	q := NewQueue(func(a, b int) bool { return a < b })
	for _, x := range nums {
		q.Push(x)
		if q.Len() > k {
			q.Pop()
		}
	}
	return q.Peek()
}
```

### [239](https://leetcode.com/problems/sliding-window-maximum/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l239/239.go; templates: Queue; imports: container/heap -->
```go
type entry struct{ value, index int }

// maxSlidingWindow removes expired entries when they reach the heap top.
// Stale interior entries can accumulate: O(n log n) time, O(n) auxiliary space.
func maxSlidingWindow(nums []int, k int) []int {
	out := make([]int, 0)
	if k <= 0 || k > len(nums) {
		return out
	}
	q := NewQueue(func(a, b entry) bool {
		if a.value != b.value {
			return a.value > b.value
		}
		return a.index > b.index
	})
	for i, v := range nums {
		q.Push(entry{v, i})
		for q.Peek().index <= i-k {
			q.Pop()
		}
		if i >= k-1 {
			out = append(out, q.Peek().value)
		}
	}
	return out
}
```

### [264](https://leetcode.com/problems/ugly-number-ii/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l264/264.go; templates: Queue; imports: container/heap -->
```go
// nthUglyNumber expands unique multiples of 2,3,5 in sorted order.
// O(n log n) time, O(n) space. int64 intermediates avoid 32-bit overflow.
func nthUglyNumber(n int) int {
	q := NewQueue(func(a, b int64) bool { return a < b }, int64(1))
	seen := map[int64]bool{1: true}
	var x int64
	for count := 0; count < n; count++ {
		x = q.Pop()
		for _, factor := range []int64{2, 3, 5} {
			y := x * factor
			if !seen[y] {
				seen[y] = true
				q.Push(y)
			}
		}
	}
	return int(x)
}
```

### [295](https://leetcode.com/problems/find-median-from-data-stream/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l295/295.go; templates: Queue; imports: container/heap -->
```go
// MedianFinder partitions values into lower(max heap) and upper(min heap).
// Lower has either the same size as upper or one additional element.
type MedianFinder struct{ lower, upper *Queue[int] }

func Constructor() MedianFinder {
	return MedianFinder{NewQueue(func(a, b int) bool { return a > b }), NewQueue(func(a, b int) bool { return a < b })}
}

// AddNum takes O(log n) time. All inserted values occupy O(n) space.
func (m *MedianFinder) AddNum(num int) {
	if m.lower.Len() == 0 || num <= m.lower.Peek() {
		m.lower.Push(num)
	} else {
		m.upper.Push(num)
	}
	if m.lower.Len() > m.upper.Len()+1 {
		m.upper.Push(m.lower.Pop())
	}
	if m.upper.Len() > m.lower.Len() {
		m.lower.Push(m.upper.Pop())
	}
}

// FindMedian takes O(1) time; at least one number must have been inserted.
func (m *MedianFinder) FindMedian() float64 {
	if m.lower.Len() > m.upper.Len() {
		return float64(m.lower.Peek())
	}
	return float64(m.lower.Peek())/2 + float64(m.upper.Peek())/2
}
```

### [347](https://leetcode.com/problems/top-k-frequent-elements/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l347/347.go; templates: Queue; imports: container/heap -->
```go
type entry struct{ value, count int }

// topKFrequent keeps the k most frequent values. Ties are deterministic.
// O(n+u log k) time, O(u+k) space for u unique values.
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, x := range nums {
		counts[x]++
	}
	q := NewQueue(func(a, b entry) bool {
		if a.count != b.count {
			return a.count < b.count
		}
		return a.value > b.value
	})
	for v, c := range counts {
		q.Push(entry{v, c})
		if q.Len() > k {
			q.Pop()
		}
	}
	out := make([]int, q.Len())
	for i := len(out) - 1; i >= 0; i-- {
		out[i] = q.Pop().value
	}
	return out
}
```

### [373](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l373/373.go; templates: Queue; imports: container/heap -->
```go
type pair struct {
	sum  int64
	i, j int
}

// kSmallestPairs merges rows nums1[i]+nums2[j] without generating the product.
// O(k log min(k,len(nums1))) time, O(min(k,len(nums1))) heap space.
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	out := make([][]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return out
	}
	q := NewQueue(func(a, b pair) bool {
		if a.sum != b.sum {
			return a.sum < b.sum
		}
		if a.i != b.i {
			return a.i < b.i
		}
		return a.j < b.j
	})
	for i := 0; i < len(nums1) && i < k; i++ {
		q.Push(pair{int64(nums1[i]) + int64(nums2[0]), i, 0})
	}
	for len(out) < k && q.Len() > 0 {
		x := q.Pop()
		out = append(out, []int{nums1[x.i], nums2[x.j]})
		if x.j+1 < len(nums2) {
			q.Push(pair{int64(nums1[x.i]) + int64(nums2[x.j+1]), x.i, x.j + 1})
		}
	}
	return out
}
```

### [378](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l378/378.go; templates: Queue; imports: container/heap -->
```go
type cell struct{ value, row, col int }

// kthSmallest merges sorted rows. O((n+k) log(n+1)) time, O(n) space.
func kthSmallest(matrix [][]int, k int) int {
	q := NewQueue(func(a, b cell) bool { return a.value < b.value })
	for r, row := range matrix {
		q.Push(cell{row[0], r, 0})
	}
	for {
		x := q.Pop()
		k--
		if k == 0 {
			return x.value
		}
		if x.col+1 < len(matrix[x.row]) {
			q.Push(cell{matrix[x.row][x.col+1], x.row, x.col + 1})
		}
	}
}
```

### [407](https://leetcode.com/problems/trapping-rain-water-ii/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l407/407.go; templates: Queue; imports: container/heap -->
```go
type cell struct{ row, col, height int }

// trapRainWater expands the lowest boundary; visited cells are marked on insertion.
// O(mn log(mn)) time, O(mn) space.
func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m == 0 {
		return 0
	}
	n := len(heightMap[0])
	if m < 3 || n < 3 {
		return 0
	}
	visited := make([][]bool, m)
	for r := range visited {
		visited[r] = make([]bool, n)
	}
	q := NewQueue(func(a, b cell) bool { return a.height < b.height })
	add := func(r, c int) {
		if !visited[r][c] {
			visited[r][c] = true
			q.Push(cell{r, c, heightMap[r][c]})
		}
	}
	for r := 0; r < m; r++ {
		add(r, 0)
		add(r, n-1)
	}
	for c := 0; c < n; c++ {
		add(0, c)
		add(m-1, c)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	water := 0
	for q.Len() > 0 {
		s := q.Pop()
		for _, d := range dirs {
			r, c := s.row+d[0], s.col+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] {
				continue
			}
			visited[r][c] = true
			water += max(0, s.height-heightMap[r][c])
			q.Push(cell{r, c, max(s.height, heightMap[r][c])})
		}
	}
	return water
}
```

### [480](https://leetcode.com/problems/sliding-window-median/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l480/480.go; templates: Queue; imports: container/heap -->
```go
type entry struct{ value, index int }

// medianSlidingWindow uses two heaps, index-based lazy expiration, and logical sizes.
// Stale interior entries may remain: O(n log n) worst-case time, O(n) space.
func medianSlidingWindow(nums []int, k int) []float64 {
	out := make([]float64, 0)
	if k <= 0 || k > len(nums) {
		return out
	}
	lower := NewQueue(func(a, b entry) bool {
		if a.value != b.value {
			return a.value > b.value
		}
		return a.index > b.index
	})
	upper := NewQueue(func(a, b entry) bool {
		if a.value != b.value {
			return a.value < b.value
		}
		return a.index < b.index
	})
	inLower := make([]bool, len(nums))
	lowSize, upSize, left := 0, 0, 0
	prune := func(q *Queue[entry]) {
		for q.Len() > 0 && q.Peek().index < left {
			q.Pop()
		}
	}
	balance := func() {
		prune(lower)
		prune(upper)
		for lowSize > upSize+1 {
			x := lower.Pop()
			lowSize--
			upSize++
			inLower[x.index] = false
			upper.Push(x)
			prune(lower)
		}
		for lowSize < upSize {
			x := upper.Pop()
			upSize--
			lowSize++
			inLower[x.index] = true
			lower.Push(x)
			prune(upper)
		}
	}
	for i, v := range nums {
		if i >= k {
			old := i - k
			if inLower[old] {
				lowSize--
			} else {
				upSize--
			}
			left = old + 1
			balance()
		}
		x := entry{v, i}
		if lowSize == 0 || v <= lower.Peek().value {
			lower.Push(x)
			inLower[i] = true
			lowSize++
		} else {
			upper.Push(x)
			upSize++
		}
		balance()
		if i >= k-1 {
			median := float64(lower.Peek().value)
			if k%2 == 0 {
				median = median/2 + float64(upper.Peek().value)/2
			}
			out = append(out, median)
		}
	}
	return out
}
```

### [621](https://leetcode.com/problems/task-scheduler/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l621/621.go; templates: Queue; imports: container/heap -->
```go
type cooling struct{ ready, count int }

// leastInterval greedily runs the largest available frequency and queues cooldowns.
// O(T log A) time for T tasks and A distinct task kinds; O(A) space.
// Idle periods are jumped over rather than simulated one unit at a time.
func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	for _, c := range tasks {
		counts[c]++
	}
	q := NewQueue(func(a, b int) bool { return a > b })
	for _, c := range counts {
		q.Push(c)
	}
	wait := make([]cooling, 0, len(counts))
	head, now := 0, 0
	for q.Len() > 0 || head < len(wait) {
		if q.Len() == 0 {
			now = max(now, wait[head].ready)
		}
		for head < len(wait) && wait[head].ready <= now {
			q.Push(wait[head].count)
			head++
		}
		remaining := q.Pop() - 1
		now++
		if remaining > 0 {
			wait = append(wait, cooling{now + n, remaining})
		}
		// Compact consumed cooldowns to keep storage bounded by active kinds.
		if head > 0 && head*2 >= len(wait) {
			wait = append(wait[:0], wait[head:]...)
			head = 0
		}
	}
	return now
}
```

### [630](https://leetcode.com/problems/course-schedule-iii/)

需要模板：`Queue`。

标准库：`container/heap`, `sort`。

<!-- solution: heap_priority_queue/l630/630.go; templates: Queue; imports: container/heap,sort -->
```go
// scheduleCourse sorts deadlines; when over budget, discard the longest course.
// O(n log n) time, O(n) space. The input is not reordered.
func scheduleCourse(courses [][]int) int {
	ordered := append([][]int(nil), courses...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i][1] < ordered[j][1] })
	q := NewQueue(func(a, b int) bool { return a > b })
	var elapsed int64
	for _, c := range ordered {
		q.Push(c[0])
		elapsed += int64(c[0])
		if elapsed > int64(c[1]) {
			elapsed -= int64(q.Pop())
		}
	}
	return q.Len()
}
```

### [632](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l632/632.go; templates: Queue; imports: container/heap -->
```go
type item struct{ value, row, col int }

// smallestRange advances the current minimum while tracking the maximum.
// O(N log k) time, O(k) space for N total values and k lists.
func smallestRange(nums [][]int) []int {
	q := NewQueue(func(a, b item) bool { return a.value < b.value })
	high := nums[0][0]
	for r, row := range nums {
		q.Push(item{row[0], r, 0})
		high = max(high, row[0])
	}
	lo, hi := q.Peek().value, high
	for {
		x := q.Pop()
		width, best := int64(high)-int64(x.value), int64(hi)-int64(lo)
		if width < best || width == best && x.value < lo {
			lo, hi = x.value, high
		}
		if x.col+1 == len(nums[x.row]) {
			break
		}
		v := nums[x.row][x.col+1]
		high = max(high, v)
		q.Push(item{v, x.row, x.col + 1})
	}
	return []int{lo, hi}
}
```

### [692](https://leetcode.com/problems/top-k-frequent-words/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l692/692.go; templates: Queue; imports: container/heap -->
```go
type entry struct {
	word  string
	count int
}

// topKFrequent returns frequency descending, then lexicographic ascending.
// O(n+u log k+k log k) heap operations, O(u+k) space; string comparisons cost length.
func topKFrequent(words []string, k int) []string {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	q := NewQueue(func(a, b entry) bool {
		if a.count != b.count {
			return a.count < b.count
		}
		return a.word > b.word
	})
	for w, c := range counts {
		q.Push(entry{w, c})
		if q.Len() > k {
			q.Pop()
		}
	}
	out := make([]string, q.Len())
	for i := len(out) - 1; i >= 0; i-- {
		out[i] = q.Pop().word
	}
	return out
}
```

### [703](https://leetcode.com/problems/kth-largest-element-in-a-stream/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l703/703.go; templates: Queue; imports: container/heap -->
```go
// KthLargest retains only the largest k values seen so far.
type KthLargest struct {
	k int
	q *Queue[int]
}

func Constructor(k int, nums []int) KthLargest {
	s := KthLargest{k: k, q: NewQueue(func(a, b int) bool { return a < b })}
	for _, x := range nums {
		s.Add(x)
	}
	return s
}

// Add takes O(log k) time; total storage is O(k).
func (s *KthLargest) Add(val int) int {
	s.q.Push(val)
	if s.q.Len() > s.k {
		s.q.Pop()
	}
	return s.q.Peek()
}
```

### [743](https://leetcode.com/problems/network-delay-time/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l743/743.go; templates: Queue; imports: container/heap -->
```go
type edge struct{ to, weight int }
type state struct{ node, distance int }

// networkDelayTime uses Dijkstra with lazy outdated-state removal.
// O(n+E log(E+1)) time, O(n+E) space; all edge weights are nonnegative.
func networkDelayTime(times [][]int, n, k int) int {
	graph := make([][]edge, n+1)
	for _, e := range times {
		graph[e[0]] = append(graph[e[0]], edge{e[1], e[2]})
	}
	inf := int(^uint(0) >> 1)
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = inf
	}
	dist[k] = 0
	q := NewQueue(func(a, b state) bool { return a.distance < b.distance }, state{k, 0})
	for q.Len() > 0 {
		s := q.Pop()
		if s.distance != dist[s.node] {
			continue
		}
		for _, e := range graph[s.node] {
			d := s.distance + e.weight
			if d < dist[e.to] {
				dist[e.to] = d
				q.Push(state{e.to, d})
			}
		}
	}
	answer := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1
		}
		answer = max(answer, dist[i])
	}
	return answer
}
```

### [767](https://leetcode.com/problems/reorganize-string/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l767/767.go; templates: Queue; imports: container/heap -->
```go
type letter struct {
	char  byte
	count int
}

// reorganizeString withholds the previous character until a different one is used.
// O(n log A) time, O(n+A) space including output for A distinct characters.
func reorganizeString(s string) string {
	counts := make(map[byte]int)
	for i := range s {
		counts[s[i]]++
	}
	q := NewQueue(func(a, b letter) bool {
		if a.count != b.count {
			return a.count > b.count
		}
		return a.char < b.char
	})
	for c, n := range counts {
		q.Push(letter{c, n})
	}
	out := make([]byte, 0, len(s))
	previous := letter{}
	for q.Len() > 0 {
		current := q.Pop()
		out = append(out, current.char)
		current.count--
		if previous.count > 0 {
			q.Push(previous)
		}
		previous = current
	}
	if len(out) != len(s) {
		return ""
	}
	return string(out)
}
```

### [857](https://leetcode.com/problems/minimum-cost-to-hire-k-workers/)

需要模板：`Queue`。

标准库：`container/heap`, `math`, `sort`。

<!-- solution: heap_priority_queue/l857/857.go; templates: Queue; imports: container/heap,math,sort -->
```go
type worker struct{ quality, wage int }

// mincostToHireWorkers scans increasing minimum wage/quality ratios.
// At each ratio a max heap retains k smallest qualities. O(n log n) time, O(n) space.
func mincostToHireWorkers(quality, wage []int, k int) float64 {
	workers := make([]worker, len(quality))
	for i := range quality {
		workers[i] = worker{quality[i], wage[i]}
	}
	sort.Slice(workers, func(i, j int) bool {
		return int64(workers[i].wage)*int64(workers[j].quality) < int64(workers[j].wage)*int64(workers[i].quality)
	})
	q := NewQueue(func(a, b int) bool { return a > b })
	var sum int64
	best := math.Inf(1)
	for _, w := range workers {
		q.Push(w.quality)
		sum += int64(w.quality)
		if q.Len() > k {
			sum -= int64(q.Pop())
		}
		if q.Len() == k {
			best = math.Min(best, float64(sum)*float64(w.wage)/float64(w.quality))
		}
	}
	return best
}
```

### [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l871/871.go; templates: Queue; imports: container/heap -->
```go
// minRefuelStops retroactively selects the largest passed station when needed.
// O(n log n) time, O(n) space. Stations must be sorted by position.
func minRefuelStops(target, startFuel int, stations [][]int) int {
	q := NewQueue(func(a, b int) bool { return a > b })
	reach := int64(startFuel)
	i, stops := 0, 0
	for reach < int64(target) {
		for i < len(stations) && int64(stations[i][0]) <= reach {
			q.Push(stations[i][1])
			i++
		}
		if q.Len() == 0 {
			return -1
		}
		reach += int64(q.Pop())
		stops++
	}
	return stops
}
```

### [912](https://leetcode.com/problems/sort-an-array/)

需要模板：`HeapSort`。

<!-- solution: heap_priority_queue/l912/912.go; templates: HeapSort; imports: none -->
```go
// sortArray performs in-place heapsort. O(n log n) time, O(1) extra space.
func sortArray(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, i, len(nums))
	}
	for end := len(nums) - 1; end > 0; end-- {
		nums[0], nums[end] = nums[end], nums[0]
		siftDown(nums, 0, end)
	}
	return nums
}
```

### [973](https://leetcode.com/problems/k-closest-points-to-origin/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l973/973.go; templates: Queue; imports: container/heap -->
```go
type point struct {
	coords   []int
	distance int64
}

// kClosest keeps k smallest squared distances. O(n log k) time, O(k) heap space.
func kClosest(points [][]int, k int) [][]int {
	q := NewQueue(func(a, b point) bool { return a.distance > b.distance })
	for _, p := range points {
		x, y := int64(p[0]), int64(p[1])
		q.Push(point{p, x*x + y*y})
		if q.Len() > k {
			q.Pop()
		}
	}
	out := make([][]int, q.Len())
	for i := range out {
		out[i] = append([]int(nil), q.Pop().coords...)
	}
	return out
}
```

### [1046](https://leetcode.com/problems/last-stone-weight/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1046/1046.go; templates: Queue; imports: container/heap -->
```go
// lastStoneWeight smashes the two largest stones. O(n log n) time, O(n) space.
func lastStoneWeight(stones []int) int {
	q := NewQueue(func(a, b int) bool { return a > b }, stones...)
	for q.Len() > 1 {
		a, b := q.Pop(), q.Pop()
		if a > b {
			q.Push(a - b)
		}
	}
	if q.Len() == 0 {
		return 0
	}
	return q.Pop()
}
```

### [1514](https://leetcode.com/problems/path-with-maximum-probability/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1514/1514.go; templates: Queue; imports: container/heap -->
```go
type edge struct {
	to          int
	probability float64
}
type state struct {
	node        int
	probability float64
}

// maxProbability uses a maximum heap; probabilities in [0,1] cannot improve along a path.
// O(n+E log(E+1)) time, O(n+E) space.
func maxProbability(n int, edges [][]int, succProb []float64, start_node, end_node int) float64 {
	graph := make([][]edge, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], edge{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], edge{e[0], succProb[i]})
	}
	best := make([]float64, n)
	best[start_node] = 1
	q := NewQueue(func(a, b state) bool { return a.probability > b.probability }, state{start_node, 1})
	for q.Len() > 0 {
		s := q.Pop()
		if s.probability < best[s.node] {
			continue
		}
		if s.node == end_node {
			return s.probability
		}
		for _, e := range graph[s.node] {
			p := s.probability * e.probability
			if p > best[e.to] {
				best[e.to] = p
				q.Push(state{e.to, p})
			}
		}
	}
	return 0
}
```

### [1631](https://leetcode.com/problems/path-with-minimum-effort/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1631/1631.go; templates: Queue; imports: container/heap -->
```go
type cell struct{ row, col, effort int }

// minimumEffortPath minimizes the maximum edge difference on a path.
// O(mn log(mn)) time, O(mn) space using Dijkstra.
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	inf := int(^uint(0) >> 1)
	dist := make([][]int, m)
	for r := range dist {
		dist[r] = make([]int, n)
		for c := range dist[r] {
			dist[r][c] = inf
		}
	}
	dist[0][0] = 0
	q := NewQueue(func(a, b cell) bool { return a.effort < b.effort }, cell{0, 0, 0})
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for q.Len() > 0 {
		s := q.Pop()
		if s.effort != dist[s.row][s.col] {
			continue
		}
		if s.row == m-1 && s.col == n-1 {
			return s.effort
		}
		for _, d := range dirs {
			r, c := s.row+d[0], s.col+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			difference := heights[r][c] - heights[s.row][s.col]
			if difference < 0 {
				difference = -difference
			}
			effort := max(s.effort, difference)
			if effort < dist[r][c] {
				dist[r][c] = effort
				q.Push(cell{r, c, effort})
			}
		}
	}
	return 0
}
```

### [1642](https://leetcode.com/problems/furthest-building-you-can-reach/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1642/1642.go; templates: Queue; imports: container/heap -->
```go
// furthestBuilding reserves ladders for the largest climbs seen so far.
// O(n log(ladders+1)) time, O(ladders+1) heap space.
func furthestBuilding(heights []int, bricks, ladders int) int {
	q := NewQueue(func(a, b int) bool { return a < b })
	for i := 1; i < len(heights); i++ {
		d := heights[i] - heights[i-1]
		if d <= 0 {
			continue
		}
		q.Push(d)
		if q.Len() > ladders {
			bricks -= q.Pop()
		}
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
}
```

### [1753](https://leetcode.com/problems/maximum-score-from-removing-stones/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1753/1753.go; templates: Queue; imports: container/heap -->
```go
// maximumScore always pairs the two largest nonempty piles.
// O(a+b+c) time, O(1) space (the heap has at most three elements).
func maximumScore(a, b, c int) int {
	q := NewQueue(func(a, b int) bool { return a > b }, a, b, c)
	score := 0
	for q.Len() > 1 {
		x, y := q.Pop(), q.Pop()
		if y == 0 {
			break
		}
		score++
		if x > 1 {
			q.Push(x - 1)
		}
		if y > 1 {
			q.Push(y - 1)
		}
	}
	return score
}
```

### [1834](https://leetcode.com/problems/single-threaded-cpu/)

需要模板：`Queue`。

标准库：`container/heap`, `sort`。

<!-- solution: heap_priority_queue/l1834/1834.go; templates: Queue; imports: container/heap,sort -->
```go
type task struct{ arrival, duration, index int }

// getOrder selects available tasks by processing time, then original index.
// O(n log n) time, O(n) space. The clock uses int64 for cumulative durations.
func getOrder(tasks [][]int) []int {
	ordered := make([]task, len(tasks))
	for i, t := range tasks {
		ordered[i] = task{t[0], t[1], i}
	}
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].arrival < ordered[j].arrival })
	q := NewQueue(func(a, b task) bool {
		if a.duration != b.duration {
			return a.duration < b.duration
		}
		return a.index < b.index
	})
	out := make([]int, 0, len(tasks))
	i := 0
	var now int64
	for i < len(ordered) || q.Len() > 0 {
		if q.Len() == 0 {
			now = max(now, int64(ordered[i].arrival))
		}
		for i < len(ordered) && int64(ordered[i].arrival) <= now {
			q.Push(ordered[i])
			i++
		}
		t := q.Pop()
		now += int64(t.duration)
		out = append(out, t.index)
	}
	return out
}
```

### [1882](https://leetcode.com/problems/process-tasks-using-servers/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1882/1882.go; templates: Queue; imports: container/heap -->
```go
type server struct {
	end        int64
	weight, id int
}

// assignTasks maintains idle weight/ID priority and busy end/weight/ID priority.
// O((n+m) log n) time, O(n+m) space including output; cumulative times use int64.
func assignTasks(servers, tasks []int) []int {
	byWeight := func(a, b server) bool {
		if a.weight != b.weight {
			return a.weight < b.weight
		}
		return a.id < b.id
	}
	idle := NewQueue(byWeight)
	busy := NewQueue(func(a, b server) bool {
		if a.end != b.end {
			return a.end < b.end
		}
		return byWeight(a, b)
	})
	for i, w := range servers {
		idle.Push(server{weight: w, id: i})
	}
	out := make([]int, len(tasks))
	var now int64
	for i, duration := range tasks {
		now = max(now, int64(i))
		if idle.Len() == 0 {
			now = max(now, busy.Peek().end)
		}
		// Release every simultaneous completion before choosing by weight.
		for busy.Len() > 0 && busy.Peek().end <= now {
			idle.Push(busy.Pop())
		}
		s := idle.Pop()
		out[i] = s.id
		s.end = now + int64(duration)
		busy.Push(s)
	}
	return out
}
```

### [1962](https://leetcode.com/problems/remove-stones-to-minimize-the-total/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l1962/1962.go; templates: Queue; imports: container/heap -->
```go
// minStoneSum removes floor(pile/2) from the largest pile each time.
// O(n+k log n) time, O(n) space.
func minStoneSum(piles []int, k int) int {
	q := NewQueue(func(a, b int) bool { return a > b }, piles...)
	total := 0
	for _, x := range piles {
		total += x
	}
	for ; k > 0 && q.Len() > 0; k-- {
		x := q.Pop()
		total -= x / 2
		q.Push(x - x/2)
	}
	return total
}
```

### [2402](https://leetcode.com/problems/meeting-rooms-iii/)

需要模板：`Queue`。

标准库：`container/heap`, `sort`。

<!-- solution: heap_priority_queue/l2402/2402.go; templates: Queue; imports: container/heap,sort -->
```go
type room struct {
	end int64
	id  int
}

// mostBooked uses an ID heap for idle rooms and an end-time/ID heap for busy rooms.
// O(m log m + (m+n) log n) time, O(m+n) space. End times may exceed int32.
func mostBooked(n int, meetings [][]int) int {
	ordered := append([][]int(nil), meetings...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i][0] < ordered[j][0] })
	idle := NewQueue(func(a, b int) bool { return a < b })
	busy := NewQueue(func(a, b room) bool {
		if a.end != b.end {
			return a.end < b.end
		}
		return a.id < b.id
	})
	for i := 0; i < n; i++ {
		idle.Push(i)
	}
	counts := make([]int, n)
	for _, m := range ordered {
		start, duration := int64(m[0]), int64(m[1])-int64(m[0])
		for busy.Len() > 0 && busy.Peek().end <= start {
			idle.Push(busy.Pop().id)
		}
		var r room
		if idle.Len() > 0 {
			r = room{start, idle.Pop()}
		} else {
			r = busy.Pop()
		}
		r.end += duration
		counts[r.id]++
		busy.Push(r)
	}
	best := 0
	for i := 1; i < n; i++ {
		if counts[i] > counts[best] {
			best = i
		}
	}
	return best
}
```

### [2530](https://leetcode.com/problems/maximal-score-after-applying-k-operations/)

需要模板：`Queue`。

标准库：`container/heap`。

<!-- solution: heap_priority_queue/l2530/2530.go; templates: Queue; imports: container/heap -->
```go
// maxKelements takes the current maximum and replaces it with ceil(x/3).
// O(n+k log n) time, O(n) space.
func maxKelements(nums []int, k int) int64 {
	q := NewQueue(func(a, b int) bool { return a > b }, nums...)
	var score int64
	for ; k > 0; k-- {
		x := q.Pop()
		score += int64(x)
		y := x / 3
		if x%3 != 0 {
			y++
		}
		q.Push(y)
	}
	return score
}
```

### [2558](https://leetcode.com/problems/take-gifts-from-the-richest-pile/)

需要模板：`Queue`。

标准库：`container/heap`, `math`。

<!-- solution: heap_priority_queue/l2558/2558.go; templates: Queue; imports: container/heap,math -->
```go
// pickGifts uses a maximum heap. O(n+k log n) time, O(n) space.
func pickGifts(gifts []int, k int) int64 {
	q := NewQueue(func(a, b int) bool { return a > b }, gifts...)
	for ; k > 0 && q.Len() > 0; k-- {
		x := q.Pop()
		r := int(math.Sqrt(float64(x)))
		// Correct rounding at perfect squares using integer comparisons.
		for r > 0 && r > x/r {
			r--
		}
		for r+1 <= x/(r+1) {
			r++
		}
		q.Push(r)
	}
	var total int64
	for q.Len() > 0 {
		total += int64(q.Pop())
	}
	return total
}
```

## 验证与原有推导

```sh
go test ./heap_priority_queue/...
```

题解文件与文档使用相同声明和适配器；修改模板后需要同时验证每题包以及由文档重新组装的代码。原有推导、详细例子和扩展题单保存在 [整理前笔记](legacy/README.before-inline-templates.md)。
