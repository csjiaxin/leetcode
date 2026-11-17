# 燃料、余额与可达距离

原 `greedy/l134`、`greedy/l871` 合并到本目录同题的 `variants/practice`，保留空函数和测试；已经完成的实现继续位于题目主目录。

## 134：失败区间整体排除

```go
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
