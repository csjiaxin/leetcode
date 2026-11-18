# 栈：嵌套解析、调用现场与分段输出

## 同文件命名模板与适配器

先复制本代码块，再复制任一题的适配器；实际题解文件同样内嵌 helper，保留官方函数签名。只使用标准库，不导入仓库本地模板包。本文所有 Go 代码块可按顺序组合成一个包；不要重复复制 helper。

```go
import (
	"strconv"
	"strings"
)

type arithmeticFrame struct{ sum, sign int }

// evaluateSum restores sum/sign at each closing parenthesis; unary '-' is allowed.
func evaluateSum(s string) int {
	stack := []arithmeticFrame{}
	sum, number, sign := 0, 0, 1
	for i := range s {
		token := s[i]
		switch {
		case token >= '0' && token <= '9':
			number = number*10 + int(token-'0')
		case token == '+' || token == '-':
			sum += sign * number
			number = 0
			sign = 1
			if token == '-' {
				sign = -1
			}
		case token == '(':
			stack = append(stack, arithmeticFrame{sum, sign})
			sum, number, sign = 0, 0, 1
		case token == ')':
			sum += sign * number
			number = 0
			frame := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum = frame.sum + frame.sign*sum
		}
	}
	return sum + sign*number
}

type decoderFrame struct {
	prefix []byte
	times  int
}

// decodeNested restores the outer prefix and repeats a completed child.
func decodeNested(s string) string {
	stack := []decoderFrame{}
	current := []byte{}
	number := 0
	for i := 0; i < len(s); i++ {
		token := s[i]
		switch {
		case token >= '0' && token <= '9':
			number = number*10 + int(token-'0')
		case token == '[':
			stack = append(stack, decoderFrame{current, number})
			current = []byte{}
			number = 0
		case token == ']':
			frame := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inner := current
			current = frame.prefix
			for repeat := 0; repeat < frame.times; repeat++ {
				current = append(current, inner...)
			}
		default:
			current = append(current, token)
		}
	}
	return string(current)
}

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

// parenthesesFold stores each open layer's sum of completed children.
func parenthesesFold(s string, wrap func(int) int) int {
	stack := []int{0}
	for _, token := range s {
		if token == '(' {
			stack = append(stack, 0)
			continue
		}
		inner := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack[len(stack)-1] += wrap(inner)
	}
	return stack[0]
}

type booleanFrame struct {
	operator          byte
	anyTrue, anyFalse bool
}

// evaluateBoolean folds each operator's direct children, never skipping syntax.
func evaluateBoolean(expression string) bool {
	stack := []booleanFrame{}
	operator := byte(0)
	result := false
	feed := func(value bool) {
		if len(stack) == 0 {
			result = value
			return
		}
		frame := &stack[len(stack)-1]
		frame.anyTrue = frame.anyTrue || value
		frame.anyFalse = frame.anyFalse || !value
	}
	for i := range expression {
		token := expression[i]
		switch token {
		case '!', '&', '|':
			operator = token
		case '(':
			stack = append(stack, booleanFrame{operator: operator})
		case 't':
			feed(true)
		case 'f':
			feed(false)
		case ')':
			frame := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			value := false
			switch frame.operator {
			case '!':
				value = frame.anyFalse
			case '&':
				value = !frame.anyFalse
			case '|':
				value = frame.anyTrue
			}
			feed(value)
		}
	}
	return result
}

// reverseNestedPairs jumps between paired delimiters and reverses traversal.
func reverseNestedPairs(s string) string {
	pair := make([]int, len(s))
	stack := []int{}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i] = j
			pair[j] = i
		}
	}
	result := make([]byte, 0, len(s))
	for i, direction := 0, 1; i >= 0 && i < len(s); i += direction {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			direction = -direction
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

// patternStackDigits flushes an increasing pending segment at each I/end.
func patternStackDigits(pattern string) string {
	pending, result := []byte{}, []byte{}
	for i := 0; i <= len(pattern); i++ {
		pending = append(pending, byte('1'+i))
		if i == len(pattern) || pattern[i] == 'I' {
			for len(pending) > 0 {
				result = append(result, pending[len(pending)-1])
				pending = pending[:len(pending)-1]
			}
		}
	}
	return string(result)
}
```

栈的共同作用是后进先出地恢复尚未完成的上下文。本目录不要求栈中数值单调；需要最近更大/更小边界的题见 [monotonic_stack](../monotonic_stack/README.md)。

## 模板与不变量

```text
进入新层：保存外层上下文，初始化当前层
读取普通 token：更新当前层的局部状态
结束当前层：计算本层结果，弹栈并合并回父层
输入结束：读取最外层结果
```

栈帧表示一个尚未完成的上下文，不一定是单个字符。乘数、运算符、外层前缀、累计分数等应根据语法分别保存；不能让所有题共用相同的“遇右括号就乘二”规则。

| 题目与当前路径 | 帧保存的内容 | 结束/弹栈时做什么 |
| --- | --- | --- |
| [394 字符串解码](https://leetcode.com/problems/decode-string/) · [本地](l394) | 外层字符串前缀、重复次数 | 内层字符串重复 k 次后拼回外层；连续数字属于一个 k |
| [856 括号的分数](https://leetcode.com/problems/score-of-parentheses/) · [本地](l856) | 当前层已经完成的兄弟表达式分数和 | 空内部为 1，非空为 `2*inner`，再加到父层 |
| [224 基本计算器](https://leetcode.com/problems/basic-calculator/) · [本地](l224) | 部分结果和符号 | 恢复外层运算上下文，注意一元负号 |
| [1106 解析布尔表达式](https://leetcode.com/problems/parsing-a-boolean-expression/) · [本地](l1106) | 布尔运算符、参数合并状态 | 对本层参数求 !、& 或 \| 后交给父层 |
| [1190 反转括号内子串](https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/) · [本地](l1190) | 预先匹配的括号端点、当前遍历方向 | 遇括号跳到配对端并反向，O(n) 输出，不反复复制内层 |
| [636 函数独占时间](https://leetcode.com/problems/exclusive-time-of-functions/) · [本地](l636) | 尚未结束的函数调用 | 结合日志时间恢复父调用；end 时间戳包含该时刻 |
| [2375 根据 DI 构造最小数字](https://leetcode.com/problems/construct-smallest-number-from-di-string/) · [本地](l2375) | 等待输出的递增候选数字段 | I 或末尾统一弹出，反转一段连续 D 关系 |

[856](https://leetcode.com/problems/score-of-parentheses/) 的具体不变量：底部额外放一个 0 作为虚拟根层；每遇 `(` 压入 0；每遇 `)` 弹出 inner，并将 `max(1,2*inner)` 加入父层。合法非空内部的分数一定大于 0，因此 inner 为 0 可以识别 `()`。栈深为尚未匹配的左括号数加一。

[2375](https://leetcode.com/problems/construct-smallest-number-from-di-string/) 是贪心分段输出，不是嵌套语法解析；[636](https://leetcode.com/problems/exclusive-time-of-functions/) 是调用栈恢复，不应机械套用括号评分的聚合操作。每题都先明确 token 边界和帧含义。

## 归并与原练习保留

- 原根目录 `394_string_decoder.go` 与测试已整体迁到 [l394](l394)，原 package 名 `stack` 不变，不影响独立目录的 Go 包编译。
- 原 `monotonic_stack/l856` 已归入 [l856/variants/from_monotonic_stack](l856/variants/from_monotonic_stack)，保留独立来源和原测试，现按同一 parenthesesFold 实现；它没有单调值不变量。
- 原 `permutation/l2375` 已归入 [l2375](l2375)，现也按本页模板实现，并保留原测试。

实际 Go 文件现与本页 helper+adapter 一致，各来源的测试仍独立运行。按包运行，例如 `go test ./stack/l856`；`go test ./stack/...` 可验证整个专题。迁移清单见 [reorganization.json](../two_pointers/reorganization.json)。


## 逐题模板适配器补全

下列函数使用本页已声明的同名子模板；不是将不同不变量强行改为同一弹栈或移动条件。

### [224](https://leetcode.com/problems/basic-calculator/)

```go
func calculate(s string) int { return evaluateSum(s) }
```

### [394](https://leetcode.com/problems/decode-string/)

```go
func decodeString(s string) string { return decodeNested(s) }
```

### [636](https://leetcode.com/problems/exclusive-time-of-functions/)

```go
func exclusiveTime(n int, logs []string) []int { return exclusiveCallTimes(n, logs) }
```

### [856](https://leetcode.com/problems/score-of-parentheses/)

```go
func scoreOfParentheses(s string) int {
	return parenthesesFold(s, func(inner int) int {
		if inner == 0 {
			return 1
		}
		return 2 * inner
	})
}
```

### [1106](https://leetcode.com/problems/parsing-a-boolean-expression/)

```go
func parseBoolExpr(expression string) bool { return evaluateBoolean(expression) }
```

### [1190](https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/)

```go
func reverseParentheses(s string) string { return reverseNestedPairs(s) }
```

### [2375](https://leetcode.com/problems/construct-smallest-number-from-di-string/)

```go
func smallestNumber(pattern string) string { return patternStackDigits(pattern) }
```
