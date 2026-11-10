package l1345

// https://leetcode.com/problems/jump-game-iv/
func minJumps(arr []int) int {

	valueToIdxs := map[int][]int{}

	for i, v := range arr {
		valueToIdxs[v] = append(valueToIdxs[v], i)
	}

	n := len(arr)
	visit := make([]bool, n)
	validAndNotVisit := func(i int) bool {
		if i >= 0 && i < n && visit[i] == false {
			return true
		}
		return false
	}

	q := []int{0}
	step := 0

	for len(q) > 0 {
		L := len(q)
		for i := 0; i < L; i++ {
			idx := q[i]
			if idx == n-1 {
				return step
			}
			visit[idx] = true

			if validAndNotVisit(idx - 1) {
				q = append(q, idx-1)
			}
			if validAndNotVisit(idx + 1) {
				q = append(q, idx+1)
			}
			for _, sameIdx := range valueToIdxs[arr[idx]] {
				if sameIdx != idx {
					if validAndNotVisit(sameIdx) {
						q = append(q, sameIdx)
					}

				}
			}
			delete(valueToIdxs, arr[idx])
		}
		step += 1
		q = q[L:]
	}
	return 0
}
