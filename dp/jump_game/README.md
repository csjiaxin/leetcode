# 跳跃题族：连续可达区间、BFS 与下降 DAG

[返回 DP 总目录](../README.md)

先定义边和目标，再选择算法。题名相同不意味着同一递推；三种模板分别维护不同不变量。

| 子模板 | 适用条件与不变量 | 题目 |
| --- | --- | --- |
| RangeJumps | 非负向前跳，当前层可达位置构成连续区间；扫描整层更新下一层最远边界 | 55、45 |
| JumpBFS | 边代价都为 1，完整位置入队即标记；首次到达得到最短步数 | 1306、1345 |
| MemoJump | 只能跳向严格更低位置，状态依赖无环；缓存从当前位置出发的最大访问数 | 1340 |

45 的一次加步发生在扫描完当前层，而不是每一步贪心落到最远位置；额外支持不可达时检测边界无法推进。1306 可双向且目标是值为零的位置，不能用最远边界。1345 同值组展开一次后删除，使所有组的总扫描量为 O(n)。

1340 需要分别向左右扫描，遇到不低的柱子就停止该方向，更远的落点也被挡住；返回的访问数包含起点。最坏 O(nd) 时间、O(n) 状态空间。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### RangeJumps

<!-- template: RangeJumps -->
```go
// RangeJumps 把连续可达区间作为 BFS 层；无法推进时返回 -1。
func RangeJumps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	end, farthest, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i > farthest {
			return -1
		}
		farthest = max(farthest, i+nums[i])
		if i == end {
			if farthest == end {
				return -1
			}
			end = farthest
			steps++
			if end >= len(nums)-1 {
				return steps
			}
		}
	}
	return -1
}
```

### JumpBFS

<!-- template: JumpBFS -->
```go
// JumpBFS：所有边代价为 1，入队即标记；next 同步枚举候选，模板过滤越界下标。
func JumpBFS(n, start int, goal func(int) bool, next func(int, func(int))) int {
	if start < 0 || start >= n {
		return -1
	}
	distance := make([]int, n)
	for i := range distance {
		distance[i] = -1
	}
	queue := []int{start}
	distance[start] = 0
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		if goal(u) {
			return distance[u]
		}
		next(u, func(v int) {
			if v >= 0 && v < n && distance[v] < 0 {
				distance[v] = distance[u] + 1
				queue = append(queue, v)
			}
		})
	}
	return -1
}
```

### MemoJump

<!-- template: MemoJump -->
```go
// MemoJump 仅适用于依赖无环的状态；缓存的零值也是合法答案。
func MemoJump[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(s S) V {
		if value, ok := memo[s]; ok {
			return value
		}
		value := evaluate(s, solve)
		memo[s] = value
		return value
	}
	return solve
}
```

## 逐题适配器

### 55

[代码](l55/55.go) · [测试](l55/55_test.go)

按连续可达区间推进，只关心是否存在到达最后位置的层。

需要模板：`RangeJumps`。

<!-- solution: dp/jump_game/l55/55.go; templates: RangeJumps -->
```go
// https://leetcode.com/problems/jump-game/
// 按连续可达区间推进，只关心是否存在到达最后位置的层。
func canJump(nums []int) bool { return RangeJumps(nums) >= 0 }
```

### 45

[代码](l45/45.go) · [测试](l45/45_test.go)

层边界扩展一次计一跳；额外支持不可达输入返回 -1。

需要模板：`RangeJumps`。

<!-- solution: dp/jump_game/l45/45.go; templates: RangeJumps -->
```go
// https://leetcode.com/problems/jump-game-ii/
// 层边界扩展一次计一跳；额外支持不可达输入返回 -1。
func jump(nums []int) int { return RangeJumps(nums) }
```

### 1306

[代码](l1306/1306.go) · [测试](l1306/1306_test.go)

目标是值为零的节点，邻居为左右跳转，不是最大下标。

需要模板：`JumpBFS`。

<!-- solution: dp/jump_game/l1306/1306.go; templates: JumpBFS -->
```go
// https://leetcode.com/problems/jump-game-iii/
// 目标是值为零的节点，邻居为左右跳转，不是最大下标。
func canReach(arr []int, start int) bool {
	return JumpBFS(len(arr), start, func(i int) bool { return arr[i] == 0 }, func(i int, yield func(int)) {
		yield(i - arr[i])
		yield(i + arr[i])
	}) >= 0
}
```

### 1345

[代码](l1345/1345.go) · [测试](l1345/1345_test.go)

同值组只展开一次，总扫描量 O(n)，避免退化到平方复杂度。

需要模板：`JumpBFS`。

<!-- solution: dp/jump_game/l1345/1345.go; templates: JumpBFS -->
```go
// https://leetcode.com/problems/jump-game-iv/
// 同值组只展开一次，总扫描量 O(n)，避免退化到平方复杂度。
func minJumps(arr []int) int {
	positions := make(map[int][]int)
	for i, value := range arr {
		positions[value] = append(positions[value], i)
	}
	return JumpBFS(len(arr), 0, func(i int) bool { return i == len(arr)-1 }, func(i int, yield func(int)) {
		yield(i - 1)
		yield(i + 1)
		for _, j := range positions[arr[i]] {
			yield(j)
		}
		delete(positions, arr[i])
	})
}
```

### 1340

[代码](l1340/1340.go) · [测试](l1340/1340_test.go)

只能严格下降，故依赖无环；遇到不矮的柱子必须停止该方向，答案包含起点。

需要模板：`MemoJump`。

<!-- solution: dp/jump_game/l1340/1340.go; templates: MemoJump -->
```go
// https://leetcode.com/problems/jump-game-v/
// 只能严格下降，故依赖无环；遇到不矮的柱子必须停止该方向，答案包含起点。
func maxJumps(arr []int, d int) int {
	solve := MemoJump(func(i int, next func(int) int) int {
		best := 1
		for _, direction := range []int{-1, 1} {
			for step := 1; step <= d; step++ {
				j := i + direction*step
				if j < 0 || j >= len(arr) || arr[j] >= arr[i] {
					break
				}
				best = max(best, 1+next(j))
			}
		}
		return best
	})
	answer := 0
	for i := range arr {
		answer = max(answer, solve(i))
	}
	return answer
}
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

```sh
go test ./dp/jump_game/...
```
