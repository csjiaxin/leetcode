package l636

import (
	"strconv"
	"strings"
)

// Standalone nested-context subtemplate and problem adapter; see README.md.
// exclusiveCallTimes accounts for the active top frame between log events.
func exclusiveCallTimes(n int, logs []string) []int {
	result := make([]int, n)
	stack := []int{}
	previous := 0
	for _, log := range logs {
		fields := strings.Split(log, ":")
		id, _ := strconv.Atoi(fields[0])
		timestamp, _ := strconv.Atoi(fields[2])
		if fields[1] == "start" {
			if len(stack) > 0 {
				result[stack[len(stack)-1]] += timestamp - previous
			}
			stack = append(stack, id)
			previous = timestamp
		} else {
			result[stack[len(stack)-1]] += timestamp - previous + 1
			stack = stack[:len(stack)-1]
			previous = timestamp + 1
		}
	}
	return result
}

func exclusiveTime(n int, logs []string) []int { return exclusiveCallTimes(n, logs) }
