# 回溯：通用解题思路、模板与 LeetCode 题单

## 同文件命名模板与适配器

先复制本代码块，再复制任一题的适配器；实际题解文件同样内嵌 helper，保留官方函数签名。只使用标准库，不导入仓库本地模板包。本文所有 Go 代码块可按顺序组合成一个包；不要重复复制 helper。

```go
import "sort"

// searchChoices restores every choice even when a successful search stops early.
// visit may record a result; returning true requests early termination.
func searchChoices[T any](visit func() bool, choices func() []T, choose, undo func(T)) bool {
	if visit() {
		return true
	}
	for _, choice := range choices() {
		choose(choice)
		found := searchChoices(visit, choices, choose, undo)
		undo(choice)
		if found {
			return true
		}
	}
	return false
}

type boardCell struct{ row, col int }
```

> 回溯不是背诵一种函数，而是：**把答案表示为一连串选择，在搜索树上尝试，失败或完成后恢复状态。**
>
> 本文只整理思路与参考代码，不修改题目目录中的实现。题目链接均使用 leetcode.com。

## 目录

- [1. 什么时候考虑回溯？](#when)
- [2. 通用解题步骤与骨架](#framework)
- [3. 子集与组合：用 start 避免重复顺序](#combination)
  - [78. 子集](#lc78)
  - [77. 组合](#lc77)
  - [39. 组合总和](#lc39)
  - [40. 组合总和 II：同层去重](#lc40)
- [4. 排列：用 used 标记当前路径](#permutation)
  - [46. 全排列](#lc46)
  - [47. 全排列 II：重复元素](#lc47)
- [5. 切割：131. 分割回文串](#partition)
- [6. 网格路径：79. 单词搜索](#grid)
- [7. 约束放置：51. N 皇后](#constraint)
- [8. 剪枝、去重与常见错误](#pitfalls)
- [9. 扩展题单与迁移点](#practice)
- [10. 推荐学习顺序与复盘清单](#route)

<a id="when"></a>
## 1. 什么时候考虑回溯？

常见信号：求所有方案、枚举组合、排列、合法切割、满足约束的摆放；或者只问是否存在方案，但很难直接计算答案。

| 问题类型 | 每层决定什么 | 递归状态 | 何时收集答案 |
| --- | --- | --- | --- |
| 子集 | 接下来选哪个元素 | `start`、`path` | 每个节点都是子集 |
| 固定长度组合 | 下一个选哪个数 | `start`、`path` | `len(path)==k` |
| 组合总和 | 下一个选哪个数 | `start`、`remain` | `remain==0` |
| 排列 | 当前位置放哪个元素 | `used`、`path` | 路径长度等于元素个数 |
| 切割 | 下一段在哪里结束 | `start`、分段路径 | 已消耗全部输入 |
| 网格路径 | 下一步走向哪里 | 位置、进度、路径访问标记 | 完成目标 |
| 约束放置 | 当前行/格放什么 | 当前行/格、约束占用信息 | 所有位置处理完 |

**复杂度通常是指数级或阶乘级，不意味着算法写错。** 如果题目要求输出所有子集或排列，输出本身就可能有 `2^n` 或 `n!` 个。

### 回溯、DFS、动态规划怎么区分？

- **DFS 是遍历方式**；回溯是通过 DFS 试探选择并恢复状态的一种搜索方法。
- 普通图遍历的 `visited` 常表示全局已访问；回溯的 `used` 往往只表示**当前路径**已使用，退出时要恢复。
- 如果不同路径到达的完整状态相同，且只求数量/最优值/是否存在，可以考虑记忆化或 DP。
- 求所有具体方案时，不能仅因两个路径的 `start` 相同就跳过其中一个：它们的前缀答案不同。

<a id="framework"></a>
## 2. 通用解题步骤与骨架

依次回答六个问题：

1. **答案长什么样？** 用 `path` 表示已经完成的部分。
2. **每层选择什么？** 一个数、一段字符串、一步方向，还是一行中的列？
3. **状态需要什么？** 如 `start`、剩余目标、`used`、列和对角线占用。
4. **何时成功，何时失败？** 明确收集答案与提前返回的条件。
5. **哪些分支一定无解或重复？** 在递归前剪枝。
6. **修改了哪些共享状态？** 递归后逐一恢复。

固定骨架（伪代码）：

```text
搜索(状态):
    如果当前节点是答案:
        记录答案的独立副本
        按题意决定是否返回

    遍历当前状态下的候选选择:
        如果选择无效、重复或不可能完成:
            跳过
        做选择：修改 path / used / 约束
        搜索(新状态)
        撤销选择：恢复共享状态
```

本页只封装选择/递归/撤销的控制流程 searchChoices；候选生成、终止条件和状态仍由逐题适配器明确给出。不同剪枝条件不混为一个判断。

### Go 的两条关键规则

```text
// path 是切片，需要复制；不能直接把可复用的 path 存进答案。
answer = append(answer, append([]int{}, path...))

// 撤销最近一次 append。
path = path[:len(path)-1]
```

`start`、`remain` 等整数按值传递，无须撤销；切片底层数据、`used` 和地图等共享状态则需要恢复。

<a id="combination"></a>
## 3. 子集与组合：用 start 避免重复顺序

组合中 `[1,2]` 与 `[2,1]` 是同一方案。只向后选择下标，保证路径下标递增，就不需要枚举不同顺序。

| 变化点 | 写法 | 原因 |
| --- | --- | --- |
| 元素只能使用一次 | `dfs(i+1)` | 下一层不能再选当前下标 |
| 元素允许重复使用 | `dfs(i)` | 下一层仍可选择当前元素 |
| 输入含重复值 | 排序后，`i>start && nums[i]==nums[i-1]` 时跳过 | 同一层不从相同值再次展开 |
| 只要 k 个元素 | 长度达到 k 才收集 | 中间节点不是完整答案 |
| 所有子集 | 进入每层时就收集 | 中间节点本身也是答案 |

<a id="lc78"></a>
### 78. 子集

[LeetCode 78](https://leetcode.com/problems/subsets/) · [本地练习](l78/78.go)

**题意**：给定互不相同的整数，返回所有子集，包括空集。

**状态与解释**：`dfs(start)` 枚举“在当前路径后，还能从哪些下标选择元素”。每个节点先保存一次，再往后扩展。

```go
func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	starts := []int{0}
	searchChoices(func() bool { result = append(result, append([]int{}, path...)); return false },
		func() []int {
			candidates := []int{}
			for i := starts[len(starts)-1]; i < len(nums); i++ {
				candidates = append(candidates, i)
			}
			return candidates
		},
		func(i int) { path = append(path, nums[i]); starts = append(starts, i+1) },
		func(i int) { path = path[:len(path)-1]; starts = starts[:len(starts)-1] })
	return result
}
```

例如 `[1,2]` 的答案是 `[]、[1]、[1,2]、[2]`。时间 `O(n·2^n)`（含复制），辅助空间 `O(n)`，不含答案。本地已有实现使用迭代扩展法，也正确；上面展示的是回溯写法。

<a id="lc77"></a>
### 77. 组合

[LeetCode 77](https://leetcode.com/problems/combinations/) · [本地练习](l77/77.go)

**题意**：从 `1..n` 中选择 k 个不同的数，顺序无关。

**变化点**：仅在长度达到 k 时收集。还缺 `need=k-len(path)` 个数，而从 i 到 n 只有 `n-i+1` 个，必须满足 `i<=n-need+1`。

```go
func combine(n int, k int) [][]int {
	result := [][]int{}
	if k < 0 || k > n {
		return result
	}
	path := []int{}
	searchChoices(func() bool {
		if len(path) == k {
			result = append(result, append([]int{}, path...))
		}
		return false
	},
		func() []int {
			candidates := []int{}
			if len(path) == k {
				return candidates
			}
			start := 1
			if len(path) > 0 {
				start = path[len(path)-1] + 1
			}
			for value := start; value <= n-(k-len(path))+1; value++ {
				candidates = append(candidates, value)
			}
			return candidates
		}, func(value int) { path = append(path, value) }, func(value int) { path = path[:len(path)-1] })
	return result
}
```

例如 `n=4,k=2` 有 `[1,2]、[1,3]、[1,4]、[2,3]、[2,4]、[3,4]`。`k>=1` 时，时间 `O(k·C(n,k))`，辅助空间 `O(k)`，不含答案；`k=0` 返回一个空组合。

<!-- MORE_PROBLEMS -->

## 目录归并：枚举选择，不以题目背景重复分类

原 `permutation/l47`、`l51`、`l526` 已实际迁入本目录；原 `permutation/l46` 与测试完整保存在 [l46/variants/from_permutation](l46/variants/from_permutation)，独立于已有可运行的 [l46](l46)。各变体保持各自包和函数名，测试不转调另一份答案。

| 模型 | 路径 | 下一层枚举范围 / 不变量 |
| --- | --- | --- |
| 不重复排列 | [46](l46) | 每层枚举全部下标，`used` 表示当前路径已选元素 |
| 重复排列 | [47](l47) | 在未使用元素中选择，并在同层去重；不能禁止不同层使用相同数值 |
| 组合与子集 | [77](l77)、[78](l78)、[90](l90) | start 之后选择；90 同层去重 |
| 可重复使用候选 | [39](l39) | 允许重复时递归继续使用当前候选下标，不机械写 i+1 |
| 字符串分段 | [131](l131) | start 是尚未划分的首位置，候选子串必须满足回文等约束 |
| 棋盘/路径约束 | [79](l79)、[51](l51) | 检查邻接和已访问状态，或列与两条对角线冲突 |
| 位置约束计数 | [526](l526) | 位置 pos 和数值 v 满足整除条件，终点加一 |

共同骨架是“选择 → 递归 → 撤销”，不是完全相同的候选集合：

```text
dfs(状态):
    若完成：保存路径副本 / 计数；返回
    对当前状态允许的候选：
        若冲突或属于同层重复：跳过
        更新 path、used、约束状态
        dfs(下一状态)
        恢复本层修改
```

关键不变量：进入和退出一次递归调用时，父层可观察到的 path/used/约束必须相同。保存答案时复制切片，避免后续撤销覆盖旧答案。剪枝必须证明后续不可能形成合法答案；全局去重和同层去重不能混淆。

本次已将实际算法统一为本页 searchChoices 适配器，并修复 51 的历史多余括号及各独立变体的未完成函数；原测试均保留。字典序下一个排列仍归 [permutation](../permutation/readme.md)，无需为只找一个相邻排列枚举 n! 个结果。完整映射见 [reorganization.json](../two_pointers/reorganization.json)。


## 逐题模板适配器补全

下列函数使用本页已声明的同名子模板；不是将不同不变量强行改为同一弹栈或移动条件。

### 39

```go
func combinationSum(candidates []int, target int) [][]int {
	nums := append([]int{}, candidates...)
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	starts := []int{0}
	sum := 0
	searchChoices(func() bool {
		if sum == target {
			result = append(result, append([]int{}, path...))
		}
		return false
	},
		func() []int {
			choices := []int{}
			if sum >= target {
				return choices
			}
			for i := starts[len(starts)-1]; i < len(nums) && sum+nums[i] <= target; i++ {
				choices = append(choices, i)
			}
			return choices
		}, func(i int) { path = append(path, nums[i]); sum += nums[i]; starts = append(starts, i) },
		func(i int) { path = path[:len(path)-1]; sum -= nums[i]; starts = starts[:len(starts)-1] })
	return result
}
```

### 46

```go
func permute(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	searchChoices(func() bool {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
		}
		return false
	}, func() []int {
		candidates := []int{}
		for i := range nums {
			if !used[i] {
				candidates = append(candidates, i)
			}
		}
		return candidates
	}, func(i int) { used[i] = true; path = append(path, nums[i]) },
		func(i int) { used[i] = false; path = path[:len(path)-1] })
	return result
}
```

### 47

```go
func permuteUnique(nums []int) [][]int {
	nums = append([]int{}, nums...)
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	searchChoices(func() bool {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
		}
		return false
	}, func() []int {
		candidates := []int{}
		for i := range nums {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			candidates = append(candidates, i)
		}
		return candidates
	}, func(i int) { used[i] = true; path = append(path, nums[i]) },
		func(i int) { used[i] = false; path = path[:len(path)-1] })
	return result
}
```

### 51

```go
func solveNQueens(n int) [][]string {
	result := [][]string{}
	if n < 0 {
		return result
	}
	if n == 0 {
		return [][]string{{}}
	}
	columns := make([]bool, n)
	down, up := make([]bool, 2*n-1), make([]bool, 2*n-1)
	path := []int{}
	searchChoices(func() bool {
		if len(path) == n {
			board := make([]string, n)
			for row, col := range path {
				line := make([]byte, n)
				for i := range line {
					line[i] = '.'
				}
				line[col] = 'Q'
				board[row] = string(line)
			}
			result = append(result, board)
		}
		return false
	}, func() []int {
		choices := []int{}
		row := len(path)
		if row == n {
			return choices
		}
		for col := 0; col < n; col++ {
			if !columns[col] && !down[row-col+n-1] && !up[row+col] {
				choices = append(choices, col)
			}
		}
		return choices
	}, func(col int) {
		row := len(path)
		columns[col] = true
		down[row-col+n-1] = true
		up[row+col] = true
		path = append(path, col)
	},
		func(col int) {
			row := len(path) - 1
			columns[col] = false
			down[row-col+n-1] = false
			up[row+col] = false
			path = path[:row]
		})
	return result
}
```

### 79

```go
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])
	var frequency [256]int
	for _, row := range board {
		for _, value := range row {
			frequency[value]++
		}
	}
	for i := range word {
		frequency[word[i]]--
		if frequency[word[i]] < 0 {
			return false
		}
	}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	path := []boardCell{}
	return searchChoices(func() bool { return len(path) == len(word) },
		func() []boardCell {
			choices := []boardCell{}
			add := func(row, col int) {
				if row >= 0 && row < rows && col >= 0 && col < cols && !visited[row][col] && board[row][col] == word[len(path)] {
					choices = append(choices, boardCell{row, col})
				}
			}
			if len(path) == 0 {
				for row := 0; row < rows; row++ {
					for col := 0; col < cols; col++ {
						add(row, col)
					}
				}
			} else {
				last := path[len(path)-1]
				add(last.row-1, last.col)
				add(last.row+1, last.col)
				add(last.row, last.col-1)
				add(last.row, last.col+1)
			}
			return choices
		}, func(cell boardCell) { visited[cell.row][cell.col] = true; path = append(path, cell) },
		func(cell boardCell) { visited[cell.row][cell.col] = false; path = path[:len(path)-1] })
}
```

### 90

```go
func subsetsWithDup(nums []int) [][]int {
	nums = append([]int{}, nums...)
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	starts := []int{0}
	searchChoices(func() bool { result = append(result, append([]int{}, path...)); return false },
		func() []int {
			candidates := []int{}
			start := starts[len(starts)-1]
			for i := start; i < len(nums); i++ {
				if i == start || nums[i] != nums[i-1] {
					candidates = append(candidates, i)
				}
			}
			return candidates
		}, func(i int) { path = append(path, nums[i]); starts = append(starts, i+1) },
		func(i int) { path = path[:len(path)-1]; starts = starts[:len(starts)-1] })
	return result
}
```

### 131

```go
func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}
	starts := []int{0}
	searchChoices(func() bool {
		if starts[len(starts)-1] == len(s) {
			result = append(result, append([]string{}, path...))
		}
		return false
	},
		func() []int {
			ends := []int{}
			start := starts[len(starts)-1]
			for end := start + 1; end <= len(s); end++ {
				valid := true
				for l, r := start, end-1; l < r; l, r = l+1, r-1 {
					if s[l] != s[r] {
						valid = false
						break
					}
				}
				if valid {
					ends = append(ends, end)
				}
			}
			return ends
		}, func(end int) {
			start := starts[len(starts)-1]
			path = append(path, s[start:end])
			starts = append(starts, end)
		},
		func(end int) { path = path[:len(path)-1]; starts = starts[:len(starts)-1] })
	return result
}
```

### 526

```go
func countArrangement(n int) int {
	// Visit the most constrained positions first; position order does not change the count.
	possible := make([][]int, n)
	for position := 1; position <= n; position++ {
		for value := 1; value <= n; value++ {
			if value%position == 0 || position%value == 0 {
				possible[position-1] = append(possible[position-1], value)
			}
		}
	}
	sort.Slice(possible, func(i, j int) bool { return len(possible[i]) < len(possible[j]) })
	used := make([]bool, n+1)
	depth, count := 0, 0
	searchChoices(func() bool {
		if depth == n {
			count++
		}
		return false
	},
		func() []int {
			choices := []int{}
			if depth == n {
				return choices
			}
			for _, value := range possible[depth] {
				if !used[value] {
					choices = append(choices, value)
				}
			}
			return choices
		}, func(value int) { used[value] = true; depth++ }, func(value int) { used[value] = false; depth-- })
	return count
}
```
