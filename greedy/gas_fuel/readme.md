# 燃料、余额与可达距离

## 本页可复用模板：燃料与可达距离

先复制以下 helper，再复制任一题适配器。实际题解在每个文件内嵌同名 helper，可独立提交，不导入仓库本地算法包。不同移动/贪心依据使用不同命名模板，不把“用了两个变量”当作同一不变量。

```go
import (
	"container/heap"
)

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
| 134 | `CircuitStart` | [实际 Go 文件](l134/134.go) |
| 871：按加油次数 DP | `RefuelByCount` | [实际 Go 文件](l871/871.go) |
| 871：最大堆贪心 | `FuelHeap + GreedyRefuel` | [实际 Go 文件](l871/variants/practice/871.go) |

为使整页 Go 示例可合并编译，重名适配器在文档中加题号/DP/Greedy 后缀；实际 Go 文件仍保留官方函数签名。链表节点等支撑类型仅在 helper 区声明一次。

原 `greedy/l134`、`greedy/l871` 合并到本目录同题的 `variants/practice`，保留独立来源及原测试；实际 Go 文件现已按本页命名模板实现。

## 134：失败区间整体排除

```text
start, balance, total := 0, 0, 0
for i := range gas {
    change := gas[i] - cost[i]
    balance += change
    total += change
    if balance < 0 {
        start = i + 1
        balance = 0
    }
}
// total < 0 则无解，否则返回 start。
```

不变量：在当前 start 到当前位置之间尚未遇到负余额；首次负余额证明这整个候选起点区间均不可行。余额等于 0 时不用排除。

## 871：区分两个合法的状态模型

| 方法 | 状态与转移 | 实现 |
| --- | --- | --- |
| 按加油次数 DP | `reach[j]` 表示加油 j 次最远能到哪里；每个站从大到小更新 j | [本目录的历史 DP 解法](l871/871.go) |
| 最大堆贪心 | 把已经可达站点加入最大堆，需要扩展距离时弹出最大油量 | [堆系列实现](../../heap_priority_queue/l871/871.go) |

最大堆版本的骨架：

```text
while reach < target:
    将 position <= reach 的未处理站点全部入堆
    若堆为空则不可达
    reach += 堆中最大油量
    stops++
```

这是延迟选择；堆内只能放已经可达的站点。两个方法的状态不同，所以保留各自转移，而不把按次数 DP 误写成局部贪心。

原题：[134. Gas Station](https://leetcode.com/problems/gas-station/)、[871. Minimum Number of Refueling Stops](https://leetcode.com/problems/minimum-number-of-refueling-stops/)。


## 其余题目的模板适配器

以下代码调用本页已经声明的 helper；复制到对应题解时使用实际 Go 文件中的官方函数名。

### 134

```go
func canCompleteCircuit(gas []int, cost []int) int { return CircuitStart(gas, cost) }
```

### 871：按加油次数 DP

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

### 871：最大堆贪心

```go
func minRefuelStopsGreedy(target int, startFuel int, stations [][]int) int {
	return GreedyRefuel(target, startFuel, stations)
}
```
