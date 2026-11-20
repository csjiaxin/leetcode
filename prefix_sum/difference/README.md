# 差分与扫描线：端点变化、前缀恢复与活跃状态

[返回前缀和目录](../readme.md) · [仓库总目录](../../README.md) · [扫描线入口](../../sweep_line/README.md) · [区间贪心](../../greedy/interval_greedy/readme.md)

本文整合差分与扫描线指南：共同的端点、边界和累计原理只解释一次，差分基础题保留原有代码入口，扫描线进阶题附学习思路与原题链接。题单不表示全部已有独立实现。原扫描线目录保留为导航，不移动源码。新增扫描模板依赖标准库 `sort`；差分泛型模板需要 Go 1.18+。

<a id="relationship"></a>

## 先选方法：两者有什么关系，分别解决什么问题？

**差分是一种记录相邻变化的表示方法；扫描线是一种按位置顺序处理事件、维护状态的算法框架。它们相关，但不等同。**

- **差分数组**：把区间加法变成两个端点增减，前缀累计恢复每个位置的值，适合批量更新后统一汇总。
- **排序端点扫描**：坐标巨大但端点少时，只保存变化位置，按坐标排序并累计。对于可加的总量，可看作稀疏差分。
- **更一般的扫描线**：状态可能是有效建筑高度集合、最短区间候选堆，不能仅靠一个差分累计值解决。

| 要解决的问题 | 维护的信息 / 选择 | 代表题 |
| --- | --- | --- |
| 多次区间加，返回所有位置的最终值 | 密集差分数组＋前缀恢复 | [1109. 航班预订统计](https://leetcode.com/problems/corporate-flight-bookings/)、[2381. 字母移位 II](https://leetcode.com/problems/shifting-letters-ii/) |
| 判断任一路段是否超载 | 当前有效人数；坐标小用数组、大用排序事件 | [1094. 拼车](https://leetcode.com/problems/car-pooling/) |
| 求最大覆盖及最早位置 | 当前覆盖数＋历史最优，严格更优才更新位置 | [1854. 人口最多的年份](https://leetcode.com/problems/maximum-population-year/) |
| 求最少资源 / 最大重叠 | 活跃数量的最大值，注意端点语义 | [253. 会议室 II](https://leetcode.com/problems/meeting-rooms-ii/)（会员题）、[2406. 将区间分为最少组数](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/) |
| 求区间并集长度 | 活跃数＋相邻事件的实际坐标距离 | 见[稀疏扫描](#sparse-sweep) |
| 求建筑轮廓 | 动态维护有效建筑的最大高度 | [218. 天际线问题](https://leetcode.com/problems/the-skyline-problem/) |
| 求每个查询点的最短覆盖区间 | 离线查询＋候选最小堆 | [1851. 包含每个查询的最小区间](https://leetcode.com/problems/minimum-interval-to-include-each-query/) |
| 每次新增预订后求最大重叠 | 重扫已知事件，或动态结构维护 | [732. 我的日程安排表 III](https://leetcode.com/problems/my-calendar-iii/) |

**选型**：先确认要维护的状态能否简单加减，再确认坐标是密集还是稀疏，最后确认是批量汇总还是每次操作后立即查询。普通差分不能直接解决区间赋值或动态最大高度。静态区间求和则优先考虑前缀和，而不是差分。区间贪心是在选择哪些区间，本文主要是在汇总区间的共同影响，目标也不同。

## 目录

- [1. 核心思路](#core)
  - [1.1 解决什么样的问题？](#problem-shape)
  - [1.2 思路来源：从逐点修改到记录相邻变化](#derivation)
  - [1.3 如何理解：开始生效与停止生效](#events)
  - [1.4 从零数组到非零数组](#initial-array)
  - [1.5 通用 Go 模板与使用边界](#template)
- [2. 1094：拼车](#lc1094)
- [3. 1109：航班预订](#lc1109)
- [4. 1854：人口最多年份](#lc1854)
- [5. 对比与易错点](#pitfalls)
- [6. 稀疏扫描、坐标压缩与覆盖长度](#sparse-sweep)
- [7. 进阶扫描线：离线查询、天际线与在线预订](#advanced-sweep)
- [8. 扩展题单](#extended-practice)
- [9. 练习顺序](#practice-order)

<a id="core"></a>

## 1. 核心思路

**一句话：对一整段位置增加相同的量，不逐点修改，只记录“从哪里开始增加、从哪里停止增加”，最后一次扫描还原结果。**

<a id="problem-shape"></a>
### 1.1 解决什么样的问题？

典型问题由三部分组成：

1. 有一排位置：数组下标、航班编号、站点、年份或时间点。
2. 每次操作影响一个**连续区间**，给区间内每个位置增加或减少相同的量。
3. 需要汇总各位置的最终值，或扫描最大值、覆盖数、是否超出容量。

| 问题描述 | 区间内累计的是什么？ | 最后求什么？ |
| --- | --- | --- |
| [1109. 航班预订统计](https://leetcode.com/problems/corporate-flight-bookings/) | 每笔订单增加的座位数 | 每个航班的总座位数 |
| [1094. 拼车](https://leetcode.com/problems/car-pooling/) | 每批乘客占用的座位数 | 是否有路段超载 |
| [1854. 人口最多的年份](https://leetcode.com/problems/maximum-population-year/) | 每个人对存活年份贡献1 | 最大人口的最早年份 |
| [1893. 检查是否区域内所有整数都被覆盖](https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/) | 每个区间贡献一次覆盖 | 目标范围内是否处处被覆盖 |

识别信号不是“题目里出现数组”，而是：**很多次连续区间的同量增减，最后统一汇总。**

例如，长度为 n 的数组有 m 次区间加法。逐个修改区间内的元素，最坏时间为 `O(m·n)`；差分把每次更新压缩为两个端点操作，最后扫描一次，时间变为 `O(m+n)`。

<a id="derivation"></a>
### 1.2 思路来源：从逐点修改到记录相邻变化

先不记公式，观察：**给一整段都加3后，相邻元素之间的差发生了什么变化？**

对零数组的闭区间 `[1,3]` 加3，下标从0开始：

```text
下标 i        0   1   2   3   4
修改前 a      0   0   0   0   0
修改后 a      0   3   3   3   0
相邻差 diff   0  +3   0   0  -3
```

定义：

```text
diff[0] = a[0]
diff[i] = a[i] - a[i-1]，i >= 1
```

区间内部相邻两个元素都加3，因此它们的差不变：

```text
(a[i]+3) - (a[i-1]+3) = a[i] - a[i-1]
```

只有两个边界发生变化：

- **左边界 l**：当前元素增加，前一个元素不增加，所以 `diff[l] += 3`。
- **右边界之后 r+1**：当前元素不增加，前一个元素增加，所以 `diff[r+1] -= 3`。

于是，修改很多元素的问题，转化成只修改两个差值的问题。

为什么做前缀和就能还原？因为相邻差相加会抵消中间项：

```text
diff[0] + diff[1] + ... + diff[i]
= a[0] + (a[1]-a[0]) + ... + (a[i]-a[i-1])
= a[i]
```

所以可以记为：**差分记录相邻变化，前缀和把变化累积回实际值。**

<a id="events"></a>
### 1.3 如何理解：开始生效与停止生效

还有一种不依赖代数公式的理解：把每次区间更新看成一项有起止时间的贡献。

对 `[1,3]` 加3，意味着：

```text
位置1：从这里开始，多累计3
位置2：没有新的变化，继续累计3
位置3：没有新的变化，继续累计3
位置4：从这里开始，取消刚才那份3
```

**右侧的减3不是把位置4的最终值变成负数，而是取消此前开始的贡献。** 扫描时要把所有事件累计起来，不能直接把单个 `diff[i]` 当作答案。

多个区间可以叠加。例如再对 `[2,4]` 加2：

```text
下标                  0   1   2   3   4   5（哨兵）
[1,3] 加3的事件       0  +3   0   0  -3   0
[2,4] 加2的事件       0   0  +2   0   0  -2
合并后的 diff         0  +3  +2   0  -3  -2
前缀累计              0   3   5   5   2   0
实际数组              0   3   5   5   2
```

下标5不属于实际数组，只用来放取消事件。最后累计回0，表示这两项额外贡献都已经结束。

#### “事件、状态、答案”分别是什么？

- **事件**：在某个位置发生的变化。例如会议开始 +1、结束 -1；乘客上车 +p、下车 -p。事件可包含位置、增减量、对象编号或高度；查询本身也可作为读取状态的事件。
- **状态**：处理到当前位置时持续维护的信息。例如当前会议数、当前乘客数、有效建筑高度集合或候选区间堆。
- **答案**：由状态提取的结果。例如当前会议数 active 的历史最大值 best，或者每个查询点的最短合法区间长度。

差分中 `diff[x]` 是 x 的净变化，不是 x 处的人数；前缀累计 `active += diff[x]` 才更新当前状态。一般扫描线则未必能只用一个整数状态：建筑结束后，还要知道其余建筑中谁最高。

会议例子采用 `[1,4)、[2,5)、[4,6)`：

| 时间 | 事件 | 净变化 | 处理后 active | 状态对应的路段 |
| --- | --- | --- | --- | --- |
| 1 | 第一场开始 | +1 | 1 | [1,2) |
| 2 | 第二场开始 | +1 | 2 | [2,4) |
| 4 | 第一场结束、第三场开始 | 0 | 2 | [4,5) |
| 5 | 第二场结束 | -1 | 1 | [5,6) |
| 6 | 第三场结束 | -1 | 0 | 6 之后 |

最大 active 为 2，因此最少需要 2 间会议室。同坐标先合并净变化，不会把时间 4 错算成同时 3 场。这也说明只需处理事件位置，不必遍历两个事件之间的每个时刻。

<a id="endpoint-rules"></a>

#### 边界约定决定贡献在哪里取消

| 区间含义 | 开始生效 | 停止生效 | 原因 |
| --- | --- | --- | --- |
| 整数闭区间 `[l,r]` | `diff[l] += delta` | `diff[r+1] -= delta` | r仍受影响，从下一个位置才取消 |
| 半开区间 `[l,r)` | `diff[l] += delta` | `diff[r] -= delta` | r本身已经不受影响 |

**扫描不变量**：处理完坐标 x 的全部净变化后，前缀累计值就是 x 处的有效总量。同坐标加减应先聚合，不凭输入事件顺序判断超载。

这也解释了拼车题为什么在下车站直接减：乘客占用的是 `[from,to)`，到达 to 时这一份占用已经结束。

- **半开区间**：同坐标净变化处理完再读取状态，等价于结束先于开始，不把同站上下车临时叠加。
- **连续坐标上的闭区间**：端点本身也算覆盖，计算端点最大重叠时应先加入、统计、再结束；不能直接用净变化后统计。退化闭区间 `[x,x]` 也覆盖 x。
- **整数闭区间**：可转为 `[l,r+1)`，但 r+1 需要检查溢出；实数区间不能机械加 1。
- **覆盖长度**：统计相邻坐标间的距离，不是每个端点的数量；单点的连续长度为 0。

<a id="initial-array"></a>
### 1.4 从零数组到非零数组

前面的例子从零开始，是为了突出每次操作的贡献。如果原数组不为零，先区分两部分：

```text
最终数组 = 原数组 + 每个位置累计获得的增量
```

例如，对原数组 `[2,5,5]` 的闭区间 `[1,2]` 加2（下标从0开始）：

```text
下标          0   1   2
原数组        2   5   5
本次增量      0   2   2
最终数组      2   7   7
```

差分可以记录**完整数组**，也可以只记录**增量**。两种方式都正确，但扫描得到的东西不同。

#### 方式一：先构建原数组差分，再修改——扫描得到最终值

第一步，把原数组 `[2,5,5]` 改写成相邻差：

```text
diff[0] = 2           // 第一个位置从0变为2
diff[1] = 5 - 2 = 3   // 从位置0到位置1，值增加3
diff[2] = 5 - 5 = 0   // 从位置1到位置2，值不变
```

此时，即使没有区间更新，扫描 `[2,3,0]` 就已经能恢复原数组：

```text
位置0：累计值 = 2           → 原值2
位置1：累计值 = 2+3 = 5     → 原值5
位置2：累计值 = 2+3+0 = 5   → 原值5
```

第二步，记录 `[1,2]` 加2：在位置1开始增加，在位置3取消。

```text
下标                   0   1   2   3（哨兵）
原数组差分             2   3   0   0
本次更新的端点事件     0  +2   0  -2
相加后的 diff          2   5   0  -2
```

这里的哨兵初始置0，只用于容纳更新的取消事件；它不属于原数组，也没有定义为“原数组在位置3归零”的差值。只扫描实际下标 `0..2` 即可。

第三步，前缀累计直接得到最终数组：

```text
位置0：累计值 = 2           → 最终值2
位置1：累计值 = 2+5 = 7     → 最终值7
位置2：累计值 = 2+5+0 = 7   → 最终值7
```

**为什么不用再加原数组？** 因为初始化 diff 时，原数组的信息已经包含在里面；更新是在这份信息上叠加的。

#### 方式二：零差分只记录更新——扫描得到增量

第一步，保留原数组不动，另外创建全零的差分数组。这里的 diff 不包含任何原值，只记录更新产生的变化。

第二步，同样在位置1加2、位置3减2：

```text
下标                   0   1   2   3（哨兵）
初始增量差分           0   0   0   0
记录更新后的 diff      0  +2   0  -2
前缀累计的增量         0   2   2   0
```

第三步，把实际位置的累计增量加到原数组上：

```text
位置0：原值2 + 累计增量0 = 2
位置1：原值5 + 累计增量2 = 7
位置2：原值5 + 累计增量2 = 7
```

**为什么必须再加原数组？** 因为这次 diff 从零开始，只知道“多了多少”，不知道原来是多少。单独扫描它只能得到 `[0,2,2]`，不能得到 `[2,7,7]`。

#### 两种方式为什么等价？

差分和前缀累加都支持加法叠加：

```text
方式一：前缀和(原数组的差分 + 增量的差分)
      = 原数组 + 增量

方式二：原数组 + 前缀和(增量的差分)
      = 原数组 + 增量
```

区别只在于：**原数组的信息，是一开始放进 diff，还是最后再加回来。**

| 对比项 | 方式一：完整数组差分 | 方式二：只记录增量 |
| --- | --- | --- |
| 初始化 | `diff[0]=a[0]`，其余实际位置为相邻差 | 全部为0 |
| 区间更新 | 左端加、右端之后减 | 左端加、右端之后减 |
| 扫描累计值表示什么？ | 最终数组的值 | 在原值基础上增加的量 |
| 如何写答案？ | `answer[i] = current` | `answer[i] = a[i] + extra` |
| 不执行任何更新时 | 扫描恢复原数组 | 扫描得到全零增量 |

**自检办法**：假设没有任何更新，你的代码是否仍能返回原数组？

- 方式一又加一次原数组，会把原值算两遍，本例错误地得到 `[4,12,12]`。
- 方式二忘记加原数组，会丢失原值，本例错误地得到 `[0,2,2]`。

下节通用 Go 模板使用方式二：变量 `extra` 明确表示“累计增量”，所以答案写成 `value + extra`。

<a id="template"></a>
### 1.5 通用 Go 模板与使用边界

下面采用方式2：原数组不变，`diff` 只记录增量。使用 Go 1.18+ 的 `T Number`，数组值与增量使用 T，下标始终使用 `int`。通过数值类型约束直接执行加减，不需要传入运算回调。

约定每次更新满足 `0 <= Left <= Right < len(nums)`，统一使用0-based整数闭区间。整数运算的中间结果与最终值都应在 T 的范围内；浮点数存在舍入误差，并应避免溢出以及 NaN/Inf。

```go
// 支持有符号整数、浮点数，以及以它们为底层类型的自定义类型。
// 不包含无符号整数，避免负增量和取消事件发生无符号下溢。
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~float32 | ~float64
}

type RangeAdd[T Number] struct {
    Left, Right int
    Delta       T
}

// 返回新数组，不修改 nums 或 updates。
func applyRangeAdds[T Number](nums []T, updates []RangeAdd[T]) []T {
    diff := make([]T, len(nums)+1) // 末尾多一个取消事件的位置
    for _, u := range updates {
        diff[u.Left] += u.Delta
        diff[u.Right+1] -= u.Delta
    }

    answer := make([]T, len(nums))
    var extra T
    for i, value := range nums {
        extra += diff[i]
        answer[i] = value + extra
    }
    return answer
}
```

**调用示例**：

```go
func example() {
    integers := applyRangeAdds(
        []int{2, 5, 5},
        []RangeAdd[int]{{Left: 1, Right: 2, Delta: 2}},
    ) // [2, 7, 7]

    decimals := applyRangeAdds(
        []float64{1.5, 2.5, 3.5},
        []RangeAdd[float64]{{Left: 0, Right: 1, Delta: -0.5}},
    ) // [1, 2, 3.5]

    _ = integers
    _ = decimals
}
```

- `~int64` 表示允许底层类型为 `int64` 的自定义类型，例如 `type Amount int64`。
- `RangeAdd` 将 `int` 下标与 T 类型的增量分开，不用 `[3]T` 混装不同含义的数据。
- 泛型不会扩大数值范围；本篇三题的题设数值范围允许使用 `int`，更大累计量应改用 `int64`。
- **下面所有 Go 代码块共用本节的 `Number`、`RangeAdd` 和 `applyRangeAdds`，放在同一个包内即可编译；提交某一道题时，需要一并复制这些声明。**

时间 `O(m+n)`，额外空间 `O(n)`，m为更新次数。

换题时，主要改变三个地方：

- **区间边界**：闭区间还是半开区间？下标从0还是1开始？
- **贡献大小**：座位数、人数、覆盖次数，还是正负位移？
- **扫描目标**：保存各位置值、判断容量，还是记录最大值？

### 什么时候不直接套用？

- **区间赋值、区间乘法**：普通加法差分不能用同样的两个端点公式直接处理；这里依赖的是可以叠加、可以用相反数取消的加法贡献。
- **每次更新后立即查询**：反复完整扫描可能退化；根据查询类型考虑树状数组或线段树。差分也可以与树状数组结合，不是只能离线使用。
- **坐标巨大但事件稀疏**：仍可使用端点事件思想，但应排序扫描事件，而不是开巨大数组。
- **只查询静态区间和**：优先考虑原数组的前缀和；差分主要是在优化区间更新，不是直接求区间和。

以下各题统一调用 `applyRangeAdds`：题目只负责把自己的区间转换为模板要求的0-based闭区间，再使用返回结果。半开区间 `[l,r)` 转成整数闭区间 `[l,r-1]`，不能直接把 r 当作模板的右端点。

<a id="lc1094"></a>

## 2. [1094. 拼车](https://leetcode.com/problems/car-pooling/)

[练习](l1094/1094.go) · [测试](l1094/1094_test.go)

**题意**：每条行程[人数,上车站,下车站]，判断任意路段是否超出capacity。

**端点**：乘客占用[from,to)，在to已下车，可给同站上车的人腾位，所以to直接减，不是to+1。

**转换**：`[from,to)` → `[from,to-1]`。初始人数全为0，模板返回每个站点出发后的载客量，无须再次做前缀和。

```go
func carPooling(trips [][]int, capacity int) bool {
    maxPosition := 0
    updates := make([]RangeAdd[int], 0, len(trips))
    for _, t := range trips {
        maxPosition = max(maxPosition, t[2])
        updates = append(updates, RangeAdd[int]{
            Left: t[1], Right: t[2]-1, Delta: t[0],
        })
    }
    passengers := applyRangeAdds(
        make([]int, maxPosition+1), updates,
    )
    for _, count := range passengers {
        if count > capacity { return false }
    }
    return true
}
```

**例子**：[[2,1,5],[3,5,7]]、容量3，位置5净变化+1，载荷从2变3，返回true，不是临时5人。

**复杂度**：m条行程、最大坐标U，时间O(m+U+1)，空间O(m+U+1)；本写法额外保存了m条模板更新。

<a id="lc1109"></a>

## 3. [1109. 航班预订统计](https://leetcode.com/problems/corporate-flight-bookings/)

[练习](l1109/1109.go) · [测试](l1109/1109_test.go)

**题意**：[first,last,seats]给编号first到last的每个航班预订seats，返回每航班总座位。

**端点**：题目1-based闭区间，转0-based后[first-1,last-1]，在last位置取消。

**转换**：`[first,last]` → `[first-1,last-1]`。初始座位数全为0，模板返回值直接就是题目答案。

```go
func corpFlightBookings(bookings [][]int, n int) []int {
    updates := make([]RangeAdd[int], 0, len(bookings))
    for _, b := range bookings {
        updates = append(updates, RangeAdd[int]{
            Left: b[0]-1, Right: b[1]-1, Delta: b[2],
        })
    }
    return applyRangeAdds(make([]int, n), updates)
}
```

**例子**：[[1,2,2],[2,3,3]]、n=3，答案[2,5,3]。航班2同时包含在两段中。

**复杂度**：时间O(m+n)，空间O(m+n)，其中模板更新列表占O(m)。模板内部的diff额外开一个哨兵，允许末航班取消事件写入diff[n]。

<a id="lc1854"></a>

## 4. [1854. 人口最多的年份](https://leetcode.com/problems/maximum-population-year/)

[练习](l1854/1854.go) · [测试](l1854/1854_test.go)

**题意**：生年计入、死亡年份不计入；求人口最多的最早年份。

**端点与汇总**：区间[birth,death)，birth加1、death减1。按年升序扫描，只有严格超过历史最大值才更新年份，保留最早并列值。

**转换**：用 `year-1950` 作为下标，`[birth,death)` → `[birth-1950,death-1950-1]`。模板返回每年的实际人口，再扫描最大值，不重复累加。

```go
func maximumPopulation(logs [][]int) int {
    const firstYear, lastYear = 1950, 2050 // 来自本题约束
    updates := make([]RangeAdd[int], 0, len(logs))
    for _, log := range logs {
        updates = append(updates, RangeAdd[int]{
            Left: log[0]-firstYear, Right: log[1]-firstYear-1, Delta: 1,
        })
    }
    population := applyRangeAdds(
        make([]int, lastYear-firstYear+1), updates,
    )
    best, answer := 0, firstYear
    for i, count := range population {
        if count > best {
            best = count
            answer = firstYear + i
        }
    }
    return answer
}
```

**例子**：[[1950,1960],[1960,1970]]，1960一出一进，人口始终1，答案1950。

**复杂度**：时间O(m+U)，空间O(m+U)，其中更新列表占O(m)。本题U为固定101个年份，因此当前写法时间和额外空间均为O(m)，不再是只用固定年份数组的O(1)空间写法。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 取消点 | 扫描时做什么 |
| --- | --- | --- |
| [1094](https://leetcode.com/problems/car-pooling/) | to | 判断容量 |
| [1109](https://leetcode.com/problems/corporate-flight-bookings/) | 0-based的right+1，即原last | 保存全部航班总量 |
| [1854](https://leetcode.com/problems/maximum-population-year/) | death | 更新最大人口最早年份 |

- 固定年份数组只用于[1854](https://leetcode.com/problems/maximum-population-year/)题设，不是任意日期的通用上限。
- 稀疏巨大坐标不要分配U数组：用map聚合同坐标变化，排序扫描，时间O(m log(m+1))、空间O(m)。
- 坐标压缩保留顺序；算覆盖长度还需乘原坐标间距，不能直接用压缩下标差。
- 实数闭区间不能机械r+1，必须明确同坐标事件顺序。
- 差分不是通用在线查询结构；动态区间更新和查询可看 [线段树](../../tree/segment)。

<a id="sparse-sweep"></a>

## 6. 稀疏坐标：从差分数组到端点排序扫描

如果坐标范围 U 很大，但只有 m 个区间，创建 O(U) 数组不划算。将同样的端点变化存成事件，排序后只扫描实际发生变化的位置，时间 O(m log(m+1))、空间 O(m)。负坐标也可直接处理。

通用步骤：选择扫描轴 → 明确端点语义 → 建立事件 → 排序并分组同坐标 → 更新状态 → 提取答案。下面保持原扫描线文档的 `MaxOverlap` / `Compress` 名称；所有 Go 代码块合并时将 `import "sort"` 放在文件的 package 声明之后、其他声明之前。

### 6.1 最大并发：半开区间模板

```go
import "sort"

type event struct { coordinate, delta int }

// intervals 为合法半开区间 [start,end)，空区间不占资源。
func MaxOverlap(intervals [][2]int) int {
    events := make([]event, 0, 2*len(intervals))
    for _, interval := range intervals {
        if interval[0] > interval[1] { panic("reversed interval") }
        if interval[0] == interval[1] { continue }
        events = append(events, event{interval[0],1}, event{interval[1],-1})
    }
    sort.Slice(events, func(i,j int) bool { return events[i].coordinate < events[j].coordinate })
    active, best := 0, 0
    for i := 0; i < len(events); {
        coordinate, delta := events[i].coordinate, 0
        for i < len(events) && events[i].coordinate == coordinate {
            delta += events[i].delta
            i++
        }
        active += delta
        if active > best { best = active }
    }
    return best
}
```

可用于 [253. 会议室 II](https://leetcode.com/problems/meeting-rooms-ii/)（会员题）。拼车则将 delta 从 ±1 改成 ±乘客数，并比较 active 与 capacity；人口题保存第一次达到严格更大 active 的坐标。若要每个数组元素的完整答案，输出规模本身就有 O(n)，不能靠只返回事件绕过。

### 6.2 坐标压缩与覆盖长度

```go
// 返回升序唯一坐标，不修改输入。
func Compress(coordinates []int) []int {
    sorted := append([]int(nil), coordinates...)
    sort.Ints(sorted)
    unique := sorted[:0]
    for _, x := range sorted {
        if len(unique) == 0 || unique[len(unique)-1] != x { unique = append(unique,x) }
    }
    return unique
}
```

时间 O(n log(n+1))、空间 O(n)。用 `sort.SearchInts(unique,x)` 获得排名；排序事件即可解题时不必额外压缩。压缩保留顺序，**不保留距离**：坐标 `[1,100,101]` 的两段实际长度是 99 和 1，不是都为 1。

求并集长度时，在处理当前坐标 x 的事件**之前**，用上一个坐标后的 active 判断 `[previous,x)` 是否被覆盖；若 active>0，则累加 `int64(x)-int64(previous)`，之后再应用 x 的净变化。也可先更新状态，再统计到下一个坐标的距离，两种约定不要混用。这里假定坐标差与总长度均可由 int64 表示。

<a id="advanced-sweep"></a>

## 7. 进阶扫描线：状态不再只是一个累计值

### 7.1 离线查询＋最小堆

[1851. 包含每个查询的最小区间](https://leetcode.com/problems/minimum-interval-to-include-each-query/)：区间按左端点排序；查询按值排序并保存原下标。扫描 query 时，将 `left<=query` 的区间加入按长度排序的最小堆；弹出堆顶 `right<query` 的失效区间，剩余堆顶即最短合法区间。

例如 `[1,4]、[2,3]`：query=2 时答案为长度 2；query=4 时 `[2,3]` 失效，答案为长度 4。这里是**整数闭区间**，长度为 `right-left+1`，不是会议的半开区间。

状态包含已开始的候选，懒删除下堆内可能残留过期区间，但回答前堆顶必须合法；并不要求每次立即清理所有失效项。n 个区间、q 个查询，时间 O(n log(n+1)+q log(q+1)+(n+q)log(n+1))、空间 O(n+q)。

**离线**指允许重排查询的处理顺序，再用原下标恢复答案；查询主要读取状态，不一定改变状态。未来更新不能提前计入当前答案。

### 7.2 天际线：活跃集合的动态最大值

[218. 天际线问题](https://leetcode.com/problems/the-skyline-problem/)：建筑开始加入、结束退出；同坐标事件分组处理后，查询有效建筑最大高度。最大高度改变才输出轮廓关键点。堆＋懒删除、计数有序集合等方式通常为 O(n log n) 时间、O(n) 空间。

不能只保留一个最大高度：最高建筑结束后，还需要知道其余有效建筑中的次高者。差分只记录“总增减”会丢失决定最大值所需的信息；重复高度还要区分建筑身份或维护频次。

### 7.3 在线预订：不要把一次扫描的复杂度当作全部操作的复杂度

[732. 我的日程安排表 III](https://leetcode.com/problems/my-calendar-iii/)：每次 Book 新增 `[start,end)`，立即返回截至本次的最大重叠数。

基线方案维护端点增量 map；每次加入 start:+1、end:-1，再排序当前已知端点并扫描。第 m 次最坏 O(m log(m+1))，m 次累计 O(m² log(m+1))，空间 O(m)。只重扫已收到的数据，仍然符合在线语义，但重复工作较多。

规模更大可用动态线段树等结构维护区间加与全局最大值。动态并不代表一定要用高级结构，而是必须核算多次更新和查询的总成本。普通区间和 / 频次问题还可考虑[树状数组](../../tree/fenwick/README.md)，动态区间维护可看[线段树](../../tree/segment/readme.md)。

<a id="extended-practice"></a>

## 8. 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [2381. 字母移位II](https://leetcode.com/problems/shifting-letters-ii/) | 闭区间增减字母位移，差分累计后对26规范化 | 负余数要加26；取消点right+1 |
| [3355. 零数组变换I](https://leetcode.com/problems/zero-array-transformation-i/) | 每条查询覆盖处可独立减0或1，差分计算覆盖能力 | 比较覆盖数>=nums[i]，不是每区间强制全部减1 |
| [1893. 检查是否区域内所有整数都被覆盖](https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/) | 整数闭区间差分，目标区间每点覆盖数至少1 | 取消点r+1，不要只比较总覆盖长度 |
| [2406. 将区间分为最少组数](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/) | 整数闭区间最大并发覆盖即最少组数 | 端点相接算重叠；大坐标用稀疏事件或最小堆 |
| [2536. 子矩阵元素加1](https://leetcode.com/problems/increment-submatrices-by-one/) | 二维矩形四角差分，再二维前缀恢复 | 额外哨兵行列；四角符号+ - - + |
| [253. 会议室 II](https://leetcode.com/problems/meeting-rooms-ii/)（会员题） | 半开区间端点扫描，取最大活跃数量 | 同时结束和开始可以复用房间 |
| [1851. 包含每个查询的最小区间](https://leetcode.com/problems/minimum-interval-to-include-each-query/) | 离线查询＋按长度排序的候选堆 | 查询前清理失效堆顶；保存原查询下标 |
| [218. 天际线问题](https://leetcode.com/problems/the-skyline-problem/) | 扫描建筑端点＋动态最高有效建筑 | 不能只用加法累计高度，重复高度要区分身份或计数 |
| [732. 我的日程安排表 III](https://leetcode.com/problems/my-calendar-iii/) | 半开预订记录端点变化，每次返回最大重叠 | 在线更新与复杂度见第 7 节，不提前使用未来预订 |

<a id="practice-order"></a>

## 9. 推荐练习顺序

- **差分基础**：[1109. 航班预订统计](https://leetcode.com/problems/corporate-flight-bookings/) → [1094. 拼车](https://leetcode.com/problems/car-pooling/) → [1854. 人口最多的年份](https://leetcode.com/problems/maximum-population-year/)。
- **条件迁移**：[1893. 整数覆盖](https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/) → [2381. 字母移位 II](https://leetcode.com/problems/shifting-letters-ii/) → [3355. 零数组变换 I](https://leetcode.com/problems/zero-array-transformation-i/)。
- **稀疏事件**：[253. 会议室 II](https://leetcode.com/problems/meeting-rooms-ii/)（会员题） → [2406. 将区间分为最少组数](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/)。
- **进阶分支**：[2536. 子矩阵元素加 1](https://leetcode.com/problems/increment-submatrices-by-one/)（二维差分）；[1851. 最短覆盖区间](https://leetcode.com/problems/minimum-interval-to-include-each-query/) → [218. 天际线](https://leetcode.com/problems/the-skyline-problem/) → [732. 在线预订](https://leetcode.com/problems/my-calendar-iii/)（扫描线与动态结构）。

**复习方法**：先写清“事件是什么、状态是什么、何时读取答案”，再确认端点语义、坐标规模和查询时序。能加减累计才考虑差分；不能只用累计和表达状态时，选择合适的候选结构，不为统一模板牺牲正确性。
