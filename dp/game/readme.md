# 博弈 DP：始终站在当前行动者视角

本目录合并原 stone 与 game，并收纳石子游戏的区间实现。共享的思想是状态求值和当前行动者视角，终止条件与收益定义仍由每题决定。

## 1. 无环状态的记忆化骨架

~~~go
func MemoGame[S comparable, V any](
    evaluate func(S, func(S) V) V,
) func(S) V {
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
~~~

缓存存在性用 map 的 ok 判断，不能把 false 或 0 当成未计算。这个骨架要求递归状态有严格推进依据；不能用它自动求解有环博弈或平局状态。

## 2. 胜负与分差是两种状态

### 必胜性

~~~text
win(state) = 存在一步合法选择，使对手进入必败状态
~~~

对于无合法操作者输的规则，终态为 false；其他规则必须重新判断。

- [1025](l1025/1025.go)：剩余整数；可先推递归，再观察奇偶结论。
- [1510](l1510/1510.go)：剩余石子数，枚举平方数。
- [464](l464/464.go)：已使用数字的 mask。入口目标已满足时为 true；总和不足目标时为 false。递归中轮到对手、而上一手已达到目标的情况不能机械套入口判断。

### 当前玩家的最终分差

~~~text
difference(state) = max(本步获得的分数 - difference(nextState))
~~~

进入下一状态后行动者变成对手，因此减去下一状态的分差。分差可以为负；初始化不能一律为 0。

| 题目 | 状态与实际本步得分 |
| --- | --- |
| [486](l486/486.go) | 剩余区间；取走端点的数值，最终分差非负即先手不败 |
| [1406](l1406/1406.go) | 后缀起点；取前 1～3 个的总分，原数值允许负数 |
| [1690](l1690/1690.go) | 剩余区间；得分是移除后剩余石子的总和，不是移除值 |
| [1140](l1140/1140.go) | 后缀起点与 M，合法选择数受 2M 限制 |

1140 若使用“当前玩家最多可得石子数 F”，转移为 suffixSum[i]-F(next)；若使用分差 D，则 Alice 的数量为 (total+D(start))/2。这两种状态含义不能混用。

## 3. 变体与历史材料

- [486 分别保存双方收益](l486/variants/score_pair/486.go)：保留原 stone 实现及原测试。
- [1690 区间迭代版](l1690/variants/interval/1690.go)：保留原 Interval 实现及原测试。
- [1690 原 stone 草稿](l1690/legacy/from_stone/1690.go)：仍为 panic("not implemented")，其测试照常运行并失败，不会被隐藏。
- [1140 历史记忆化版](l1140/legacy/from_stone/1140.go)、[877 历史片段](l877/877)均保留。
- [原 game 完整学习笔记](legacy/game.readme.md)与[原 stone 笔记](legacy/stone.readme.md)继续可查。

## 套题检查

1. 当前状态轮到谁？状态值属于谁？
2. 合法动作是否由当前剩余区间、M 或 mask 完整决定？
3. 下一状态为何更小，能否保证无环？
4. 返回的是胜负、分差还是当前玩家总得分？
5. 缓存维度与每个状态的动作数是多少？用了记忆化不代表自动是 O(n)。

目录合并仅保留和组织已有方法，没有用统一转移覆盖所有历史实现。
