# 博弈 DP：缓存状态，始终说明当前行动者

[返回 DP 总目录](../README.md)

共享的是无环状态求值骨架，不是所有游戏共用同一个收益公式。MemoGame 的 map 用 ok 区分未计算与合法的 false/0，每次构造新缓存，不能自动处理有环博弈或平局。

| 模型 | 适配器必须定义的内容 | 题目 |
| --- | --- | --- |
| 必胜性 | 存在一步能使对手必败；终态规则不可省略 | [1025](https://leetcode.com/problems/divisor-game/)、[1510](https://leetcode.com/problems/stone-game-iv/)、[464](https://leetcode.com/problems/can-i-win/) |
| 当前玩家最终分差 | 本步得分减去下一状态的对手分差；初值允许为负 | [486](https://leetcode.com/problems/predict-the-winner/)、[1406](https://leetcode.com/problems/stone-game-iii/)、[1690](https://leetcode.com/problems/stone-game-vii/)、[1140](https://leetcode.com/problems/stone-game-ii/) |
| IntervalTable 后序填表 | 保留相同区间含义，以长度替代递归调度 | [1690](https://leetcode.com/problems/stone-game-vii/) 区间变体、[486](https://leetcode.com/problems/predict-the-winner/) 双方收益变体 |

例如 [486](https://leetcode.com/problems/predict-the-winner/) 拿走端点得到的是该端点数值；[1690](https://leetcode.com/problems/stone-game-vii/) 拿走端点得到的却是剩余石子总和，两者虽然都是区间分差，得分项不同。[1140](https://leetcode.com/problems/stone-game-ii/) 下面使用分差 D，所以先手实得分数为 `(总分+D)/2`；不能把 D 直接当作先手总分。

[464](https://leetcode.com/problems/can-i-win/) 的入口 desiredTotal<=0 为真；递归中一手达到目标是当前玩家直接获胜，不把它当成轮到对手的同一个入口特判。状态须包含完整限制：used mask、剩余区间或后缀起点与 M。

原 stone 的 [1690](https://leetcode.com/problems/stone-game-vii/) panic 练习已经按同一模板补全，原测试保留。递归版与区间迭代、双方收益版仍作为独立方法展示。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### MemoGame

<!-- template: MemoGame -->
```go
// MemoGame 为 evaluate 的结果添加缓存；每次调用创建独立缓存。
// evaluate 如需递归，应调用外部声明的缓存函数；状态转移必须无环。
func MemoGame[S comparable, V any](evaluate func(S) V) func(S) V {
	memo := make(map[S]V)
	return func(state S) V {
		if value, ok := memo[state]; ok {
			return value
		}
		value := evaluate(state)
		memo[state] = value
		return value
	}
}
```

递归函数先在外部声明，例如 `var win func(int) bool`，再赋值为 `win = MemoGame(...)`；分差函数 `diff` 同理。`evaluate` 通过闭包调用外部的 `win` 或 `diff`，递归必须经过这个带缓存的函数，不能直接调用原始 `evaluate`。完成赋值后再调用入口状态。

### IntervalTable

<!-- template: IntervalTable -->
```go
// IntervalTable 按闭区间长度递增；step 只能读取严格更短区间。
// base 初始化单点；空区间不在表中，由适配器自行处理。
func IntervalTable[T any](n int, base func(i int) T, step func(l, r int, dp [][]T) T) [][]T {
	dp := make([][]T, n)
	for i := range dp {
		dp[i] = make([]T, n)
		dp[i][i] = base(i)
	}
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			dp[l][r] = step(l, r, dp)
		}
	}
	return dp
}
```

## 逐题适配器

### [1025](https://leetcode.com/problems/divisor-game/)

[代码](l1025/1025.go) · [测试](l1025/1025_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1025/1025.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/divisor-game/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func divisorGame(n int) bool {
	var win func(int) bool
	win = MemoGame(func(remaining int) bool {
		for x := 1; x < remaining; x++ {
			if remaining%x == 0 && !win(remaining-x) {
				return true
			}
		}
		return false // 没有能让对手输的合法动作，当前玩家必败。
	})
	return win(n)
}
```

### [1510](https://leetcode.com/problems/stone-game-iv/)

[代码](l1510/1510.go) · [测试](l1510/1510_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1510/1510.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/stone-game-iv/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func winnerSquareGame(n int) bool {
	var win func(int) bool
	win = MemoGame(func(remaining int) bool {
		for root := 1; root <= remaining/root; root++ {
			if !win(remaining - root*root) {
				return true
			}
		}
		return false
	})
	return win(n)
}
```

### [464](https://leetcode.com/problems/can-i-win/)

[代码](l464/464.go) · [测试](l464/464_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l464/464.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/can-i-win/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	var win func(int) bool
	win = MemoGame(func(mask int) bool {
		// 已选数字之和由 mask 唯一决定，不必加入缓存键。
		remaining := desiredTotal
		for x := 1; x <= maxChoosableInteger; x++ {
			if mask&(1<<(x-1)) != 0 {
				remaining -= x
			}
		}
		for x := maxChoosableInteger; x >= 1; x-- {
			bit := 1 << (x - 1)
			if mask&bit == 0 && (x >= remaining || !win(mask|bit)) {
				return true
			}
		}
		return false
	})
	return win(0)
}
```

### [486](https://leetcode.com/problems/predict-the-winner/)

[代码](l486/486.go) · [测试](l486/486_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l486/486.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/predict-the-winner/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func predictTheWinner(nums []int) bool {
	type interval struct{ left, right int }
	var diff func(interval) int
	diff = MemoGame(func(s interval) int {
		if s.left > s.right {
			return 0
		}
		left := nums[s.left] - diff(interval{s.left + 1, s.right})
		right := nums[s.right] - diff(interval{s.left, s.right - 1})
		if left > right {
			return left
		}
		return right
	})
	return diff(interval{0, len(nums) - 1}) >= 0 // 平局也算先手赢。
}
```

### [1406](https://leetcode.com/problems/stone-game-iii/)

[代码](l1406/1406.go) · [测试](l1406/1406_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1406/1406.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/stone-game-iii/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func stoneGameIII(stoneValue []int) string {
	var diff func(int) int
	diff = MemoGame(func(start int) int {
		if start == len(stoneValue) {
			return 0
		}
		// 不能用 0 初始化：必须取石子，即使所有选择都亏分。
		taken := stoneValue[start]
		best := taken - diff(start+1)
		for count := 2; count <= 3 && start+count <= len(stoneValue); count++ {
			taken += stoneValue[start+count-1]
			candidate := taken - diff(start+count)
			if candidate > best {
				best = candidate
			}
		}
		return best
	})
	result := diff(0)
	if result > 0 {
		return "Alice"
	}
	if result < 0 {
		return "Bob"
	}
	return "Tie"
}
```

### [1690](https://leetcode.com/problems/stone-game-vii/)

[代码](l1690/1690.go) · [测试](l1690/1690_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1690/1690.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/stone-game-vii/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func stoneGameVII(stones []int) int {
	prefix := make([]int, len(stones)+1)
	for i, value := range stones {
		prefix[i+1] = prefix[i] + value
	}
	type interval struct{ left, right int }
	var diff func(interval) int
	diff = MemoGame(func(s interval) int {
		if s.left >= s.right {
			return 0
		}
		// 得到的是移除后剩余区间总和，而不是被移除石子的值。
		left := prefix[s.right+1] - prefix[s.left+1] - diff(interval{s.left + 1, s.right})
		right := prefix[s.right] - prefix[s.left] - diff(interval{s.left, s.right - 1})
		if left > right {
			return left
		}
		return right
	})
	return diff(interval{0, len(stones) - 1})
}
```

### [1140](https://leetcode.com/problems/stone-game-ii/)

[代码](l1140/1140.go) · [测试](l1140/1140_test.go)

状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1140/1140.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/stone-game-ii/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}
	type state struct{ start, m int }
	var diff func(state) int
	diff = MemoGame(func(s state) int {
		if s.start == n {
			return 0
		}
		// 原题全为正数，能全拿时直接结束。
		if 2*s.m >= n-s.start {
			return suffix[s.start]
		}
		best, initialized := 0, false
		for x := 1; x <= 2*s.m && s.start+x <= n; x++ {
			nextM := s.m
			if x > nextM {
				nextM = x
			}
			taken := suffix[s.start] - suffix[s.start+x]
			candidate := taken - diff(state{s.start + x, nextM})
			if !initialized || candidate > best {
				best, initialized = candidate, true
			}
		}
		return best
	})
	// 总分固定：先手分数 = (总分 + 先手分差) / 2。
	return (suffix[0] + diff(state{0, 1})) / 2
}
```

### [1690](https://leetcode.com/problems/stone-game-vii/) · 历史独立版本

[代码](l1690/legacy/from_stone/1690.go) · [测试](l1690/legacy/from_stone/1690_test.go)

与主版本相同的当前玩家分差状态，补全原 panic 练习。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1690/legacy/from_stone/1690.go; templates: MemoGame -->
```go
// https://leetcode.com/problems/stone-game-vii/
// 与主版本相同的当前玩家分差状态，补全原 panic 练习。
func stoneGameVII(stones []int) int {
	prefix := make([]int, len(stones)+1)
	for i, value := range stones {
		prefix[i+1] = prefix[i] + value
	}
	type interval struct{ left, right int }
	var diff func(interval) int
	diff = MemoGame(func(s interval) int {
		if s.left >= s.right {
			return 0
		}
		// 得到的是移除后剩余区间总和，而不是被移除石子的值。
		left := prefix[s.right+1] - prefix[s.left+1] - diff(interval{s.left + 1, s.right})
		right := prefix[s.right] - prefix[s.left] - diff(interval{s.left, s.right - 1})
		if left > right {
			return left
		}
		return right
	})
	return diff(interval{0, len(stones) - 1})
}
```

### [1140](https://leetcode.com/problems/stone-game-ii/) · 历史独立版本

[代码](l1140/legacy/from_stone/1140.go)

历史记忆化版本保留为独立包，统一使用同一无环求值骨架。

需要模板：`MemoGame`。

<!-- solution: dp/game/l1140/legacy/from_stone/1140.go; templates: MemoGame -->
```go
// 历史记忆化版本保留为独立包，统一使用同一无环求值骨架。
func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}
	type state struct{ start, m int }
	var diff func(state) int
	diff = MemoGame(func(s state) int {
		if s.start == n {
			return 0
		}
		// 原题全为正数，能全拿时直接结束。
		if 2*s.m >= n-s.start {
			return suffix[s.start]
		}
		best, initialized := 0, false
		for x := 1; x <= 2*s.m && s.start+x <= n; x++ {
			nextM := s.m
			if x > nextM {
				nextM = x
			}
			taken := suffix[s.start] - suffix[s.start+x]
			candidate := taken - diff(state{s.start + x, nextM})
			if !initialized || candidate > best {
				best, initialized = candidate, true
			}
		}
		return best
	})
	// 总分固定：先手分数 = (总分 + 先手分差) / 2。
	return (suffix[0] + diff(state{0, 1})) / 2
}
```

### [1690](https://leetcode.com/problems/stone-game-vii/) · interval 变体

[代码](l1690/variants/interval/1690.go) · [测试](l1690/variants/interval/1690_test.go)

区间迭代版：移除端点后取得的是剩余石子总和，减去对手分差。

需要模板：`IntervalTable`。

<!-- solution: dp/game/l1690/variants/interval/1690.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/stone-game-vii/
// 区间迭代版：移除端点后取得的是剩余石子总和，减去对手分差。
func stoneGameVII(stones []int) int {
	n := len(stones)
	if n < 2 {
		return 0
	}
	prefix := make([]int, n+1)
	for i, v := range stones {
		prefix[i+1] = prefix[i] + v
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		return max(prefix[r+1]-prefix[l+1]-dp[l+1][r], prefix[r]-prefix[l]-dp[l][r-1])
	})
	return dp[0][n-1]
}
```

### [486](https://leetcode.com/problems/predict-the-winner/) · score_pair 变体

[代码](l486/variants/score_pair/486.go) · [测试](l486/variants/score_pair/486_test.go)

双方收益版本保留 E 状态；子区间当前行动者变为对手，因此角色互换。

需要模板：`IntervalTable`。

<!-- solution: dp/game/l486/variants/score_pair/486.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/predict-the-winner/
// 双方收益版本保留 E 状态；子区间当前行动者变为对手，因此角色互换。
type E struct{ firstMax, secondMax int }

func predictTheWinner(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	dp := IntervalTable(len(nums), func(i int) E { return E{nums[i], 0} }, func(l, r int, dp [][]E) E {
		left, right := nums[l]+dp[l+1][r].secondMax, nums[r]+dp[l][r-1].secondMax
		if left >= right {
			return E{left, dp[l+1][r].firstMax}
		}
		return E{right, dp[l][r-1].firstMax}
	})
	result := dp[0][len(nums)-1]
	return result.firstMax >= result.secondMax
}
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

```sh
go test ./dp/game/...
```
