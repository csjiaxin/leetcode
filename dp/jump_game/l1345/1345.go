package l1345

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
