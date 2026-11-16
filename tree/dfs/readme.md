# 二叉树 DFS：设计递归返回状态，逐步进入树形 DP

## 目录

- [1. 统一 PostOrder 模板](#framework)
- [2. 刷题顺序与状态速查](#overview)
- [3. 104 最大深度](#lc104)
- [4. 110 平衡二叉树](#lc110)
- [5. 543 二叉树直径](#lc543)
- [6. 236 最近公共祖先](#lc236)
- [7. 124 最大路径和](#lc124)
- [8. 337 打家劫舍 III](#lc337)
- [9. 979 分配硬币](#lc979)
- [10. 968 监控二叉树](#lc968)
- [11. 从递归到树形 DP](#tree-dp)
- [12. 复杂度与练习方式](#practice)
- [13. State 设计方法与示例](#state-design)

所有逐题答案均使用同一份 PostOrder。独立提交时，复制 TreeNode（平台未提供时）、PostOrder 和对应解题函数即可。需要 Go 1.21+（泛型需 1.18+，内置 min/max 需 1.21+）。

<a id="framework"></a>

## 1. 统一框架：先定义返回值，再写递归

```text
取得左子树的信息
    ↓
取得右子树的信息
    ↓
在当前节点合并
    ↓
向父节点返回它需要的状态
```

这类题以**后序处理**为主：不是简单地访问节点，而是先解决子问题，再合并子树信息。

每题先回答三个问题：

1. **父节点需要子树提供什么信息？** 把返回值写成一句严格定义。
2. **空节点应该返回什么？** 返回值应能自然参与合并，而不是凭直觉填写 0。
3. **递归返回值与最终答案是否同义？** 子树给父节点的信息，未必就是子树内的完整答案。

```go
// 结构示意：State、emptyState、merge 是按题意定义的占位符。
func dfs(node *TreeNode) State {
	if node == nil {
		return emptyState
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	return merge(node, left, right)
}
```

### 1.1 可覆盖这八类题的统一模板：返回完整子树状态

把原来通过全局变量累计的答案也放进 State，就可以统一成一个后序折叠（postorder fold）。**通用的是控制流程，不是所有题共享同一组状态字段或转移公式。**

```go
// Go 1.18+；TreeNode 定义见下方。
// empty 定义空树状态；merge 定义当前节点如何合并左右子树状态。
func PostOrder[S any](
	node *TreeNode,
	empty S,
	merge func(node *TreeNode, left, right S) S,
) S {
	if node == nil {
		return empty
	}

	left := PostOrder(node.Left, empty, merge)
	right := PostOrder(node.Right, empty, merge)

	return merge(node, left, right)
}
```

每题只需填四项：`State`、`empty`、`merge`、`根状态如何取答案`。empty 建议使用值类型；若包含 map、切片或指针，不要无意修改多棵空子树共用的数据。

| 题目 | State | empty | merge / 根部取答案 |
|---|---|---|---|
| [104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | height | 0 | max(L,R)+1；直接返回 |
| [110](https://leetcode.com/problems/balanced-binary-tree/) | 高度或 -1 | 0 | 任一子树失败或高度差超过 1 则 -1；根非负即平衡 |
| [543](https://leetcode.com/problems/diameter-of-binary-tree/) | {height,diameter} | {0,0} | 高度取较大值+1；直径取 max(L.diameter,R.diameter,L.height+R.height) |
| [236](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | 候选 *TreeNode | nil | 当前是目标则返回当前；左右均非空返回当前；否则返回非空候选 |
| [124](https://leetcode.com/problems/binary-tree-maximum-path-sum/) | {down,best,valid} | {0,0,false} | down 只接一支；best 比较当前双支路径与有效子树 best；根取 best |
| [337](https://leetcode.com/problems/house-robber-iii/) | {skip,take} | {0,0} | 分别计算不选与选的最优值；根取 max(skip,take) |
| [979](https://leetcode.com/problems/distribute-coins-in-binary-tree/) | {balance,moves} | {0,0} | balance=Val-1+L.balance+R.balance；moves=L.moves+R.moves+abs(L.balance)+abs(R.balance) |
| [968](https://leetcode.com/problems/binary-tree-cameras/) | {A,B,C} 三种代价 | {INF,0,0} | 按第 10.2 节状态转移；根取 min(A,B) |

124 中 valid 区分“空树没有非空路径”和“合法路径和为 0”，避免把空树的 best=0 当作答案。也可用安全的负无穷表示空树 best，但要小心溢出。

### 1.2 示例：543 不用全局变量也能套模板

```go
func diameterWithState(root *TreeNode) int {
	type State struct {
		height   int
		diameter int
	}
	state := PostOrder(root, State{}, func(node *TreeNode, l, r State) State {
		return State{
			height:   1 + max(l.height, r.height),
			diameter: max(l.diameter, r.diameter, l.height+r.height),
		}
	})
	return state.diameter
}
```

本文所有解法统一使用 PostOrder，543、124、979 也将完整答案保存在 State 中，不使用外部累计变量。实际刷题时也可改写为局部 dfs，控制流程与状态含义不变。通用遍历总是先处理左右子树，因此不提供按题意提前剪枝；110 的失败状态、236 的目标命中只在 merge 中合并。每个节点仍只访问一次。

**适用范围：** 这八道题都能由左右子树状态自底向上合并，因此可使用此模板。它不是所有二叉树 DFS 的万能解法：根到叶路径约束常需向下传参，换根 DP 常需两遍遍历，不能机械套用。

下文各题模板独立使用，统一采用 LeetCode 的 TreeNode；max/min 使用 Go 1.21+ 内置函数。

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

<a id="overview"></a>

## 2. 刷题顺序与状态速查

| 顺序 | 题目 | 返回给父节点的信息 | 空节点返回 | 最终答案 |
|---|---|---|---|---|
| 1 | [104. 二叉树的最大深度](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | 子树高度，按节点数计算 | 0 | 根的高度 |
| 2 | [110. 平衡二叉树](https://leetcode.com/problems/balanced-binary-tree/) | 高度，或 -1 表示不平衡 | 0 | 根返回值是否非负 |
| 3 | [543. 二叉树的直径](https://leetcode.com/problems/diameter-of-binary-tree/) | {height,diameter} | {0,0} | 根的 diameter，按边计 |
| 4 | [236. 二叉树的最近公共祖先](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | 找到的目标节点或已确定的 LCA | nil | 根调用返回的节点 |
| 5 | [124. 二叉树中的最大路径和](https://leetcode.com/problems/binary-tree-maximum-path-sum/) | {down,best,valid} | {0,0,false} | 根的 best |
| 6 | [337. 打家劫舍 III](https://leetcode.com/problems/house-robber-iii/) | 不选 / 选当前节点的两个最优值 | (0,0) | max(skip,take) |
| 7 | [979. 在二叉树中分配硬币](https://leetcode.com/problems/distribute-coins-in-binary-tree/) | {balance,moves} | {0,0} | 根的 moves |
| 8 | [968. 监控二叉树](https://leetcode.com/problems/binary-tree-cameras/) | {camera,covered,uncovered} 三种代价 | {INF,0,0} | min(camera,covered)；另有贪心对照 |

<a id="lc104"></a>

## 3. [104. 二叉树的最大深度](https://leetcode.com/problems/maximum-depth-of-binary-tree/)：返回子树高度

练习：[目录](l104/) · [函数](l104/104.go) · [测试](l104/104_test.go)

**状态定义：** dfs(node) 返回以 node 为根的子树高度，按节点数计算。空树高度为 0。

```go
func maxDepth(root *TreeNode) int {
	return PostOrder(root, 0, func(node *TreeNode, left, right int) int {
		return 1 + max(left, right)
	})
}
```

父节点只需要左右高度，当前高度为 `max(left,right)+1`。本题返回值就是最终答案。

**边界：** 空树为 0，叶子为 1。先约定高度按节点数还是边数，后面 543 依赖这个约定。

<a id="lc110"></a>

## 4. [110. 平衡二叉树](https://leetcode.com/problems/balanced-binary-tree/)：高度 + 失败状态

练习：[目录](l110/) · [函数](l110/110.go) · [测试](l110/110_test.go)

仅有左右高度不足以判断整棵子树平衡：左右子树内部也必须平衡。因此用 -1 表示“这棵子树不平衡”，非负数表示真实高度。

```go
func isBalanced(root *TreeNode) bool {
	height := PostOrder(root, 0, func(node *TreeNode, left, right int) int {
		if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
			return -1
		}
		return 1 + max(left, right)
	})
	return height != -1
}
```

这里返回的是“高度或失败”，最终答案是布尔值。每个节点只求一次高度，避免在每层重新计算子树高度而退化到 O(n²)。空树平衡。

<a id="lc543"></a>

## 5. [543. 二叉树的直径](https://leetcode.com/problems/diameter-of-binary-tree/)：返回高度，更新直径

练习：[目录](l543/) · [函数](l543/543.go) · [测试](l543/543_test.go)

**返回 State：** height 是从当前节点向下走的最大节点数；diameter 是子树内部最大路径的边数。

**当前答案候选：** `left.height+right.height`，与左右子树已有 diameter 比较。只有 height 可以作为一条支路接入父节点，diameter 仅向上传递作答案比较。

```go
func diameterOfBinaryTree(root *TreeNode) int {
	type State struct {
		height, diameter int
	}
	state := PostOrder(root, State{}, func(node *TreeNode, left, right State) State {
		return State{
			height:   1 + max(left.height, right.height),
			diameter: max(left.diameter, right.diameter, left.height+right.height),
		}
	})
	return state.diameter
}
```

```text
      1
     / \
    2   3
   / \
  4   5

节点 2：左右高度 1、1，答案候选 2，向父返回高度 2。
节点 1：左右高度 2、1，答案候选 3，向父返回高度 3。
最终直径 3 条边：4 → 2 → 1 → 3（或从 5 出发）。
```

不能返回 `left+right+1` 给父节点：那已经包含两条分支，父节点再接入会形成分叉，不再是一条简单路径。直径也不一定经过根，每个节点都要尝试更新答案。

<a id="lc236"></a>

## 6. [236. 二叉树的最近公共祖先](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)：返回目标节点或祖先候选

练习：[目录](l236/) · [函数](l236/236.go) · [测试](l236/236_test.go)

标准题目保证 p、q 是树中两个不同节点。比较的是**节点指针身份**，不是 Val。

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return PostOrder(root, (*TreeNode)(nil), func(node, left, right *TreeNode) *TreeNode {
		if node == p || node == q {
			return node
		}
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	})
}
```

返回值语义：

- nil：子树中没有找到目标。
- p 或 q：找到一个目标，或者该目标本身就是两者的祖先。
- 其他节点：已经在子树内部确定 LCA，向上传递它。

左右均非空且当前节点不是目标时，说明目标分居两侧，当前节点就是 LCA；只有一侧非空则向上传递。merge 遇到当前节点为目标时，优先返回当前节点：若另一个目标在其后代中，当前节点就是祖先。这里左右子树已经遍历完成，并不是递归入口处提前返回。

**适用前提：** p、q 均存在。如果允许缺失，单个返回指针不足以证明两者都存在，应额外返回目标命中数量等状态。

<a id="lc124"></a>

## 7. [124. 二叉树中的最大路径和](https://leetcode.com/problems/binary-tree-maximum-path-sum/)：返回单支贡献，更新完整路径

练习：[目录](l124/) · [函数](l124/124.go) · [测试](l124/124_test.go)

**返回 State：** down 是从当前节点出发向下的一条非空路径的最大和，可能为负；best 是子树内任意非空路径的最大和；valid 区分空树与合法结果。父节点只通过 down 接入支路，通过 best 继承已完成答案。

**当前答案候选：** 当前节点 + 左侧非负 down + 右侧非负 down，并与左右有效 best 比较。

```go
// 题目保证 root 非空，路径至少包含一个节点。
func maxPathSum(root *TreeNode) int {
	type State struct {
		down, best int
		valid      bool
	}
	state := PostOrder(root, State{}, func(node *TreeNode, left, right State) State {
		leftGain, rightGain := max(0, left.down), max(0, right.down)
		best := node.Val + leftGain + rightGain
		if left.valid {
			best = max(best, left.best)
		}
		if right.valid {
			best = max(best, right.best)
		}
		return State{
			down:  node.Val + max(leftGain, rightGain),
			best:  best,
			valid: true,
		}
	})
	return state.best
}
```

```text
       -10
       / \
      9   20
         /  \
        15   7

节点 20：左右贡献 15、7。
更新答案：15+20+7=42。
向父返回：20+max(15,7)=35，不能返回 42。

节点 -10：左右贡献 9、35。
答案候选：9-10+35=34，小于已有的 42。
最终答案：42，并不经过根。
```

**全负数：** 当前 best 从包含 node.Val 的路径开始，仅比较 valid 子树的 best，不能无条件与空树的 0 比较。`max(0,child.down)` 表示不接入负贡献子树，不代表可以丢弃当前节点。

### 543 与 124 的统一模板

| 含义 | 543 直径 | 124 最大路径和 |
|---|---|---|
| 完整返回状态 | {height,diameter} | {down,best,valid} |
| 可延伸字段 | height | down |
| 是否舍弃负贡献 | 不需要 | 用 max(0,child.down) 舍弃 |
| 当前答案候选 | L.height+R.height，按边计 | node.Val+leftGain+rightGain |
| 延伸字段更新 | 1+max(L.height,R.height) | node.Val+max(leftGain,rightGain) |
| 内部答案继承 | 同时比较左右 diameter | 仅比较 valid 子树的 best，支持全负树 |

**返回给父节点只能延伸一条支路；当前答案可以连接左右两条支路。** 任意简单路径都有一个深度最浅的节点，在该节点合并左右分支，因此遍历所有节点不会漏掉最优路径。

<a id="lc337"></a>

## 8. [337. 打家劫舍 III](https://leetcode.com/problems/house-robber-iii/)：返回两个条件最优值，正式进入树形 DP

练习：[目录](l337/) · [函数](l337/337.go) · [测试](l337/337_test.go)

父节点是否被选择，会改变当前节点是否允许被选择。不能只返回一个不带条件的最优值，需要同时返回：

- skip：不选当前节点时，整棵子树可获得的最大金额。
- take：选当前节点时，整棵子树可获得的最大金额。

```go
func rob(root *TreeNode) int {
	type State struct {
		skip, take int
	}
	state := PostOrder(root, State{}, func(node *TreeNode, left, right State) State {
		return State{
			skip: max(left.skip, left.take) + max(right.skip, right.take),
			take: node.Val + left.skip + right.skip,
		}
	})
	return max(state.skip, state.take)
}
```

```text
        3
       / \
      2   3
       \   \
        3   1

叶子 3 → (skip=0,take=3)
节点 2 → (skip=3,take=2)
右侧 3 → (skip=1,take=3)
根节点 → (skip=6,take=7)
答案 7。
```

空节点 `(0,0)` 是方便合并的约定，不代表真的选择了空节点。根没有父节点限制，因此取两种状态的最大值。

已有实现：[337.go](../../dp/house_robber/l337/337.go)。

<a id="lc979"></a>

## 9. [979. 在二叉树中分配硬币](https://leetcode.com/problems/distribute-coins-in-binary-tree/)：返回净盈余，累计跨边流量

练习：[目录](l979/) · [函数](l979/979.go) · [测试](l979/979_test.go)

每个节点最终需要 1 枚硬币。定义：

```text
balance(node) = 子树硬币总数 - 子树节点数
             = node.Val - 1 + leftBalance + rightBalance
```

- balance > 0：子树必须向父节点交出这些硬币。
- balance < 0：子树必须从父节点获得这些硬币。
- balance == 0：子树内部供需平衡，但内部可能仍发生移动。

```go
func distributeCoins(root *TreeNode) int {
	type State struct {
		balance, moves int
	}
	state := PostOrder(root, State{}, func(node *TreeNode, left, right State) State {
		return State{
			balance: node.Val - 1 + left.balance + right.balance,
			moves:   left.moves + right.moves + abs(left.balance) + abs(right.balance),
		}
	})
	return state.moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

例如 `[3,0,0]`，两个叶子各返回 -1，根累计移动 2 次，再返回 `3-1-1-1=0`。

**最优性：** 一条边切开树后，子树与外界只能通过这条边交换硬币，因此至少有 `abs(balance)` 枚硬币必须跨过该边；按净流量交换就能达到下界。每条父子边只在父节点处理子状态时计费一次。

**balance 不是答案：** 标准题目保证硬币总数等于节点总数，因此根状态的 balance 为 0，但 moves 可能很大。完整 State 同时保存二者，最终取 moves。

<a id="lc968"></a>

## 10. [968. 监控二叉树](https://leetcode.com/problems/binary-tree-cameras/)：从后序贪心走向状态 DP

练习：[目录](l968/) · [函数](l968/968.go) · [测试](l968/968_test.go)

### 10.1 三状态后序贪心

摄像头覆盖自身、父节点和直接孩子，不覆盖更远节点。

返回状态必须明确区分：

- uncovered：子树其他节点已覆盖，但当前节点未覆盖，需要父节点装摄像头。
- covered：整棵子树已覆盖，当前节点没有摄像头。
- camera：整棵子树已覆盖，当前节点有摄像头。

空节点返回 covered：它不需要被覆盖，也不能要求父节点安装摄像头。

```go
// 对照用贪心版本，也通过 PostOrder 返回完整信息，不使用外部计数。
func minCameraCoverGreedy(root *TreeNode) int {
	const (
		uncovered = iota
		covered
		camera
	)
	type State struct {
		status, cameras int
	}
	state := PostOrder(root, State{status: covered}, func(node *TreeNode, left, right State) State {
		count := left.cameras + right.cameras
		if left.status == uncovered || right.status == uncovered {
			return State{status: camera, cameras: count + 1}
		}
		if left.status == camera || right.status == camera {
			return State{status: covered, cameras: count}
		}
		return State{status: uncovered, cameras: count}
	})
	if state.status == uncovered {
		return state.cameras + 1
	}
	return state.cameras
}
```

判断顺序不能交换：若一个孩子未覆盖、另一个孩子有摄像头，当前节点仍必须安装摄像头，因为兄弟节点的摄像头不能覆盖另一个孩子。

```text
叶子：两个空孩子均 covered → 返回 uncovered，不急着装。
叶子父节点：发现 uncovered 孩子 → 安装摄像头。
再上一层：孩子有 camera → 自己 covered。
根：若最终 uncovered，单独补一个摄像头。
```

直觉：给未覆盖叶子的父节点装，比给叶子装能覆盖更多上层节点；自底向上优先满足已经无法由更高祖先直接覆盖的孩子。

### 10.2 对照学习：显式状态 DP

贪心只返回最终选定的状态；DP 为每种边界条件都保存最小代价。定义：

- A：当前节点有摄像头，整棵子树覆盖的最少摄像头数。
- B：当前节点无摄像头，但被至少一个孩子的摄像头覆盖，整棵子树已覆盖。
- C：当前节点未覆盖，子树其他节点已覆盖，等待父节点摄像头。

空节点约定 `(A,B,C)=(INF,0,0)`：A 不可行；B/C 的 0 是方便缺失孩子参与合并的边界约定，不应按真实节点理解。

```text
A = 1 + min(L.A,L.B,L.C) + min(R.A,R.B,R.C)
B = min(L.A + min(R.A,R.B), R.A + min(L.A,L.B))
C = L.B + R.B
答案 = min(root.A,root.B)
```

解释：

- A 自己有摄像头，因此孩子可以等待当前节点覆盖。
- B 自己无摄像头，孩子不能等待它；并且至少一个孩子必须有摄像头。
- C 自己未覆盖，孩子既不能有摄像头（否则自己被覆盖），也不能等待自己覆盖，只能都处于 B。
- 根没有父节点，不能使用 C。

INF 选取大于任何合法答案的安全值，例如本题节点数上限之外的 `1<<20`，并保证加法不会溢出。这里固定为“未覆盖”状态的 C 与有些资料中的“允许依赖父节点覆盖”状态不同，不能混用转移公式。

统一 PostOrder 的 DP 答案如下，对应 camera=A、covered=B、uncovered=C：

```go
func minCameraCover(root *TreeNode) int {
	const inf = 1 << 20 // 题目至多 1000 个节点，且加法不会溢出。
	type State struct {
		camera, covered, uncovered int
	}
	empty := State{camera: inf}
	state := PostOrder(root, empty, func(node *TreeNode, left, right State) State {
		return State{
			camera: min(inf, 1+min(left.camera, left.covered, left.uncovered)+
				min(right.camera, right.covered, right.uncovered)),
			covered: min(inf, left.camera+min(right.camera, right.covered),
				right.camera+min(left.camera, left.covered)),
			uncovered: min(inf, left.covered+right.covered),
		}
	})
	return min(state.camera, state.covered)
}
```

<a id="tree-dp"></a>

## 11. 从递归到树形 DP：什么时候需要多个状态？

递归是执行方式，DP 是状态设计与转移方法，两者不是互斥分类。

```text
104：一个结构信息（高度）
110：结构信息 + 失败标记
543 / 124：向上传递的信息 ≠ 子树中的完整答案
236：候选节点的合并
337：父节点决策影响孩子，返回一组条件最优值
979：返回局部收支，答案累计在边上
968：边界关系影响覆盖，设计有限状态，再比较贪心与 DP
```

**多状态的信号：** 同一棵子树，在“父节点选不选 / 是否能提供覆盖”等不同外部条件下，最优值不同。此时不要让 dfs 只返回一个最大值，而要返回各条件下的最优值。

树上每个子树只有一个父节点；状态一旦计算即可向上传递，通常不需要额外 memo。若递归反复计算孙子等重叠子问题，才需要记忆化或重新设计返回状态。

### 通用自检清单

- 状态是否只描述当前子树，并且足够让父节点合并？
- 空节点返回值能否正确推导出叶子的状态？
- 路径题返回给父节点时，是否错误地携带了两条支路？
- 高度按节点数还是边数，是否与答案单位一致？
- 是否处理全负值、单节点、链状树、答案不经过根的情况？
- 完整答案是否保存在当前调用的 State 内？本文不使用外部可变答案，避免多次调用污染。
- 摄像头的空节点状态、不可行代价和根节点收尾是否正确？

<a id="practice"></a>

## 12. 复杂度与建议练习方式

八题目录均包含独立的 TreeNode、相同的 PostOrder、标准解题函数和表驱动测试。所有答案均已基于 PostOrder 实现；其他目录中的解法保持不变。每题包可独立测试，重复放置模板是为了便于单题练习和提交。

```sh
# 运行八个题目包的全部测试。
GO111MODULE=off go test -v ./tree/dfs/...

# 完成某题后执行对应测试。
GO111MODULE=off go test -v ./tree/dfs/l104
```

以上模板每个节点最多处理一次，时间均为 **O(n)**，递归栈为 **O(h)**：平衡树 O(log n)，链状树 O(n)。统一 PostOrder 总是遍历整棵树，236 不在遇到目标时提前结束子树遍历；这不影响候选合并的正确性。968 的三状态 DP 每个节点只有常数个转移，也为 O(n)。

递归栈不算 O(1) 空间；极深树需要留意运行环境的栈限制，必要时改为显式栈后序遍历。

推荐按 **104 → 110 → 543 → 236 → 124 → 337 → 979 → 968** 刷题：

1. 写一句状态定义，不看代码先确定返回值。
2. 写空节点状态，再手算叶子状态。
3. 用三节点树推导合并公式。
4. 单独写出“返回值”和“答案更新式”，确认是否同义。
5. 最后补单节点、偏斜树与该题特有边界测试。

重点将 **543 与 124 并排练习**，再将 **337 与 968 的条件状态并排比较**。

**口诀：先定义返回，空树定边界；左右取信息，当前做合并；单支交父亲，双支更答案；条件若不同，多状态来存。**

<a id="state-design"></a>

## 13. State 如何设计：从父节点的问题反推

State 是**一棵子树处理完后，交给父节点的信息摘要**。父节点不需要重新遍历子树，只依靠摘要就能完成自己的计算。

不要先问“结构体放几个字段”，先问：**我是父节点，要解决当前问题，需要向左右子树分别询问什么？**

### 13.1 先找最小必要信息，再检查是否丢失关键区别

[104. 最大深度](https://leetcode.com/problems/maximum-depth-of-binary-tree/)：父节点只问左右有多高，因此一个 height 就够了，合并为 `1+max(left,right)`。不必返回节点列表或每层宽度；一个字段时直接返回 int 即可。

检查状态是否足够，可以用这个标准：

> 如果两棵子树返回相同的 State，父节点能否把它们当作同一种情况处理？如果不能，说明漏掉了信息。

例如 [110. 平衡二叉树](https://leetcode.com/problems/balanced-binary-tree/)，下面两棵子树高度都是 3：

```text
子树 A：         子树 B：
    o                o
   / \              /
  o   o            o
 /                /
o                o
```

A 平衡，B 不平衡。只返回高度 3，父节点无法区分内部是否失衡，因此完整状态应是：

```go
// 状态与转移片段，供理解，不是独立函数。
type BalanceState struct {
	height   int
	balanced bool
}

// height = 1 + max(left.height, right.height)
// balanced = left.balanced && right.balanced && abs(left.height-right.height) <= 1
```

随后才能把它压缩成“非负数表示高度，-1 表示不平衡”。**特殊值是状态编码，不是状态定义本身。**

### 13.2 分开保存：向上延伸的信息，与子树内部的答案

[543. 二叉树的直径](https://leetcode.com/problems/diameter-of-binary-tree/)需要两类信息：

- height：从当前节点向下最多走多远，父节点可以继续接入。
- diameter：子树内部已经出现的最长路径，可能不经过当前节点，不能作为一条支路接入父节点。

```text
height   = 1 + max(L.height,R.height)
diameter = max(L.diameter,R.diameter,L.height+R.height)
```

第二个公式枚举了路径的全部位置：完全在左侧、完全在右侧、经过当前节点。

| 实现形式 | 向上返回 | 完整答案存放位置 |
|---|---|---|
| 闭包维护答案 | height | 外部 answer |
| 完整 State | {height,diameter} | State.diameter |

两者等价。“只返回高度”不代表算法只需要高度，而是完整答案被保存在另一个位置。第 1.2 节提供了完整 State 实现。

### 13.3 用 valid 区分空状态与合法的零答案

[124. 最大路径和](https://leetcode.com/problems/binary-tree-maximum-path-sum/)同样拆成 down 与 best，但全负数要求区分空路径和合法路径。

以下完整示例可以直接结合第 1 节的 PostOrder 使用：

```go
// 题目保证 root 非空。
func maxPathSumWithState(root *TreeNode) int {
	type State struct {
		down  int  // 从当前节点出发，向下的单支非空路径最大和
		best  int  // 子树内任意非空路径的最大和
		valid bool // 子树是否非空，best 是否有效
	}
	state := PostOrder(root, State{}, func(node *TreeNode, l, r State) State {
		leftGain, rightGain := max(0, l.down), max(0, r.down)
		best := node.Val + leftGain + rightGain
		if l.valid {
			best = max(best, l.best)
		}
		if r.valid {
			best = max(best, r.best)
		}
		return State{
			down:  node.Val + max(leftGain, rightGain),
			best:  best,
			valid: true,
		}
	})
	return state.best
}
```

空状态是 `{down:0,best:0,valid:false}`。只有一个 -3 节点时，不能直接计算 `max(0,0,-3)`，否则错误地允许空路径。

**0 可以表示“不接入某条支路”，但不能无条件表示“合法的最大路径和”。** 这里 down=0 的空状态可直接参与贡献比较，best=0 则必须由 valid 保护。

### 13.4 外部条件影响选择：返回每种条件下的答案

[337. 打家劫舍 III](https://leetcode.com/problems/house-robber-iii/)考虑下面的子树：

```text
    5
   /
  4
```

单独求最优值是 5，但若父节点已被选择，当前节点 5 就不能选，只能取得 4。只返回 5 会丢失父节点需要的信息，因此同时保留：

```text
skip：不选当前节点时，整棵子树的最大金额。
take：选当前节点时，整棵子树的最大金额。

skip = max(L.skip,L.take) + max(R.skip,R.take)
take = node.Val + L.skip + R.skip
```

不能在子树内部提前把 skip / take 合并成一个最大值；必须让父节点根据自己的选择使用对应状态。根没有父节点限制，才取 `max(skip,take)`。

[968. 监控二叉树](https://leetcode.com/problems/binary-tree-cameras/)也是同样的逻辑，只是父子间的约束从“不能相邻选择”变成“谁提供覆盖”：

| DP 字段 | 边界条件 | 字段保存什么 |
|---|---|---|
| camera（第 10.2 节 A） | 当前有摄像头，子树全覆盖 | 满足条件的最少摄像头数 |
| covered（B） | 当前无摄像头，由孩子覆盖，子树全覆盖 | 满足条件的最少摄像头数 |
| uncovered（C） | 仅当前未覆盖，等待父节点覆盖 | 满足条件的最少摄像头数 |

**三状态贪心返回一个已经选定的状态；状态 DP 同时返回每种条件的最小代价。** 状态名称类似，但返回值含义不同，不要混用状态编号与代价。

### 13.5 State 不一定是最优值，也可能是供需或候选

[979. 分配硬币](https://leetcode.com/problems/distribute-coins-in-binary-tree/)中，父节点问的是“子树多了多少硬币，或缺了多少”，而不是“子树最多有多少硬币”：

```text
balance = node.Val - 1 + L.balance + R.balance
moves   = L.moves + R.moves + abs(L.balance) + abs(R.balance)
```

balance 是通过父边交换的净数量；moves 是子树内部成本。当前节点合并时，增加左右两条边上的必要交换次数。

[236. 最近公共祖先](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)则返回候选节点身份，而不是数值最优解。在标准题保证目标均存在的前提下，父节点只需要知道目标是否在左右子树，或是否已经找到祖先候选。若目标可能缺失，就需额外保留命中数量或标记。

### 13.6 五步设计法：定义、交互、完整答案、边界、根收尾

1. **限定子问题范围：** `dfs(node)` 只描述以 node 为根的子树。
2. **寻找父子交互：** 父节点接路径、限制选择、提供覆盖，还是交换资源？据此定义向上信息或条件状态。
3. **保存内部答案：** 向上信息是否能代表子树内所有答案？不能就增加 best / cost，或用外部变量累计。
4. **定义空状态并验证叶子：** 状态应能自然合并，不能把不可行状态误写成免费的 0。
5. **明确根收尾：** 根没有父节点，哪些状态可以选，哪些必须排除？

| 题目 | 空状态 | 叶子应得到 | 根如何取答案 |
|---|---|---|---|
| [104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | 高度 0 | 高度 1 | height |
| [110](https://leetcode.com/problems/balanced-binary-tree/) | {0,true} | {1,true} | balanced，或编码后高度非负 |
| [543](https://leetcode.com/problems/diameter-of-binary-tree/) | {height:0,diameter:0} | {1,0} | diameter |
| [124](https://leetcode.com/problems/binary-tree-maximum-path-sum/) | {0,0,false} | {Val,Val,true} | best |
| [337](https://leetcode.com/problems/house-robber-iii/) | {skip:0,take:0} | {0,Val} | max(skip,take) |
| [979](https://leetcode.com/problems/distribute-coins-in-binary-tree/) | {balance:0,moves:0} | {Val-1,0} | moves，合法输入下 balance=0 |
| [968](https://leetcode.com/problems/binary-tree-cameras/) DP | {A:INF,B:0,C:0} | {1,INF,0} | min(A,B)，不允许等待父节点 |

968 的 INF 表示不可行，做加法时结果可能大于 INF，仍视为不可行；实现可以截断到 INF。不能用最大 int 直接相加导致溢出。

如果空状态推导不出正确的叶子状态，优先重新检查状态定义与合并公式，而不是到处补叶子特判。

### 13.7 最终记忆：把字段分成三类

```text
1. 给父节点接着使用的信息
   height、down、balance、祖先候选

2. 子树内部已经完成的答案
   diameter、best、moves

3. 不同边界条件下的答案
   skip / take、camera / covered / uncovered
```

不是每道题都需要三类。**先问父节点需要什么，再看子树内部要保留什么，最后检查不同外部条件是否必须分开。**
