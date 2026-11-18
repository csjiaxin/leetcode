# 博弈 DP：从当前行动者出发，判断胜负与计算分差

## 目录

- [1. 两种模型与共同不变量](#models)
- [2. 统一的记忆化模板](#template)
- [3. 胜负型：1025、1510、464](#win)
- [4. 分差型：486、1406、1690、1140](#difference)
- [5. 状态设计、复杂度与易错点](#pitfalls)
- [6. 练习顺序、题目链接与测试](#practice)

<a id="models"></a>

## 1. 两种模型与共同不变量

双方都采取最优策略，不是模拟一条看起来合理的操作序列，而是枚举当前玩家的选择，并假定对手也会选对自己最有利的操作。

**最重要的不变量：状态始终站在当前行动者的角度。递归一层，玩家身份交换，但状态含义不变。** 在这些规则对双方对称的题目中，通常不必额外记录 Alice/Bob 或轮次。

### 胜负型

```text
win(state) = 存在一种合法选择：
                 能立即获胜
                 或 !win(nextState)

如果没有任何这样的选择，则当前状态必败。
```

没有合法操作通常判负；立即达到目标的规则需要在递归前判定，不能一律把“目标已达到”解释成当前玩家获胜。例如 [464](https://leetcode.com/problems/can-i-win/) 中，选到足以达标的数字时，当前玩家立即赢，不应再把回合交给对手。

### 分差型

```text
diff(state) = 当前玩家最终得分 - 对手最终得分 的最大可能值

diff(state) = max(本次得分 - diff(nextState))
```

为什么减？进入下一状态后，`diff(nextState)` 是对手的优势。本次得 5 分，之后对手领先 2 分，这次选择给当前玩家带来的最终分差是 `5-2=3`。

没有剩余操作时分差为 0。`diff` 可以为负数，意味着最优策略下仍然会输；它不是“当前玩家最多能拿多少分”。

这些题的状态转移均单向推进：剩余数量下降、区间缩短、后缀起点增大，或者 mask 中已选位增多。因此状态图无环，可把递归搜索变成 DAG 上的 DP。

<a id="template"></a>

## 2. 统一的记忆化模板

七道题共用下面的 `MemoGame[S, V]`，需要 Go 1.18+。模板只负责缓存；终止条件、合法操作与当前行动者视角由回调负责。

```go
// MemoGame 仅缓存无环状态的求值；每次调用创建独立缓存。
func MemoGame[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
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

- `S` 是完整状态，可以是 int、mask 或只含可比较字段的结构体。
- `V` 是返回值：胜负型用 bool，分差型用 int。
- `evaluate(state, solve)` 通过传入的 `solve(next)` 递归，确保子状态也使用同一缓存。
- 使用 map 的 `ok` 判断是否计算过：false、0、负数都可能是合法答案，不能拿它们当“未计算”。
- 每个题解函数内部创建新的求解器，输入数组、目标值等固定规则由闭包捕获，避免跨题目调用污染缓存。
- 它不解决有环博弈，不提供和局环检测，也不是并发安全的缓存。若转移可能回到尚在求值的状态，仅靠这个模板会无限递归。

为什么不强行把所有题写成同一个 `max`？胜负型需要“存在必败后继”并可短路，分差型要比较所有候选值；统一的是记忆化框架，而不是抹平它们的游戏规则。

<a id="win"></a>

## 3. 胜负型：[1025](https://leetcode.com/problems/divisor-game/)、[1510](https://leetcode.com/problems/stone-game-iv/)、[464](https://leetcode.com/problems/can-i-win/)

### [1025](https://leetcode.com/problems/divisor-game/)：剩余数字作为状态

```go
func divisorGame(n int) bool {
    win := MemoGame(func(remaining int, solve func(int) bool) bool {
        for x := 1; x < remaining; x++ {
            if remaining%x == 0 && !solve(remaining-x) {
                return true
            }
        }
        return false
    })
    return win(n)
}
```

`n=1` 没有合法除数，当前玩家失败；只要能把对方送入必败状态就立即返回 true。原题还可以证明“偶数必胜，奇数必败”，这里保留搜索版本训练模型，测试用奇偶结论独立对照。

### [1510](https://leetcode.com/problems/stone-game-iv/)：枚举取走的平方数

```go
func winnerSquareGame(n int) bool {
    win := MemoGame(func(remaining int, solve func(int) bool) bool {
        for root := 1; root <= remaining/root; root++ {
            if !solve(remaining-root*root) {
                return true
            }
        }
        return false
    })
    return win(n)
}
```

`remaining=0` 时循环不执行，返回 false；只剩一个平方数时，可以一次取完，把对手送到 0。

### [464](https://leetcode.com/problems/can-i-win/)：mask 记录已经选过的数字

选择 1～maxChoosableInteger 中尚未使用的数字，累计和达到 desiredTotal 者立即获胜。

- 位 `x-1` 为 1 表示数字 x 已使用。
- 同一个 mask 唯一决定已用数字之和，所以目标剩余值可以重算，不必重复放入缓存键。
- `desiredTotal <= 0` 时先手直接满足目标；所有数字总和仍不足目标时直接返回 false，否则“无法达标”会被误当成普通的最后一步取胜。
- 每个可选数字 x：若 x 已足够达标，或选择后让对手必败，则当前状态必胜。

完整实现见 [464.go](../l464/464.go)。虽然它用了位掩码，但仍然是“集合状态＋记忆化胜负搜索”，不是靠已选数字的排列顺序建状态。

<a id="difference"></a>

## 4. 分差型：[486](https://leetcode.com/problems/predict-the-winner/)、[1406](https://leetcode.com/problems/stone-game-iii/)、[1690](https://leetcode.com/problems/stone-game-vii/)、[1140](https://leetcode.com/problems/stone-game-ii/)

### [486](https://leetcode.com/problems/predict-the-winner/)：从区间两端拿，定义区间分差

```go
func predictTheWinner(nums []int) bool {
    type interval struct{ left, right int }
    diff := MemoGame(func(s interval, solve func(interval) int) int {
        if s.left > s.right { return 0 }
        left := nums[s.left] - solve(interval{s.left+1, s.right})
        right := nums[s.right] - solve(interval{s.left, s.right-1})
        if left > right { return left }
        return right
    })
    return diff(interval{0, len(nums)-1}) >= 0
}
```

平局算先手赢，因此判断 `>= 0` 而不是 `> 0`。先算好左右候选再比较，不依赖全局轮次。自底向上时按区间长度从小到大计算，就是区间 DP。

### [1406](https://leetcode.com/problems/stone-game-iii/)：每次取前面 1～3 个，定义后缀分差

```go
func stoneGameIII(stoneValue []int) string {
    diff := MemoGame(func(start int, solve func(int) int) int {
        if start == len(stoneValue) { return 0 }
        taken := stoneValue[start]
        best := taken - solve(start+1)
        for count := 2; count <= 3 && start+count <= len(stoneValue); count++ {
            taken += stoneValue[start+count-1]
            candidate := taken - solve(start+count)
            if candidate > best { best = candidate }
        }
        return best
    })
    result := diff(0)
    if result > 0 { return "Alice" }
    if result < 0 { return "Bob" }
    return "Tie"
}
```

状态 `start` 表示剩余后缀，`taken` 是本次取走的总分。存在负数，不能把 best 初始化为 0 并允许“不取”；必须从至少一个合法选择初始化。自底向上时，start 从 n-1 递减到 0。

### [1690](https://leetcode.com/problems/stone-game-vii/)：区间分差，但本次得分是剩余总和

先构造 `prefix[i+1] = prefix[i] + stones[i]`，区间总和为 `prefix[r+1]-prefix[l]`。

```text
移除左端：剩余总和 = prefix[r+1] - prefix[l+1]
候选分差 = 剩余总和 - diff(l+1, r)

移除右端：剩余总和 = prefix[r] - prefix[l]
候选分差 = 剩余总和 - diff(l, r-1)
```

对两种选择取最大。单个石子移除后剩余总和为 0，所以 `l >= r` 返回 0。这与 [486](https://leetcode.com/problems/predict-the-winner/) 的结构相同，区别只在“本次得分”，不能误写成被移除石子的值。完整模板调用见 [1690.go](../l1690/1690.go)。

### [1140](https://leetcode.com/problems/stone-game-ii/)：把 M 放入状态，最后从分差还原得分

状态为 `(start, M)`，本次可取 `1 <= X <= min(2M, n-start)`，下一状态是 `(start+X, max(M,X))`。

```text
本次得分 = suffix[start] - suffix[start+X]
diff(start,M) = max(本次得分 - diff(start+X, max(M,X)))
```

相同 start、不同 M 的可选操作不同，因此不能只用 start 缓存。原题每堆石子数均为正数：若 `2M >= n-start`，可直接全拿，分差为剩余总和。

题目问 Alice 拿到的数量，而不是分差。总和 T 固定，令分差 D=A-B，又有 T=A+B，因此：

```text
Alice = (T + D) / 2
```

完整实现见 [1140.go](../l1140/1140.go)。这个还原公式适用于所有分数都归属某一方的固定总分游戏，如 [486](https://leetcode.com/problems/predict-the-winner/)、[1406](https://leetcode.com/problems/stone-game-iii/)、[1140](https://leetcode.com/problems/stone-game-ii/)；不能直接拿石子原始总和去还原 [1690](https://leetcode.com/problems/stone-game-vii/) 的分数，因为 [1690](https://leetcode.com/problems/stone-game-vii/) 会重复对剩余石子计分。

<a id="pitfalls"></a>

## 5. 状态设计、复杂度与易错点

| 题目 | 状态 S | 返回值 V | 终止／答案 | 当前实现最坏时间 | 辅助空间 |
| --- | --- | --- | --- | --- | --- |
| [1025](https://leetcode.com/problems/divisor-game/) | 剩余 n | bool | 无合法操作为 false | O(n²) | O(n) |
| [1510](https://leetcode.com/problems/stone-game-iv/) | 剩余 n | bool | 0 为 false | O(n√n) | O(n) |
| [464](https://leetcode.com/problems/can-i-win/) | 已用数字 mask | bool | 能达标则立即赢 | O(k·2^k) | O(2^k) |
| [486](https://leetcode.com/problems/predict-the-winner/) | [l,r] | int 分差 | 空区间为 0；最终 >=0 | O(n²) | O(n²) |
| [1406](https://leetcode.com/problems/stone-game-iii/) | 后缀起点 i | int 分差 | i=n 为 0；按正负映射名字 | O(n) | O(n) |
| [1690](https://leetcode.com/problems/stone-game-vii/) | [l,r] | int 分差 | l>=r 为 0 | O(n²) | O(n²) |
| [1140](https://leetcode.com/problems/stone-game-ii/) | (i,M) | int 分差 | i=n 为 0；最终还原先手分数 | O(n³) | O(n²) |

按哈希表平均 O(1) 操作分析。[464](https://leetcode.com/problems/can-i-win/) 中 k 是最大可选数字，枚举已用数字之和及下一步均为 O(k)，不是 O(k²)；[1140](https://leetcode.com/problems/stone-game-ii/) 有 O(n²) 个状态，每个最多枚举 O(n) 步，提前全拿能剪枝，但不把上述上界直接写成 O(n²)。空间包括记忆化表和递归栈；递归深度可能为 O(n)，[464](https://leetcode.com/problems/can-i-win/) 为 O(k)。

### 检查清单

1. **当前行动者视角统一**：子状态返回的是对手的结果，所以胜负取反、分差相减。
2. **缓存键包含所有影响未来选择的信息**：[1140](https://leetcode.com/problems/stone-game-ii/) 必须有 M；[464](https://leetcode.com/problems/can-i-win/) 不需要重复记录由 mask 决定的累计和。
3. **没有额外的 visited 回溯**：缓存永久保留已求值状态；搜索所有分支不等于共享路径枚举。
4. **别把双方收益加起来**：最大化的是当前玩家相对优势，而不是合作的总得分。
5. **注意终止语义与平局规则**：[1025](https://leetcode.com/problems/divisor-game/)、[1510](https://leetcode.com/problems/stone-game-iv/) 无法操作判负；[486](https://leetcode.com/problems/predict-the-winner/) 平局算先手赢；[1406](https://leetcode.com/problems/stone-game-iii/) 平局返回 Tie。
6. **不能跳过负收益操作**：[1406](https://leetcode.com/problems/stone-game-iii/) 即使拿取值为负也必须行动。
7. **区分得分定义**：[486](https://leetcode.com/problems/predict-the-winner/) 是拿走的值；[1690](https://leetcode.com/problems/stone-game-vii/) 是移除后剩余总和；[1140](https://leetcode.com/problems/stone-game-ii/) 是拿走的堆总和。
8. **胜负型可短路，分差型通常需比较所有操作**：胜负找到一个能赢的选择就足够；不要把这一点误用成只比较分差的第一个候选。
9. **不修改输入**：所有题解只读数组；缓存只在当前题解调用内使用。若求解器捕获的规则或数组发生变化，应重新创建求解器。
10. **深度与规模**：递归便于理解；[1510](https://leetcode.com/problems/stone-game-iv/)、[1406](https://leetcode.com/problems/stone-game-iii/) 等大规模一维状态可改为自底向上，避免深递归与 map 常数。无环状态才适用本模板。

### 串起已有知识

- 记忆化搜索：把重复局面合并，每个状态只求值一次。
- 区间 DP：[486](https://leetcode.com/problems/predict-the-winner/)、[1690](https://leetcode.com/problems/stone-game-vii/) 的依赖是更短区间，可按长度递增填表。
- 后缀 DP：[1406](https://leetcode.com/problems/stone-game-iii/) 的依赖是更大起点，可从右向左填表。
- 状态压缩 DP：[464](https://leetcode.com/problems/can-i-win/) 把已使用集合编码成 mask；不同选择顺序到达同一集合，共享缓存。
- DAG DP：这些状态单调变化，可递归按需计算，也可找出合法的自底向上顺序。

<a id="practice"></a>

## 6. 练习顺序、题目链接与测试

| 顺序 | LeetCode | 本地题解 | 测试 | 核心训练点 |
| --- | --- | --- | --- | --- |
| 1 | [1025. 除数博弈](https://leetcode.com/problems/divisor-game/) | [1025.go](../l1025/1025.go) | [测试](../l1025/1025_test.go) | 根据后继判断必胜必败 |
| 2 | [1510. 石子游戏 IV](https://leetcode.com/problems/stone-game-iv/) | [1510.go](../l1510/1510.go) | [测试](../l1510/1510_test.go) | 枚举平方数，寻找必败后继 |
| 3 | [464. 我能赢吗](https://leetcode.com/problems/can-i-win/) | [464.go](../l464/464.go) | [测试](../l464/464_test.go) | mask 集合状态＋记忆化 |
| 4 | [486. 预测赢家](https://leetcode.com/problems/predict-the-winner/) | [486.go](../l486/486.go) | [测试](../l486/486_test.go) | 当前玩家的区间最优分差 |
| 5 | [1406. 石子游戏 III](https://leetcode.com/problems/stone-game-iii/) | [1406.go](../l1406/1406.go) | [测试](../l1406/1406_test.go) | 取前 1～3 个，后缀分差 |
| 6 | [1690. 石子游戏 VII](https://leetcode.com/problems/stone-game-vii/) | [1690.go](../l1690/1690.go) | [测试](../l1690/1690_test.go) | 剩余区间总和计分 |
| 7 | [1140. 石子游戏 II](https://leetcode.com/problems/stone-game-ii/) | [1140.go](../l1140/1140.go) | [测试](../l1140/1140_test.go) | 随操作变化的范围 M |

建议先抓主线：**[1025](https://leetcode.com/problems/divisor-game/) → [486](https://leetcode.com/problems/predict-the-winner/) → [1406](https://leetcode.com/problems/stone-game-iii/) → [464](https://leetcode.com/problems/can-i-win/)**，依次掌握胜负判断、区间分差、后缀分差和集合博弈，再补齐 [1510](https://leetcode.com/problems/stone-game-iv/)、[1690](https://leetcode.com/problems/stone-game-vii/)、[1140](https://leetcode.com/problems/stone-game-ii/) 的规则变化。

```text
dp/game/
├── readme.md
├── l1025/  (1025.go, 1025_test.go)
├── l1510/  (1510.go, 1510_test.go)
├── l464/   (464.go, 464_test.go)
├── l486/   (486.go, 486_test.go)
├── l1406/  (1406.go, 1406_test.go)
├── l1690/  (1690.go, 1690_test.go)
└── l1140/  (1140.go, 1140_test.go)
```

每题均保留相同的 MemoGame 模板，便于独立阅读、练习和提交；文档中的示例也统一使用该模板。

```sh
go test -count=1 -cover ./dp/game/...
```

测试内容：官方示例、终止状态、平局、负数、输入不变与重复调用；[1025](https://leetcode.com/problems/divisor-game/) 对照奇偶结论，[1510](https://leetcode.com/problems/stone-game-iv/) 对照自底向上胜负表，[464](https://leetcode.com/problems/can-i-win/) 小规模枚举对照，四道分差题分别用显式 Alice/Bob 极大极小搜索做固定种子随机对照。另含 [1510](https://leetcode.com/problems/stone-game-iv/) 的 100000 石子、[1406](https://leetcode.com/problems/stone-game-iii/) 的 50000 项输入，以及缓存对 false/0 值的正确复用检查。

**一句话：永远问“轮到我时，这个局面对我意味着什么”；能让对手输就赢，能取得的分差就是本次得分减去对手之后的优势。**
