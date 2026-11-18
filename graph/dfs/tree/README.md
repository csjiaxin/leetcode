# 树 DFS：子树返回值与后序汇总

[返回 DFS 目录](../README.md)

适用条件是输入确实为树。无向树中跳过父节点就能避免返回；一般图需要另外的 visited。

**不变量**：dfs 返回时，该节点的整棵子树已经处理完成；返回值只代表父节点需要的信息。

```go
// TreeFold 用于无向树；parent 防止沿父边回走，不适用于任意有环图。
func TreeFold[S any](graph [][]int, node, parent int, start func(int) S, merge func(S, S) S, finish func(int, S)) S {
	state := start(node)
	for _, child := range graph[node] {
		if child != parent {
			state = merge(state, TreeFold(graph, child, node, start, merge, finish))
		}
	}
	if finish != nil {
		finish(node, state)
	}
	return state
}
```

例如子树大小只改变初始化和合并。模板与适配组合使用：

```go
func subtreeSize(graph [][]int, root int) int {
    return TreeFold(graph, root, -1,
        func(int) int { return 1 },
        func(current, child int) int { return current + child },
        nil,
    )
}
```

[1519](https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/) 的标签频次数组适配（实际主文件已经内联同一 TreeFold）：

```go
// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	answer := make([]int, n)
	if n == 0 {
		return answer
	}
	TreeFold(graph, 0, -1, func(node int) [26]int {
		var counts [26]int
		counts[labels[node]-'a'] = 1
		return counts
	}, func(left, right [26]int) [26]int {
		for i := range left {
			left[i] += right[i]
		}
		return left
	}, func(node int, counts [26]int) { answer[node] = counts[labels[node]-'a'] })
	return answer
}
```

| 题目 | 原始代码 | 子树状态／边贡献 |
| --- | --- | --- |
| [1443. 收集树上所有苹果的最少时间](https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/) | [邻接表](l1443/1443) · [显式树变体](l1443/variants/explicit_tree/1443_tree) | 返回子树是否需要访问；需要时累计往返边代价 2 |
| [1519. 子树中标签相同的节点数](https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/) | [实现](l1519/1519.go) | 返回标签频次；合并孩子后记录当前标签数量 |
| [1466. 重新规划路线](https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/) | [笔记](l1466/1466) | 建双向遍历边，同时记录原方向是否需要反转 |
| [2477. 到达首都的最少油耗](https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/) | [笔记](l2477/2477) | 返回子树人数；非根边贡献为向上取整的 人数/seats |
| [98. 验证二叉搜索树](https://leetcode.com/problems/validate-binary-search-tree/) | [笔记](l98/98) | 传递开区间上下界，或中序检查顺序 |
| [101. 对称二叉树](https://leetcode.com/problems/symmetric-tree/) | [空笔记](l101/101) | 成对比较镜像节点；保留原空文件待练习 |
| [114. 二叉树展开为链表](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/) | [笔记](l114/114) | 需要改写指针，不能仅返回子树统计量 |

最后三题保留原 DFS 笔记供查找；二叉树比较、结构改写与上面的子树大小模板分别有自己的状态。递归空间为 O(h)，退化链中 h=n，不能一律写为 O(log n)。
