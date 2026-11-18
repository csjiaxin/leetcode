# 双序列前缀 DP：LCS、编辑距离、匹配计数与答案构造

本系列复用现有的 `dp/strings`。七道双前缀题使用同一套 `Table[T]`，各题主文件内均有一份完整本地实现；[template.go](template.go) 保留为参考实现与测试入口，不是题解的运行时依赖。[5](https://leetcode.com/problems/longest-palindromic-substring/)、[516](https://leetcode.com/problems/longest-palindromic-subsequence/)、[1312](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/) 另外使用本页第 8 节的闭区间 `IntervalTable`。题目目录和函数签名保持原样。

建议顺序：**[1143](https://leetcode.com/problems/longest-common-subsequence/) → [1035](https://leetcode.com/problems/uncrossed-lines/) → [583](https://leetcode.com/problems/delete-operation-for-two-strings/) → [72](https://leetcode.com/problems/edit-distance/) → [97](https://leetcode.com/problems/interleaving-string/) → [115](https://leetcode.com/problems/distinct-subsequences/) → [1092](https://leetcode.com/problems/shortest-common-supersequence/)**。

## 1. 适用条件：同时处理两个序列的前缀

当问题需要比较、匹配、编辑或者合并两个序列，并且当前状态能由更短前缀推导时，可以尝试这个模板。

例如 `A="abcde"`、`B="ace"`，允许跳过字符、保持剩余字符的相对顺序，就可以匹配出公共子序列 `"ace"`。如果只考虑 `A="abc"`、`B="ac"`，又得到一个结构相同、规模更小的问题。

因此用两个前缀长度描述状态：

```text
dp[i][j] = 只考虑 A[:i]、B[:j] 时，对应问题的状态值

A[:i] 对应 [0,i)，当前字符是 A[i-1]
B[:j] 对应 [0,j)，当前字符是 B[j-1]
```

这里 `i`、`j` 是前缀长度，取值可以为 0。表格大小为 `(m+1) × (n+1)`；第 0 行、第 0 列代表空前缀。

这些题的字符串均按原题 ASCII 字符范围处理，使用 Go 的字节下标。若扩展到任意 Unicode 字符，应先转为 `[]rune`，再按字符数建表。[1035](https://leetcode.com/problems/uncrossed-lines/) 直接对整数数组使用相同的前缀定义。

## 2. 不变量与计算顺序

每个状态通常从下列位置中选取合法来源：

```text
        j-1                    j
i-1     diagonal               up
         ↘                      ↓
i       left         →         当前 dp[i][j]
```

| 来源 | 已经求出的状态 | 从来源走到当前格子的含义 |
| --- | --- | --- |
| `up` | `dp[i-1][j]` | 新增处理 A 的一个元素 |
| `left` | `dp[i][j-1]` | 新增处理 B 的一个元素 |
| `diagonal` | `dp[i-1][j-1]` | 同时新增处理两边的一个元素 |

**不变量：处理 `dp[i][j]` 时，上方、左方、左上方的状态均已计算完成。**

先初始化空前缀边界，再让 `i` 从 1 到 m、`j` 从 1 到 n 填表即可。这些依赖构成 DAG，行优先遍历就是一种合法计算顺序。

每题只采用题目允许的来源。例如 [115](https://leetcode.com/problems/distinct-subsequences/) 可以跳过源字符，却不能跳过目标字符，因此不会使用 `left`；[97](https://leetcode.com/problems/interleaving-string/) 每次从其中一个来源串消费字符，不使用 `diagonal`。

## 3. 固定代码模板

参考实现见 [template.go](template.go)，下方完整声明也已内联到各双前缀题的主文件中。模板需要 Go 1.18+ 的泛型语法；运行本仓库时使用根目录 `go.mod` 指定的 Go 版本。

```go
func Table[T any](
    m, n int,
    base func(i, j int) T,
    step func(i, j int, up, left, diagonal T) T,
) [][]T {
    dp := make([][]T, m+1)
    for i := range dp {
        dp[i] = make([]T, n+1)
    }

    for i := 0; i <= m; i++ {
        dp[i][0] = base(i, 0)
    }
    for j := 1; j <= n; j++ {
        dp[0][j] = base(0, j)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = step(i, j,
                dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
        }
    }
    return dp
}
```

套题时只修改四处：

1. **状态值 `T` 与含义**：长度、代价、计数使用整数，可行性使用 bool。
2. **`base`**：空前缀的答案。
3. **`step`**：允许哪些来源，如何根据当前元素合并。
4. **答案提取**：通常返回 `dp[m][n]`；[1092](https://leetcode.com/problems/shortest-common-supersequence/) 还需要沿表回溯构造字符串。

`m`、`n` 必须非负，回调通过闭包读取原序列。`base` 仅负责第 0 行和第 0 列，`step` 仅在 `i>0 && j>0` 时调用。字符下标 `i-1`、`j-1` 因而不会出现在空前缀边界。

回调为 O(1) 时，模板时间与空间均为 **O((m+1)(n+1))**，两个序列非空时通常简写为 O(mn)。[1092](https://leetcode.com/problems/shortest-common-supersequence/) 回溯额外使用 O(m+n) 时间和输出缓冲区。

## 4. 套用示例：[1143](https://leetcode.com/problems/longest-common-subsequence/) 最长公共子序列

定义 `dp[i][j]` 为两个前缀的最长公共子序列长度。任一前缀为空时长度为 0。

下面的文档示例直接调用第 3 节定义的本地 `Table`。把两个代码块放在同一文件即可使用；实际题解同样内联模板，无需导入仓库中的包。

```go
func longestCommonSubsequence(text1, text2 string) int {
    dp := Table[int](len(text1), len(text2),
        func(i, j int) int { return 0 },
        func(i, j, up, left, diagonal int) int {
            if text1[i-1] == text2[j-1] {
                return diagonal + 1
            }
            if up > left {
                return up
            }
            return left
        },
    )
    return dp[len(text1)][len(text2)]
}
```

两个末尾字符相同时，将它们匹配，接在更短前缀的答案后；不相同时，分别尝试跳过其中一个字符，取较大的结果。[1035](https://leetcode.com/problems/uncrossed-lines/) 只需把字符比较替换成整数比较，就能使用同一转移。

## 5. 七题分别修改什么

下面的 `up`、`left`、`diagonal` 均指模板传入的三个前驱值，`same` 表示两个前缀的末尾元素相同。

| 题目 | `dp[i][j]` 的含义 | 空前缀 `base` | 正常格子 `step` |
| --- | --- | --- | --- |
| [1143](https://leetcode.com/problems/longest-common-subsequence/) / [1035](https://leetcode.com/problems/uncrossed-lines/) | 两个前缀的 LCS 长度 | 全部为 0 | `same` 时 `diagonal+1`；否则 `max(up,left)` |
| [583](https://leetcode.com/problems/delete-operation-for-two-strings/) | 使两个前缀相同的最少删除次数 | `i+j` | `same` 时 `diagonal`；否则 `min(up,left)+1` |
| [72](https://leetcode.com/problems/edit-distance/) | 将 A 的前缀变成 B 的前缀的最少编辑次数 | `i+j` | `same` 时 `diagonal`；否则 `min(up,left,diagonal)+1` |
| [97](https://leetcode.com/problems/interleaving-string/) | 两个来源前缀能否组成目标的前 `i+j` 个字符 | 空对空为 true；单串前缀必须与目标对应前缀相同 | `(up && A[i-1]==C[i+j-1]) \|\| (left && B[j-1]==C[i+j-1])` |
| [115](https://leetcode.com/problems/distinct-subsequences/) | 从源前缀中选择字符、组成目标前缀的方案数 | 目标为空为 1；仅源为空为 0 | `same` 时 `up+diagonal`；否则 `up` |
| [1092](https://leetcode.com/problems/shortest-common-supersequence/) | 同时包含两个前缀的最短公共超序列长度 | `i+j` | `same` 时 `diagonal+1`；否则 `min(up,left)+1` |

### [583](https://leetcode.com/problems/delete-operation-for-two-strings/) 与 [72](https://leetcode.com/problems/edit-distance/)：允许的操作决定来源

[583](https://leetcode.com/problems/delete-operation-for-two-strings/) 只能删除。末尾不相同时，要么删除 A 的末尾，对应 `up+1`；要么删除 B 的末尾，对应 `left+1`。也可以先求 LCS 长度 L，再用 `m+n-2L` 求删除次数。

[72](https://leetcode.com/problems/edit-distance/) 允许插入、删除和替换。把 A 变为 B 时，`up+1` 表示删除，`left+1` 表示插入，`diagonal+1` 表示替换。例如 `"a"` 与 `"b"`，[583](https://leetcode.com/problems/delete-operation-for-two-strings/) 需要两次删除，[72](https://leetcode.com/problems/edit-distance/) 只需要一次替换。

### [97](https://leetcode.com/problems/interleaving-string/)：来源不同，两个选择都要考虑

先检查 `len(A)+len(B)==len(C)`。处理到 `(i,j)` 时，总共消费了 `i+j` 个来源字符，因此目标字符是 `C[i+j-1]`。

如果两个来源串的当前字符都匹配目标，两个方向都要尝试，使用逻辑或合并。只走第一个匹配方向的贪心可能失败。实现为两个单串边界预计算匹配标记，避免在每个边界格子重新比较长前缀。

### [115](https://leetcode.com/problems/distinct-subsequences/)：方案计数相加，空目标有一种形成方式

末尾相同时有两个互斥来源：跳过当前源字符，或者使用它匹配当前目标字符。因此是 `up+diagonal`，不是取最大值。即使形成的目标字符串相同，只要选择的源下标不同，就计为不同方案。

初始化 `dp[i][0]=1`：无论源前缀有多长，全部不选，都能以一种方式形成空目标。`dp[0][j>0]=0`。

[原题](https://leetcode.com/problems/distinct-subsequences/)保证最终答案不超过 `2^31-1`，但某些中间前缀计数仍可能超过 int64。本实现将中间值保存为 `min(真实计数, 2^31-1)`，并在相加前检查上限。所有转移都是非负加法，所以这个截断不改变题目保证范围内的最终答案，也避免依赖整数溢出。若扩展到允许更大最终答案的输入，需要相应扩大计数表示。

### [1092](https://leetcode.com/problems/shortest-common-supersequence/)：长度表之后还要恢复实际答案

从 `(m,n)` 向 `(0,0)` 回溯：

1. 末尾字符相同：放入一次该字符，同时减少 i、j。
2. 末尾字符不同：选择长度更小的前驱，放入该方向消费的字符。
3. 两个前驱长度相同：任选一条；最短答案可能不唯一。
4. 某一前缀为空：把另一前缀剩余字符加入。

回溯从右往左产生字符，先追加到字节缓冲区，最后整体反转。这样恢复答案为 O(m+n)，不用在循环中反复向字符串头部拼接。测试验证两个输入都是结果的子序列、且结果长度最短，不要求所有情况返回同一个字符串。

## 6. 题目链接、实现和测试

| 顺序 | LeetCode | 本地实现 | 测试 | 训练目标 |
| --- | --- | --- | --- | --- |
| 1 | [1143. 最长公共子序列](https://leetcode.com/problems/longest-common-subsequence/) | [1143.go](l1143/1143.go) | [1143_test.go](l1143/1143_test.go) | 匹配与跳过，求最大长度 |
| 2 | [1035. 不相交的线](https://leetcode.com/problems/uncrossed-lines/) | [1035.go](l1035/1035.go) | [1035_test.go](l1035/1035_test.go) | 从字符串迁移到整数序列 |
| 3 | [583. 两个字符串的删除操作](https://leetcode.com/problems/delete-operation-for-two-strings/) | [583.go](l583/583.go) | [583_test.go](l583/583_test.go) | 删除代价与 LCS 的联系 |
| 4 | [72. 编辑距离](https://leetcode.com/problems/edit-distance/) | [72.go](l72/72.go) | [72_test.go](l72/72_test.go) | 插入、删除、替换的最小代价 |
| 5 | [97. 交错字符串](https://leetcode.com/problems/interleaving-string/) | [97.go](l97/97.go) | [97_test.go](l97/97_test.go) | 两个前缀合并的可行性 |
| 6 | [115. 不同的子序列](https://leetcode.com/problems/distinct-subsequences/) | [115.go](l115/115.go) | [115_test.go](l115/115_test.go) | 使用与跳过的方案计数 |
| 7 | [1092. 最短公共超序列](https://leetcode.com/problems/shortest-common-supersequence/) | [1092.go](l1092/1092.go) | [1092_test.go](l1092/1092_test.go) | 长度 DP 与答案回溯 |

```text
dp/strings/
├── readme.md
├── template.go
├── template_test.go
├── l1143/  (1143.go, 1143_test.go)
├── l1035/  (1035.go, 1035_test.go)
├── l583/   (583.go, 583_test.go)
├── l72/    (72.go, 72_test.go)
├── l97/    (97.go, 97_test.go)
├── l115/   (115.go, 115_test.go)
└── l1092/  (1092.go, 1092_test.go)
```

目录中还保留其他已有的字符串题和历史片段；上面只列出本系列文件。

在仓库根目录运行：

```sh
go test -count=1 ./dp/strings/...
go test -count=1 -cover ./dp/strings/...

# 单独练习一道题
go test -v ./dp/strings/l115
```

测试包括官方示例、空前缀、重复字符、顺序约束，以及计数与构造类的边界。公共模板还验证空维度和三个前驱的传递；[1092](https://leetcode.com/problems/shortest-common-supersequence/) 另用小规模穷举对照最短长度及子序列关系。

每题主 `{题号}.go` 已包含题解及其完整本地模板。提交主解法到 LeetCode 时，保留题目函数和辅助模板、去掉仓库包声明即可，无需再导入或复制外部模板。同一包只保留一份模板定义；例如 l516 的多个源文件不应重复声明 `IntervalTable`。若单独选择其他变体，只需带上它实际使用的辅助函数。

## 7. 其他已有字符串 DP 资料

- [132. 分割回文串 II](https://leetcode.com/problems/palindrome-partitioning-ii/)
- [647. 回文子串](https://leetcode.com/problems/palindromic-substrings/)
- [516. 最长回文子序列](https://leetcode.com/problems/longest-palindromic-subsequence/)
- [原有 LeetCode 题单](https://leetcode.com/list/55afh7m7/)

## 8. 本目录其他字符串题的子模板

前面的 `Table` 承诺上、左、左上三个前驱，覆盖本系列七题。现有回文实现及相关的切分、模式匹配题，按下面的依赖分别选择骨架。

| 子系列 | 状态与固定转移 | 本地实现 / 相关原题 |
| --- | --- | --- |
| 连续回文区间 | `pal[l][r] = s[l]==s[r] && (r-l<2 \|\| pal[l+1][r-1])`；按长度递增或 l 递减 | [5](https://leetcode.com/problems/longest-palindromic-substring/) · [代码](l5/5.go)、[647](https://leetcode.com/problems/palindromic-substrings/) |
| 回文子序列 | 同字符接上内部答案；不同字符跳过一端取最大 | [516](https://leetcode.com/problems/longest-palindromic-subsequence/) · [代码](l516/516.go)、[反转串 LCS 变体](l516/516_reverse_lcs.go) |
| 最少插入成为回文 | 同字符取内部答案；不同字符取较小相邻区间再加 1 | [1312](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/) · [代码](l1312/1312.go) |
| 回文前缀切分 | 预处理 pal；对合法末段 `[j,i)` 枚举分割点，累计段数或判断可行性 | [132](https://leetcode.com/problems/palindrome-partitioning-ii/)、[1745](https://leetcode.com/problems/palindrome-partitioning-iv/) |
| 带权删除 | 同字符取对角；否则比较删除当前 A 或 B 字符后的累计 ASCII 代价 | [712](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/) |
| 通配符匹配 | 普通字符/问号走对角；星号选择匹配零个或继续消费文本字符 | [44](https://leetcode.com/problems/wildcard-matching/) |
| 正则匹配 | 星号作用于前一个模式元素，需要考虑 `dp[i][j-2]` 和继续匹配该元素 | [10](https://leetcode.com/problems/regular-expression-matching/) |

### 闭区间模板：按长度填表

[5](https://leetcode.com/problems/longest-palindromic-substring/)、[516](https://leetcode.com/problems/longest-palindromic-subsequence/)、[1312](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/) 使用同一份本地 `IntervalTable`。它与前缀 `Table` 的依赖不同：需要去掉左端、去掉右端和内部区间，而不是上、左、左上。下面先声明一次模板，再给出三题适配；组合模板与任一适配即可形成独立示例。示例的内置 min/max 需要 Go 1.21+。

```go
// IntervalTable 按区间长度填表；insideEmpty 是空内部区间的单位值。
func IntervalTable[T any](n int, insideEmpty T, base func(i int) T, step func(l, r int, withoutLeft, withoutRight, inside T) T) [][]T {
	dp := make([][]T, n)
	for i := range dp {
		dp[i] = make([]T, n)
		dp[i][i] = base(i)
	}
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			inside := insideEmpty
			if length > 2 {
				inside = dp[l+1][r-1]
			}
			dp[l][r] = step(l, r, dp[l+1][r], dp[l][r-1], inside)
		}
	}
	return dp
}
```

`insideEmpty` 是长度为 2 时的空内部区间：回文判定取 true，长度/操作数取 0；`base` 负责单字符。回调接收的三个状态分别是 dp[l+1][r]、dp[l][r-1]、dp[l+1][r-1]。外层按长度递增，保证依赖已完成。

#### [5](https://leetcode.com/problems/longest-palindromic-substring/)：最长连续回文子串

```go
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, best := 0, 1
	IntervalTable(len(s), true, func(i int) bool { return true }, func(l, r int, withoutLeft, withoutRight, inside bool) bool {
		valid := s[l] == s[r] && inside
		if valid && r-l+1 > best {
			start, best = l, r-l+1
		}
		return valid
	})
	return s[start : start+best]
}
```

#### [516](https://leetcode.com/problems/longest-palindromic-subsequence/)：最长回文子序列

```go
func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := IntervalTable(len(s), 0, func(i int) int { return 1 }, func(l, r, withoutLeft, withoutRight, inside int) int {
		if s[l] == s[r] {
			return inside + 2
		}
		return max(withoutLeft, withoutRight)
	})
	return dp[0][len(s)-1]
}
```

#### [1312](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/)：最少插入成为回文

```go
func minInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := IntervalTable(len(s), 0, func(i int) int { return 0 }, func(l, r, withoutLeft, withoutRight, inside int) int {
		if s[l] == s[r] {
			return inside
		}
		return min(withoutLeft, withoutRight) + 1
	})
	return dp[0][len(s)-1]
}
```

[516](https://leetcode.com/problems/longest-palindromic-subsequence/) 的单字符答案为 1、空内部区间为 0；[1312](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/) 的单字符和空区间均不需要插入。不能只复用遍历顺序而沿用错误的单位元或代价。

切分问题中，若 `pieces[i]` 表示前 i 个字符的最少回文段数，则 `pieces[0]=0`，合法末段转移为 `pieces[j]+1`；非空串的最少切割次数为段数减 1。正则匹配需要两列之前的状态，保留独立填表骨架，不机械套只接收三个邻居的 `Table` 回调。

更多区间依赖与分割点枚举见 [区间 DP](../interval/README.md)。扩展原题链接提供状态迁移方向，是否已有可运行实现以实际题目目录与测试为准。
