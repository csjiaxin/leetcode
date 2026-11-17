package l1306

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

// https://leetcode.com/problems/jump-game-iii/
// 目标是值为零的节点，邻居为左右跳转，不是最大下标。
func canReach(arr []int, start int) bool {
	return JumpBFS(len(arr), start, func(i int) bool { return arr[i] == 0 }, func(i int, yield func(int)) {
		yield(i - arr[i])
		yield(i + arr[i])
	}) >= 0
}
