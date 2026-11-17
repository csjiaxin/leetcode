# 所有点对最短路：Floyd-Warshall

[返回图算法目录](../README.md)

适用于节点数量允许 O(n³) 的所有点对路径问题，可以含负边。存在负环时，受影响的路径没有有限最短值。

**状态来源**：把中间节点是否允许使用作为阶段。第 k 阶段完成后，dist[i][j] 表示只允许 0..k 作为中间节点的最短距离。因此 k 必须是最外层循环。

```go
// 初始化：自身为 0；直接边取最小；不可达用独立 reachable 标记。
for k := 0; k < n; k++ {
    for i := 0; i < n; i++ {
        if !reachable[i][k] { continue }
        for j := 0; j < n; j++ {
            if !reachable[k][j] { continue }
            candidate := dist[i][k] + dist[k][j] // 相加前检查溢出
            if !reachable[i][j] || candidate < dist[i][j] {
                dist[i][j], reachable[i][j] = candidate, true
            }
        }
    }
}
```

需要修改的部分是建图、代价合成与最终读取答案；不能随意交换三层循环。时间 O(n³)，空间 O(n²)。完成后若 dist[v][v]<0，说明存在负环。

| 文件／题目 | 状态与现状 |
| --- | --- |
| [floyd-warshall.go](floyd-warshall.go) | 通用实现原样保留，原文件缺少 Edge／checkedAdd 依赖，尚不能独立编译 |
| [1334. 阈值距离内邻居最少的城市](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) · [原始片段](l1334/1334) | 求点对距离后数不超过阈值的邻居，再处理编号优先级 |
| [399. 除法求值](https://leetcode.com/problems/evaluate-division/) · [未完成草稿](l399/399.go) | 路径合成是乘法、反向边取倒数；原草稿有语法缺口，单独成包保留 |

下面保留原始资料链接：

https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
