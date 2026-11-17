# 路径回溯与递归构造

[返回 DFS 目录](../README.md)

## 当前路径的选择必须恢复

适用于枚举路径、切分或排列。**不变量**：path 和 used 只描述当前递归路径，兄弟分支开始前恢复到同一状态。

```go
func search(state int) {
    if done(state) {
        answers = append(answers, append([]int(nil), path...))
        return
    }
    for _, choice := range choices(state) {
        if !allowed(choice) { continue }
        choose(choice)       // 加入 path；需要时标记 used
        search(next(state, choice))
        undo(choice)         // 与 choose 成对恢复
    }
}
```

这是需要填入题意的骨架：`done/choices/allowed/choose/next/undo` 由调用方定义。保存答案必须复制切片；若提前返回，也要确保共享路径状态已经恢复。

| 题目 | 原始代码 | 状态与变化点 |
| --- | --- | --- |
| [79. 单词搜索](https://leetcode.com/problems/word-search/) | [笔记](l79/79) | `(r,c,index)`；同一格不能在当前单词路径重复使用，回退后可用于其他路径 |
| [131. 分割回文串](https://leetcode.com/problems/palindrome-partitioning/) | [笔记](l131/131) | 下一个切分起点；枚举回文片段并 push/pop |
| [95. 不同的二叉搜索树 II](https://leetcode.com/problems/unique-binary-search-trees-ii/) | [独立实现](l95/95.go) | 根值划分左右区间，返回所有子树，做笛卡尔积组合 |

95 的主要模型是递归返回候选集合，不需要机械添加 used；它与“共享一个 path 然后撤销”是两种组织方式。组合中可能共享子树指针，后续如果要独立修改各棵结果树，需要额外复制。

## 95：按区间返回候选集合

这里使用无环区间状态的 MemoDFS，不把路径 used 状态机械加进分治构造。先声明公共类型和模板，再组合适配函数；实际 95.go 已包含它们，独立提交不需要仓库导入。

```go
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func MemoDFS[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(state S) V {
		if value, ok := memo[state]; ok {
			return value
		}
		value := evaluate(state, solve)
		memo[state] = value
		return value
	}
	return solve
}
```

```go
// https://leetcode.com/problems/unique-binary-search-trees-ii/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	solve := MemoDFS(func(interval [2]int, recur func([2]int) []*TreeNode) []*TreeNode {
		left, right := interval[0], interval[1]
		if left > right {
			return []*TreeNode{nil}
		}
		result := make([]*TreeNode, 0)
		for root := left; root <= right; root++ {
			for _, l := range recur([2]int{left, root - 1}) {
				for _, r := range recur([2]int{root + 1, right}) {
					result = append(result, &TreeNode{Val: root, Left: l, Right: r})
				}
			}
		}
		return result
	})
	return solve([2]int{1, n})
}
```

复杂度取决于搜索树与输出总量；剪枝、预计算回文表等优化必须保持可行解不遗漏。它不能套用 [FloodFill](../grid/readme.md) 的全局访问后不撤销规则。
