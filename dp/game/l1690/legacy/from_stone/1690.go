package l1690

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

// https://leetcode.com/problems/stone-game-vii/
// 与主版本相同的当前玩家分差状态，补全原 panic 练习。
func stoneGameVII(stones []int) int {
	prefix := make([]int, len(stones)+1)
	for i, value := range stones {
		prefix[i+1] = prefix[i] + value
	}
	type interval struct{ left, right int }
	diff := MemoGame(func(s interval, solve func(interval) int) int {
		if s.left >= s.right {
			return 0
		}
		// 得到的是移除后剩余区间总和，而不是被移除石子的值。
		left := prefix[s.right+1] - prefix[s.left+1] - solve(interval{s.left + 1, s.right})
		right := prefix[s.right] - prefix[s.left] - solve(interval{s.left, s.right - 1})
		if left > right {
			return left
		}
		return right
	})
	return diff(interval{0, len(stones) - 1})
}
