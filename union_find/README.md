# 并查集：集合合并与无向连通性

[返回总目录](../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice) · [保留的原分类笔记](#legacy-notes)

## 目录

- [1. 核心思路](#core)
- [2. 547：省份数量](#lc547)
- [3. 684：冗余连接](#lc684)
- [4. 1319：连接网络](#lc1319)
- [5. 721：账户合并](#lc721)
- [6. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

Find得到集合代表，Union合并不同代表。parent[root]=root；size仅在根有意义。路径压缩＋按大小合并的摊还操作复杂度为O(α(n))。

```text
初始化每个节点独立 → 找两个根 → 根相同则无需合并
根不同：小树挂大树 → 分量数减1
```

下面每题含自己的find/union，方便单独复制。它们是必要的集合操作，不再包装题目无关回调。只解决无向连通或等价关系，不支持一般有向可达、在线删边和最短路。

<a id="lc547"></a>

## 2. [547. 省份数量](https://leetcode.com/problems/number-of-provinces/)

[练习](l547/547.go) · [测试](l547/547_test.go)

**题意**：城市邻接矩阵中统计无向连通分量。

**状态**：components初始n，只有不同集合合并时减1。遍历上三角避免重复处理。

```go
package template

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	find := func(x int) int {
		for x != parent[x] {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	components := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			a, b := find(i), find(j)
			if a == b {
				continue
			}
			if size[a] < size[b] {
				a, b = b, a
			}
			parent[b] = a
			size[a] += size[b]
			components--
		}
	}
	return components
}
```

**例子**：[[1,1,0],[1,1,0],[0,0,1]]，前两点合并，答案2。

**复杂度**：保守时间O(n²α(n))，扫描矩阵本身O(n²)，空间O(n)。α很小但不等于严格常数。

<a id="lc684"></a>

## 3. [684. 冗余连接](https://leetcode.com/problems/redundant-connection/)

[练习](l684/684.go) · [测试](l684/684_test.go)

**题意**：树多加一条边，返回能删除且输入中最靠后的那条环上边。

**变化**：按输入顺序加入；第一次发现端点已连通的边，就是环上最后出现的边。

```go
package template

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent, size := make([]int, n+1), make([]int, n+1)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	find := func(x int) int {
		for x != parent[x] {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, e := range edges {
		a, b := find(e[0]), find(e[1])
		if a == b {
			return []int{e[0], e[1]}
		}
		if size[a] < size[b] {
			a, b = b, a
		}
		parent[b] = a
		size[a] += size[b]
	}
	return nil
}
```

**例子**：[[1,2],[1,3],[2,3]]返回[2,3]。题目编号1..n，数组需n+1。

**复杂度**：时间O(nα(n))，空间O(n)。不能用这一逻辑判断一般有向环。

<a id="lc1319"></a>

## 4. [1319. 连通网络的操作次数](https://leetcode.com/problems/number-of-operations-to-make-network-connected/)

[练习](l1319/1319.go) · [测试](l1319/1319_test.go)

**题意**：可以重新连接冗余网线，求使n台计算机连通的最少操作。

**推导**：线少于n-1必无解；否则冗余线足以连接所有分量，答案components-1。

```go
package template

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	find := func(x int) int {
		for x != parent[x] {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	components := n
	for _, e := range connections {
		a, b := find(e[0]), find(e[1])
		if a == b {
			continue
		}
		if size[a] < size[b] {
			a, b = b, a
		}
		parent[b] = a
		size[a] += size[b]
		components--
	}
	return components - 1
}
```

**例子**：n=4、[[0,1],[0,2],[1,2]]，一个冗余边可连接孤立3，答案1。

**复杂度**：时间O(n+mα(n))，空间O(n)。n=1且无线缆答案0。

<a id="lc721"></a>

## 5. [721. 账户合并](https://leetcode.com/problems/accounts-merge/)

[练习](l721/721.go) · [测试](l721/721_test.go)

**题意**：共享邮箱的账户属于同一人，合并后邮箱排序；同名但不共享邮箱不能直接合并。

**建模**：账户下标为DSU节点。owner记录邮箱首次出现的账户；再次出现时合并两个账户。最后把每个唯一邮箱归入最终根。

```go
package template

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	find := func(x int) int {
		for x != parent[x] {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(a, b int) {
		a, b = find(a), find(b)
		if a == b {
			return
		}
		if size[a] < size[b] {
			a, b = b, a
		}
		parent[b] = a
		size[a] += size[b]
	}
	owner := make(map[string]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			if j, ok := owner[email]; ok {
				union(i, j)
			} else {
				owner[email] = i
			}
		}
	}
	groups := make(map[int][]string)
	for email, i := range owner {
		root := find(i)
		groups[root] = append(groups[root], email)
	}
	answer := [][]string{}
	for root, emails := range groups {
		sort.Strings(emails)
		group := []string{accounts[root][0]}
		answer = append(answer, append(group, emails...))
	}
	return answer
}
```

**例子**：[A,a,b]与[A,b,c]共享b，合为[A,a,b,c]；另一条[A,d]保持独立。重复邮箱在owner中只出现一次。

**复杂度**：A个账户、M个邮箱出现次数、U个唯一邮箱，忽略字符串比较长度时平均时间O(A+Mα(A)+U log(U+1))，空间O(A+U)。字符串哈希和比较另随邮箱长度变化。返回组间顺序任意，组内邮箱有序。

<a id="pitfalls"></a>

## 6. 对比与易错点

| 题目 | 节点 | 合并失败意味着什么 | 答案 |
| --- | --- | --- | --- |
| 547 | 城市 | 已连通 | 分量数 |
| 684 | 无向图节点 | 新边闭合环 | 该边 |
| 1319 | 计算机 | 存在冗余网线 | 线足够时分量数-1 |
| 721 | 账户 | 本来同一人 | 按根归并邮箱 |

- 必须合并根，不是把任意parent[a]指向b。
- 只有根不同时才减分量数，重复边不能重复减。
- 根只是代表元，不保证编号最小；姓名和业务信息须按题意保存。
- DSU不能提供原图路径，也不能替代 [有向判环](../graph/topo_sort/readme.md)。

已有分类：[union_find](.)。


<a id="legacy-notes"></a>

## 原分类笔记（保留）

https://leetcode.com/discuss/general-discussion/1072418/Disjoint-Set-Union-(DSU)Union-Find-A-Complete-Guide
https://leetcode.com/list/5lhmb4mj/


924. Minimize Malware Spread
https://leetcode.com/problems/minimize-malware-spread/discuss/181129/C%2B%2BPython-Union-Found

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [128. 最长连续序列](https://leetcode.com/problems/longest-consecutive-sequence/) | 数字映射节点，与相邻数字合并，维护集合大小 | 值需去重；哈希起点扫描通常更简单 |
| [990. 等式方程的可满足性](https://leetcode.com/problems/satisfiability-of-equality-equations/) | 先合并所有等式，再检查不等式两端不能同根 | 不能按输入顺序边合并边判所有不等式 |
| [947. 移除最多的同行或同列石头](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/) | 共享行列形成连通分量，每分量最终留一颗 | 行和列ID需分离；答案石头数-石头连通分量数 |
| [1202. 交换字符串中的元素](https://leetcode.com/problems/smallest-string-with-swaps/) | 可交换下标合并，分量内部字符排序后分配给升序下标 | 排序是每个连通分量内部，不是全字符串 |
| [959. 由斜杠划分区域](https://leetcode.com/problems/regions-cut-by-slashes/) | 每格拆四个三角形，按斜杠方向与邻格边界合并 | 空格连接四块，正反斜杠连接方式不同 |
| [1697. 检查边长度限制的路径是否存在](https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/) | 边和查询按阈值排序，离线增量合并满足条件的边 | 严格小于limit，不是小于等于；恢复查询原顺序 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
