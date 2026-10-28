// Package decodestring solves LeetCode 394. Decode String.
//
// 思路：用栈保存进入 '[' 之前的状态（重复次数 + 已解码的前缀），
// 遇到 ']' 时弹栈，把当前片段重复 k 次后拼接到前缀后面。
// 例：3[a2[c]] -> accaccacc
package decodestring

import "strings"

// frame 记录进入一层 '[' 之前的上下文。
type frame struct {
	times  int    // 当前括号内内容需要重复的次数
	prefix string // '[' 之前已经解码好的字符串
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func decodeString(s string) string {
	var (
		stack   []frame
		times   int
		current strings.Builder
	)

	for _, c := range s {
		switch {
		case isDigit(c):
			// 支持多位数字，如 "12[a]"。
			times = times*10 + int(c-'0')
		case c == '[':
			stack = append(stack, frame{times: times, prefix: current.String()})
			times = 0
			current.Reset()
		case c == ']':
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			repeated := strings.Repeat(current.String(), top.times)
			current.Reset()
			current.WriteString(top.prefix)
			current.WriteString(repeated)
		default:
			current.WriteRune(c)
		}
	}
	return current.String()
}
