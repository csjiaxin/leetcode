# 贪心算法（Greedy）

## 分类学习指南：先证明决策，再选择模板

**贪心不是一种固定的数据结构，也没有能覆盖所有题目的单一代码模板。** 排序、双指针、单调栈、堆、扫描都可能承载贪心；真正共通的是：证明某个选择不会损失最优解，随后舍弃不必要的其他选择。

本指南新增的 Go helper 使用独立名称；`sort`、`container/heap` 及 `FuelHeap` 在后面的“本页可复用模板”中定义。代码按各题数据范围使用 `int`，更大范围需要改用 `int64` 并检查累加溢出。排序类模板通常会原地修改输入，需要保留原数组时先复制。

### 导航与分类速查

- [通用解题步骤与证明](#greedy-process)
- [A. 字典序：31、556、1053 的贪心推导与模板](#greedy-lexicographic)
- [B. 排序匹配与单位收益](#greedy-sorting)
- [C. 区间选择、刺点与覆盖扩展](#greedy-intervals)
- [D. 可达范围与分层跳跃](#greedy-reach)
- [E. 余额扫描与失败区间排除](#greedy-balance)
- [F. 堆：延迟选择与反悔](#greedy-heap)
- [G. 单调栈：撤销较差前缀](#greedy-stack)
- [H. 相邻约束：双向扫描](#greedy-two-pass)
- [I. 贪心作为答案二分的判定](#greedy-check)
- [学习路线与反例检查](#greedy-practice)

| 类别 | 局部决策 | 证明抓手 | 代表题 |
| --- | --- | --- | --- |
| 字典序 | 最晚改变前缀，再最小化变化幅度 | 第一个不同位置决定大小 | 31、556、1053 |
| 排序匹配 | 最小够用资源满足最小需求 | 交换后不减少匹配数 | 455、881 |
| 单位收益 | 等容量单位优先取收益高的 | 一对一替换收益不减 | 1710 |
| 区间选择 / 刺点 | 优先最早结束 / 在最早右端点刺点 | 为后续保留最多空间 | 435、452 |
| 覆盖扩展 | 当前可接入区间中延伸最远 | 相同选择数，覆盖更远者支配更近者 | 1024、1326 |
| 可达边界 | 扫描整层后跳到下一层最远边界 | BFS 层数下界与可达范围不变量 | 55、45 |
| 失败排除 | 前缀首次失败，跳过整个候选区间 | 证明中间任一起点也会失败 | 134 |
| 堆延迟选择 / 反悔 | 需要时取最大收益 / 超限时撤销最大代价 | 替换后状态不劣 | 871、1642、630 |
| 单调栈 | 若允许且必要，撤销更差的前缀末尾 | 字典序优先＋可行性约束 | 402、316、1081、321 |
| 双向扫描 | 分别满足左右约束，逐点取下界最大值 | 任意可行解都必须满足两个下界 | 135 |
| 二分 check | 固定上限下，每组尽量装满 | 非负、顺序固定时，贪心分段数最少 | 1011、410 |

<a id="greedy-process"></a>

### 通用解题步骤：不能省略“为什么”

1. **定义目标与约束**：最大数量、最小代价、字典序最小？能否重排、撤销、重复使用资源？
2. **确定决策单位**：当前数位、一个区间、一层可达位置，还是某个候选任务？
3. **提出选择规则**：写清主关键字和相等时的处理，而不是笼统地“排序后贪心”。
4. **尝试反例**：重复值、端点相等、局部收益大但后续代价大，是否会击穿规则？
5. **证明安全性**：交换论证、领先 / 支配关系、失败区间排除，或者必要下界与可行解一致。
6. **维护状态并处理边界**：说明状态含义、何时更新、何时答案确定、无解如何返回。

通用的只是推理框架，而不是可直接套题的统一函数：

```text
明确剩余问题与不变量
while 还有决策：
    找到当前合法候选
    根据已经证明的规则选择、延迟选择或替换候选
    更新状态，保持不变量
检查完整可行性，返回结果
```

贪心并不意味着代码绝不撤销操作：堆反悔和单调栈都会替换暂定选择，但每次替换都由确定规则支配，不像回溯那样枚举所有分支。

<a id="greedy-lexicographic"></a>

### A. 字典序贪心：为什么能直接决定交换位置？

**字典序只看第一个不同位置，后面的任何改动都无法弥补这个位置的大小差异。** 因此依次优化：保留最长前缀 → 优化首次变化的数值 → 优化剩余后缀。

#### A1. 31 下一个排列：最小幅度地变大

[31. 下一个排列](https://leetcode.com/problems/next-permutation/) 要求在所有严格更大的排列中取最小值：

1. **最晚修改**：从右向左找第一个 `a[i] < a[i+1]`。右侧已非递增，是该后缀的最大排列；只重排右侧不可能变大，所以 i 是最右可增大的位置。改更左处会更大，不可能更优。
2. **最小增幅**：在右侧找严格大于 `a[i]` 的最小值。后缀非递增，从右向左遇到的第一个更大值就是答案。
3. **最小后缀**：交换后前缀已保证变大，将后缀调整为升序，令整体尽量小。交换后的后缀仍非递增，直接反转即可。

```text
1 3 | 5 4 2   后缀已最大，必须改变 3
1 4 | 5 3 2   换成最小的更大值 4，而不是 5
1 4 | 2 3 5   反转后缀为升序，得到下一个排列
```

找不到 i 表示已经是最大排列。31 要求重置为最小排列；下面的返回值同时告诉调用者是否存在严格更大的排列。[556. 下一个更大元素 III](https://leetcode.com/problems/next-greater-element-iii/) 对数位应用同一模板，无更大结果或超出 32 位有符号整数范围则返回 -1；可用 `int64` 累加检查。

```go
// 原地修改；有下一个排列返回 true，否则重置为最小排列并返回 false。
func GreedyNextPermutation(a []int) bool {
    i := len(a) - 2
    for i >= 0 && a[i] >= a[i+1] { i-- }
    advanced := i >= 0
    if advanced {
        j := len(a) - 1
        for a[j] <= a[i] { j-- }
        a[i], a[j] = a[j], a[i]
    }
    for l, r := i+1, len(a)-1; l < r; l, r = l+1, r-1 {
        a[l], a[r] = a[r], a[l]
    }
    return advanced
}
```

时间 O(n)，额外空间 O(1)，支持重复元素。[31 本地实现](../permutation/l31/31.go)、[556 本地实现](../permutation/l556/556.go)。

#### A2. 1053：一次交换，最大限度地接近原排列但严格变小

[1053. 交换一次的先前排列](https://leetcode.com/problems/previous-permutation-with-one-swap/) 的结果不一定是所有排列中的直接前驱，因为最多只能交换一次：

1. **最晚修改**：从右向左找第一个 `a[i] > a[i+1]`。右侧非递减，内部交换不能使它变小，必须改 i 或更左处；选择 i 保留最长前缀。
2. **最小降幅**：从右侧取严格小于 `a[i]` 的最大值。例如 `1 9 4 6 7 → 1 7 4 6 9`，换 7 比换 6 更接近原排列。
3. **重复候选取最左位置**：`3 1 1` 与左边的 1 交换得 `1 3 1`，与右边的交换得 `1 1 3`。前者把原来的大值放得更靠左，后缀更大。

```go
func GreedyPreviousOneSwap(a []int) []int {
    i := len(a) - 2
    for i >= 0 && a[i] <= a[i+1] { i-- }
    if i < 0 { return a }
    j := len(a) - 1
    for a[j] >= a[i] { j-- }
    for j > i+1 && a[j] == a[j-1] { j-- }
    a[i], a[j] = a[j], a[i]
    return a
}
```

时间 O(n)，额外空间 O(1)。**交换后不能再反转后缀**，否则违反一次交换限制。[1053 本地实现](../permutation/l1053/1053.go)。

| 决策 | 31 / 556 | 1053 |
| --- | --- | --- |
| 目标 | 严格变大且尽量小 | 一次交换后严格变小且尽量大 |
| 位置 | 最右可增大处 | 最右可减小处 |
| 候选值 | 严格更大值中的最小值 | 严格更小值中的最大值 |
| 后缀 | 反转为最小排列 | 不可再重排；重复候选取最左位置 |

<a id="greedy-sorting"></a>

### B. 排序匹配与单位收益：不要混成“排序就对了”

**匹配模板**：两边升序，拿当前最小资源尝试满足最小需求；不够就丢弃资源，够用就同时推进。下方 `MatchSorted` 要求调用前已排序。

```go
func GreedyAssign(requirements, resources []int) int {
    sort.Ints(requirements)
    sort.Ints(resources)
    return MatchSorted(requirements, resources)
}
```

- [455. 分发饼干](https://leetcode.com/problems/assign-cookies/)：资源不够满足最小需求，就不可能满足更大需求；够用时匹配最小需求，可以通过交换论证保留最优匹配数。时间 O(n log n + m log m)，扫描 O(n+m)。
- [881. 救生艇](https://leetcode.com/problems/boats-to-save-people/)：每次必须安排最重者；若能与最轻者配对就一起走，否则最重者只能独乘。**每船最多两人**是关键，不能推广到每船可坐任意多人。

**单位收益模板**：下方 `TakeByValue` 按单箱收益降序，每次拿尽可能多的当前类型。适用于 [1710. 卡车上的最大单元数](https://leetcode.com/problems/maximum-units-on-a-truck/)，因为每个箱子都恰好占一个车位，替换一个低收益箱子不会增加容量消耗。时间 O(t log t)，t 为类型数。

反例边界：一般 0/1 背包的物品重量不同，不能套单位收益贪心；容量 4，物品 `(重量,价值)` 为 `(3,5)、(2,3)、(2,3)`，按价值密度先选前者只能得 5，但后两者能得 6。

<a id="greedy-intervals"></a>

### C. 区间：选择、刺点、覆盖，是三个不同目标

**C1. 最多保留不重叠区间**：[435. 无重叠区间](https://leetcode.com/problems/non-overlapping-intervals/)。按右端点升序，保留最早结束者；相同保留数下，结束更早不会减少未来选择。下方 `SelectByEnd(intervals, func(s,e int) bool { return s >= e })` 返回最多保留数，删除数为总数减保留数。

**C2. 最少点刺穿全部闭区间**：[452. 用最少数量的箭引爆气球](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)。把箭放在最早结束区间的右端点：任何解都必须有一箭命中它，将这一箭右移到该右端点，不会漏掉尚未处理且原来被该箭命中的区间。复用 `SelectByEnd`，新箭条件改为 `start > arrow`。

两者排序时间均为 O(n log n)。**435 中端点相等可以接续，452 中端点相等可以共用箭**。

**C3. 最少区间覆盖 `[0,target]`**：[1024. 视频拼接](https://leetcode.com/problems/video-stitching/)、[1326. 灌溉花园的最少水龙头数目](https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/)。按左端点升序，在当前覆盖边界之前可接入的区间中，选右端点最远的；其他选择无法用相同的一步覆盖得更远。

```go
// intervals 中每项为合法闭区间 [left,right]；覆盖目标为 [0,target]。
func GreedyCover(intervals [][]int, target int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    end, i, count := 0, 0, 0
    for end < target {
        farthest := end
        for i < len(intervals) && intervals[i][0] <= end {
            if intervals[i][1] > farthest { farthest = intervals[i][1] }
            i++
        }
        if farthest == end { return -1 }
        end = farthest
        count++
    }
    return count
}
```

1326 先将第 i 个水龙头转为 `[max(0,i-ranges[i]), min(n,i+ranges[i])]`。时间 O(n log n)；不是“选择最早结束”，因为目标变成了**用最少区间延伸覆盖范围**。更多内容见 [区间贪心](interval_greedy/readme.md)。

<a id="greedy-reach"></a>

### D. 可达范围：边界支配与隐式 BFS

适用前提：从 i 可以前进到 `i+1..i+nums[i]` 中任意位置，跳跃长度非负；可达位置形成连续前缀。任意图的可达集合不一定连续，不能仅用一个边界替代 BFS。

- [55. 跳跃游戏](https://leetcode.com/problems/jump-game/)：只扫描可达位置，维护 `farthest = max(farthest, i+nums[i])`。若 `i > farthest`，后面都无法到达。
- [45. 跳跃游戏 II](https://leetcode.com/problems/jump-game-ii/)：维护 `end`（当前跳数对应的最远边界）、`farthest`（下一层最远边界）；扫描到 end 才增加步数，见下方 `RangeJumps`。

```text
扫描当前层中的每个位置，合并其下一跳能到达的范围
到达当前层终点时：
    若无法扩展，则无解
    steps++
    end = farthest
```

每层穷尽当前跳数的所有可达选择，因此首次覆盖终点所需层数最少，等价于压缩状态后的 BFS。不是每次跳到最远下标：`[2,3,1,1,4]` 中先跳到下标 1 比先跳到下标 2 更好。时间 O(n)，额外空间 O(1)。

<a id="greedy-balance"></a>

### E. 余额扫描：一次失败排除整个区间

[134. 加油站](https://leetcode.com/problems/gas-station/) 用 `delta[i] = gas[i]-cost[i]` 表示净收益，见下方 `CircuitStart`。

```text
total += delta[i]，tank += delta[i]
若 tank < 0：start = i+1，tank = 0
扫描结束：total < 0 则无解，否则返回 start
```

证明：若从 start 出发首次在 i 失败，则 start 到任意 `j-1` 的余额非负，所以从中间 j 出发到 i 的余额，只会比从 start 出发更小或相等，仍为负。由此整个 `[start,i]` 的起点一次性排除。总净收益非负，再配合这种起点选择，才能保证完整环路可行。

时间 O(n)，额外空间 O(1)。余额等于 0 不失败；不能只检查总量后任意返回一个起点。见 [燃料与余额专题](gas_fuel/readme.md)。

<a id="greedy-heap"></a>

### F. 堆贪心：延迟选择和反悔不是同一种状态

#### F1. 延迟选择最大收益

[871. 最低加油次数](https://leetcode.com/problems/minimum-number-of-refueling-stops/)：站点按位置升序，先把所有已经可达但未使用的站点加入最大堆。只有需要扩展可达距离时，才弹出最大油量，见下方 `GreedyRefuel`。

```text
while reach < target：
    所有 position <= reach 的未处理站点入最大堆
    堆为空则无解
    reach += 弹出的最大油量
    stops++
```

相同增加一次加油，取已可达候选中的最大油量能到得最远，不会损失后续候选。这里依赖油箱容量无限等原题条件；只允许取已可达站点，不能提前借用未来站点。时间 O(n log n)，额外空间 O(n)。

类似题：[1642. 可以到达的最远建筑](https://leetcode.com/problems/furthest-building-you-can-reach/)，可用最小堆把梯子保留给最大的几次正高度差，其余用砖块。局部决策相近，但状态和代码不能直接照搬加油模板。

#### F2. 超限时撤销最大代价

[630. 课程表 III](https://leetcode.com/problems/course-schedule-iii/)：按截止时间升序，暂时选入课程；若超时，删除已选课程中耗时最长的一门。

```go
func GreedySchedule(courses [][]int) int {
    sort.Slice(courses, func(i, j int) bool {
        return courses[i][1] < courses[j][1]
    })
    selected := &FuelHeap{} // 复用下方最大堆；这里存课程时长，不是油量。
    heap.Init(selected)
    elapsed := 0
    for _, c := range courses {
        elapsed += c[0]
        heap.Push(selected, c[0])
        if elapsed > c[1] {
            elapsed -= heap.Pop(selected).(int)
        }
    }
    return selected.Len()
}
```

每门课为 `[正时长, 截止时间]`。按截止时间处理确保旧约束不会因为新截止时间而变紧；超限时减少数量不可避免，删最长者给后续留下最多时间。只需删一门：被删时长至少等于刚加入课程的时长，因此总时长不超过加入前的值，恢复可行。交换论证可维持前缀内最大可选数量，并在该数量下最小化总时长。

时间 O(n log n)，额外空间 O(n)。这是允许替换暂定选择的贪心，不是“选了就永不后悔”。

<a id="greedy-stack"></a>

### G. 单调栈：让更优元素尽量靠前，但必须还能完成任务

[402. 移掉 K 位数字](https://leetcode.com/problems/remove-k-digits/)：保留相对顺序，删除恰好 k 位，使结果最小。若栈顶比当前数字大且还有删除额度，删除栈顶能让首次不同位置变小，是更优的前缀；若数字非递减导致额度剩余，则从末尾删除。

```go
// num 为非空十进制数字串，0 <= k <= len(num)。
func GreedyRemoveDigits(num string, k int) string {
    stack := make([]byte, 0, len(num))
    for i := 0; i < len(num); i++ {
        digit := num[i]
        for k > 0 && len(stack) > 0 && stack[len(stack)-1] > digit {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, digit)
    }
    stack = stack[:len(stack)-k]
    first := 0
    for first < len(stack) && stack[first] == '0' { first++ }
    if first == len(stack) { return "0" }
    return string(stack[first:])
}
```

例如 `1432219, k=3`：读到 3 时移除 4，读到 2 时移除 3，读到 1 时移除前面的 2，最终为 `1219`。每个数字最多入栈、出栈一次，时间 O(n)，额外空间 O(n)。注意不是保证最终栈始终单调：删除额度耗尽后，不能再弹栈。

**同族但弹栈可行条件不同**：

- [316. 去除重复字母](https://leetcode.com/problems/remove-duplicate-letters/)、[1081. 不同字符的最小子序列](https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/)：每个不同字符恰好一次；弹出较大字符前，还必须确认它后面会再次出现，并维护是否已在栈内。不能把 402 的删除额度判断原样搬过来。
- [321. 拼接最大数](https://leetcode.com/problems/create-maximum-number/)：固定选取长度时用反向比较的单调栈求最大子序列，还要枚举两个数组各取多少个，并按**剩余后缀**字典序合并，不能只比较当前首字符。

更一般的推理骨架是：

```text
for 当前元素 x：
    while 栈顶导致目标更差，并且弹出后仍能完成任务：
        弹出栈顶，更新剩余资源 / 状态
    若 x 满足加入条件，则入栈
处理剩余删除额度、长度约束及输出格式
```

相关实现见 [单调栈目录](../monotonic_stack/README.md)。

<a id="greedy-two-pass"></a>

### H. 双向扫描：合并两个方向的必要下界

[135. 分发糖果](https://leetcode.com/problems/candy/)：每人至少 1 颗，评分高于相邻者时糖果也必须更多。

- 从左到右，计算只满足左邻约束所需的最少糖果 `left[i]`。
- 从右到左，计算只满足右邻约束所需的最少糖果 `right[i]`。
- 每点取 `max(left[i], right[i])`。任意合法解至少达到这个下界，而该结果又同时满足左右约束，因此全局总和最小。

```go
func GreedyCandy(ratings []int) int {
    left := make([]int, len(ratings))
    for i := range ratings {
        left[i] = 1
        if i > 0 && ratings[i] > ratings[i-1] {
            left[i] = left[i-1] + 1
        }
    }
    total, right := 0, 1
    for i := len(ratings)-1; i >= 0; i-- {
        if i+1 < len(ratings) && ratings[i] > ratings[i+1] {
            right++
        } else {
            right = 1
        }
        value := left[i]
        if right > value { value = right }
        total += value
    }
    return total
}
```

时间 O(n)，额外空间 O(n)，空输入返回 0。相同评分没有“必须更多”的要求，连续递增的条件必须是严格大于。这个模板适用于可拆成左右单向下界、取最大值后仍可行的相邻约束，**不是所有双向约束都能直接取 max**。

<a id="greedy-check"></a>

### I. 贪心用于答案二分：固定上限后，计算最少分组数

[1011. 在 D 天内送达包裹的能力](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)、[410. 分割数组的最大值](https://leetcode.com/problems/split-array-largest-sum/) 的共性：非负数组、保持原顺序、每组连续，组和不超过给定上限。

```go
// nums 非负，limit 非负；返回最少非空连续组数，以及是否存在单元素超限。
func GreedyGroupsUnderLimit(nums []int, limit int) (int, bool) {
    if len(nums) == 0 { return 0, true }
    groups, sum := 1, 0
    for _, x := range nums {
        if x > limit { return 0, false }
        if sum > limit-x { // 避免 sum+x 溢出。
            groups++
            sum = 0
        }
        sum += x
    }
    return groups, true
}
```

**证明**：每组尽量向右延伸；任何可行分法的第一组都不能越过贪心第一组的终点。非负性保证对剩余部分同样成立，归纳得到贪心以相同组数覆盖的前缀不短于任何其他分法，因此组数最少。

外层二分最小可行上限，`check(limit)` 为 `ok && groups <= k`。上限越大，所需组数不会增加，满足二分的单调性。通常搜索区间为 `[max(nums), sum(nums)]`，总和可能需要 `int64`。410 要求恰好 k 组：当 `1 <= k <= n` 时，非负数组可继续拆组且不增大组和，所以“至多 k 组”判定仍成立。1011 则允许少于 D 天完成。

单次 check 时间 O(n)，额外空间 O(1)，外层总时间 O(n log S)，S 为整数搜索区间规模。若允许任意重排、出现负数或允许非连续分组，这个证明与模板可能失效。见 [二分答案目录](../binary_search/readme.md)。

<a id="greedy-practice"></a>

### 学习路线、反例与模板选型清单

```text
入门交换论证：455 → 1710 → 881
区间目标区分：435 → 452 → 1024 → 1326
边界与失败排除：55 → 45 → 134
字典序：31 → 556 → 1053
延迟选择与反悔：871 → 1642 → 630
单调栈与前缀：402 → 316 / 1081 → 321
约束下界：135
二分答案＋贪心：1011 → 410
```

遇到新题，先问：

- **排序依据是否有证明？** 不是所有区间都按右端点排序，也不是所有收益问题都按价值密度排序。
- **相等时怎么办？** 端点能否共用、比较是否严格、重复候选选左还是右，都会改变结果。
- **候选是否真的可行？** 不可达站点不能提前入候选堆；已超出可达范围的位置不能用于扩展边界。
- **为什么不需要记录更多状态？** 找不到支配关系或交换证明时，考虑 DP、搜索、图算法，而不是强行贪心。
- **能否用小规模暴力验证猜想？** 比较全部排列、全部区间子集、全部任务子集，可快速找到反例；实验能帮助排错，但不能替代证明。

上述题单是代表性应用，不代表每题都已在 `greedy/` 下实现。跨目录实现保持原位置；本次只扩充学习文档，不移动代码。原有 helper 与逐题适配器保留如下。

## 本页可复用模板：贪心与可达范围

先复制以下 helper，再复制任一题适配器。实际题解在每个文件内嵌同名 helper，可独立提交，不导入仓库本地算法包。不同移动/贪心依据使用不同命名模板，不把“用了两个变量”当作同一不变量。

```go
import (
	"container/heap"
	"sort"
)

func RangeJumps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	end, farthest, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i > farthest {
			return -1
		}
		farthest = max(farthest, i+nums[i])
		if i == end {
			if farthest == end {
				return -1
			}
			end = farthest
			steps++
			if end >= len(nums)-1 {
				return steps
			}
		}
	}
	return -1
}

func MatchSorted(requirements, resources []int) int {
	matched := 0
	for _, resource := range resources {
		if matched < len(requirements) && resource >= requirements[matched] {
			matched++
		}
	}
	return matched
}

func TakeByValue(items [][]int, capacity int) int {
	sort.Slice(items, func(i, j int) bool { return items[i][1] > items[j][1] })
	total := 0
	for _, item := range items {
		take := min(capacity, item[0])
		total += take * item[1]
		capacity -= take
		if capacity == 0 {
			break
		}
	}
	return total
}

func SelectByEnd(intervals [][]int, separate func(start, end int) bool) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	chosen, end := 0, 0
	for _, interval := range intervals {
		if chosen == 0 || separate(interval[0], end) {
			chosen++
			end = interval[1]
		}
	}
	return chosen
}

func CircuitStart(gas, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	total, tank, start := 0, 0, 0
	for i, value := range gas {
		net := value - cost[i]
		total += net
		tank += net
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

func RefuelByCount(startFuel int, stations [][]int) []int64 {
	reach := make([]int64, len(stations)+1)
	for j := 1; j < len(reach); j++ {
		reach[j] = -1
	}
	reach[0] = int64(startFuel)
	for i, station := range stations {
		for stops := i + 1; stops >= 1; stops-- {
			if reach[stops-1] >= int64(station[0]) {
				reach[stops] = max(reach[stops], reach[stops-1]+int64(station[1]))
			}
		}
	}
	return reach
}

type FuelHeap []int

func (h FuelHeap) Len() int { return len(h) }

func (h FuelHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h FuelHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *FuelHeap) Push(value any) { *h = append(*h, value.(int)) }

func (h *FuelHeap) Pop() any {
	values := *h
	last := values[len(values)-1]
	*h = values[:len(values)-1]
	return last
}

func GreedyRefuel(target, startFuel int, stations [][]int) int {
	reach := int64(startFuel)
	index, stops := 0, 0
	candidates := &FuelHeap{}
	heap.Init(candidates)
	for reach < int64(target) {
		for index < len(stations) && int64(stations[index][0]) <= reach {
			heap.Push(candidates, stations[index][1])
			index++
		}
		if candidates.Len() == 0 {
			return -1
		}
		reach += int64(heap.Pop(candidates).(int))
		stops++
	}
	return stops
}
```

| 题目 | 采用的局部模板 | 实际实现 |
| --- | --- | --- |
| [55](https://leetcode.com/problems/jump-game/) | `RangeJumps` | [实际 Go 文件](l55/55.go) |
| [45](https://leetcode.com/problems/jump-game-ii/) | `RangeJumps` | [实际 Go 文件](l45/45.go) |
| [455](https://leetcode.com/problems/assign-cookies/) | `MatchSorted` | [实际 Go 文件](l455/455.go) |
| [1710](https://leetcode.com/problems/maximum-units-on-a-truck/) | `TakeByValue` | [实际 Go 文件](l1710/1710.go) |
| [435](https://leetcode.com/problems/non-overlapping-intervals/) | `SelectByEnd` | [实际 Go 文件](interval_greedy/l435/435.go) |
| [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) | `SelectByEnd` | [实际 Go 文件](interval_greedy/l452/452.go) |
| [134](https://leetcode.com/problems/gas-station/) | `CircuitStart` | [实际 Go 文件](gas_fuel/l134/134.go) |
| [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/)：按加油次数 DP | `RefuelByCount` | [实际 Go 文件](gas_fuel/l871/871.go) |
| [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/)：最大堆贪心 | `FuelHeap + GreedyRefuel` | [实际 Go 文件](gas_fuel/l871/variants/practice/871.go) |

为使整页 Go 示例可合并编译，重名适配器在文档中加题号/DP/Greedy 后缀；实际 Go 文件仍保留官方函数签名。链表节点等支撑类型仅在 helper 区声明一次。

## 目录

- [1. 核心思想](#core)
- [2. 如何证明贪心正确](#proof)
- [3. 常见模式与速查表](#patterns)
- [4. 与动态规划、答案二分的关系](#connections)
- [5. 常见错误](#pitfalls)
- [6. 逐题说明与解释](#problems)
  - [455. 分发饼干](https://leetcode.com/problems/assign-cookies/) · [本节](#lc455)
  - [1710. 卡车上的最大单元数](https://leetcode.com/problems/maximum-units-on-a-truck/) · [本节](#lc1710)
  - [55. 跳跃游戏](https://leetcode.com/problems/jump-game/) · [本节](#lc55)
  - [45. 跳跃游戏 II](https://leetcode.com/problems/jump-game-ii/) · [本节](#lc45)
  - [435. 无重叠区间](https://leetcode.com/problems/non-overlapping-intervals/) · [本节](#lc435)
  - [452. 用最少数量的箭引爆气球](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) · [本节](#lc452)
  - [134. 加油站](https://leetcode.com/problems/gas-station/) · [本节](#lc134)
  - [871. 最低加油次数](https://leetcode.com/problems/minimum-number-of-refueling-stops/) · [本节](#lc871)
- [7. 练习与测试](#tests)

<a id="core"></a>

## 1. 核心思想

**每一步做当前最有利的选择，并证明这种选择不会妨碍得到全局最优解。** 贪心通常不撤销已经确定的选择，但“当前最有利”必须有依据，不能凭直觉。

做题时先回答：

1. 优化目标是什么：最多满足人数、最少跳跃次数、最少删除数量，还是最少加油次数？
2. 当前决策有哪些候选，按什么规则选择？
3. 为什么该选择不比其他选择差？能否与某个最优方案交换？
4. 需要维护什么状态：最远可达位置、区间终点、累计余额或候选最大堆？
5. 边界相等意味着可行还是冲突？

不是所有“最大值 / 最小值”问题都能贪心。先尝试构造反例，再寻找证明。

<a id="proof"></a>

## 2. 如何证明贪心正确

### 交换论证

从一个最优方案出发，将某个选择换为贪心选择，证明方案仍合法且目标不变差。于是存在采用该贪心选择的最优方案，可继续处理剩余问题。

- [455](https://leetcode.com/problems/assign-cookies/)：最小的够用饼干留给当前胃口最小的孩子，不浪费更大的饼干。
- [1710](https://leetcode.com/problems/maximum-units-on-a-truck/)：每箱占一个位置，用更高单元数的箱子替换较低者不会减少收益。
- [435](https://leetcode.com/problems/non-overlapping-intervals/)：把最优安排的第一个区间替换为结束最早的区间，给后续留下的空间不会减少。

### 维护支配关系或不变量

- [55](https://leetcode.com/problems/jump-game/)：所有已处理的可达位置都被纳入最远可达范围；更远的可达边界支配更近的边界。
- [45](https://leetcode.com/problems/jump-game-ii/)：扫描完整的一跳可达层，维护下一层最远边界，等价于压缩后的 BFS 分层。

### 排除不可能候选

[134](https://leetcode.com/problems/gas-station/) 从 start 首次在 i 处累计余额为负，意味着 start 到 i 都不能作为答案，可整体排除，而不是只排除 start。

### 延迟决策

[871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) 不急着确定在哪个站加油。先记录所有可达站的油量，需要扩展距离时才选择其中最大的油量。

<a id="patterns"></a>

## 3. 常见模式与速查表

| 模式 | 题目 / 原题链接 | 决策或排序规则 | 维护状态 | 关键边界 |
| --- | --- | --- | --- | --- |
| 排序匹配 | [455](https://leetcode.com/problems/assign-cookies/) · [本节](#lc455)（[LeetCode](https://leetcode.com/problems/assign-cookies/)） | 两边升序，最小够用资源匹配最小需求 | 孩子、饼干双指针 | 饼干大小等于胃口也能满足 |
| 单位收益 | [1710](https://leetcode.com/problems/maximum-units-on-a-truck/) · [本节](#lc1710)（[LeetCode](https://leetcode.com/problems/maximum-units-on-a-truck/)） | 单箱单元数降序 | 剩余车位、总单元数 | 可只取某类型的一部分箱子 |
| 可达范围 | [55](https://leetcode.com/problems/jump-game/) · [本节](#lc55)（[LeetCode](https://leetcode.com/problems/jump-game/)） | 只从可达位置扩展 | farthest | i>farthest 时失败 |
| 分层扩展 | [45](https://leetcode.com/problems/jump-game-ii/) · [本节](#lc45)（[LeetCode](https://leetcode.com/problems/jump-game-ii/)） | 扫完整层后增加一步 | end、farthest、steps | 不在终点再增加跳跃次数 |
| 区间选择 | [435](https://leetcode.com/problems/non-overlapping-intervals/) · [本节](#lc435)（[LeetCode](https://leetcode.com/problems/non-overlapping-intervals/)） | 结束时间升序 | 已保留区间的末尾 | start>=end 可保留 |
| 区间覆盖 | [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) · [本节](#lc452)（[LeetCode](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)） | 结束位置升序，箭放在右端点 | 当前箭位置、箭数 | start>arrow 才需要新箭 |
| 环形余额 | [134](https://leetcode.com/problems/gas-station/) · [本节](#lc134)（[LeetCode](https://leetcode.com/problems/gas-station/)） | 局部余额为负则重置起点 | start、tank、total | tank=0 不应重置 |
| 最大堆选优 | [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) · [本节](#lc871)（[LeetCode](https://leetcode.com/problems/minimum-number-of-refueling-stops/)） | 已可达站中选油量最大者 | reach、站点指针、最大堆 | 站点位置<=reach 即可达 |

<a id="connections"></a>

## 4. 与动态规划、答案二分的关系

- **动态规划**：保留多个可能的状态，再比较最优结果。贪心需要证明其他候选可以安全舍弃。例如 [55](https://leetcode.com/problems/jump-game/) 可从可达 DP 简化为一个最远边界；[871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) 可从按加油次数的 DP 改为最大堆贪心。
- **答案二分**：外层找答案，内层 check 可能使用贪心。[1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) 按原顺序每天尽量装满；[410](https://leetcode.com/problems/split-array-largest-sum/) 对非负数组尽量延长每组。这些贪心用于计算固定上限下的最少天数 / 组数。
- **限制改变，证明可能失效**：[1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) 若允许任意重排且物品不可拆分，就变成装箱可行性，不能用顺序贪心保证正确；[1710](https://leetcode.com/problems/maximum-units-on-a-truck/) 每箱占用相同车位才能按单箱收益排序，不能直接推广到一般 0/1 背包。

<a id="pitfalls"></a>

## 5. 常见错误

- 只写出“每次选最大 / 最小”，却没有证明；应说清交换后为什么不会破坏后续可行性。
- 混淆 [435](https://leetcode.com/problems/non-overlapping-intervals/) 和 [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)：前者端点相接不冲突，后者端点相接可以共用一箭。
- 跳跃时直接选当前能跳到的最远下标，而非比较整个可达层能扩展出的下一层边界。
- 使用不可达位置扩展 farthest，例如 [1,0,100] 不能因为末位很大就返回 true。
- [134](https://leetcode.com/problems/gas-station/) 只检查总油量足够就返回任意起点；总量足够意味着有解，但起点仍须正确寻找。
- [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) 选择最远站或最大的未到达站；必须从已经可达的站点中选择最大油量。
- 排序比较器用减法，可能溢出；直接比较终点或收益大小。
- 忽略排序会修改输入。测试为每个用例复制切片，二维输入复制每一行。

<a id="problems"></a>

## 6. 逐题说明与解释

按目录顺序练习。各题已有实现与测试；下文补充算法解释，代码可参考本页 helper、适配器及对应 Go 文件。

<a id="lc455"></a>

### [455. 分发饼干](https://leetcode.com/problems/assign-cookies/)

- **练习**：[455.go](l455/455.go) · [测试](l455/455_test.go)
- **题意**：每个孩子有胃口值，每块饼干有大小；每人至多分一块，大小不小于胃口即可满足，求最多满足人数。例如 g=[1,2,3]、s=[1,1]，答案 1。
- **决策与状态**：胃口、饼干都升序。用最小的剩余饼干尝试满足胃口最小的孩子；太小则跳过饼干，够用则匹配成功、两个指针一起移动。
- **为什么正确**：太小的饼干无法满足当前孩子，更无法满足之后胃口更大的孩子。够用时用它匹配，不会浪费更大的资源；任何采用更大饼干的最优方案，都可以交换为当前较小饼干而不减少满足人数。
- **复杂度与边界**：排序及扫描时间 `O(g log(g+1)+s log(s+1))`；除排序栈外额外空间 `O(1)`。饼干为空、重复值、恰好满足要覆盖。

<a id="lc1710"></a>

### [1710. 卡车上的最大单元数](https://leetcode.com/problems/maximum-units-on-a-truck/)

- **练习**：[1710.go](l1710/1710.go) · [测试](l1710/1710_test.go)
- **题意**：boxTypes[i]=[箱数,每箱单元数]，卡车最多装 truckSize 个箱子，求最多单元数。例如 [[1,3],[2,2],[3,1]]、truckSize=4，答案 8。
- **决策与状态**：按每箱单元数降序。每次取 `min(当前类型箱数,剩余车位)` 箱，累加收益，减少车位。
- **为什么正确**：所有箱子都占一个车位，把较低收益箱子换为尚未选取的较高收益箱子不会降低总收益，所以应先选高收益者。不是按“箱数×单箱收益”排序。
- **复杂度与边界**：m 种类型，时间 `O(m log(m+1))`，除排序栈外额外空间 `O(1)`。可以只取某类型的一部分；车位多于总箱数时不必装满。

<a id="lc55"></a>

### [55. 跳跃游戏](https://leetcode.com/problems/jump-game/)

- **练习**：[55.go](l55/55.go) · [测试](l55/55_test.go)
- **题意**：nums[i] 是从 i 最多能跳多远，判断能否从 0 到最后下标。[2,3,1,1,4] 为 true，[3,2,1,0,4] 为 false。
- **决策与状态**：维护 farthest。遍历 i，先判断 i 是否可达：若 `i>farthest` 则失败，否则用 `i+nums[i]` 更新 farthest；覆盖终点即可成功。
- **为什么正确**：从一个位置可以选择不超过最大跳跃长度的任意距离，已知可达位置形成从 0 开始的连续范围。只需保留最远边界，不必记录每条路径。
- **复杂度与边界**：时间 `O(n)`，额外空间 `O(1)`。单元素 [0] 已在终点；零可被跨越，但不可从断点之后的位置更新 farthest。

<a id="lc45"></a>

### [45. 跳跃游戏 II](https://leetcode.com/problems/jump-game-ii/)

- **练习**：[45.go](l45/45.go) · [测试](l45/45_test.go)
- **题意**：与 [55](https://leetcode.com/problems/jump-game/) 相同，但题设保证终点可达，要求最少跳跃次数。[2,3,1,1,4] 的答案为 2。
- **决策与状态**：维护 steps、当前层右边界 end、下一层最远边界 farthest。遍历到 n-2，每个位置先更新 farthest；当 i=end 时说明当前层扫描结束，steps 加一，end 更新为 farthest。
- **为什么正确**：当前层覆盖在已有跳数内能到达的位置。扫描完整层才得到下一跳能够覆盖的最远位置，首次覆盖终点的层数就是最少跳数，等价于 BFS。
- **复杂度与边界**：时间 `O(n)`，额外空间 `O(1)`。不要在最后一个位置再加一步，单元素答案为 0；不能简单每次跳到当前最远位置，[2,3,1,1,4] 应先跳到下标 1 而非 2。

<a id="lc435"></a>

### [435. 无重叠区间](https://leetcode.com/problems/non-overlapping-intervals/)

- **练习**：[435.go](interval_greedy/l435/435.go) · [测试](interval_greedy/l435/435_test.go)
- **题意**：删除最少区间，使剩余区间互不重叠；端点相接允许保留。[[1,2],[2,3],[3,4],[1,3]] 的答案为 1。
- **决策与状态**：按终点升序，保留结束最早的区间；维护最后保留的 end，后续仅当 `start>=end` 时保留并更新 end。答案为总数减最多保留数。
- **为什么正确**：将最优安排的第一个区间换成结束最早的区间，后续原本可选区间仍可选，数量不会减少；对剩余部分重复即可。不应优先选择起点最早或长度最短的区间。
- **复杂度与边界**：时间 `O(n log(n+1))`，除排序栈外额外空间 `O(1)`。重复区间只能留一个；负数端点、包含关系、相同终点也适用。

<a id="lc452"></a>

### [452. 用最少数量的箭引爆气球](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

- **练习**：[452.go](interval_greedy/l452/452.go) · [测试](interval_greedy/l452/452_test.go)
- **题意**：气球覆盖闭区间 [start,end]；一箭可射穿所有覆盖该坐标的气球，求最少箭数。[[1,2],[2,3]] 共用坐标 2，只需一箭。
- **决策与状态**：按右端点升序，将第一箭放在最早右端点 arrow；后续若 `start<=arrow` 已被覆盖，不移动箭；仅当 `start>arrow` 时增加一箭，放在当前右端点。
- **为什么正确**：最早结束的气球必须被射中，将射中它的箭右移至其右端点，不会失去原来同时覆盖的其他气球，因为这些气球的右端点都不更早。这个位置为后续覆盖留下最多机会。
- **复杂度与边界**：时间 `O(n log(n+1))`，除排序栈外额外空间 `O(1)`。端点相接可以共用箭；一串相邻重叠不代表有共同交点；比较端点不要用相减排序，以免溢出。

<a id="lc134"></a>

### [134. 加油站](https://leetcode.com/problems/gas-station/)

- **练习**：[134.go](gas_fuel/l134/variants/practice/134.go) · [测试](gas_fuel/l134/variants/practice/134_test.go)
- **题意**：环形路线上第 i 站有 gas[i] 升油，到下一站消耗 cost[i]，油箱无限大、初始为空。返回能绕一圈的起点，无解返回 -1；题设有解则唯一。
- **决策与状态**：令 net=gas[i]-cost[i]，维护 total（全程净油量）、tank（从当前候选起点的累计净油量）、start。若 tank<0，设置 start=i+1 并清零 tank。遍历后 total<0 则无解，否则返回 start。
- **为什么正确**：从 start 到 i 首次失败，之前各前缀余额均非负。改从其中任意站 j 出发，到 i 的余额为“原失败余额减去此前非负前缀”，仍为负，故 start 到 i 全部可排除。总量非负时，最终候选可完成环路。
- **复杂度与边界**：时间 `O(n)`，额外空间 `O(1)`。tank=0 不重置；不可按 gas 最大选起点；注意局部可行不等于总量足够。

<a id="lc871"></a>

### [871. 最低加油次数](https://leetcode.com/problems/minimum-number-of-refueling-stops/)

- **练习**：[871.go](gas_fuel/l871/variants/practice/871.go) · [测试](gas_fuel/l871/variants/practice/871_test.go)
- **题意**：到 target 的路上有按位置递增的站点 [position,fuel]，每单位距离耗一升油，初始 startFuel，油箱无限大。每站可取全部油，求最少停靠次数，无法到达则 -1。
- **决策与状态**：reach 初始为 startFuel。将所有 `position<=reach` 且尚未入堆的站点油量加入最大堆；若 reach<target，取堆中最大油量，加到 reach 并增加停靠数。空堆且未到目标则失败。
- **为什么正确**：在当前必须增加一次加油的时刻，选择已可达站中最多的油量，使同样一次停靠获得的可达范围最大，不会比选更小油量差。候选站原本就在可达路径上，所以这是延迟确定过去在哪站加油，而非实际倒车。
- **复杂度与边界**：n 个站，时间 `O(n log(n+1))`，额外空间 `O(n)`。到站时恰好没油仍能加；到目标时无油也算成功。累计可达距离可能超过 32 位整数，建议 reach 使用 int64。Go 的 container/heap 默认为按 Less 定义优先级，最大堆需让较大油量优先。

<a id="tests"></a>

## 7. 练习与测试

本系列新增独立目录：`greedy/l455`、`l1710`、`l55`、`l45`、`l435`、`l452`、`l134`、`l871`。每题包含 `题号.go` 与 `题号_test.go`，测试覆盖示例与关键边界。

**当前函数体为空**，带返回值的函数会报 missing return，暂时不能编译或通过测试。完成某题后，在仓库根目录运行：

```sh
go test ./greedy/l455 -v
```

八题都完成后，仅运行本系列的八个包：

```sh
go test ./greedy/l455 ./greedy/l1710 ./greedy/l55 ./greedy/l45 ./greedy/l435 ./greedy/l452 ./greedy/l134 ./greedy/l871
```

不要直接使用 `go test ./greedy/...` 验证未完成的练习；该路径也会包含其他已有子目录。

已有的 `greedy/gas_fuel/`、`dp/jump_game/`、扩展名为空的草稿及区间专题内容全部保留；本系列不清空旧实现，也不覆盖其他练习。


## 其余题目的模板适配器

以下代码调用本页已经声明的 helper；复制到对应题解时使用实际 Go 文件中的官方函数名。

### [55](https://leetcode.com/problems/jump-game/)

```go
func canJump(nums []int) bool { return RangeJumps(nums) >= 0 }
```

### [45](https://leetcode.com/problems/jump-game-ii/)

```go
func jump(nums []int) int { return RangeJumps(nums) }
```

### [455](https://leetcode.com/problems/assign-cookies/)

```go
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	return MatchSorted(g, s)
}
```

### [1710](https://leetcode.com/problems/maximum-units-on-a-truck/)

```go
func maximumUnits(boxTypes [][]int, truckSize int) int { return TakeByValue(boxTypes, truckSize) }
```

### [435](https://leetcode.com/problems/non-overlapping-intervals/)

```go
func eraseOverlapIntervals(intervals [][]int) int {
	keep := SelectByEnd(intervals, func(start, end int) bool { return start >= end })
	return len(intervals) - keep
}
```

### [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

```go
func findMinArrowShots(points [][]int) int {
	return SelectByEnd(points, func(start, end int) bool { return start > end })
}
```

### [134](https://leetcode.com/problems/gas-station/)

```go
func canCompleteCircuit(gas []int, cost []int) int { return CircuitStart(gas, cost) }
```

### [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/)：按加油次数 DP

```go
func minRefuelStopsDP(target int, startFuel int, stations [][]int) int {
	for stops, reach := range RefuelByCount(startFuel, stations) {
		if reach >= int64(target) {
			return stops
		}
	}
	return -1
}
```

### [871](https://leetcode.com/problems/minimum-number-of-refueling-stops/)：最大堆贪心

```go
func minRefuelStopsGreedy(target int, startFuel int, stations [][]int) int {
	return GreedyRefuel(target, startFuel, stations)
}
```
