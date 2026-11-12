# 并查集 DSU：合并与连通性

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [Go 通用骨架](#code)
- [推演与题目变化点](#mapping)
- [复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

对象形成互不相交集合，操作主要是合并集合、判断两个对象是否属于同一集合。常见于无向连通性、等价关系和 Kruskal 最小生成树。

不适用：一般有向可达性、有向判环、普通在线删边。DSU 只保存集合关系，不保留实际路径，也不能直接回答两点最短路。

**不变量**：parent 构成若干棵树，根满足 parent[root]=root；同集合的 Find 相同。size 只在根上表示集合大小，非根的 size 无需更新。components 只在两个不同集合真正合并时减一。

固定骨架是路径压缩＋按大小合并；输入如何建模为节点、一次失败合并意味着什么，由题目决定。

<a id="code"></a>

## 2. Go 通用骨架

```go
package template

type DSU struct {
    parent []int
    size []int
    components int
}

func NewDSU(n int) *DSU {
    d:=&DSU{parent:make([]int,n),size:make([]int,n),components:n}
    for i:=0;i<n;i++ { d.parent[i]=i; d.size[i]=1 }
    return d
}

// 路径折半压缩；x 必须在合法编号范围。
func(d *DSU) Find(x int) int {
    for x!=d.parent[x] {
        d.parent[x]=d.parent[d.parent[x]]
        x=d.parent[x]
    }
    return x
}

// 返回是否发生了真正的合并。
func(d *DSU) Union(a,b int) bool {
    ra,rb:=d.Find(a),d.Find(b)
    if ra==rb { return false }
    if d.size[ra]<d.size[rb] { ra,rb=rb,ra }
    d.parent[rb]=ra
    d.size[ra]+=d.size[rb]
    d.components--
    return true
}

func(d *DSU) Connected(a,b int) bool { return d.Find(a)==d.Find(b) }
func(d *DSU) Components() int { return d.components }
func(d *DSU) Size(x int) int { return d.size[d.Find(x)] }
```

<a id="mapping"></a>

## 3. 推演与题目变化点

n=5，边依次为 (0,1)、(1,2)、(0,2)、(3,4)：

| 操作 | 是否合并 | components | 含义 |
| --- | --- | --- | --- |
| 初始化 | — | 5 | 每点独立 |
| Union(0,1) | true | 4 | 两点连通 |
| Union(1,2) | true | 3 | 0、1、2 同集合 |
| Union(0,2) | false | 3 | 已连通，不减少分量 |
| Union(3,4) | true | 2 | 最终两个分量 |

| 题目 | 输入映射 | 何时 Union | 答案 / 变化点 |
| --- | --- | --- | --- |
| [547. 省份数量](https://leetcode.com/problems/number-of-provinces/) | 城市为节点，邻接矩阵 | i<j 且矩阵为1 | Components；遍历矩阵成本 O(n²) 不能忽略 |
| [684. 冗余连接](https://leetcode.com/problems/redundant-connection/) | 无向边，标签为1..n | 按输入顺序处理边 | Union 返回 false 的边产生环；分配 n+1 或转0-based |
| [1319. 连通网络的操作次数](https://leetcode.com/problems/number-of-operations-to-make-network-connected/) | 计算机和网线 | 合并每条边 | 边少于 n-1 直接 -1，否则 Components-1 |
| [721. 账户合并](https://leetcode.com/problems/accounts-merge/) | 邮箱映射整数ID | 同账户内邮箱连边 | 按根收集邮箱并排序；同名不意味着同集合 |

1319 需要的边数条件不能省略：分量数告诉你缺几条连接，但必须有足够多可重接的冗余边。边总数至少 n-1 时，冗余边数量足够完成合并。

684 的判环仅针对无向边逐条加入；有向图中两点已在同一无向集合，不代表新增有向边会形成有向环。

<a id="pitfalls"></a>

## 4. 复杂度与陷阱

- 初始化 O(n)，m 次 Find/Union 总时间 O(m·α(n)) 的摊还量级，空间 O(n)。α(n) 很小，但不是宣称每次操作严格 O(1)。
- 必须合并根 ra、rb，不能直接把 parent[a] 指向 b，否则可能拆散已有结构或形成错误链接。
- 不要重复减 components，重复边、已连通边不会减少集合数量。
- 根编号只是代表元，不一定是最小节点、最早节点或题目要求的名字；额外属性须单独维护。
- size 查询要先 Find。路径压缩改变树形，但不改变实际集合，不能用 parent 边当原图路径。
- 带权 DSU、可回滚 DSU 是额外变体：回滚版本通常不直接使用路径压缩，不能不加分析照搬。

已有练习：[union_find](../../union_find/)、[最小生成树](../../graph/MinimumSpanningTree/)。有向依赖见 [拓扑排序](../topological_sort/README.md)。
