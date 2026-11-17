package l986

// AdvanceTwo 扫描两个独立序列；step 每次至少推进一个指针，不能后退。
func AdvanceTwo(n, m int, step func(i, j int) (int, int)) (int, int) {
	i, j := 0, 0
	for i < n && j < m {
		i, j = step(i, j)
	}
	return i, j
}

// https://leetcode.com/problems/interval-list-intersections/
// 输入各自有序且内部不相交；输出当前交集后推进结束更早的区间。
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	answer := [][]int{}
	AdvanceTwo(len(firstList), len(secondList), func(i, j int) (int, int) {
		left, right := max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])
		if left <= right {
			answer = append(answer, []int{left, right})
		}
		if firstList[i][1] < secondList[j][1] {
			return i + 1, j
		}
		if firstList[i][1] > secondList[j][1] {
			return i, j + 1
		}
		return i + 1, j + 1
	})
	return answer
}
