package main

import "fmt"

// https://leetcode.com/problems/evaluate-division/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ids := make(map[string]int)
	for _, equation := range equations {
		for _, name := range equation {
			if _, ok := ids[name]; !ok {
				ids[name] = len(ids)
			}
		}
	}
	n := len(ids)
	ratios := make([][]float64, n)
	reachable := make([][]bool, n)
	for i := range ratios {
		ratios[i] = make([]float64, n)
		reachable[i] = make([]bool, n)
		ratios[i][i], reachable[i][i] = 1, true
	}
	for i, equation := range equations {
		a, b := ids[equation[0]], ids[equation[1]]
		ratios[a][b], ratios[b][a] = values[i], 1/values[i]
		reachable[a][b], reachable[b][a] = true, true
	}
	// 原题关系一致，已有可达比值无需重新优化，只补出未知的关系。
	FloydClosure(ratios, reachable, func(a, b float64) (float64, bool) { return a * b, true }, func(a, b float64) bool { return false })
	answer := make([]float64, len(queries))
	for i, query := range queries {
		a, aOK := ids[query[0]]
		b, bOK := ids[query[1]]
		answer[i] = -1
		if aOK && bOK && reachable[a][b] {
			answer[i] = ratios[a][b]
		}
	}
	return answer
}

// FloydClosure 按允许的中间节点递增；extend 失败时停止并返回 false。
func FloydClosure[T any](value [][]T, reachable [][]bool, extend func(T, T) (T, bool), better func(T, T) bool) bool {
	for k := range value {
		for i := range value {
			if !reachable[i][k] {
				continue
			}
			for j := range value {
				if !reachable[k][j] {
					continue
				}
				candidate, ok := extend(value[i][k], value[k][j])
				if !ok {
					return false
				}
				if !reachable[i][j] || better(candidate, value[i][j]) {
					value[i][j] = candidate
					reachable[i][j] = true
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2, 3}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
}
