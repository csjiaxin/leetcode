# 拓扑排序：依赖顺序与DAG动态规划

## 本地模板与参考实现

207 与六题主线的每个主文件均内联完整的 Kahn／ZeroIndegree，没有仓库包导入或跨包薄封装。[template/kahn.go](template/kahn.go) 保留为参考实现；[参考测试](template/kahn_test.go)验证全部前置到齐才出队、重边计数、入参切片所有权、环阻塞与显式种子。

| 代码 | 使用方式 |
| --- | --- |
| [普通 Kahn](indegree_kahn.go) / [三色 DFS](dfs.go) | 直接返回完整拓扑序；DFS 结果需要反转后序 |
| [Kahn 参考模板](template/kahn.go) | 固定调度，按题目设置 next、onPop、relax 与 DP 初值；各题内有完整副本 |
| 210、329、1462、1857、2050、2115 | 题目实现和原测试完整保留 |
| [207](l207/207.go) | 本地 Kahn 适配，只检查处理数量 |
| [2192](l2192/2192) | 祖先集合传播的原始无扩展名笔记 |
| [1203](l1203/1203) | 分组拓扑的原始无扩展名笔记 |

207、210、2050 的原始片段移到各题 variants/legacy；[2115 旧 DFS 实现](l2115/variants/legacy/2115.go)独立成包，避免与通用模板的全局名字相撞。普通零入度队列表示所有依赖都已满足；2115 的外部原料没有供应时，即使入度为零也不能加入 seeds。

本轮主线：[210](https://leetcode.com/problems/course-schedule-ii/) → [2115](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/) → [1462](https://leetcode.com/problems/course-schedule-iv/) → [2050](https://leetcode.com/problems/parallel-courses-iii/) → [1857](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/) → [329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)。

直接阅读[第 6 节：共同模板与六题适配](#dag-dp)。文档只声明一次 `Kahn`／`ZeroIndegree`，随后各题给出适配函数；组合第 6.1 节声明与任一适配即可编译。实际每题主文件已经包含完整本地模板，复制该文件提交时不需要额外仓库依赖，同一包不要重复定义模板。

[返回仓库首页](../../README.md)

[扩展练习：迁移点与边界](#extended-practice) · [保留的原分类笔记](#legacy-notes)

## 目录

- [1. 核心思路](#core)
- [2. 207：能否完成课程](#lc207)
- [3. 210：返回课程顺序](#lc210)
- [4. 2050：并行课程耗时](#lc2050)
- [5. 对比与易错点](#pitfalls)
- [6. 六题主线：拓扑排序＋DAG 动态规划](#dag-dp)
  - [共同模板：起点＋后继枚举＋状态传播](#dag-template)
  - [六题套用与状态转移](#dag-adapters)
  - [复杂度、边界与练习顺序](#dag-practice)

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
func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
        indegree[p[0]]++
    }
    processed := Kahn(indegree, ZeroIndegree(indegree),
        func(u int, yield func(int)) {
            for _, v := range graph[u] { yield(v) }
        }, nil, nil,
    )
    return processed == numCourses
}
```

**例子**：2门课、[[1,0],[0,1]]构成环，无法启动，返回false。若环外有独立节点，也仍须最终判数。

**复杂度**：时间O(V+E)，空间O(V+E)。

<a id="lc210"></a>

## 3. [210. 课程表 II](https://leetcode.com/problems/course-schedule-ii/)

[练习](l210/210.go) · [测试](l210/210_test.go)

**题意**：返回任意合法课程顺序，有环时返回空数组。

**变化点**：与207相同建图，通过模板的 `onPop` 收集出队顺序，即拓扑序。

完整的统一模板版本见[第 6.2 节的 210 示例](#dag-adapters)：`onPop` 收集顺序，`relax` 传 nil，不再单独维护队列。

**例子**：[[1,0],[2,0],[3,1],[3,2]]可返回[0,1,2,3]或[0,2,1,3]。不要排序最终答案，那可能破坏依赖。

**复杂度**：时间O(V+E)，空间O(V+E)。若要求字典序最小，用最小堆选择当前零入度节点。

<a id="lc2050"></a>

## 4. [2050. 并行课程 III](https://leetcode.com/problems/parallel-courses-iii/)

[练习](l2050/2050.go) · [测试](l2050/2050_test.go)

**题意**：课程可无限并行，每门有时长，所有前置完成才能开始，求全部完成的最短时间。relations=[前置,后继]且编号1-based。

**状态**：finish[v]表示最早完成时间，初始time[v]。每条边u→v更新finish[v]=max(finish[v],finish[u]+time[v])，入度归零后v才入队。

完整的统一模板版本见第 6.2 节的 2050 示例：`relax` 更新最早完成时间，`onPop` 更新全局答案，模板统一管理入度和队列。

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

已有分类：[graph/topo_sort](.)。


<a id="dag-dp"></a>

## 6. 六题主线：拓扑排序＋DAG 动态规划

拓扑排序不只是求一个处理顺序，更是在**按照依赖顺序计算每个节点的答案**。

```text
建立边：前置节点 → 后继节点
    ↓
初始化状态，将初始可用节点入队
    ↓
取出 u，此时 u 已经收齐所有前置节点的信息
    ↓
对每条 u → v：
    用 u 的最终状态更新 v 的候选状态
    将 v 的剩余入度减 1
    入度变为 0 时，v 才能入队
```

**核心不变量：节点出队时，它的所有前置节点都已处理完，状态已经最终确定。** 与网格 Flood Fill 相比，这里不是“第一次遇到就扩展”：多前驱节点必须等全部依赖完成。入度归零负责解锁，DP 负责累计信息，二者不能混为一谈。

<a id="dag-template"></a>

### 6.1 共同模板：起点＋后继枚举＋状态传播

这个模板统一处理六题，不强制所有状态都是一个数字。DP 可以是整数、布尔集合或 26 维数组，存放在调用方闭包中；`relax` 负责传播，`onPop` 负责收集已经确定的答案。

```go
// Kahn 返回实际处理的节点数；节点编号为 [0, len(indegree))。
// seeds 是互不重复的初始可用节点，且其入度必须为 0。
// next 必须同步、逐条枚举后继，与 indegree 的计数完全一致。
func Kahn(
    indegree []int,
    seeds []int,
    next func(u int, yield func(v int)),
    onPop func(u int),
    relax func(u, v int),
) int {
    // 不修改调用方入度和初始队列，便于重复调用。
    remaining := append([]int(nil), indegree...)
    queue := append([]int(nil), seeds...)
    for head := 0; head < len(queue); head++ {
        u := queue[head]
        if onPop != nil {
            onPop(u)
        }
        next(u, func(v int) {
            // 更新发生在解锁之前，v 出队时才能收齐所有贡献。
            if relax != nil {
                relax(u, v)
            }
            remaining[v]--
            if remaining[v] == 0 {
                queue = append(queue, v)
            }
        })
    }
    return len(queue)
}

func ZeroIndegree(indegree []int) []int {
    seeds := make([]int, 0)
    for u, degree := range indegree {
        if degree == 0 {
            seeds = append(seeds, u)
        }
    }
    return seeds
}
```

| 参数 | 职责 | 六题中的变化 |
| --- | --- | --- |
| `indegree` | 每个节点尚需满足的前置依赖数 | 显式边计数，或统计更小的相邻格子 |
| `seeds` | 最初就可用的节点 | 普通题为全部零入度节点；2115 仅为 supplies |
| `next` | 同步枚举 u 的所有后继 | 邻接表，或 329 的隐式四邻居 |
| `onPop` | 读取已确定的状态、收集答案 | 收集拓扑序／配方，或更新全局最优值；可为 nil |
| `relax` | 把 u 的最终状态传播给 v | 最大值、集合并、26 维逐项最大值；可为 nil |

为什么不强制写成 `Kahn[S]`？六题状态形状不同，且 210、2115 根本不需要额外 DP。将状态存放在闭包中，模板只负责依赖调度，能避免为统一而增加无用状态。`onPop` 不应破坏即将向后继传播的最终状态。

普通邻接表可适配成下面的 `next`；329 则改成枚举四邻居：

```go
next := func(u int, yield func(int)) {
    for _, v := range graph[u] {
        yield(v)
    }
}
```

**模板的前提**：入度是尚未满足的依赖数；传播操作在全部前驱到齐后能得到最终状态。重复边若保留，入度与 next 都必须计入；若去重，则两边一起去重。seeds 不能重复，尤其不能让同一 supply 被消费两次。模板不检查非法编号或不一致的图，建图方负责这些约束。

**判环的适用范围**：当 seeds 是全图所有零入度节点时，`processed < 节点数` 才等价于图中有环。2115 只用 supplies 作为 seeds，未处理节点也可能仅仅缺少原料，不能据此断言有环。

<a id="dag-adapters"></a>

### 6.2 六题套用与状态转移

以下六个函数都包含完整建图／状态初始化，配合第 6.1 节模板即可使用。它们是不同题目的独立实现；本地代码与测试入口如下。

| 题目 | 题解 | 测试 |
| --- | --- | --- |
| 210 | [210.go](l210/210.go) | [210_test.go](l210/210_test.go) · [模板与边界](l210/template_test.go) |
| 2115 | [2115.go](l2115/2115.go) | [2115_test.go](l2115/2115_test.go) |
| 1462 | [1462.go](l1462/1462.go) | [1462_test.go](l1462/1462_test.go) |
| 2050 | [2050.go](l2050/2050.go) | [2050_test.go](l2050/2050_test.go) · [边界](l2050/boundary_test.go) |
| 1857 | [1857.go](l1857/1857.go) | [1857_test.go](l1857/1857_test.go) |
| 329 | [329.go](l329/329.go) | [329_test.go](l329/329_test.go) |

定向测试（只运行这六题，不依赖父目录的其他练习）：

```sh
go test -count=1 -cover ./graph/topo_sort/l210 ./graph/topo_sort/l2115 ./graph/topo_sort/l1462 ./graph/topo_sort/l2050 ./graph/topo_sort/l1857 ./graph/topo_sort/l329
```

测试包括示例、判环、重复边、输入不变及重复调用；1462 用 200 个随机 DAG 对照 Floyd 可达性，329 用 300 个随机矩阵对照记忆化 DFS。

#### ① [210. 课程表 II](https://leetcode.com/problems/course-schedule-ii/)：收集出队顺序

输入 `[a,b]` 表示 b 是 a 的前置课程，建边 **b → a**。零入度课程全入队，包括孤立课程。不需要 DP，`relax` 传 nil。

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
        indegree[p[0]]++
    }
    order := make([]int, 0, numCourses)
    processed := Kahn(indegree, ZeroIndegree(indegree),
        func(u int, yield func(int)) {
            for _, v := range graph[u] { yield(v) }
        },
        func(u int) { order = append(order, u) }, nil,
    )
    if processed != numCourses { return []int{} }
    return order
}
```

#### ② [2115. 从给定原材料中找到所有可以做出的菜](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/)：可用性传播

将 recipes、ingredients、supplies 中的所有名字映射为整数 ID。对每个配方 r 的每个原料 x 建边 **x → r**，`indegree[r]` 等于所需原料数量。准备 `recipeNames map[int]string`，以识别出队节点是否为配方。

```go
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    ids := map[string]int{}
    intern := func(name string) int {
        if id, ok := ids[name]; ok { return id }
        id := len(ids)
        ids[name] = id
        return id
    }
    for _, name := range recipes { intern(name) }
    for _, list := range ingredients {
        for _, name := range list { intern(name) }
    }
    for _, name := range supplies { intern(name) }
    graph := make([][]int, len(ids))
    indegree := make([]int, len(ids))
    recipeNames := map[int]string{}
    for i, name := range recipes {
        v := ids[name]
        recipeNames[v] = name
        for _, ingredient := range ingredients[i] {
            u := ids[ingredient]
            graph[u] = append(graph[u], v)
            indegree[v]++
        }
    }
    seeds := make([]int, 0)
    seen := make([]bool, len(ids))
    addSeed := func(u int) {
        if !seen[u] { seen[u] = true; seeds = append(seeds, u) }
    }
    // 原题 supplies 不与 recipes 重名，不能加入其他零入度原料。
    for _, name := range supplies { addSeed(ids[name]) }
    // 额外支持空配方；原题没有这种输入。
    for i, name := range recipes {
        if len(ingredients[i]) == 0 { addSeed(ids[name]) }
    }
    answer := make([]string, 0)
    Kahn(indegree, seeds,
        func(u int, yield func(int)) {
            for _, v := range graph[u] { yield(v) }
        },
        func(u int) {
            if name, ok := recipeNames[u]; ok { answer = append(answer, name) }
        }, nil,
    )
    return answer
}
```

缺失原料即使图中入度为 0，也不代表可用。每种可用原料只需发送一次“已满足”事件，不涉及库存消耗；解锁的配方会继续作为其他配方的原料。题设每个配方至少需要一种原料；若扩展为空配方，需显式把这些配方加入初始可用节点。

#### ③ [1462. 课程表 IV](https://leetcode.com/problems/course-schedule-iv/)：传播祖先集合

本题输入 `[a,b]` 表示 **a → b**，与 210 的输入解释不同。定义 `before[v][x]` 表示 x 是 v 的直接或间接前置课程，初始化全 false。

```go
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    before := make([][]bool, numCourses)
    for v := range before { before[v] = make([]bool, numCourses) }
    for _, e := range prerequisites {
        graph[e[0]] = append(graph[e[0]], e[1])
        indegree[e[1]]++
    }
    Kahn(indegree, ZeroIndegree(indegree),
        func(u int, yield func(int)) {
            for _, v := range graph[u] { yield(v) }
        }, nil,
        func(u, v int) {
            before[v][u] = true
            for x := 0; x < numCourses; x++ {
                before[v][x] = before[v][x] || before[u][x]
            }
        },
    )
    answer := make([]bool, len(queries))
    for i, q := range queries { answer[i] = before[q[1]][q[0]] }
    return answer
}
```

沿每条边把 `祖先(u) ∪ {u}` 并入 v；不能只标记直接依赖。本题保证 DAG，这种模板不负责一般有环图的传递闭包；有环时应改用其他可达性算法。

#### ④ [2050. 并行课程 III](https://leetcode.com/problems/parallel-courses-iii/)：DAG 最长路径

relations 的 `[a,b]` 表示 a → b，建图时减一转成 0-based。定义 `finish[v]` 为课程 v 的最早完成时间，初始为自身耗时。

```go
func minimumTime(n int, relations [][]int, time []int) int {
    graph := make([][]int, n)
    indegree := make([]int, n)
    for _, e := range relations {
        u, v := e[0]-1, e[1]-1
        graph[u] = append(graph[u], v)
        indegree[v]++
    }
    finish := make([]int64, n)
    for v, duration := range time { finish[v] = int64(duration) }
    var answer int64
    processed := Kahn(indegree, ZeroIndegree(indegree),
        func(u int, yield func(int)) {
            for _, v := range graph[u] { yield(v) }
        },
        func(u int) {
            if finish[u] > answer { answer = finish[u] }
        },
        func(u, v int) {
            candidate := finish[u] + int64(time[v])
            if candidate > finish[v] { finish[v] = candidate }
        },
    )
    if processed != n { return -1 }
    return int(answer)
}
```

`finish[v] = max(finish[v], finish[u]+time[v])`。v 必须等待**所有**前置完成，因此开始时间取前置完成时间的最大值，而不是最小值或总和。两门前置耗时 3、2，后继耗时 5，答案是 `max(3,2)+5=8`。队列的层数不能替代课程耗时。

#### ⑤ [1857. 有向图中最大颜色值](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/)：26 维路径状态

定义 `dp[v][c]`：所有以 v 为终点的路径中，颜色 c 出现次数的最大值。初始化每个节点自身颜色为 1；其他维度为 0。

```go
func largestPathValue(colors string, edges [][]int) int {
    n := len(colors)
    graph := make([][]int, n)
    indegree := make([]int, n)
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        indegree[e[1]]++
    }
    dp := make([][26]int, n)
    for v := range dp { dp[v][colors[v]-'a'] = 1 }
    answer := 0
    processed := Kahn(indegree, ZeroIndegree(indegree),
        func(u int, yield func(int)) {
            for _, v := range graph[u] { yield(v) }
        },
        func(u int) {
            for _, value := range dp[u] {
                if value > answer { answer = value }
            }
        },
        func(u, v int) {
            for c := 0; c < 26; c++ {
                candidate := dp[u][c]
                if c == int(colors[v]-'a') { candidate++ }
                if candidate > dp[v][c] { dp[v][c] = candidate }
            }
        },
    )
    if processed != n { return -1 }
    return answer
}
```

每一维分别取最大值，不同维度可能来自不同路径，不能把它们相加当作一条路径。也不能把不同前驱的同色计数相加，因为路径不是合并多条分支。当前写法已经在转移时计入 v 的颜色，出队时不要再加一次。即使无环部分已有答案，只要图中任何位置存在环，仍返回 -1。

#### ⑥ [329. 矩阵中的最长递增路径](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)：隐式网格 DAG

把格子编号为 `u = r*cols+c`。从较小格子向相邻的较大格子建边；严格递增保证不可能回到原值，因此无环。无需真的分配邻接表，`next` 按需枚举后继。

下面是可与共同模板组合使用的完整题目函数：

```go
func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    rows, cols := len(matrix), len(matrix[0])
    n := rows * cols
    directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    next := func(u int, yield func(int)) {
        r, c := u/cols, u%cols
        for _, d := range directions {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
                matrix[nr][nc] > matrix[r][c] {
                yield(nr*cols + nc)
            }
        }
    }
    indegree := make([]int, n)
    for u := 0; u < n; u++ {
        next(u, func(v int) { indegree[v]++ })
    }
    length := make([]int, n)
    for u := range length {
        length[u] = 1
    }
    answer := 0
    Kahn(indegree, ZeroIndegree(indegree), next,
        func(u int) {
            if length[u] > answer {
                answer = length[u]
            }
        },
        func(u, v int) {
            if length[u]+1 > length[v] {
                length[v] = length[u]+1
            }
        },
    )
    return answer
}
```

这里的路径长度按格子数计算，初始化为 1；相等值之间不能连边。输入矩阵不修改，不需要路径 visited，也不会像 Flood Fill 那样第一次遇到邻居就确定其答案：必须收齐所有较小前驱的贡献。

<a id="dag-practice"></a>

### 6.3 复杂度、边界与练习顺序

共同模板只控制遍历；总复杂度还要加上每条边的状态传播成本。

| 顺序 | 题目 | 初始队列 | 状态／汇总 | 时间复杂度 |
| --- | --- | --- | --- | --- |
| 1 | 210 课程表 II | 所有零入度课程 | 出队顺序，处理数判环 | O(V+E) |
| 2 | 2115 配方解锁 | 仅给定 supplies | 剩余依赖计数 | O(V+E)，按已映射名字计 |
| 3 | 1462 课程表 IV | 所有零入度课程 | 祖先布尔集合 | O(V²+VE+Q) |
| 4 | 2050 并行课程 III | 所有零入度课程 | 最早完成时间 | O(V+E) |
| 5 | 1857 最大颜色值 | 所有零入度节点 | 26 维路径最大次数 | O(26(V+E)) |
| 6 | 329 最长递增路径 | 没有更小邻居的格子 | 最长路径格子数 | O(mn) |

空间：一般为 O(V+E)；1462 另有 O(V²) 的集合表，1857 另有 O(26V) 状态；329 不存显式边，只需 O(mn)。2115 的字符串映射按通常哈希操作计，若计入字符串读取/哈希，还需考虑总字符量。

重点主线：**210 → 2050 → 1857 → 329**，从“求处理顺序”到“沿顺序计算标量、向量与隐式图状态”。2115 强化初始可用性的语义，1462 强化集合传播。

检查清单：

- 明确每道题输入的边方向，不照搬上一题的数组下标。
- 初始化所有合法起点与状态，包括孤立点；2115 不能无条件加入所有零入度名字。
- 每条边都先更新状态，再减少入度；只在入度恰好归零时入队。
- 全部零入度起点的模式下，最后用处理数判环；存在环时不能返回部分 DP 结果。
- 只有依赖已经无环、前驱状态最终确定，才能直接做这种 DAG DP；不能用它代替一般图最短路或一般可达性算法。
- 模板复制了入度和 seeds，但 DP、答案等回调状态仍由调用方重新初始化。

建议测试：无边、多源、多汇、菱形依赖、孤立节点、环外另有无环分量、自环、重复边一致性；配方缺原料但无环、循环配方、原料解锁多级配方；矩阵全等、单格、单行与蛇形递增。

**一句话：入度决定什么时候可以算，状态转移决定算什么；节点出队时，它已经收齐所有前驱的信息。**

<a id="legacy-notes"></a>

## 原分类笔记（保留）

https://leetcode.com/tag/topological-sort/
https://leetcode.com/discuss/general-discussion/1078072/introduction-to-topological-sort

拓扑排序针对有向图，且只有**有向无环图（DAG）**才存在合法的拓扑序。

TODO
2115. Find All Possible Recipes from Given Supplies

Topo sort:
1 DFS
2 degree/kahn's algorithm

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [2115. 从给定原材料中找到所有可以做出的菜](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/) | 原材料为初始可用源，按缺失依赖数解锁菜品 | 不是所有零入度名字都可用，只能从给定supplies传播 |
| [802. 找到最终的安全状态](https://leetcode.com/problems/find-eventual-safe-states/) | 反图配合原图出度，从终止节点逆向删除 | 结果按编号排序；普通入度为0不是安全状态条件 |
| [1462. 课程表IV](https://leetcode.com/problems/course-schedule-iv/) | 拓扑顺序传播祖先集合，回答先修可达性 | 需要间接依赖；图规模小也可Floyd闭包 |
| [1203. 项目管理](https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/) | 分别构建组间DAG和组内项目DAG，两层排序 | 无组项目分配独立组；组间重复边计度需一致 |
| [1857. 有向图中最大颜色值](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/) | 拓扑DP每点维护26种颜色的路径最大次数 | 有环返回-1；并非统计整个分量颜色总数 |
| [2392. 给定条件下构造矩阵](https://leetcode.com/problems/build-a-matrix-with-conditions/) | 行约束和列约束分别拓扑排序，合并坐标 | 任一约束有环就无解，两个顺序独立 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
