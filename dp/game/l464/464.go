package l464

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
		remainning := desiredTotal
		for x := 1; x <= maxChoosableInteger; x++ {
			bit := 1 << (x - 1)
			if mask&bit != 0 {
				remainning -= x
			}
		}

		for x := maxChoosableInteger; x >= 1; x-- {
			bit := 1 << (x - 1)
			if mask&bit == 0 && (x >= remainning || !win(mask|bit)) {
				return true
			}
		}
		return false
	})
	return win(0)
}
