# 差分＋前缀扫描：批量区间变化

[返回总目录](../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice)

## 目录

- [1. 核心思路](#core)
- [2. 1094：拼车](#lc1094)
- [3. 1109：航班预订](#lc1109)
- [4. 1854：人口最多年份](#lc1854)
- [5. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

整段增加delta，只记录开始生效与停止生效两个事件，最后前缀扫描恢复当前总量。

| 区间 | 起点 | 取消位置 |
| --- | --- | --- |
| 整数闭区间[l,r] | diff[l]+=delta | diff[r+1]-=delta |
| 半开区间[l,r) | diff[l]+=delta | diff[r]-=delta |

**不变量**：处理完坐标x的全部净变化后，前缀累计值就是x处的有效总量。同坐标加减应先聚合，不凭输入事件顺序判断超载。

适合批量离线更新。以下分别给出各题实现，不把闭区间与半开区间混在同一函数里。

<a id="lc1094"></a>

## 2. [1094. 拼车](https://leetcode.com/problems/car-pooling/)

[练习](l1094/1094.go) · [测试](l1094/1094_test.go)

**题意**：每条行程[人数,上车站,下车站]，判断任意路段是否超出capacity。

**端点**：乘客占用[from,to)，在to已下车，可给同站上车的人腾位，所以to直接减，不是to+1。

```go
package template

func carPooling(trips [][]int, capacity int) bool {
	maxPosition := 0
	for _, t := range trips {
		maxPosition = max(maxPosition, t[2])
	}
	diff := make([]int64, maxPosition+1)
	for _, t := range trips {
		diff[t[1]] += int64(t[0])
		diff[t[2]] -= int64(t[0])
	}
	var passengers int64
	for _, change := range diff {
		passengers += change
		if passengers > int64(capacity) {
			return false
		}
	}
	return true
}
```

**例子**：[[2,1,5],[3,5,7]]、容量3，位置5净变化+1，载荷从2变3，返回true，不是临时5人。

**复杂度**：m条行程、最大坐标U，时间O(m+U)，空间O(U)。

<a id="lc1109"></a>

## 3. [1109. 航班预订统计](https://leetcode.com/problems/corporate-flight-bookings/)

[练习](l1109/1109.go) · [测试](l1109/1109_test.go)

**题意**：[first,last,seats]给编号first到last的每个航班预订seats，返回每航班总座位。

**端点**：题目1-based闭区间，转0-based后[first-1,last-1]，在last位置取消。

```go
package template

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int64, n+1)
	for _, b := range bookings {
		diff[b[0]-1] += int64(b[2])
		diff[b[1]] -= int64(b[2])
	}
	answer := make([]int, n)
	var seats int64
	for i := 0; i < n; i++ {
		seats += diff[i]
		answer[i] = int(seats)
	}
	return answer
}
```

**例子**：[[1,2,2],[2,3,3]]、n=3，答案[2,5,3]。航班2同时包含在两段中。

**复杂度**：时间O(m+n)，空间O(n)。diff额外开一个哨兵，允许末航班取消事件写入diff[n]。

<a id="lc1854"></a>

## 4. [1854. 人口最多的年份](https://leetcode.com/problems/maximum-population-year/)

[练习](l1854/1854.go) · [测试](l1854/1854_test.go)

**题意**：生年计入、死亡年份不计入；求人口最多的最早年份。

**端点与汇总**：区间[birth,death)，birth加1、death减1。按年升序扫描，只有严格超过历史最大值才更新年份，保留最早并列值。

```go
package template

func maximumPopulation(logs [][]int) int {
	// 题设年份位于1950到2050之间。
	var diff [2051]int
	for _, log := range logs {
		diff[log[0]]++
		diff[log[1]]--
	}
	population, best, answer := 0, 0, 1950
	for year := 1950; year <= 2050; year++ {
		population += diff[year]
		if population > best {
			best = population
			answer = year
		}
	}
	return answer
}
```

**例子**：[[1950,1960],[1960,1970]]，1960一出一进，人口始终1，答案1950。

**复杂度**：时间O(m+U)，空间O(U)，本题U为固定101个年份，额外空间可视为O(1)。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 取消点 | 扫描时做什么 |
| --- | --- | --- |
| 1094 | to | 判断容量 |
| 1109 | 0-based的right+1，即原last | 保存全部航班总量 |
| 1854 | death | 更新最大人口最早年份 |

- 固定年份数组只用于1854题设，不是任意日期的通用上限。
- 稀疏巨大坐标不要分配U数组：用map聚合同坐标变化，排序扫描，时间O(m log(m+1))、空间O(m)。
- 坐标压缩保留顺序；算覆盖长度还需乘原坐标间距，不能直接用压缩下标差。
- 实数闭区间不能机械r+1，必须明确同坐标事件顺序。
- 差分不是通用在线查询结构；动态区间更新和查询可看 [线段树](../tree/segment)。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [2381. 字母移位II](https://leetcode.com/problems/shifting-letters-ii/) | 闭区间增减字母位移，差分累计后对26规范化 | 负余数要加26；取消点right+1 |
| [3355. 零数组变换I](https://leetcode.com/problems/zero-array-transformation-i/) | 每条查询覆盖处可独立减0或1，差分计算覆盖能力 | 比较覆盖数>=nums[i]，不是每区间强制全部减1 |
| [1893. 检查是否区域内所有整数都被覆盖](https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/) | 整数闭区间差分，目标区间每点覆盖数至少1 | 取消点r+1，不要只比较总覆盖长度 |
| [2406. 将区间分为最少组数](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/) | 整数闭区间最大并发覆盖即最少组数 | 端点相接算重叠；大坐标用稀疏事件或最小堆 |
| [2536. 子矩阵元素加1](https://leetcode.com/problems/increment-submatrices-by-one/) | 二维矩形四角差分，再二维前缀恢复 | 额外哨兵行列；四角符号+ - - + |
| [732. 我的日程安排表III](https://leetcode.com/problems/my-calendar-iii/) | 半开预订记录端点变化，每次返回最大重叠 | 这是在线更新；简单有序事件重扫可行但非最优，规模大考虑动态线段树 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
