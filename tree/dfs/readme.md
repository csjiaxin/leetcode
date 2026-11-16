# 树上路径 DFS：前序状态传递＋回溯

本文的六道题解统一使用同一个 `PathDFS[S]` 模板：遍历由模板负责，题目差异放在 `enter` 回调中，共享路径状态由返回的 `undo` 撤销。示例依赖各题的 `TreeNode` 和[第 5 节的共同模板](#template)，不同题目的同名函数应放在各自包中。

## 目录

- [1. 从后序 DFS 切换到前序 DFS](#intro)
- [2. 值传递：路径和不需要手动回溯](#value-state)
- [3. 共享路径：进入时选择，离开时撤销](#backtracking)
- [4. 从数组前缀和，迁移到树上路径](#prefix-state)
- [5. 统一思考框架与共同模板](#template)
- [6. 建议练习顺序与题目目录](#practice)

<a id="intro"></a>

## 1. 从后序 DFS 切换到前序 DFS

| 模式 | 核心问题 | 状态方向 | 典型场景 |
| --- | --- | --- | --- |
| 后序汇总 | 左右子树能给当前节点什么信息？ | 孩子 → 父节点 | 高度、直径、硬币分配 |
| 前序传递 | 到达当前节点时，祖先路径提供了什么信息？ | 父节点 → 孩子 | 路径和、路径记录、祖先约束 |

这里的“前序”强调：**先用当前节点更新路径状态，再把更新后的状态传给孩子**，不一定要输出前序遍历序列。

把递归函数理解为：

```text
dfs(node, state)
state 表示从根到 node 的父节点这一段路径的信息。
进入 node 后，先把 node 纳入状态，再处理答案和孩子。
```

也可以约定传入状态已经包含当前节点，但必须全程统一，不能重复计入或漏掉当前节点。

<a id="value-state"></a>

## 2. 值传递：路径和不需要手动回溯

以 [LeetCode 112「路径总和」](https://leetcode.com/problems/path-sum/) 为例，只统计根到叶子的路径：

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    found := false
    PathDFS(root, int64(0), func(node *TreeNode, sum int64) (int64, func()) {
        sum += int64(node.Val)
        if node.Left == nil && node.Right == nil && sum == int64(targetSum) {
            found = true
        }
        // 值状态各层独立，无需撤销。
        return sum, nil
    })
    return found
}
```

上例用 `int64` 直接表示状态；当前 [112.go](l112/112.go) 也可以用 `type S struct { sum int }` 包装路径和，通过 `S{sum: parentState.sum + node.Val}` 返回新状态。两种形式都遵循同一个模板；只含整数的结构体也是值复制，不需要 `undo`。若路径累加值可能超过 `int` 范围，应把字段改为 `int64`。

- `sum` 是整数，作为参数传递时复制，每层递归各有自己的值。
- 左子树修改自己的 `sum`，不会影响右子树，不需要 `sum -= node.Val`。
- 根到叶子的题目必须判断叶子，不能在中间节点满足条件时直接算作答案。
- 路径状态也可以是最大值、位掩码、剩余目标、上下界等，不只限于求和。

<a id="backtracking"></a>

## 3. 共享路径：进入时选择，离开时撤销

以 [LeetCode 113「路径总和 II」](https://leetcode.com/problems/path-sum-ii/) 为例，需要保存具体路径：

```go
func pathSum(root *TreeNode, targetSum int) [][]int {
    result := make([][]int, 0)
    var path []int
    PathDFS(root, int64(0), func(node *TreeNode, sum int64) (int64, func()) {
        sum += int64(node.Val)
        path = append(path, node.Val)
        if node.Left == nil && node.Right == nil && sum == int64(targetSum) {
            result = append(result, append([]int(nil), path...))
        }
        // 返回闭包，而不是现在执行截短：模板会在左右子树遍历完后调用它。
        return sum, func() {
            path = path[:len(path)-1]
        }
    })
    return result
}
```

**回溯的不变量：退出当前层 `PathDFS(node, ...)` 时，共享路径恢复到进入它之前的状态。**

注意两个 Go 细节：

1. 上例的 `path` 是闭包共享变量，必须撤销，否则左子树的路径会污染右子树。
2. 保存答案时必须复制切片元素；直接保存 `path` 会共享底层数组，后续遍历可能覆盖已保存答案。

如果把 `path []int` 作为递归参数传入，切片头会复制，孩子的长度变化不会改变父调用的长度，因此通常不需要手动缩短父调用的切片。但底层数组仍可能共享，保存答案时依然需要复制。

<a id="prefix-state"></a>

## 4. 从数组前缀和，迁移到树上路径

[LeetCode 437「路径总和 III」](https://leetcode.com/problems/path-sum-iii/) 允许路径从任意节点开始，但只能沿父到子的方向向下延伸，不要求结束于叶子。

定义 `P(node)` 为根到当前节点的路径和，则：

```text
某条以当前节点结尾的向下路径和
= 当前前缀和 - 路径起点之前的前缀和

当前前缀和 - 旧前缀和 = target
旧前缀和 = 当前前缀和 - target
```

这和 [LeetCode 560「和为 K 的子数组」](https://leetcode.com/problems/subarray-sum-equals-k/) 完全对应，区别在于哈希表的有效范围：

| 数组 | 树 |
| --- | --- |
| 前缀来自当前下标之前的位置 | 前缀只能来自当前节点的祖先路径 |
| 从左向右扫描，旧前缀一直有效 | 离开子树必须撤销，不能带到兄弟分支 |

### Go 模板：前缀和＋频次表＋回溯

以下 `pathSum` 与上一节是不同题目的独立实现，不应同时放在同一包中：

```go
func pathSum(root *TreeNode, targetSum int) int {
    count := 0
    freq := map[int64]int{0: 1}
    target := int64(targetSum)
    PathDFS(root, int64(0), func(node *TreeNode, prefix int64) (int64, func()) {
        prefix += int64(node.Val)
        // 1. 查询祖先前缀；答案累计后不撤销。
        count += freq[prefix-target]
        // 2. 登记当前前缀，供后代查询。
        freq[prefix]++
        // 3. 模板递归左右孩子，之后执行这个 undo。
        return prefix, func() {
            freq[prefix]--
            if freq[prefix] == 0 {
                delete(freq, prefix)
            }
        }
    })
    return count
}
```

### 为什么顺序必须是“查 → 加 → 递归 → 撤”？

- **先查再加**：查询时只允许祖先前缀；先加入当前前缀会在 `target == 0` 时错误计入空路径。
- **初始化 `{0: 1}`**：表示根之前的空前缀，让从根开始的合法路径也能被统计。
- **使用频次而不是集合**：不同祖先位置可能有相同前缀和，对应不同的合法路径起点。
- **退出时撤销**：左子树的前缀不是右子树的祖先，不能参与右子树的查询。
- **频次为零时删除键**：保持表只存当前路径的有效前缀，空间才能维持在 O(h)。
- **路径和使用 `int64`**：单个节点值可能不大，累加的路径和仍可能超出 32 位范围。

例如目标和为 0，链 `0 → 0`：第一个节点贡献 1 条路径，第二个贡献 2 条，共 3 条。布尔集合无法记录这种重复前缀的次数。

期望时间复杂度为 **O(n)**；递归栈和有效前缀表的辅助空间为 **O(h)**，其中 h 是树高，最坏为 O(n)。

<a id="template"></a>

## 5. 统一思考框架与共同模板

### 5.1 可复用的 Go 遍历模板

与后序的 `merge(node, left, right) S` 对应，前序的核心是 `enter(node, parentState)`：先得到当前路径状态，再递归孩子，最后撤销共享状态。

下面的模板假设调用方已定义 `TreeNode`，需要 Go 1.18+：

```go
func PathDFS[S any](
    node *TreeNode,
    parentState S,
    enter func(node *TreeNode, parentState S) (state S, undo func()),
) {
    if node == nil {
        return
    }

    state, undo := enter(node, parentState)
    // 保证离开当前节点时恢复共享状态。
    if undo != nil {
        defer undo()
    }

    PathDFS(node.Left, state, enter)
    PathDFS(node.Right, state, enter)
}
```

`enter` 统一完成“更新局部状态、查询/记录答案、登记共享状态”；`undo` 只撤销共享路径状态，**不撤销已经累计的答案**。没有共享状态时返回 `nil`。

**返回值的含义**：`return state, nil` 表示只传递值状态；`return state, func() { ... }` 表示还需要退出清理。回调返回不会结束遍历，模板接下来仍会访问左右孩子。`defer undo()` 登记在当前这层 `PathDFS` 中，等左右孩子都返回后才执行；不要在 `enter` 中用 `defer` 提前清理。

六题都使用这个模板。112、113、437 的完整回调见前文；以下补齐其余三题：

#### 129：把路径当作十进制数字

```go
func sumNumbers(root *TreeNode) int {
    total := 0
    PathDFS(root, 0, func(node *TreeNode, value int) (int, func()) {
        value = value*10 + node.Val
        if node.Left == nil && node.Right == nil {
            total += value
        }
        return value, nil
    })
    return total
}
```

只在叶子处累计，不能把中间前缀数字也算入答案。按题目约束，数字与答案可用 `int`。

#### 1448：传递祖先最大值

```go
func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    count := 0
    PathDFS(root, root.Val, func(node *TreeNode, maximum int) (int, func()) {
        if node.Val >= maximum {
            count++
            maximum = node.Val
        }
        return maximum, nil
    })
    return count
}
```

先用祖先最大值判断，再把更新后的最大值交给孩子。根节点必然是好节点；用根值初始化而不是 0，可以正确处理全负数树。

#### 1457：用位掩码传递出现次数的奇偶性

```go
func pseudoPalindromicPaths(root *TreeNode) int {
    count := 0
    PathDFS(root, 0, func(node *TreeNode, mask int) (int, func()) {
        mask ^= 1 << uint(node.Val)
        if node.Left == nil && node.Right == nil && mask&(mask-1) == 0 {
            count++
        }
        return mask, nil
    })
    return count
}
```

题目中的数字为 1～9；每遇到一个数字就翻转对应位。能重排为回文，当且仅当至多一个数字出现奇数次，也就是掩码至多有一个置位位。`mask == 0` 的全偶数情况也满足条件。

### 六题状态对照

| 题目 | 传给孩子的 state | enter 中的查询/记录 | 共享状态登记与 undo |
| --- | --- | --- | --- |
| 112 | 当前路径和 | 在叶子判断目标和 | 无，返回 nil |
| 129 | 当前数字 `parent*10+Val` | 在叶子累加数字 | 无，返回 nil |
| 1448 | 更新后的祖先最大值 | 用更新前的最大值判断好节点 | 无，返回 nil |
| 113 | 当前路径和 | 路径 append 后，在合法叶子复制路径 | append 路径；退出时截短 |
| 1457 | 更新后的奇偶位掩码 | 在叶子判断至多一个置位位 | 无，返回 nil |
| 437 | 当前前缀和 | 先累计 `freq[prefix-target]` | 再增加频次；退出时减少并删除零频键 |

初始状态按题目选择：求和/数字/位掩码通常是 0；1448 可在根非空时用 `root.Val`；437 的共享频次表预先放入 `{0: 1}`。

**模板的边界**：这个版本统一遍历所有节点，不提供找到答案后短路或剪枝；112 的专用实现可以提前返回，更简洁。若需要剪枝，可扩展回调返回值，但剪枝前仍必须恢复已修改的共享状态。练习时优先理解骨架，不必强行套泛型。

### 5.2 对应的递归执行顺序


```text
进入节点：
    更新当前路径状态
    利用祖先信息查询/记录答案
    把当前状态登记给后代使用

递归孩子：
    dfs(left, state)
    dfs(right, state)

离开节点：
    撤销共享状态，恢复调用前的现场
```

并非每道题都需要全部步骤。先判断状态归属：

- **按值传递的局部状态**：如整数、布尔值、仅含值字段的结构体，通常自动隔离。
- **共享可变状态**：如闭包路径切片、map、指针指向的数据，需要考虑恢复。
- 结构体或切片“按值传递”不等于深拷贝；内部的 map、切片、指针仍可能共享数据。
- 是否有返回值不是前序/后序的判断标准：437 也返回计数，但核心路径状态仍是自顶向下传递的。

**一句话：前序把祖先状态带给孩子，回溯让状态只属于当前路径；前缀状态＋哈希表把“枚举路径起点”变成“查询匹配的祖先前缀”。**

<a id="practice"></a>

## 6. 建议练习顺序与题目目录

各目录中的 Go 文件均包含独立的 `TreeNode`、LeetCode 链接、相同的 `PathDFS[S]` 模板及基于回调的题解，可单独阅读和使用。模板在每题中保留一份，便于提交到 LeetCode；对应的 `*_test.go` 覆盖示例、边界与分支状态隔离，113 额外检查路径快照，437 额外用固定种子的随机树与暴力枚举对照。

统一使用第 5 节的模板：112、129、1448、1457 返回值状态及 `nil`；113 返回路径和及截短路径的 `undo`；437 返回前缀和及撤销频次的 `undo`。这版 112 不短路，最坏时间与其余计数题一样为 O(n)；113 还需计入复制输出路径的开销。

运行这组测试（仓库未配置 go.mod）：

```sh
GO111MODULE=off go test ./tree/dfs/...
```

| 顺序 | LeetCode 题目 | 本地练习 | 重点 |
| --- | --- | --- | --- |
| 1 | [112 路径总和](https://leetcode.com/problems/path-sum/) | [l112/112.go](l112/112.go) | 整数状态传递，根到叶子的限制 |
| 2 | [129 求根节点到叶节点数字之和](https://leetcode.com/problems/sum-root-to-leaf-numbers/) | [l129/129.go](l129/129.go) | 状态更新 `value*10 + node.Val` |
| 3 | [1448 统计二叉树中好节点的数目](https://leetcode.com/problems/count-good-nodes-in-binary-tree/) | [l1448/1448.go](l1448/1448.go) | 传递祖先最大值 |
| 4 | [113 路径总和 II](https://leetcode.com/problems/path-sum-ii/) | [l113/113.go](l113/113.go) | 路径回溯与答案复制 |
| 5 | [1457 二叉树中的伪回文路径](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/) | [l1457/1457.go](l1457/1457.go) | 位掩码记录出现次数的奇偶性 |
| 6 | [437 路径总和 III](https://leetcode.com/problems/path-sum-iii/) | [l437/437.go](l437/437.go) | 任意起点的向下路径，前缀和＋频次表 |

```text
tree/dfs/
├── readme.md
├── l112/
│   ├── 112.go
│   └── 112_test.go
├── l129/
│   ├── 129.go
│   └── 129_test.go
├── l1448/
│   ├── 1448.go
│   └── 1448_test.go
├── l113/
│   ├── 113.go
│   └── 113_test.go
├── l1457/
│   ├── 1457.go
│   └── 1457_test.go
└── l437/
    ├── 437.go
    └── 437_test.go
```
