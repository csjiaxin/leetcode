# 所有点对最短路：Floyd-Warshall

[返回图算法目录](../README.md)

适用于节点数量允许 O(n³) 的所有点对路径问题，可以含负边。存在负环时，受影响的路径没有有限最短值。

**状态来源**：把中间节点是否允许使用作为阶段。第 k 阶段完成后，dist[i][j] 表示只允许 0..k 作为中间节点的最短距离。因此 k 必须是最外层循环。

固定三层循环可以抽成下面的 `FloydClosure`。调用方初始化方阵 value 与同尺寸 reachable，再提供路径合成函数 extend 和严格改善函数 better。extend 返回false表示无法表示合成结果，helper 立即停止；调用方决定如何报告错误。泛型需要Go 1.18及以上。

```go
package template

// FloydClosure 按允许的中间节点递增；extend 失败时停止并返回 false。
func FloydClosure[T any](value [][]T, reachable [][]bool, extend func(T, T) (T, bool), better func(T, T) bool) bool {
	for k := range value {
		for i := range value {
			if !reachable[i][k] {
				continue
			}
			for j := range value {
				if !reachable[k][j] {
					continue
				}
				candidate, ok := extend(value[i][k], value[k][j])
				if !ok {
					return false
				}
				if !reachable[i][j] || better(candidate, value[i][j]) {
					value[i][j] = candidate
					reachable[i][j] = true
				}
			}
		}
	}
	return true
}
```

需要修改的部分是建图、代价合成与最终读取答案；不能随意交换三层循环。时间 O(n³)，空间 O(n²)。对加法最短路，完成后若 dist[v][v]<0，说明存在负环。泛型helper只固定中间节点阶段，具体合成与比较规则仍需证明符合题目的路径语义，不能把任何回调都当成最短路。

| 文件／题目 | 状态与现状 |
| --- | --- |
| [floyd-warshall.go](floyd-warshall.go) | 已补齐 Edge／checkedAdd，FloydWarshall 调用同文件 FloydClosure，检查溢出和负环，可独立成包编译 |
| [1334. 阈值距离内邻居最少的城市](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) · [原始片段](l1334/1334) | 求点对距离后数不超过阈值的邻居，再处理编号优先级 |
| [399. 除法求值](https://leetcode.com/problems/evaluate-division/) · [实现](l399/399.go) · [原混合草稿归档](l399/legacy/399_mixed_original) | 已实现清晰的 calcEquation API；路径合成用乘法、反向边取倒数，同文件复制完整 FloydClosure，可独立提交 |

## 加法最短路：初始化、合成与负环检查

把公共helper与下方适配代码放在同一个Go文件中，只保留一次 `package`，并在所有声明前加入 `import "fmt"`。Edge、checkedAdd与FloydClosure各保留一份。实际通用文件已经包含完整实现。

初始化自身距离0且可达，直接边取较小权重；不可达用独立bool表示，不拿合法距离充当INF。只有两段路径都可达时才合成，checkedAdd在相加前检查int64溢出。负边可以处理；检测到负环时，该适配函数返回错误，不把受影响距离当作有限最短路。

```go
type Edge struct {
	From, To int
	Weight   int64
}

// FloydWarshall 返回有向图全源最短距离及可达性；边端点须在 [0,n) 内。
func FloydWarshall(n int, edges []Edge) ([][]int64, [][]bool, error) {
	if n < 0 {
		return nil, nil, fmt.Errorf("n must be nonnegative")
	}
	dist := make([][]int64, n)
	reachable := make([][]bool, n)
	for i := range dist {
		dist[i] = make([]int64, n)
		reachable[i] = make([]bool, n)
		reachable[i][i] = true
	}
	for _, edge := range edges {
		u, v := edge.From, edge.To
		if u < 0 || u >= n || v < 0 || v >= n {
			return nil, nil, fmt.Errorf("edge endpoint out of range")
		}
		if !reachable[u][v] || edge.Weight < dist[u][v] {
			dist[u][v] = edge.Weight
			reachable[u][v] = true
		}
	}
	if !FloydClosure(dist, reachable, checkedAdd, func(a, b int64) bool { return a < b }) {
		return nil, nil, fmt.Errorf("distance overflow")
	}
	for i := range dist {
		if dist[i][i] < 0 {
			return nil, nil, fmt.Errorf("negative cycle detected")
		}
	}
	return dist, reachable, nil
}

func checkedAdd(a, b int64) (int64, bool) {
	const highest = int64(1<<63 - 1)
	const lowest = int64(-1 << 63)
	if b > 0 && a > highest-b || b < 0 && a < lowest-b {
		return 0, false
	}
	return a + b, true
}
```

## 399：一致比值关系的适配

原题保证给定除法关系一致、值为正。把变量映射成节点，a/b=w对应正向w与反向1/w；已知变量的自身比值为1。路径合成改为乘法。因为一致关系下所有可达路径给出相同比值，better固定返回false，只补充尚未知晓的关系；这不表示在求最大乘积或最短距离，不能直接套到矛盾关系上。

下面只展示适配函数，与同一份FloydClosure合并即可，无需额外导入包。实际399主文件已包含helper和演示main；原来混入城市草稿的内容保存在[原混合草稿归档](l399/legacy/399_mixed_original)。

```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ids := make(map[string]int)
	for _, equation := range equations {
		for _, name := range equation {
			if _, ok := ids[name]; !ok {
				ids[name] = len(ids)
			}
		}
	}
	n := len(ids)
	ratios := make([][]float64, n)
	reachable := make([][]bool, n)
	for i := range ratios {
		ratios[i] = make([]float64, n)
		reachable[i] = make([]bool, n)
		ratios[i][i], reachable[i][i] = 1, true
	}
	for i, equation := range equations {
		a, b := ids[equation[0]], ids[equation[1]]
		ratios[a][b], ratios[b][a] = values[i], 1/values[i]
		reachable[a][b], reachable[b][a] = true, true
	}
	// 原题关系一致，已有可达比值无需重新优化，只补出未知的关系。
	FloydClosure(ratios, reachable, func(a, b float64) (float64, bool) { return a * b, true }, func(a, b float64) bool { return false })
	answer := make([]float64, len(queries))
	for i, query := range queries {
		a, aOK := ids[query[0]]
		b, bOK := ids[query[1]]
		answer[i] = -1
		if aOK && bOK && reachable[a][b] {
			answer[i] = ratios[a][b]
		}
	}
	return answer
}
```

未知变量的查询返回-1，即使查询是未知的x/x；已知的a/a返回1。设不同变量数为V，构图后闭包时间O(V³)、矩阵空间O(V²)，每次查询为均摊O(1)。

下面保留原始资料链接：

https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
