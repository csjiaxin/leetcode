package l134

// CircuitStart 某段累计净油量首次为负时，排除该段内所有候选起点。
func CircuitStart(gas, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	total, tank, start := 0, 0, 0
	for i, value := range gas {
		net := value - cost[i]
		total += net
		tank += net
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

// https://leetcode.com/problems/gas-station/
// 总量判断是否有解，局部首次失败负责排除候选起点。
func canCompleteCircuit(gas []int, cost []int) int { return CircuitStart(gas, cost) }
