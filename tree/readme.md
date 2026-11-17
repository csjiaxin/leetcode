# 树：按信息流向与维护的不变量选择模板

| 子系列 | 状态流向 / 共同骨架 | 题目与实现 |
| --- | --- | --- |
| [后序返回与树形 DP](postorder/readme.md) | `left=solve(leftChild); right=solve(rightChild); return merge(node,left,right)` | 104、110、543、236、124、337、979、968 |
| [祖先路径状态与回溯](dfs/readme.md) | `enter → 左右孩子 → undo`；退出恢复共享路径信息 | 112、113、129、1448、1026、1457、437 |
| [动态区间聚合](segment/readme.md) | 修改叶子后向上合并，查询合并互不重叠的覆盖节点 | 307 |
| 平衡搜索树 | 搜索路径、旋转与各自的平衡修复规则 | [AVL](avltree/tree.go)、[红黑树](redblacktree/tree.go) |

`dp/house_robber/l337` 的树形状态已经合并到 [337 的独立 DP 变体](postorder/l337/variants/dp/337.go)。主实现和变体都保留自己的 TreeNode、状态结构和测试。

## 两种 DFS 模板如何衔接

后序不变量：返回父节点时，返回值准确描述当前整棵子树，不能把“经过当前点的最终答案”误当成“能继续接到父节点的状态”。

路径 DFS 不变量：进入孩子前，状态只代表根到当前节点的路径；孩子返回后，共享路径、频次表等恢复到进入它之前。标量按值传递通常无需撤销，map 和复用的切片内容仍需恢复。

这两类不能仅因为都写了递归就合并成一个状态转移。各题的节点类型保留在自己的包中，通用骨架与修改点见相应系列文档。

AVL 与红黑树有各自的嵌套 Go 模块，根目录的 `go test ./...` 不会自动覆盖；应在对应模块单独执行测试。两者的平衡不变量不同，保留独立实现。

## 原有资料

[Vertical Order Traversal](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/260502/C%2B%2B-BFSDFS)
