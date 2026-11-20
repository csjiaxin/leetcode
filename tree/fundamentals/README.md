# 二叉树基础：层序、迭代中序、BST 与结构构造

[返回总目录](../../README.md) · [树系列导航](../readme.md)

本指南补齐现有[后序 DP](../postorder/readme.md)、[路径 DFS](../dfs/readme.md)之外的高频模型。下方为 Go 参考模板，不表示所有题目已有独立题解；平台提交时沿用平台 TreeNode 定义。

## 分类与识别

| 需求 | 模型与关键状态 | 代表题 |
| --- | --- | --- |
| 按层输出 / 每层最后一个节点 | BFS 队列，固定当前层长度 | 102、199、103 |
| 有序遍历 / 第 k 小 | BST 中序，栈保存尚未完成的祖先 | 94、230、173 |
| 验证 BST | 整棵左子树更小、右子树更大 | 98 |
| 遍历序列重建结构 | 根定位＋左右子树区间 | 105、106 |
| 序列化与反序列化 | 数据＋结构标记必须可逆 | 297 |
| 原地展开 | 保存原右子树，再连接左展开链和右展开链 | 114 |

## 模板一：按层 BFS

```go
type TreeNode struct { Val int; Left, Right *TreeNode }

func LevelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil { return result }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        level := make([]int, 0, size)
        for i := 0; i < size; i++ {
            node := queue[i]
            level = append(level, node.Val)
            if node.Left != nil { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
            queue[i] = nil // 释放已消费节点引用。
        }
        result = append(result, level)
        queue = queue[size:]
    }
    return result
}
```

循环开始队列是一整层；本轮追加的孩子必须下一轮处理。102 保存整层，199 取每层最后节点，103 在保存结果时改变索引方向，不必改变依赖关系。时间 O(n)，队列有效元素为 O(w)，w 为最大层宽；输出占 O(n)。切片若长期保存全部历史而只增加 head，会占 O(n) 辅助空间；此处消费后舍弃前缀并清除引用。

## 模板二：显式栈中序与 BST 验证

```go
func Inorder(root *TreeNode) []int {
    result := []int{}
    stack := []*TreeNode{}
    current := root
    for current != nil || len(stack) > 0 {
        for current != nil { stack = append(stack, current); current = current.Left }
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.Val)
        current = current.Right
    }
    return result
}

// 原题不允许重复值，因此中序必须严格递增。
func IsValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    current := root
    previous, hasPrevious := 0, false
    for current != nil || len(stack) > 0 {
        for current != nil { stack = append(stack, current); current = current.Left }
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if hasPrevious && current.Val <= previous { return false }
        previous, hasPrevious = current.Val, true
        current = current.Right
    }
    return true
}
```

栈保存沿左链下探时尚未访问自身的祖先；弹栈访问后转向右子树。时间 O(n)，栈 O(h)，最坏退化 O(n)；Inorder 另占 O(n) 输出。hasPrevious 避免使用可能与合法值冲突的极小值哨兵。

230 在弹出第 k 个节点时返回，时间 O(h+k)、栈 O(h)（k 合法时）。不要先完整排序整个 BST 再查；173 可将相同遍历状态保存到迭代器，Next 摊还 O(1)、单次最坏 O(h)。

## 构造与编码的通用步骤

**105 前序＋中序建树**：要求节点值互不相同。预建中序值→位置映射；每个递归子问题用中序半开区间 `[left,right)`，按前序游标取根，由中序位置切分左右子树；先建左，再建右。平均时间 O(n)、映射 O(n)、栈 O(h)。重复值时位置不唯一，不能直接照搬单位置映射。空区间返回 nil。

**106 后序＋中序**：从后序末尾取根时，先建右再建左；不要沿用前序的递归顺序。

**297 序列化**：可用前序＋空标记（如 `#`）＋分隔符，反序列化共享 token 游标；遇空标记返回 nil，否则解析值并依次构造左右子树。时间和空间按节点数及实际字符串长度分析。只保存节点值、不保存空标记一般不能唯一恢复结构；多位数、负数、空树和非法输入应明确约定。

**114 展开**：不能覆盖右指针后才想起保存原右子树。朴素每个节点反复寻找左链尾可能退化 O(n²)；可用显式栈前序 O(n) 时间、O(h) 栈空间，或进一步学习原地连接的摊还证明。

## 检查与路线

- 不允许环与共享子节点的普通树假设应明确；一般图不能直接使用无 visited 的树遍历。
- BST 不只是比较左右孩子，约束作用于整棵子树。
- h 为树高，不一定是 log n；“平衡树”条件不能自行假设。
- 空树、单节点、退化链、重复值、极端整数、左右子树缺失均要测试。

练习：102 → 199 → 103 → 94 → 98 → 230 → 105 → 106 → 297 → 114 → 173。普通面试先练这些，再深入手写 AVL / 红黑树。
