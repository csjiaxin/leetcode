package l134

// https://leetcode.com/problems/gas-station/
func canCompleteCircuit(gas []int, cost []int) int {
	start, tank, total := 0, 0, 0
	for i := range gas {
		net := gas[i] - cost[i]
		tank += net
		total += net
		if tank < 0 {
			// No station from start through i can reach station i+1.
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
