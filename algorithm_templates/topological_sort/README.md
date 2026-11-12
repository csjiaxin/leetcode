# 拓扑排序：依赖顺序与DAG动态规划

[返回总目录](../README.md)

## 目录

- [1. 核心思路](#core)
- [2. 207：能否完成课程](#lc207)
- [3. 210：返回课程顺序](#lc210)
- [4. 2050：并行课程耗时](#lc2050)
- [5. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

边u→v表示u先于v。Kahn算法维护剩余入度，只有前置依赖全部完成的节点才能入队。

```text
建邻接表、入度 → 全部零入度节点入队 → 弹出u → 删除出边
v的入度降至0时入队 → 最后处理数量<n说明有环
```

孤立节点也要处理。入度和邻接表必须一致：重复边若保留就两处都计入，若去重就两处都去重。每题实现独立完整。

<a id="lc207"></a>

## 2. [207. 课程表](https://leetcode.com/problems/course-schedule/)

[练习](l207/207.go) · [测试](l207/207_test.go)

**题意**：判断能否完成所有课程。输入[a,b]表示先修b再修a。

**变化点**：建边b→a，只需计处理数，不必保存顺序。

```go
package template

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indegree[p[0]]++
	}
	queue := []int{}
	for i, d := range indegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(queue) == numCourses
}
```

**例子**：2门课、[[1,0],[0,1]]构成环，无法启动，返回false。若环外有独立节点，也仍须最终判数。

**复杂度**：时间O(V+E)，空间O(V+E)。

<a id="lc210"></a>

## 3. [210. 课程表 II](https://leetcode.com/problems/course-schedule-ii/)

[练习](l210/210.go) · [测试](l210/210_test.go)

**题意**：返回任意合法课程顺序，有环时返回空数组。

**变化点**：与207相同建图，队列的出队顺序就是拓扑序；本实现队列数组最终即该顺序。

```go
package template

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indegree[p[0]]++
	}
	order := []int{}
	for i, d := range indegree {
		if d == 0 {
			order = append(order, i)
		}
	}
	for head := 0; head < len(order); head++ {
		u := order[head]
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				order = append(order, v)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}
```

**例子**：[[1,0],[2,0],[3,1],[3,2]]可返回[0,1,2,3]或[0,2,1,3]。不要排序最终答案，那可能破坏依赖。

**复杂度**：时间O(V+E)，空间O(V+E)。若要求字典序最小，用最小堆选择当前零入度节点。

<a id="lc2050"></a>

## 4. [2050. 并行课程 III](https://leetcode.com/problems/parallel-courses-iii/)

[练习](l2050/2050.go) · [测试](l2050/2050_test.go)

**题意**：课程可无限并行，每门有时长，所有前置完成才能开始，求全部完成的最短时间。relations=[前置,后继]且编号1-based。

**状态**：finish[v]表示最早完成时间，初始time[v]。每条边u→v更新finish[v]=max(finish[v],finish[u]+time[v])，入度归零后v才入队。

```go
package template

func minimumTime(n int, relations [][]int, time []int) int {
	graph := make([][]int, n)
	indegree := make([]int, n)
	finish := make([]int64, n)
	for i, t := range time {
		finish[i] = int64(t)
	}
	for _, e := range relations {
		u, v := e[0]-1, e[1]-1
		graph[u] = append(graph[u], v)
		indegree[v]++
	}
	queue := []int{}
	for i, d := range indegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	var answer int64
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		answer = max(answer, finish[u])
		for _, v := range graph[u] {
			finish[v] = max(finish[v], finish[u]+int64(time[v]))
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 题设保证DAG；防御性检查。
	if len(queue) != n {
		return -1
	}
	return int(answer)
}
```

**例子**：1和2都先于3，时长[3,2,5]，两前置可并行，结果max(3,2)+5=8，不是10。

**复杂度**：时间O(V+E)，空间O(V+E)。不是BFS层数；不同课程时长必须用DAG最长路径。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 边方向 | 汇总 |
| --- | --- | --- |
| 207 | prerequisites[1]→[0] | 处理数是否等于n |
| 210 | prerequisites[1]→[0] | 拓扑序 |
| 2050 | relations[0]-1→[1]-1 | 最早完成时间最大值 |

- 所有零入度节点一起初始化，不能遗漏孤立点。
- 有环时可能先处理无环部分，因此必须最后判数量。
- 仅在入度降为0时入队，不能每删一条边就入队。
- DSU只处理集合关系，不能替代有向判环。

已有分类：[graph/topo_sort](../../graph/topo_sort/)。
