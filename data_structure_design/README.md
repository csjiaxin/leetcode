# 数据结构设计：从操作复杂度反推组合

[返回总目录](../README.md)

这是面试设计题的学习指南与参考模板，不代表题单均已有独立实现。先列出 API、操作频率、顺序要求和复杂度，再组合结构；不要先猜使用哪一个容器。

## 已实现的基础题

| 题目 | 独立实现 | 测试 | 核心复杂度 |
| --- | --- | --- | --- |
| [146. LRU 缓存](https://leetcode.com/problems/lru-cache/) | [146.go](l146/146.go) | [146_test.go](l146/146_test.go) | Get / Put 平均 O(1) |
| [155. 最小栈](https://leetcode.com/problems/min-stack/) | [155.go](l155/155.go) | [155_test.go](l155/155_test.go) | Push 摊还 O(1)，其他 O(1) |
| [232. 用栈实现队列](https://leetcode.com/problems/implement-queue-using-stacks/) | [232.go](l232/232.go) | [232_test.go](l232/232_test.go) | 各操作摊还 O(1)，单次搬运最坏 O(n) |
| [380. O(1) 时间插入、删除和获取随机元素](https://leetcode.com/problems/insert-delete-getrandom-o1/) | [380.go](l380/380.go) | [380_test.go](l380/380_test.go) | 插入平均且摊还 O(1)，删除平均 O(1)，随机取值 O(1) |

四题使用独立的 `l题号` 包和官方 `Constructor` / 方法签名，不依赖仓库公共 helper。复制提交时省略 package 行。LRU 和随机集合先调用构造函数；含指针、map 或切片的对象使用后不要按值复制，全部实现均非并发安全。队列和最小栈的零值也可使用。

按官方约定，栈的 Pop / Top / GetMin、队列的 Pop / Peek、集合的 GetRandom 仅在非空时调用；不为无效调用设计额外返回值。LRU 额外支持容量为 0，负容量会 panic。

```sh
go test ./data_structure_design/... -count=1 -cover
go test -race ./data_structure_design/...
```

测试包含官方示例、重复操作和固定种子的随机模型对照：LRU 检查 map 与链表访问顺序；最小栈对照完整扫描最小值；队列对照数组 FIFO；随机集合检查紧凑数组与下标映射双射。GetRandom 均匀性来自“每个值恰好对应一个下标＋均匀抽取下标”，测试不使用可能偶发失败的统计阈值。981、460 仍是后续练习，295 复用已有实现。

空间按逻辑容量分析：LRU O(capacity)，其余 O(n)。Go 切片 / map 删除后不保证立即缩小底层分配，长期反复增删时实际保留内存可能取决于历史峰值，而非当前元素数。

## 分类与不变量

| 题目 | 组合 | 关键不变量 |
| --- | --- | --- |
| 146 LRU 缓存 | 哈希表＋双向链表 | 每个 key 唯一；访问移到最近端，淘汰最旧端 |
| 155 最小栈 | 每层记录当前最小值 | 栈顶同时保存当前值和整个前缀最小值 |
| 380 O(1) 插入删除随机集合 | 数组＋值到下标映射 | 删除时尾元素填洞并修正映射 |
| 232 用栈实现队列 | 输入栈＋输出栈 | 仅当输出栈为空时倒入，保证 FIFO |
| 981 基于时间的键值存储 | key→有序历史＋二分 | 每个 key 的时间戳递增时才可直接追加 |
| 460 LFU 缓存 | key 映射＋频率桶＋桶内时间顺序 | 维护最小非空频率，同频淘汰最旧者 |
| 295 数据流中位数 | 双堆 | 大小平衡，左半最大值不大于右半最小值 |

295 已有[实现](../heap_priority_queue/l295/295.go)，无需重复建题解。

## 模板一：LRU（Go 标准库双向链表）

以下代码块在同一文件中使用；官方题目构造函数名可按签名改为 Constructor。零值不可直接使用，应通过 NewLRU 初始化；非并发安全。

```go
import "container/list"

type cacheEntry struct { key, value int }
type LRUCache struct {
    capacity int
    order *list.List
    index map[int]*list.Element
}
func NewLRU(capacity int) *LRUCache {
    if capacity < 0 { panic("negative capacity") }
    return &LRUCache{capacity, list.New(), make(map[int]*list.Element)}
}
func (c *LRUCache) Get(key int) int {
    element, ok := c.index[key]
    if !ok { return -1 }
    c.order.MoveToFront(element)
    return element.Value.(cacheEntry).value
}
func (c *LRUCache) Put(key, value int) {
    if element, ok := c.index[key]; ok {
        element.Value = cacheEntry{key, value}
        c.order.MoveToFront(element)
        return
    }
    if c.capacity == 0 { return }
    c.index[key] = c.order.PushFront(cacheEntry{key, value})
    if c.order.Len() > c.capacity {
        oldest := c.order.Back()
        delete(c.index, oldest.Value.(cacheEntry).key)
        c.order.Remove(oldest)
    }
}
```

哈希表负责定位，链表负责 O(1) 移动 / 删除；只用数组维护访问顺序会让中间移动退化为 O(n)。Get / Put 平均 O(1)，空间 O(capacity)。此 API 用 -1 表示不存在，遵循 146；通用容器可返回 `(value,ok)`。

## 模板二：每层存前缀最小值

```go
type minEntry struct { value, minimum int }
type MinStack struct { entries []minEntry }
func (s *MinStack) Push(value int) {
    minimum := value
    if len(s.entries) > 0 && s.entries[len(s.entries)-1].minimum < minimum {
        minimum = s.entries[len(s.entries)-1].minimum
    }
    s.entries = append(s.entries, minEntry{value, minimum})
}
func (s *MinStack) Pop() { s.entries = s.entries[:len(s.entries)-1] }
func (s *MinStack) Top() int { return s.entries[len(s.entries)-1].value }
func (s *MinStack) GetMin() int { return s.entries[len(s.entries)-1].minimum }
```

按题意 Pop / Top / GetMin 只在非空时调用。Push 受切片扩容影响为摊还 O(1)，其他操作 O(1)，空间 O(n)。重复最小值天然保留；若用单独最小值栈，则入栈比较应包含相等。

## 其他题的实现骨架与边界

- **380**：插入将值追加末尾并记录下标；删除前把末尾值覆盖目标位置并更新其映射，再缩短切片、删除旧 key；随机选择均匀下标。插入平均且摊还 O(1)，删除平均 O(1)，数组随机访问 O(1)，空间 O(n)。
- **232**：入队只压输入栈；出队 / 查看时仅在输出栈空时整体倒入。单次倒入最坏 O(n)，每个元素最多搬一次，所以每次操作摊还 O(1)，空间 O(n)。
- **981**：历史序列二分最后一个 `timestamp<=query`；若写入乱序，需要有序插入或其他结构，不能直接宣称追加 O(1)。
- **460**：频率变化时从旧桶移除并加入新桶；同频内部仍需维护时间顺序。完成 LRU 后再练。

## 检查与练习

先定义“存在 / 不存在”、容量为零、重复插入、更新是否算访问、空操作、并发语义。map 的复杂度是平均值，动态数组追加是摊还值，不要统一写成最坏 O(1)。

练习：155 → 232 → 380 → 146 → 981 → 295 → 460。
