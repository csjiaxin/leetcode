# 贪心算法（Greedy）

## 目录

- [1. 核心思想](#core)
- [2. 如何证明贪心正确](#proof)
- [3. 常见模式与速查表](#patterns)
- [4. 与动态规划、答案二分的关系](#connections)
- [5. 常见错误](#pitfalls)
- [6. 逐题说明与解释](#problems)
  - [455. 分发饼干](#lc455)
  - [1710. 卡车上的最大单元数](#lc1710)
  - [55. 跳跃游戏](#lc55)
  - [45. 跳跃游戏 II](#lc45)
  - [435. 无重叠区间](#lc435)
  - [452. 用最少数量的箭引爆气球](#lc452)
  - [134. 加油站](#lc134)
  - [871. 最低加油次数](#lc871)
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
| 排序匹配 | [455](#lc455)（[LeetCode](https://leetcode.com/problems/assign-cookies/)） | 两边升序，最小够用资源匹配最小需求 | 孩子、饼干双指针 | 饼干大小等于胃口也能满足 |
| 单位收益 | [1710](#lc1710)（[LeetCode](https://leetcode.com/problems/maximum-units-on-a-truck/)） | 单箱单元数降序 | 剩余车位、总单元数 | 可只取某类型的一部分箱子 |
| 可达范围 | [55](#lc55)（[LeetCode](https://leetcode.com/problems/jump-game/)） | 只从可达位置扩展 | farthest | i>farthest 时失败 |
| 分层扩展 | [45](#lc45)（[LeetCode](https://leetcode.com/problems/jump-game-ii/)） | 扫完整层后增加一步 | end、farthest、steps | 不在终点再增加跳跃次数 |
| 区间选择 | [435](#lc435)（[LeetCode](https://leetcode.com/problems/non-overlapping-intervals/)） | 结束时间升序 | 已保留区间的末尾 | start>=end 可保留 |
| 区间覆盖 | [452](#lc452)（[LeetCode](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)） | 结束位置升序，箭放在右端点 | 当前箭位置、箭数 | start>arrow 才需要新箭 |
| 环形余额 | [134](#lc134)（[LeetCode](https://leetcode.com/problems/gas-station/)） | 局部余额为负则重置起点 | start、tank、total | tank=0 不应重置 |
| 最大堆选优 | [871](#lc871)（[LeetCode](https://leetcode.com/problems/minimum-number-of-refueling-stops/)） | 已可达站中选油量最大者 | reach、站点指针、最大堆 | 站点位置<=reach 即可达 |

<a id="connections"></a>

## 4. 与动态规划、答案二分的关系

- **动态规划**：保留多个可能的状态，再比较最优结果。贪心需要证明其他候选可以安全舍弃。例如 [55](https://leetcode.com/problems/jump-game/) 可从可达 DP 简化为一个最远边界；[871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) 可从按加油次数的 DP 改为最大堆贪心。
- **答案二分**：外层找答案，内层 check 可能使用贪心。[1011](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) 按原顺序每天尽量装满；[410](https://leetcode.com/problems/split-array-largest-sum/) 对非负数组尽量延长每组。这些贪心用于计算固定上限下的最少天数 / 组数。
- **限制改变，证明可能失效**：1011 若允许任意重排且物品不可拆分，就变成装箱可行性，不能用顺序贪心保证正确；1710 每箱占用相同车位才能按单箱收益排序，不能直接推广到一般 0/1 背包。

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

按目录顺序练习。每题提供空函数和测试；下文仅描述算法，不填充实现。

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
- **题意**：与 55 相同，但题设保证终点可达，要求最少跳跃次数。[2,3,1,1,4] 的答案为 2。
- **决策与状态**：维护 steps、当前层右边界 end、下一层最远边界 farthest。遍历到 n-2，每个位置先更新 farthest；当 i=end 时说明当前层扫描结束，steps 加一，end 更新为 farthest。
- **为什么正确**：当前层覆盖在已有跳数内能到达的位置。扫描完整层才得到下一跳能够覆盖的最远位置，首次覆盖终点的层数就是最少跳数，等价于 BFS。
- **复杂度与边界**：时间 `O(n)`，额外空间 `O(1)`。不要在最后一个位置再加一步，单元素答案为 0；不能简单每次跳到当前最远位置，[2,3,1,1,4] 应先跳到下标 1 而非 2。

<a id="lc435"></a>

### [435. 无重叠区间](https://leetcode.com/problems/non-overlapping-intervals/)

- **练习**：[435.go](l435/435.go) · [测试](l435/435_test.go)
- **题意**：删除最少区间，使剩余区间互不重叠；端点相接允许保留。[[1,2],[2,3],[3,4],[1,3]] 的答案为 1。
- **决策与状态**：按终点升序，保留结束最早的区间；维护最后保留的 end，后续仅当 `start>=end` 时保留并更新 end。答案为总数减最多保留数。
- **为什么正确**：将最优安排的第一个区间换成结束最早的区间，后续原本可选区间仍可选，数量不会减少；对剩余部分重复即可。不应优先选择起点最早或长度最短的区间。
- **复杂度与边界**：时间 `O(n log(n+1))`，除排序栈外额外空间 `O(1)`。重复区间只能留一个；负数端点、包含关系、相同终点也适用。

<a id="lc452"></a>

### [452. 用最少数量的箭引爆气球](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

- **练习**：[452.go](l452/452.go) · [测试](l452/452_test.go)
- **题意**：气球覆盖闭区间 [start,end]；一箭可射穿所有覆盖该坐标的气球，求最少箭数。[[1,2],[2,3]] 共用坐标 2，只需一箭。
- **决策与状态**：按右端点升序，将第一箭放在最早右端点 arrow；后续若 `start<=arrow` 已被覆盖，不移动箭；仅当 `start>arrow` 时增加一箭，放在当前右端点。
- **为什么正确**：最早结束的气球必须被射中，将射中它的箭右移至其右端点，不会失去原来同时覆盖的其他气球，因为这些气球的右端点都不更早。这个位置为后续覆盖留下最多机会。
- **复杂度与边界**：时间 `O(n log(n+1))`，除排序栈外额外空间 `O(1)`。端点相接可以共用箭；一串相邻重叠不代表有共同交点；比较端点不要用相减排序，以免溢出。

<a id="lc134"></a>

### [134. 加油站](https://leetcode.com/problems/gas-station/)

- **练习**：[134.go](l134/134.go) · [测试](l134/134_test.go)
- **题意**：环形路线上第 i 站有 gas[i] 升油，到下一站消耗 cost[i]，油箱无限大、初始为空。返回能绕一圈的起点，无解返回 -1；题设有解则唯一。
- **决策与状态**：令 net=gas[i]-cost[i]，维护 total（全程净油量）、tank（从当前候选起点的累计净油量）、start。若 tank<0，设置 start=i+1 并清零 tank。遍历后 total<0 则无解，否则返回 start。
- **为什么正确**：从 start 到 i 首次失败，之前各前缀余额均非负。改从其中任意站 j 出发，到 i 的余额为“原失败余额减去此前非负前缀”，仍为负，故 start 到 i 全部可排除。总量非负时，最终候选可完成环路。
- **复杂度与边界**：时间 `O(n)`，额外空间 `O(1)`。tank=0 不重置；不可按 gas 最大选起点；注意局部可行不等于总量足够。

<a id="lc871"></a>

### [871. 最低加油次数](https://leetcode.com/problems/minimum-number-of-refueling-stops/)

- **练习**：[871.go](l871/871.go) · [测试](l871/871_test.go)
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
