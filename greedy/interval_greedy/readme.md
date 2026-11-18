1	[56](https://leetcode.com/problems/merge-intervals/). 合并区间	按起点排序，合并重叠区间
2	[57](https://leetcode.com/problems/insert-interval/). 插入区间	分成“左侧、重叠部分、右侧”处理
3	[986](https://leetcode.com/problems/interval-list-intersections/). 区间列表的交集	双指针，计算两个区间的交集
4	[435](https://leetcode.com/problems/non-overlapping-intervals/). 无重叠区间	按终点排序，优先保留结束早的区间
5	[452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/). 用最少数量的箭引爆气球	区间贪心，在合适的公共位置射箭
6	[1094](https://leetcode.com/problems/car-pooling/). 拼车	差分／扫描线，维护同时在车上的人数
7	[2406](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/). 将区间分为最少组数	最大同时重叠数量；可用扫描线或最小堆
8	[218](https://leetcode.com/problems/the-skyline-problem/). 天际线问题	扫描线＋最大堆，维护当前有效建筑的最高高度
# 区间贪心：按结束位置排序

原 `greedy/l435`、`greedy/l452` 合并到本目录。两个问题都按结束位置升序扫描，保留当前选择的最右端点，差别在于相接端点是否需要开始一个新选择。

```text
按 end 升序排序
第一次选择初始化 end
依次扫描区间 [start,finish]：
    若与当前选择满足“分离”条件：
        选择次数加 1
        end = finish
```

| 问题 | 分离条件 | 返回值 | 证明要点 |
| --- | --- | --- | --- |
| [435. 无重叠区间](https://leetcode.com/problems/non-overlapping-intervals/) | `start >= end` | 总数减最多保留数 | 更早结束给未来留下最多空间 |
| [452. 最少箭数](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) | `start > end` | 选择次数 | 箭放在最早右端点，可以覆盖所有包含该点的气球 |

不变量：已经处理的区间都已合法安排，当前端点在同样选择数下尽可能靠左。空输入应单独返回 0；排序可能修改输入。

原有区间题资料保留如下。

---



## 本页可复用模板：按结束位置选择区间

先复制以下 helper，再复制任一题适配器。实际题解在每个文件内嵌同名 helper，可独立提交，不导入仓库本地算法包。不同移动/贪心依据使用不同命名模板，不把“用了两个变量”当作同一不变量。

```go
import (
	"sort"
)

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
```

| 题目 | 采用的局部模板 | 实际实现 |
| --- | --- | --- |
| [435](https://leetcode.com/problems/non-overlapping-intervals/) | `SelectByEnd` | [实际 Go 文件](l435/435.go) |
| [452](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) | `SelectByEnd` | [实际 Go 文件](l452/452.go) |

为使整页 Go 示例可合并编译，重名适配器在文档中加题号/DP/Greedy 后缀；实际 Go 文件仍保留官方函数签名。链表节点等支撑类型仅在 helper 区声明一次。


## 其余题目的模板适配器

以下代码调用本页已经声明的 helper；复制到对应题解时使用实际 Go 文件中的官方函数名。

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

