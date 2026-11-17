package l464

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
	win := MemoGame(func(mask int, solve func(int) bool) bool {
		// 已选数字之和由 mask 唯一决定，不必加入缓存键。
		remaining := desiredTotal
		for x := 1; x <= maxChoosableInteger; x++ {
			if mask&(1<<(x-1)) != 0 {
				remaining -= x
			}
		}
		for x := maxChoosableInteger; x >= 1; x-- {
			bit := 1 << (x - 1)
			if mask&bit == 0 && (x >= remaining || !solve(mask|bit)) {
				return true
			}
		}
		return false
	})
	return win(0)
}
