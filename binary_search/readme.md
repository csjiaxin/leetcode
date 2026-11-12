# 二分查找

## 目录

- [1. 核心思想](#core)
- [2. 基础查找：闭区间](#basic)
- [3. 边界查找：第一个 true](#boundary)
  - [3.1 应用：有序重复数组的首尾下标](#duplicates)
    - [第一个等于目标的位置：lower](#first-equal)
    - [最后一个等于目标的位置：upper - 1](#last-equal)
    - [组合求解、示例与边界](#duplicate-example)
  - [3.2 firstTrue 条件速查表](#first-true-cheatsheet)
- [4. 答案二分](#answer)
- [5. 特殊结构](#special)
- [6. 常见错误](#pitfalls)
- [7. 运行测试](#tests)
- [8. 逐题说明与解释：firstTrue 与 check](#explanations)
  - [704. 二分查找](#lc704)
  - [34. 第一个和最后一个位置](#lc34)
  - [153. 旋转数组最小值](#lc153)
  - [33. 搜索旋转数组](#lc33)
  - [81. 搜索旋转数组 II](#lc81)
  - [162. 寻找峰值](#lc162)
  - [1901. 寻找峰值 II](#lc1901)
  - [875. 爱吃香蕉的珂珂](#lc875)
  - [1011. D 天送达包裹](#lc1011)
  - [410. 分割数组的最大值](#lc410)
  - [1760. 袋子里最少数目的球](#lc1760)
  - [2024. 考试的最大困扰度](#lc2024)
  - [2498. 青蛙过河 II](#lc2498)
  - [1170. 比较字符串最小字母出现频次](#lc1170)
  - [1608. 特殊数组](#lc1608)
  - [1855. 下标对的最大距离](#lc1855)
  - [2089. 找出目标下标](#lc2089)
  - [729. 我的日程安排表 I](#lc729)
  - [888. 公平的糖果交换](#lc888)

<a id="core"></a>

## 1. 核心思想

每次排除一半不可能包含答案的搜索空间。数组不一定整体有序，但必须证明丢弃一半后，剩余区间仍包含答案。

做题前先回答：

1. 二分什么：数组下标，还是答案值？
2. 为什么可以排除一半：单调性、局部有序性，还是峰值存在性？
3. 找什么：任意目标、第一个满足条件的位置、最后一个满足条件的位置？
4. 搜索区间统一为闭区间 `[l,r]`，循环条件为 `l<=r`。
5. 为什么不会漏解，为什么区间一定收缩？

<a id="basic"></a>

## 2. 基础查找：闭区间 [l,r]

```go
func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := l + (r-l)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return -1
}
```

`l <= r` 保证单元素区间仍被检查；中点已被排除，更新时使用 `mid±1`。重复值只保证找到任意一个。时间 `O(log n)`，额外空间 `O(1)`。

<a id="boundary"></a>

## 3. 边界查找：统一为第一个 true

把条件转换为 `false false | true true`，寻找分界点。

```go
func firstTrue(n int, predicate func(int) bool) int {
    l, r := 0, n-1
    for l <= r {
        mid := l + (r-l)/2
        if predicate(mid) {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}
```

循环不变量：下标小于 l 的位置全为 false，下标大于 r 的位置全为 true，闭区间 `[l,r]` 尚待确定。满足条件时，中点已被归入 true 区域，可继续找更靠左的位置，因此使用 `r=mid-1`。退出时 `l=r+1`，l 就是第一个 true；全 false 返回 n，全 true 返回 0，空数组返回 0。

对于升序数组，使用上述闭区间模板（各题独立实现，不依赖标准库二分）：

```go
lower := firstTrue(len(nums), func(i int) bool { return nums[i] >= target })
upper := firstTrue(len(nums), func(i int) bool { return nums[i] > target })
```

注意：下表中的 `[lower,upper)` 是结果范围的表达，不是二分循环的搜索区间。搜索过程始终使用闭区间。

| 需求 | 结果 |
| --- | --- |
| 第一个 >= target | lower |
| 第一个 > target | upper |
| 最后一个 < target | lower - 1 |
| 最后一个 <= target | upper - 1 |
| target 出现次数 | upper - lower |
| target 的全部位置 | [lower,upper) |

存在性判断：`lower < len(nums) && nums[lower] == target`。访问 `n` 或 `-1` 哨兵位置前必须检查边界。

应用：[1170](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/) 的严格大于计数是 `n-upper`；[2089](https://leetcode.com/problems/find-target-indices-after-sorting-array/) 返回 `[lower,upper)`；[1855](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/) 在非递增数组中找第一个小于目标的位置；[729](https://leetcode.com/problems/my-calendar-i/) 找插入点后检查前后邻居，但数组插入仍需要 `O(n)`。

<a id="duplicates"></a>

### 3.1 应用：有序重复数组的首尾下标

前提是数组**非递减有序**，允许重复值。目标是找到第一个和最后一个**等于** target 的下标，不存在时返回 `[-1,-1]`。

无需再记另一套二分模板：直接使用本节的闭区间 `firstTrue` 两次，分别找第一个 `>= target` 的位置 lower，以及第一个 `> target` 的位置 upper。

<a id="first-equal"></a>

#### 第一个等于目标的位置：lower

```go
lower := firstTrue(len(nums), func(i int) bool {
    return nums[i] >= target
})
```

条件 `nums[i]>=target` 为 `false→true`。lower 是第一个不小于目标的位置：目标存在时，它就是目标首下标；不存在时，它可能指向更大的值或等于 n，不能直接当作有效下标访问。

<a id="last-equal"></a>

#### 最后一个等于目标的位置：upper - 1

```go
upper := firstTrue(len(nums), func(i int) bool {
    return nums[i] > target
})
```

条件 `nums[i]>target` 同样为 `false→true`。upper 是第一个大于目标的位置，其前一个位置 upper-1 就是最后一个不大于目标的位置。目标存在时，它就是目标末下标。

**两次搜索都找第一个 true，也都返回 l**，无需另外编写“最后一个 true、返回 r”的循环。只有 predicate 在相等时的结果不同：

| 比较结果 | >= target：寻找 lower | > target：寻找 upper |
| --- | --- | --- |
| nums[mid] < target | false，l=mid+1 | false，l=mid+1 |
| nums[mid] == target | true，r=mid-1 | false，l=mid+1 |
| nums[mid] > target | true，r=mid-1 | true，r=mid-1 |
| 退出后返回 | l，即 lower | l，即 upper |

直接比较 `>target`，不要用 `>=target+1`，以免 target 为整数最大值时溢出。

<a id="duplicate-example"></a>

#### 组合求解、示例与边界

复用第 3 节定义的 `firstTrue`：

```go
func searchRange(nums []int, target int) []int {
    lower := firstTrue(len(nums), func(i int) bool {
        return nums[i] >= target
    })
    upper := firstTrue(len(nums), func(i int) bool {
        return nums[i] > target
    })

    if lower == upper {
        return []int{-1, -1}
    }
    return []int{lower, upper - 1}
}
```

**为什么 lower==upper 就表示不存在？** 排序后，`[lower,upper)` 中的值同时满足 `>=target` 和 `<=target`，因此恰好全部等于 target。两边界相等意味着这段区间为空。该判断无需访问数组，因此空数组、lower=n 或 upper=0 都能安全处理。

以 `nums=[1,2,2,2,4]、target=2` 为例：

```text
下标        0      1      2      3      4
nums        1      2      2      2      4
>= target   false  true   true   true   true
                   ↑ lower=1
> target    false  false  false  false  true
                                        ↑ upper=4
首尾下标 = [lower, upper-1] = [1,3]
出现次数 = upper-lower = 3
```

两次二分的闭区间变化：

| 条件 | l | r | mid | 中点条件 | 更新 |
| --- | --- | --- | --- | --- | --- |
| >= target | 0 | 4 | 2 | true | r=1 |
| >= target | 0 | 1 | 0 | false | l=1 |
| >= target | 1 | 1 | 1 | true | r=0，退出，返回 l=1 |
| > target | 0 | 4 | 2 | false | l=3 |
| > target | 3 | 4 | 3 | false | l=4 |
| > target | 4 | 4 | 4 | true | r=3，退出，返回 l=4 |

| nums | target | lower | upper | 首尾下标 |
| --- | --- | --- | --- | --- |
| [] | 2 | 0 | 0 | [-1,-1] |
| [2] | 2 | 0 | 1 | [0,0] |
| [2,2,2] | 2 | 0 | 3 | [0,2] |
| [1,3,3] | 2 | 1 | 1 | [-1,-1] |
| [2,2] | 1 | 0 | 0 | [-1,-1] |
| [2,2] | 3 | 2 | 2 | [-1,-1] |

这正是 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) 的解法。两次二分总时间 `O(log n)`，额外空间 `O(1)`；不要找到一个 target 后再线性扩展，全部相等时会退化到 `O(n)`。

注意：`[lower,upper)` 仅表示结果区间，搜索始终使用闭区间 `[l,r]`。若原数组无序，排序后得到的是**排序后下标**（如 [2089](https://leetcode.com/problems/find-target-indices-after-sorting-array/)），不是原数组位置。

<a id="first-true-cheatsheet"></a>

### 3.2 firstTrue 条件速查表

下面统一复用第 3 节的闭区间模板：`p=firstTrue(N,predicate)`。实际搜索范围为 `[0,N-1]`，predicate 必须先 false 后 true；返回 N 表示范围内没有 true。表中的“答案转换”是在得到分界点后进行的，不改变二分模板。

| 题目 / 用途 | N 与搜索对象 | predicate：何时为 true | 分界点后的答案转换 | 必要检查 |
| --- | --- | --- | --- | --- |
| [34](#lc34)（[LeetCode](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)）：首个目标位置 | N=n，数组下标 i | `nums[i]>=target` | 得到 lower | 与 upper 一起判断目标是否存在 |
| [34](#lc34)（[LeetCode](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)）：末个目标位置 | N=n，数组下标 i | `nums[i]>target` | 得到 upper，末下标为 upper-1 | lower=upper 时返回 `[-1,-1]`，不能直接访问 upper-1 |
| [153](#lc153)（[LeetCode](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)）：旋转数组最小值 | N=n，数组下标 i | `nums[i]<=last` | 返回 nums[p] | 非空、无重复；last 固定为原数组末值；必有 p<n |
| [1170](#lc1170)（[LeetCode](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/)）：严格大于查询频次的数量 | N=w，排序后的频次数组下标 j | `freqs[j]>q` | 返回 w-p | p=w 时返回 0；相等不计入 |
| [1608](#lc1608)（[LeetCode](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/)）：内层 count(x) | N=n，排序后的数组下标 i | `nums[i]>=x` | count(x)=n-p | p=n 表示数量为 0 |
| [1608](#lc1608)（[LeetCode](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/)）：外层候选 x | N=n+1，候选值 x∈[0,n] | `count(x)<=x` | p 即候选 x | 必须检查 count(p)=p，否则返回 -1；此处必有 p<=n |
| [1855](#lc1855)（[LeetCode](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/)）：最后一个满足大小条件的下标 | N=n，nums2 的下标 j | `nums2[j]<nums1[i]` | 最后满足位置 j=p-1，用 j-i 更新最大距离 | nums2 非递增；仅在 j>=i 时更新；p=0 时没有满足位置 |
| [2089](#lc2089)（[LeetCode](https://leetcode.com/problems/find-target-indices-after-sorting-array/)）：排序后的全部目标下标 | N=n，排序后的数组下标 i，调用两次 | `nums[i]>=target`；`nums[i]>target` | 得到 lower、upper，输出 lower 到 upper-1 | lower=upper 时输出空数组；下标属于排序后数组 |
| [729](#lc729)（[LeetCode](https://leetcode.com/problems/my-calendar-i/)）：日程插入位置 | N=n，按起点排序的区间下标 j | `bookings[j].start>=start` | p 为插入位置，检查邻居后插入 | p>0 才检查前驱，p<n 才检查后继；p=n 是合法尾插位置 |

**记忆三步：**

1. **写出 false→true 的条件**：不是看原数组是否递增，而是看 predicate 的结果是否单调。例如 [1855](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/) 的数组非递增，但“小于目标”仍然是先 false 后 true。
2. **先求分界点，再换算答案**：求首个位置用 p，求最后一个满足位置通常用“第一个不满足的位置减一”，求后缀数量用 N-p。
3. **识别哨兵并检查存在性**：p=N 可以表示无匹配，也可以表示合法的尾部插入点；p-1=-1 不能访问。按题意判断，不要把分界点直接当作有效元素下标。

**注意区分：** 普通精确查找、含重复值的旋转数组查找和峰值查找，不要直接套用表中的 predicate；必须先证明条件为 false→true。答案二分也可以使用这一思想，但搜索的是答案值域，不能将速度、容量等答案值直接当成数组下标。

<a id="answer"></a>

## 4. 答案二分

将最优化问题转换为可行性判断：确定答案范围，定义 `check(x)`，证明单调性，再找分界点。

### 最小可行值：false → true

```go
// 初始 [l,r] 覆盖答案范围，并存在可行值。
for l <= r {
    mid := l + (r-l)/2
    if check(mid) {
        r = mid - 1
    } else {
        l = mid + 1
    }
}
// l 即最小可行值。
```

[1760. 袋子里最少数目的球](https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/)：每袋容量上限为 x，一个有 v 个球的袋子需要 `(v-1)/x` 次操作。x 越大，操作越少。搜索范围 `[1,max(nums)]`。

[875. 爱吃香蕉的珂珂](https://leetcode.com/problems/koko-eating-bananas/)：速度越大，耗时越少；每堆耗时用整数向上取整。

[1011. 在 D 天内送达包裹的能力](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)、[410. 分割数组的最大值](https://leetcode.com/problems/split-array-largest-sum/)：容量或最大子数组和越大，需要的天数或组数越少。贪心判断时从左到右尽量装满，搜索范围为 `[max(nums),sum(nums)]`。其中 410 题的数组为非负数且 `k<=n`，所以可将不足 k 组继续拆分为恰好 k 个非空组。

### 最大可行值：true → false

```go
// 初始 [l,r] 覆盖答案范围，并存在可行值。
for l <= r {
    mid := l + (r-l)/2
    if check(mid) {
        l = mid + 1
    } else {
        r = mid - 1
    }
}
// r 即最大可行值。
```

两个模板都用下中位数和 `mid±1`，不需要上中位数。最小可行值返回 l，最大可行值返回 r；若范围内没有可行值，分别得到初始上界加一、初始下界减一，需要调用方识别该哨兵。

复杂度通常是 `O(check复杂度 × log(答案范围))`。累加、乘法或向上取整需留意溢出，必要时使用 `int64`。

能二分不代表二分最优：[2024. 考试的最大困扰度](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/) 用滑动窗口可做到 `O(n)`；[2498. 青蛙过河 II](https://leetcode.com/problems/frog-jump-ii/) 可直接取跨一个石头的最大距离，做到 `O(n)`。本目录提供空函数与测试，供练习答案二分。

<a id="special"></a>

## 5. 特殊结构

### 旋转数组：[33](https://leetcode.com/problems/search-in-rotated-sorted-array/)、[81](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)、[153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

[33. 搜索旋转排序数组](https://leetcode.com/problems/search-in-rotated-sorted-array/)：先识别有序半区，再判断 target 是否在该半区。

[153. 寻找旋转排序数组中的最小值](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)：固定原数组末值 last，条件 `nums[mid]<=last` 构成先 false 后 true 的序列；用 `firstTrue` 查找分界点，即最小值下标。注意 last 固定不随 r 移动。

[81. 搜索旋转排序数组 II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)：检查中点不是目标后，若 `nums[mid]==nums[r]`，不能确定哪半边有序，但可以安全 `r--`。最坏时间退化为 `O(n)`。

### 峰值：[162](https://leetcode.com/problems/find-peak-element/)、[1901](https://leetcode.com/problems/find-a-peak-element-ii/)

[162. 寻找峰值](https://leetcode.com/problems/find-peak-element/)：在相邻元素不等、两端外侧视为负无穷的题设下，若左邻居更大，执行 `r=mid-1`；否则若右邻居更大，执行 `l=mid+1`；否则 mid 就是峰值，直接返回。访问邻居前检查数组边界。这不是单调布尔序列，而是保留“区间内存在峰”的性质。

[1901. 寻找峰值 II](https://leetcode.com/problems/find-a-peak-element-ii/)：二分列，找当前列最大值，再与同一行的相邻列比较方向。最终返回列的最大值；单列也需正确处理。复杂度 `O(rows × log(cols+1))`。测试应验证峰值性质，不应限定某个坐标。

<a id="pitfalls"></a>

## 6. 常见错误

- 混用模板：本目录统一使用闭区间，数组初始 `r=n-1`、循环 `l<=r`、更新 `l=mid+1` 或 `r=mid-1`；答案值域上界按题意确定，不是统一减一。不要混入 `r=mid` 或 `l=mid`。日程题本身的 `[start,end)` 时间区间语义不变，与二分搜索区间无关。
- 相等就返回，导致无法找到最左或最右边界。
- 没证明 check 的单调性；仅凭题目问最值就使用二分。
- 在 `l<=r` 循环中使用 `l=mid` 或 `r=mid`，导致区间无法严格收缩。
- 忽略空数组、单元素、全部相等、目标不存在、全真/全假与整数溢出。
- 将排序或插入成本忽略：排序后查找总体不一定是 `O(log n)`。

<a id="tests"></a>

## 7. 运行测试

当前函数体为空，带返回值的函数会报 `missing return`，暂时无法编译或通过测试。完成对应题目实现后，在仓库根目录执行（需要已有 Go module 配置）：

```sh
go test ./binary_search/...
```

只运行单题：`go test ./binary_search/l81 -v`。

<a id="explanations"></a>

## 8. 逐题说明与解释：firstTrue 与 check

每题位于 `l题号/题号.go`，测试位于 `l题号/题号_test.go`。以下只修改文档说明，题目函数体继续留空。统一约定：

- `check` 描述判定条件；传给 `firstTrue` 的条件必须为 **false→true**。check 若是 true→false，则搜索 `!check`，再减一。
- 数组下标调用 `firstTrue(n, check)`，复用第 3 节闭区间模板。返回 n 是哨兵，不一定能直接访问。
- 答案值域 `[lo,hi]` 使用偏移：`answer=lo+firstTrue(hi-lo+1, func(offset int) bool { return check(lo+offset) })`。offset 是下标，lo+offset 才是真实答案值；lo、hi 为初始固定边界。
- 特殊结构不能只因命名为 check 就宣称单调。81 先线性定位旋转断点再用 firstTrue；162、1901 保留基于方向判定的闭区间二分，不把非单调 check 传给 firstTrue。

下列片段用于理解模板组合，不写入练习 `.go` 文件。

<a id="lc704"></a>

### [704. 二分查找](https://leetcode.com/problems/binary-search/)

- **题意**：在严格递增数组中找 target，下标不存在时返回 -1。例如 `[-1,0,3,5,9,12]、target=9`，答案为 4。
- **check**：`check(i)=nums[i]>=target`，条件为 false→true。
- **firstTrue 与答案**：`p=firstTrue(n,check)`；若 `p<n && nums[p]==target`，返回 p，否则返回 -1。
- **解释**：先找第一个不小于目标的位置，再检查是否相等，将精确查找转成边界查找。空数组时返回哨兵 0，短路判断不会访问越界。
- **复杂度与边界**：时间 `O(log n)`，额外空间 `O(1)`。检查空数组、单元素和目标在数组范围外。

<a id="lc34"></a>

### [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

- **题意**：非递减数组允许重复，返回目标的首尾下标；不存在返回 `[-1,-1]`。例如 `[5,7,7,8,8,10]、target=8`，答案 `[3,4]`。
- **check**：`checkLower(i)=nums[i]>=target`，`checkUpper(i)=nums[i]>target`，两者均为 false→true。
- **firstTrue 与答案**：`lower=firstTrue(n,checkLower)`，`upper=firstTrue(n,checkUpper)`。lower=upper 时返回 `[-1,-1]`，否则返回 `[lower,upper-1]`，详见 [3.1 节](#duplicates)。
- **解释**：目标值连续出现，首尾位置就是两个分界点。普通查找相等即返回，只能命中任意一个，不能保证边界。
- **复杂度与边界**：时间 `O(log n)`，额外空间 `O(1)`。通过 `lower==upper` 判断目标不存在，无需访问哨兵下标；不要线性扩展左右边界。

<a id="lc153"></a>

### [153. 寻找旋转排序数组中的最小值](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

- **题意**：严格递增数组旋转后，找最小元素。例如 `[3,4,5,1,2]` 返回 1；未旋转也需要支持。
- **check**：固定原数组最后一个值 last，定义 `check(i)=nums[i]<=last`，旋转点之前为 false、之后为 true。
- **firstTrue 与答案**：`p=firstTrue(n,check)`，返回 `nums[p]`。末元素本身一定满足，因此 p<n，可以安全访问。
- **解释**：旋转后前半段都大于 last，后半段都小于等于 last，形成 `false→true`；未旋转时全部为 true，得到下标 0。
- **复杂度与边界**：时间 `O(log n)`，额外空间 `O(1)`。要求非空且无重复值；last 不能随搜索右边界移动。

<a id="lc33"></a>

### [33. 搜索旋转排序数组](https://leetcode.com/problems/search-in-rotated-sorted-array/)

- **题意**：在无重复值的旋转有序数组中找 target。例如 `[4,5,6,7,0,1,2]、target=0` 返回 4。
- **check 与第一阶段**：空数组直接返回 -1。固定 `last=nums[n-1]`，定义 `checkPivot(i)=nums[i]<=last`，令 `pivot=firstTrue(n,checkPivot)` 找到最小值位置。
- **check 与第二阶段**：把数组视为虚拟有序数组 `value(i)=nums[(pivot+i)%n]`，定义 `check(i)=value(i)>=target`，令 `p=firstTrue(n,check)`。
- **答案**：若 `p<n && value(p)==target`，返回原数组下标 `(pivot+p)%n`，否则返回 -1。
- **解释**：两次 firstTrue 的 check 都是 false→true。先找旋转点，再恢复虚拟有序访问，无需真的旋转或复制数组。
- **复杂度与边界**：时间 `O(log n)`，额外空间 `O(1)`。注意两元素、未旋转、目标不存在及两端比较符号。

<a id="lc81"></a>

### [81. 搜索旋转排序数组 II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)

- **题意**：[33. 搜索旋转排序数组](https://leetcode.com/problems/search-in-rotated-sorted-array/) 的允许重复值版本，只需返回是否存在目标。例如 `[1,0,1,1,1]、target=0` 返回 true。
- **先处理不单调性**：重复值下 `nums[i]<=last` 不一定 false→true，例如 `[1,2,0,1]` 的结果为 true、false、true、true，不单调；`[1,0,1,1,1]` 更是全部 true，无法识别真实断点，不能照搬 33 的旋转点 check。
- **预处理**：空数组返回 false。线性找第一个 `nums[i]<nums[i-1]` 的下标作为 pivot；没有下降则 pivot=0，得到真正的有序起点。
- **check**：虚拟有序访问 `value(i)=nums[(pivot+i)%n]`，定义 `check(i)=value(i)>=target`，此时才是 false→true。
- **firstTrue 与答案**：`p=firstTrue(n,check)`；返回 `p<n && value(p)==target`。
- **解释与复杂度**：为统一 firstTrue，先线性找断点，整体时间 `O(n)`，额外空间 `O(1)`。第 5 节原生旋转二分可在无歧义时折半，但最坏同样 `O(n)`；不能为了套模板声称本版本为 `O(log n)`。

<a id="lc162"></a>

### [162. 寻找峰值](https://leetcode.com/problems/find-peak-element/)

- **题意**：相邻元素不等，找到任意一个严格大于两侧邻居的位置；数组外侧视为负无穷。例如 `[1,2,3,1]` 返回 2。
- **check**：定义 `checkPeak(i)` 判断当前位置是否严格大于所有存在的相邻元素；另用左右邻居大小决定搜索方向。
- **firstTrue 的限制**：多峰时 checkPeak 不单调；`nums[i]>nums[i+1]` 也可能反复真假，不能传给 firstTrue 并声称寻找第一个 true。此题保留方向型闭区间二分。
- **闭区间调用逻辑**：在 `[0,n-1]`、`l<=r` 中，若 checkPeak(mid) 为 true 就返回 mid；否则左邻居更大时 `r=mid-1`，否则 `l=mid+1`。访问邻居前检查边界。
- **解释**：沿更大的邻居方向上坡，总能在转折处或数组端点遇到峰。不需要全数组有序，也不要求比较条件全局单调。
- **复杂度与边界**：时间 `O(log n)`，额外空间 `O(1)`。要求非空；单元素、严格递增、严格递减都存在峰，多峰时测试只验证峰值性质。

<a id="lc1901"></a>

### [1901. 寻找峰值 II](https://leetcode.com/problems/find-a-peak-element-ii/)

- **题意**：在相邻元素不等的正整数矩阵中，找严格大于上下左右邻居的元素坐标，边界外视为 -1。
- **check**：扫描列 c 取得最大值所在行 row，定义 `checkPeak(c)` 为该元素也严格大于存在的左右邻居；上下邻居由列最大值及相邻不等保证。
- **firstTrue 的限制**：不同列的 checkPeak 可能反复真假，不满足 firstTrue 的单调前提；不能直接用 firstTrue 查峰值列。
- **闭区间调用逻辑**：二分列 `[0,cols-1]`，`l<=r`。checkPeak(mid) 为 true 时返回 `[row,mid]`；否则若同一行左邻居更大，`r=mid-1`，否则 `l=mid+1`。每轮仅扫描当前列一次，复用 row，不重复计算。
- **解释**：列最大值在相邻不等的条件下严格大于上下邻居，只需处理水平比较。朝更大的水平邻居移动，可保留存在峰值的一半。
- **复杂度与边界**：时间 `O(rows × log(cols+1))`，额外空间 `O(1)`。单列也必须扫描并返回最大值，不能依赖循环更新时才保存坐标。

<a id="lc875"></a>

### [875. 爱吃香蕉的珂珂](https://leetcode.com/problems/koko-eating-bananas/)

- **题意**：每小时只吃一堆，速度为每小时 k 根，求 h 小时内吃完所有香蕉的最小整数速度。例如 `[3,6,7,11]、h=8`，答案 4。
- **check**：`check(speed)` 判断 `sum((pile-1)/speed+1)<=h`；按堆累加，超过预算即返回 false。速度范围 `[1,M]`，M=max(piles)，check 为 false→true。
- **firstTrue 与答案**：`answer=1+firstTrue(M, func(offset int) bool { return check(1+offset) })`。搜索的是偏移 `[0,M-1]`，真实速度为 1+offset，因此不会除零。
- **解释**：速度越快，总时间越少，check 为 `false→true`。同一小时不能把剩余时间用于另一堆，所以必须逐堆向上取整，而不是对香蕉总数取整。
- **复杂度与边界**：时间 `O(n log M)`，M 为最大堆大小；额外空间 `O(1)`。题设 `h>=n` 保证有解；使用剩余小时递减或 int64 累加避免溢出。

<a id="lc1011"></a>

### [1011. 在 D 天内送达包裹的能力](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)

- **题意**：包裹必须按原顺序装船，每天总重量不能超过容量，求 days 天内运完的最小容量。例如 `[1,2,3,4,5,6,7,8,9,10]、days=5`，答案 15。
- **check**：`check(capacity)` 从左到右装包裹，当前天装不下就开始下一天，返回所需天数是否 `<=days`。定义固定边界 lo=max(weights)、hi=sum(weights)，check 为 false→true。
- **firstTrue 与答案**：`answer=lo+firstTrue(hi-lo+1, func(offset int) bool { return check(lo+offset) })`。hi 一定可行；单个包裹重量不会超过传入容量。
- **解释**：固定容量时，每天尽量装满不会增加未来工作量，贪心可得到最少天数；容量越大，需要天数不会增加，满足单调性。
- **复杂度与边界**：时间 `O(n log(S-M+2))`，S 为总重量、M 为最大重量（加 2 也覆盖只有一个候选值的情况）；额外空间 `O(1)`。一天时答案为总和，天数充足时答案为最大重量。不能排序包裹。

<a id="lc410"></a>

### [410. 分割数组的最大值](https://leetcode.com/problems/split-array-largest-sum/)

- **题意**：把非负数组划分为恰好 k 个连续非空子数组，最小化这些子数组和的最大值。例如 `[7,2,5,10,8]、k=2`，答案 18，对应 `[7,2,5]` 和 `[10,8]`。
- **check**：`check(limit)` 贪心累计非负元素，加入下一个值会超过上限时新开一组，返回组数是否 `<=k`。固定 lo=max(nums)、hi=sum(nums)，check 为 false→true。
- **firstTrue 与答案**：`answer=lo+firstTrue(hi-lo+1, func(offset int) bool { return check(lo+offset) })`。hi 必可行；全零时 lo=hi=0，仍搜索一个候选，答案为 0。
- **解释**：非负数保证贪心开组得到最少组数；若不足 k 组，在 `k<=n` 下还能继续拆成 k 个非空组而不增加最大组和。允许上限越大，所需组数越少。
- **复杂度与边界**：时间 `O(n log(S-M+2))`，S 为总和、M 为最大元素，额外空间 `O(1)`。全零、k=1、k=n 都要处理；若允许负数，上述贪心依据不成立。

<a id="lc1760"></a>

### [1760. 袋子里最少数目的球](https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/)

- **题意**：每次把一个袋子拆为两个非空袋子，最多操作 maxOperations 次，最小化最终单袋球数的最大值。例如 `[9]、maxOperations=2`，答案 3。
- **check**：`check(limit)` 判断 `sum((v-1)/limit)<=maxOperations`；v 个球至少需要 `ceil(v/limit)` 袋，因此操作次数为 `ceil(v/limit)-1=(v-1)/limit`。定义 M=max(nums)，check 为 false→true。
- **firstTrue 与答案**：`answer=1+firstTrue(M, func(offset int) bool { return check(1+offset) })`。真实容量范围 `[1,M]`，不会除零，M 必可行。
- **解释**：拆成 b 袋恰需 b-1 次操作；容量越大，每袋需要的拆分次数越少，check 为 `false→true`。
- **复杂度与边界**：时间 `O(n log M)`，额外空间 `O(1)`。操作次数为零时答案是最大袋大小；limit 从 1 开始，避免除零，可在预算耗尽时提前退出。

<a id="lc2024"></a>

### [2024. 考试的最大困扰度](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/)

- **题意**：最多翻转 k 个 T/F 字符，求能得到的最长连续相同字符长度。例如 `TTFF、k=2`，答案 4。
- **check**：`check(length)` 判断是否存在该长度的窗口满足 `min(countT,countF)<=k`；使用固定长度滑动窗口线性检查，并定义 check(0)=true。长度越大越难，check 为 true→false。
- **firstTrue 与答案**：`p=firstTrue(n+1, func(length int) bool { return !check(length) })` 找第一个不可行长度，返回 `p-1`。传入的 !check 为 false→true，不能直接传 check。
- **解释**：把少数字符翻成多数字符最省操作。长度 L 可行时，更短子段也可行；若 `[0,n]` 全可行，firstTrue 返回哨兵 n+1，减一后正确得到 n。
- **复杂度与边界**：二分版本时间 `O(n log(n+1))`，额外空间 `O(1)`。k=0 时寻找已有最长连续段；长度 0 视为可行。直接可变窗口能做到 `O(n)`，二分用于练习最大可行值模板。

<a id="lc2498"></a>

### [2498. 青蛙过河 II](https://leetcode.com/problems/frog-jump-ii/)

- **题意**：石头坐标严格递增，青蛙从第一块到最后一块再返回，每块石头至多访问一次（起点是往返端点），最小化整条路线中的最大跳跃距离。例如 `[0,2,5,6,7]`，答案 5。
- **check**：`check(limit)` 先检查 `stones[1]-stones[0]<=limit`，再检查所有 i>=2 是否满足 `stones[i]-stones[i-2]<=limit`。这个统一判断覆盖 n=2；令 D=stones[n-1]-stones[0]，check 为 false→true。
- **firstTrue 与答案**：`answer=firstTrue(D+1,check)`，真实距离与候选下标相同，范围 `[0,D]`，无需偏移。D 必可行，结果不超过 D。
- **解释**：任意连续三块石头中，若中间石头被去程使用，回程必须跨过它；反之亦然。因此最大跳跃至少覆盖每个两步间距。去程走交替石头、回程走剩余石头，就能达到这个下界，故条件充分且必要。
- **复杂度与边界**：二分版本时间 `O(n log(D+1))`，额外空间 `O(1)`，D 为首尾距离。要求至少两块石头。直接取两步间距最大值（n=2 单独处理）可做到 `O(n)`。

<a id="lc1170"></a>

### [1170. 比较字符串最小字母出现频次](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/)

- **题意**：f(s) 表示字符串中最小字母的出现次数。对每个 query，统计满足 `f(word)>f(query)` 的 word 数量。例如 queries=`["bbb","cc"]`、words=`["a","aa","aaa","aaaa"]`，答案 `[1,2]`。
- **check**：计算所有 word 的频次并排序。对每个查询频次 q，定义 `check(j)=freqs[j]>q`，排序后为 false→true，相等属于 false。
- **firstTrue 与答案**：`upper=firstTrue(w,check)`，答案为 `w-upper`。upper=w 表示没有符合条件的 word，upper=0 表示全部符合。
- **解释**：排序后频次严格大于 q 的值构成后缀。相等不能计入答案，所以相等时必须继续向右。f(s) 不是出现最多的字母频次，例如 `abbbb` 的 f 是 1。
- **复杂度与边界**：设 L 为全部字符串总长度、w 为 word 数量、q 为查询数量，时间 `O(L+w log(w+1)+q log(w+1))`；除输出外额外空间 `O(w)`（固定字符计数及排序栈不超过此量级）。全部相等、全部大于、全部不大于都要检查。

<a id="lc1608"></a>

### [1608. 特殊数组的特征值](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/)

- **题意**：找非负整数 x，使数组中恰好有 x 个元素 `>=x`；不存在返回 -1。例如 `[3,5]` 的答案为 2，`[0,0]` 无解。
- **内层 check**：排序后，对固定 x 定义 `checkIndex(i)=nums[i]>=x`，通过 `count(x)=n-firstTrue(n,checkIndex)` 计算元素中 `>=x` 的数量。
- **外层 check**：定义 `check(x)=count(x)<=x`，为 false→true；候选值为 `[0,n]`。
- **firstTrue 与答案**：`x=firstTrue(n+1,check)`；仅当 `count(x)==x` 时返回 x，否则返回 -1。因为 count(n)<=n，外层必有 true，不会返回 n+1；分界点不一定就是零点。
- **解释**：x 增加时 count 不会增加，因此 `count(x)-x` 严格递减，最多有一个零点；但零点可能不存在，不能把普通分界点直接当答案。
- **复杂度与边界**：排序 `O(n log(n+1))`，双层二分 `O(log²(n+1))`；原地排序的额外栈空间通常 `O(log n)`。全零、重复值、答案恰为 n 都需覆盖。

<a id="lc1855"></a>

### [1855. 下标对中的最大距离](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/)

- **题意**：两个数组均非递增，寻找满足 `i<=j` 且 `nums1[i]<=nums2[j]` 的最大 `j-i`，无有效对返回 0。例如 `[55,30,5,4,2]` 与 `[100,20,10,10,5]`，答案 2。
- **check**：枚举 i，定义 `check(j)=nums2[j]<nums1[i]`，表示“不满足大小条件”。nums2 非递增，所以 check 为 false→true。
- **firstTrue 与答案**：`p=firstTrue(n,check)`，最后一个满足位置为 `j=p-1`；仅在 `j>=i` 时用 `j-i` 更新最大距离，答案初值为 0。寻找最后一个满足，等价于寻找第一个不满足再减一。
- **解释**：p=0 表示所有元素都太小，p=n 表示全部满足大小条件；最终仍须检查 `j>=i`，落实下标约束。i>=n 时不存在合法 j，可直接跳过。
- **复杂度与边界**：二分版本时间 `O(m log(n+1))`，额外空间 `O(1)`。两数组长度不同、重复值、无有效对需覆盖。也可利用双指针优化到 `O(m+n)`。

<a id="lc2089"></a>

### [2089. 找出数组排序后的目标下标](https://leetcode.com/problems/find-target-indices-after-sorting-array/)

- **题意**：数组排序后返回所有值等于 target 的下标。例如 `[1,2,5,2,3]、target=2`，答案 `[1,2]`；不是返回原数组的下标。
- **check**：先排序，定义 `checkLower(i)=nums[i]>=target` 和 `checkUpper(i)=nums[i]>target`，均为 false→true。
- **firstTrue 与答案**：`lower=firstTrue(n,checkLower)`，`upper=firstTrue(n,checkUpper)`，输出 lower 到 upper-1。与 [3.1 节](#duplicates) 相同，本题需先排序并返回整个下标列表。
- **解释**：有序数组中重复目标连续出现，两个边界之间恰好全部等于 target。不存在时 lower=upper，输出空数组。
- **复杂度与边界**：时间 `O(n log(n+1)+k)`，k 为输出长度；输出空间 `O(k)`，另有排序栈空间。全部等于目标、目标不在范围内都要覆盖。不要求二分时，也可直接计数小于和等于 target 的元素做到 `O(n)`。

<a id="lc729"></a>

### [729. 我的日程安排表 I](https://leetcode.com/problems/my-calendar-i/)

- **题意**：依次预订时间段 `[start,end)`，只有与已成功预订的时间段不重叠时才插入并返回 true。例如 `[10,20)` 成功，随后 `[15,25)` 失败、`[20,30)` 成功。
- **check**：按 start 维护有序且互不重叠的区间，定义 `check(j)=bookings[j].start>=start`，为 false→true。
- **firstTrue 与答案**：`i=firstTrue(n,check)` 得到插入位置。若前一段结束时间 `>start` 或后一段开始时间 `<end` 则返回 false，否则插入并返回 true。
- **边界含义**：i=0 表示在最前插入，i=n 表示在末尾插入；访问邻居前分别检查 i>0 和 i<n。
- **解释**：已有区间互不重叠，非相邻区间不可能绕过邻居与新区间重叠，因此只需检查前后两段。查找使用闭区间，但日程本身仍为题设的左闭右开时间段，相邻端点相等允许预订。
- **复杂度与边界**：每次查找 `O(log(n+1))`，切片插入最坏 `O(n)`，存储空间 `O(n)`。失败预订不能改变状态；同起点、包含关系、乱序插入和边界相接都要测试。

<a id="lc888"></a>

### [888. 公平的糖果交换](https://leetcode.com/problems/fair-candy-swap/)

- **题意**：Alice 和 Bob 各交换一袋糖果，使双方总数相同，返回交换的两袋大小；题目保证存在答案，多个答案任取其一。例如 `[1,1]` 与 `[2,2]` 返回 `[1,2]`。
- **check**：令总数为 A、B，交换量为 a、b，由 `A-a+b=B-b+a` 得 `b=a-(A-B)/2`。对 Bob 的数组排序；枚举每个 a，定义 `check(i)=bobSizes[i]>=b`，为 false→true。
- **firstTrue 与答案**：`p=firstTrue(len(bobSizes),check)`；若 `p<len(bobSizes) && bobSizes[p]==b`，返回 `[a,bobSizes[p]]`，否则继续枚举。题设保证存在答案；防御性处理可在 A-B 为奇数或无解时返回空结果。
- **解释**：代数转换把枚举两袋的 `O(mn)` 降为枚举一袋并搜索另一袋。只需排序 Bob；总数相等时也必须交换两袋相同大小的糖果，不能直接返回空结果。
- **复杂度与边界**：时间 `O(n log(n+1)+m log(n+1))`，额外空间为排序栈 `O(log n)`。总和用 int64 避免溢出；测试应验证交换后总数相等，不固定唯一答案。
