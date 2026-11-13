# KMP／前缀函数系列

目标：固定一个模板，解决多类字符串与序列匹配问题。

核心思想：匹配失败时，利用已经匹配的前后缀信息，避免从头重新比较。

## 目录

- [一、练习路线](#一练习路线)
- [二、先理解前缀表](#二先理解前缀表)
  - [1. 什么是前缀、后缀？](#1-什么是前缀后缀)
  - [2. 为什么要排除整个字符串？](#2-为什么要排除整个字符串)
  - [3. pi[i] 不是只看字符 s[i]](#3-pii-不是只看字符-si)
  - [4. 前缀和后缀可以重叠吗？](#4-前缀和后缀可以重叠吗)
  - [5. 为什么返回 s[:ml]，不是 s[:ml+1]？](#5-为什么返回-sml不是-sml1)
  - [6. 一句话总结与自测](#6-一句话总结与自测)
- [三、两个可复用函数](#三两个可复用函数)
  - [1. 先分清两个函数各自做什么](#1-先分清两个函数各自做什么)
  - [2. Go 模板：构造前缀表](#2-go-模板构造前缀表)
    - [设计思路：从定义推导，包括 j 的含义与回退](#设计思路先从定义出发不急着写代码)
    - [Go 实现](#go-实现)
    - [2.1 先明确：每一轮到底在计算什么？](#21-先明确每一轮到底在计算什么)
    - [2.2 初始化：pi[0] 为什么不用算？](#22-初始化pi0-为什么不用算)
    - [2.3 第一轮：i = 1](#23-第一轮i--1计算-ab-的答案)
    - [2.4 第二轮：i = 2](#24-第二轮i--2计算-aba-的答案)
    - [2.5 第三轮：i = 3](#25-第三轮i--3计算-abab-的答案)
    - [2.6 第四轮：i = 4](#26-第四轮i--4计算-ababa-的答案)
    - [2.7 第五轮：i = 5（连续回退，重点）](#27-第五轮i--5计算-ababaa-的答案重点)
    - [2.8 把所有轮次放在一起](#28-把所有轮次放在一起)
    - [2.9 为什么这三个细节不能省？](#29-为什么这三个细节不能省)
  - [3. Go 模板：查找全部匹配位置](#3-go-模板查找全部匹配位置)
    - [设计思路：j 的含义与失配回退](#设计思路j-的含义与失配回退)
    - [第一步：前四个字符匹配成功](#第一步前四个字符匹配成功)
    - [第二步：之前的成功比较，能留下什么？](#第二步之前的成功比较能留下什么)
    - [第三步：这个 2 从哪里来？](#第三步这个-2-从哪里来)
    - [第四步：为什么文本下标 i 不回退？](#第四步为什么文本下标-i-不回退)
    - [Go 实现：findAll](#go-实现findall)
    - [完整匹配后，为什么还要回退？](#完整匹配后为什么还要回退)
  - [4. 复杂度与字符单位](#4-复杂度与字符单位)
  - [5. 最核心的一句话](#5-最核心的一句话)
- [四、每道题要提炼的结论](#四每道题要提炼的结论)
- [五、扩展题单与分类总结](#五扩展题单与分类总结)
  - [A. 基础查找与循环文本](#a-基础查找与循环文本)
  - [B. 全部匹配位置与整数序列](#b-全部匹配位置与整数序列)
  - [C. 前缀表、border 链与周期](#c-前缀表border-链与周期)
  - [D. KMP 状态与其他数据结构结合](#d-kmp-状态与其他数据结构结合)
  - [E. 看到题目时，如何选择模板？](#e-看到题目时如何选择模板)
  - [F. 推荐练习顺序与阶段目标](#f-推荐练习顺序与阶段目标)
- [六、练习检查清单](#六练习检查清单)

## 一、练习路线

| 顺序 | 题目 | 核心训练点 |
| --- | --- | --- |
| 1 | [28. 找出字符串中第一个匹配项的下标](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | 实现基础 KMP 匹配 |
| 2 | [1392. 最长快乐前缀](https://leetcode.com/problems/longest-happy-prefix/) | 理解最长相同真前缀、后缀，即前缀表的含义 |
| 3 | [459. 重复的子字符串](https://leetcode.com/problems/repeated-substring-pattern/) | 利用前缀表判断字符串的周期 |
| 4 | [686. 重复叠加字符串匹配](https://leetcode.com/problems/repeated-string-match/) | 在重复文本中寻找模式串 |
| 5 | [214. 最短回文串](https://leetcode.com/problems/shortest-palindrome/) | 通过反转、拼接，把回文问题转成前后缀匹配 |
| 6 | [3036. 匹配模式数组的子数组数目 II](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/) | 将相邻数字的大小关系编码，再进行序列匹配 |
| 7 | [3008. 找出数组中的美丽下标 II](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/) | KMP 找全部匹配位置，再结合二分或双指针 |

注：匹配模式数组的子数组数目 II 的题号是 3036。

先做 28 → 1392 → 459，再做 686 → 214；最后两道综合题暂时放后面。

## 二、先理解前缀表

先不看代码，也不考虑 KMP 怎么匹配。先弄懂一个问题：

> 一个字符串的开头和结尾，最多有多长是一样的？但不能拿整个字符串和自己比较。

前缀表就是记录这个答案的。

### 1. 什么是前缀、后缀？

以 `"abab"` 为例：

- **前缀**：从开头取连续的一段，例如 `"a"`、`"ab"`、`"aba"`、`"abab"`。
- **后缀**：取以结尾结束的连续一段，字符顺序不变，例如 `"b"`、`"ab"`、`"bab"`、`"abab"`。

比较相同长度的前缀和后缀：

| 长度 | 前缀 | 后缀 | 相等吗？ |
| --- | --- | --- | --- |
| 1 | `"a"` | `"b"` | 否 |
| 2 | `"ab"` | `"ab"` | **是** |
| 3 | `"aba"` | `"bab"` | 否 |
| 4 | `"abab"` | `"abab"` | 是，但**不允许选整个字符串** |

排除最后一行后，最长相等的前后缀是 `"ab"`，长度为 **2**。

### 2. 为什么要排除整个字符串？

因为任何字符串都和自己相等。如果允许选整个字符串：

```text
"abc"   → 答案总是 3
"abab"  → 答案总是 4
"hello" → 答案总是 5
```

这样就没有意义了。所以只考虑**比当前字符串短的前后缀**。
“真前缀”“真后缀”中的“真”，表达的就是排除整个字符串这一限制。
没有非空的相等真前后缀时，长度记为 0。

注意：必须比**当前正在考察的字符串**短，不只是比原始字符串短。

### 3. pi[i] 不是只看字符 s[i]

正式定义：

```text
pi[i] = s[0:i+1] 中，最长的、相等的真前缀与后缀的长度
```

这里使用 Go 的左闭右开切片表示法，`s[0:i+1]` 包含下标 `i`。

**i 表示我们考察到哪里，而不是只考察哪一个字符。**

对于 `s = "abab"`：

```text
pi[0] → 考察 "a"
pi[1] → 考察 "ab"
pi[2] → 考察 "aba"
pi[3] → 考察 "abab"
```

前缀表实际上是在每个位置问一次：

> 对目前这一段字符串，它的开头和结尾，最长有多少字符相同？

#### pi[0]：考察 "a"

只有一个字符，不能选整个 `"a"`，所以没有非空的相等真前后缀：

```text
pi[0] = 0
```

#### pi[1]：考察 "ab"

只需要比较长度为 1 的前后缀：前缀 `"a"`，后缀 `"b"`，不相等。
不能选择 `"ab"`，因为它是**当前字符串的全部**。

```text
pi[1] = 0
```

#### pi[2]：考察 "aba"

| 长度 | 前缀 | 后缀 | 相等吗？ |
| --- | --- | --- | --- |
| 1 | `"a"` | `"a"` | **是** |
| 2 | `"ab"` | `"ba"` | 否 |

最长相等的是 `"a"`，长度为 1：

```text
当前字符串：a b a
前缀：      a
后缀：          a

pi[2] = 1
```

#### pi[3]：考察 "abab"

| 长度 | 前缀 | 后缀 | 相等吗？ |
| --- | --- | --- | --- |
| 1 | `"a"` | `"b"` | 否 |
| 2 | `"ab"` | `"ab"` | **是** |
| 3 | `"aba"` | `"bab"` | 否 |

最长相等的是 `"ab"`，长度为 2，所以 `pi[3] = 2`。

把四次结果放在一起：

```text
i          0      1       2        3
考察范围   "a"    "ab"    "aba"    "abab"
pi[i]      0      0       1        2

pi = [0, 0, 1, 2]
```

**每个位置存的，都是“从开头到这里这一段”的答案。**

### 4. 前缀和后缀可以重叠吗？

**可以！不要求它们互不相交。**

例如 `"aaa"`：

```text
字符位置：  0 1 2
字符串：    a a a
前缀 "aa"： ✓ ✓
后缀 "aa"：   ✓ ✓
```

长度为 2 的前缀和后缀都是 `"aa"`。虽然共用了中间那个字符，但完全合法。
它们都比整个 `"aaa"` 短，所以最长相等真前后缀长度为 2。

因此：

```text
s  = "aaaa"
pi = [0, 1, 2, 3]
```

不是必须找到“两段分离的字符串”，只需要比较**开头取出的内容和结尾取出的内容**。

### 5. 为什么返回 s[:ml]，不是 s[:ml+1]？

因为 `ml` 是**长度（字符数量）**，不是最后一个字符的下标。
这里的示例都是 ASCII 字符，字符数量与 Go 字节数量一致。

例如找到 `"ab"`，长度为 2：

```text
下标： 0 1 2 3
字符： a b a b
       └─┘
       取 2 个字符
```

Go 的切片右边界不包含在内：

```go
s := "abab"
s[:2] // 取下标 0、1 → "ab"
s[:3] // 取下标 0、1、2 → "aba"
```

可以记住这个对应关系：

```text
长度为 ml
最后一个字符的下标是 ml-1（ml > 0 时）
切片的右边界是 ml
```

所以对应前缀应该取 `s[:ml]`。当 `ml == 0` 时，`s[:0]` 就是空字符串。

### 6. 一句话总结与自测

> pi[i]：把字符串截到下标 i，看看这一段的开头和结尾，最长能有多少字符一样；整段自己不算。

先不用急着写算法，可以通过枚举长度手算，确认理解了定义。

自测：对于 `"ababa"`，最长相等真前缀、后缀分别是什么？`pi[4]` 是多少？

<details>
<summary>展开查看答案</summary>

前缀和后缀都是 `"aba"`，长度为 3，因此 `pi[4] = 3`。
两段可以重叠。完整前缀表为 `[0, 0, 1, 2, 3]`。

</details>

## 三、两个可复用函数

### 1. 先分清两个函数各自做什么

- **prefixFunction(pattern)**：只分析模式串自己，逐个计算 `pi[i]`，最后返回整个 `pi` 数组。
- **findAll(text, pattern)**：使用这张前缀表，在文本里找模式串，返回全部匹配起点。

```text
prefixFunction(pattern)  是构建过程
pi                      是构建出来的整张表
pi[i]                   是表中的一个结果
```

例如 `prefixFunction("abab")` 的构建过程：

| i | 当前考察的字符串 pattern[:i+1] | 最长相等真前后缀 | pi[i] |
| --- | --- | --- | --- |
| 0 | `"a"` | 无 | 0 |
| 1 | `"ab"` | 无 | 0 |
| 2 | `"aba"` | `"a"` | 1 |
| 3 | `"abab"` | `"ab"` | 2 |

最后返回 `[0, 0, 1, 2]`。此时还没有去文本里查找匹配，只是在分析模式串自身。

最难的不是记住两个函数名，而是理解：**j 是什么，以及为什么 j = pi[j-1]。**

### 2. Go 模板：构造前缀表

#### 设计思路：先从定义出发，不急着写代码

我们的目标是构造整个 pi 数组，而不是在文本里查找模式串。
`pi[i]` 表示 `pattern[:i+1]` 的最长相等真前后缀长度。

也可以把这个过程理解为“模式串自己和自己匹配”：
把模式串开头当作要匹配的内容，逐步检查它能否与当前结尾接上。
下面先推导为什么可以这样做，再通过完整执行过程观察“能延长”和“需要回退”两种情况。

**第一步：最直接的做法是什么？**

对每个 i，枚举允许的前后缀长度，从长到短逐个比较。
例如对 `"ababa"`，依次比较长度 4、3、2、1 的前后缀，第一个相等的就是答案。
这样符合定义，但会反复比较已经看过的字符。

所以我们要问：**增加一个新字符后，能不能利用前一段已经算好的结果？**

**第二步：新答案一定来自旧前后缀的延长。**

假设当前字符串的某个相等真前后缀长度为 L，且 L > 0。
将这两段各去掉最后一个字符，剩下的两段仍相等，长度为 L-1；
它们分别是旧字符串 `pattern[:i]` 的前缀与后缀。

```text
新前缀 = 旧前缀 + 一个字符
新后缀 = 旧后缀 + 新字符 pattern[i]
```

因此，当前答案不是凭空出现的，而是由旧字符串的某个相等真前后缀延长而来。
当 L = 1 时，旧的两段都是空串，也包含在这个思路中。

**第三步：从旧的最长候选开始，使用 j 记录候选长度。**

既然希望得到最长答案，就先尝试旧字符串的最长相等真前后缀：

```text
j = pi[i-1]
```

在 prefixFunction 中，**j 是计算当前答案时的工作变量，表示准备尝试延长的候选前后缀长度**。
它不是当前字符的下标，也不一定等于最终的 pi[i]。

| 变量 | 职责 |
| --- | --- |
| i | 决定正在计算哪一格 pi[i]，也指向本轮新加入的字符 |
| pi | 保存已经算好的历史答案 |
| j | 本轮使用的候选长度，失配时缩短，接上新字符后增加 1 |

例如从 `"abab"` 加入一个 `'a'`，变成 `"ababa"`：

```text
旧字符串："abab"
旧前缀：  "ab"
旧后缀：  "ab"
j = 2
```

**第四步：j 为什么既是长度，又能当下标？**

长度为 j 的前缀占据下标 0 到 j-1，所以下一个字符的下标恰好是 j。

```text
下标： 0 1   2 3 4
字符： a b | a b a
       └─┘   ↑   ↑
       长度2 j   i
```

当 `i = 4，j = 2` 时：

```text
前缀一侧："ab" + pattern[j] = "ab" + 'a'
后缀一侧："ab" + pattern[i] = "ab" + 'a'
```

已有的两段 `"ab"` 相等，不需要重比，只比较 `pattern[j]` 和 `pattern[i]`。
如果相等，就把候选长度从 j 延长到 j+1，即执行 `j++`。

所以并不是用 j 随便访问一个字符，而是在访问**当前候选前缀后面、等待接上的字符**。

**第五步：延长失败，为什么使用 j = pi[j-1]？**

这里最难理解的是：

> 我们明明在计算整个 `"ababaa"`，为什么失败后，却去查 `"aba"` 自己的前后缀？

先不背公式，用具体字符把这个关系展开。

**（1）目前已经知道什么？**

正在给旧字符串 `"ababa"` 加上新字符 `'a'`。
旧字符串的最长相等真前后缀是 `"aba"`，长度为 3：

```text
下标：    0 1 2 3 4
旧字符串：a b a b a
前缀：    a b a
后缀：        a b a

新字符：'a'，位于 pattern[5]
j = 3
```

已经有两段 `"aba"` 相等，先试试能不能把它们延长一个字符：

```text
前缀一侧："aba" + pattern[3] = "aba" + 'b' → "abab"
后缀一侧："aba" + pattern[5] = "aba" + 'a' → "abaa"
```

不相等，所以旧的长度 3 候选不能延长。
注意：**这次失败的是“延长成长度 4”，不是所有更短长度都失败。**

**（2）接下来到底在找什么？**

我们要找：旧字符串 `"ababa"` 有没有更短的相等前后缀，可以拿来继续尝试延长？

先用最直观的方式，逐个检查更短的候选：

| 候选长度 | 旧字符串的前缀 | 旧字符串的后缀 | 能保留吗？ |
| --- | --- | --- | --- |
| 2 | `"ab"` | `"ba"` | 不行，它们本来就不相等 |
| 1 | `"a"` | `"a"` | 可以，尝试延长 |
| 0 | 空串 | 空串 | 可以，作为最后的候选 |

因此应该从长度 3 直接跳到长度 1。
但不逐个检查，怎么直接知道下一个是 1？

**（3）答案藏在已经相等的两段 "aba" 里面。**

我们已经知道：

```text
旧字符串的开头 = "aba"
旧字符串的结尾 = "aba"
```

对于任何比 3 小的候选长度 k：

- 整个旧字符串的前 k 个字符，就是开头那段 `"aba"` 的前 k 个字符。
- 整个旧字符串的后 k 个字符，就是结尾那段 `"aba"` 的后 k 个字符。
- 两段内容完全相同，都等于 `"aba"`。

所以：

> 比较整个旧字符串的较短前缀与后缀，等价于比较 `"aba"` 自己的前缀与后缀。

| 长度 | 整个 "ababa" 的前缀 / 后缀 | "aba" 自己的前缀 / 后缀 |
| --- | --- | --- |
| 2 | `"ab"` / `"ba"` | `"ab"` / `"ba"` |
| 1 | `"a"` / `"a"` | `"a"` / `"a"` |

**比较的内容完全一样！**

所以不必重新检查整个 `"ababa"`，只要问：
“`"aba"` 自己的最长相等真前后缀长度是多少？”

答案是 1，而这个结果早就算过了：

```text
pattern[:3] = "aba"
pi[2] = 1
```

**（4）为什么统一写成 pi[j-1]？**

```text
候选长度：j = 3
候选内容：pattern[:j] = "aba"
最后下标：j-1 = 2
```

pi 使用结束下标作为索引，所以 `pi[j-1]` 就是：

> 当前候选这段字符串，自己还有多长的相等真前后缀？

赋给 j：

```go
j = pi[j-1] // j = pi[2] = 1
```

意思是：**放弃接不上的长候选，换成下一个仍然合法的短候选。**
不是把长度随便减小，而是沿着已经确认前后缀相等的长度往下找。
不能直接用 j--，因为长度 2 的前后缀 `"ab"` 与 `"ba"` 本来就不相等。

上面的“比较内容完全相同”对所有更短的候选长度都成立，
所以沿 pi 链从长到短找，不会漏掉可能的答案。

**（5）换成长度 1 后，接着做什么？**

现在保留的旧前后缀都是 `"a"`，继续尝试接入同一个新字符：

```text
前缀一侧："a" + pattern[1] = "a" + 'b' → "ab"
后缀一侧："a" + pattern[5] = "a" + 'a' → "aa"
```

还是接不上，继续回退：

```go
j = pi[j-1] // j = pi[0] = 0
```

长度 0 表示没有旧字符可保留，但新字符 `'a'` 自己还能与 pattern[0] 匹配。
因此 j 从 0 增加到 1，最后得到 `pi[5] = 1`。

```text
候选回退：3 → 1 → 0
新字符接上：       0 → 1

"ababaa" 的最长相等真前后缀 = "a"
```

本轮 i 一直为 5，始终在处理同一个新字符 `'a'`，只有 j 在变化。

**把这条关系连起来：**

```text
旧字符串的开头和结尾，都是同一个长度 j 的字符串
                        ↓
寻找更短的相等前后缀
                        ↓
只需看这段长度 j 的字符串自己，有什么相等真前后缀
                        ↓
结果已经保存在 pi[j-1] 中
```

**第六步：什么时候保存 pi[i]？**

- 找到能接上的候选：执行 j++，得到新长度。
- 退到 j == 0：仍需比较 pattern[i] 与 pattern[0]；相等得到长度 1，否则为 0。
- 最后执行 pi[i] = j，保存本轮答案。

```text
读取旧答案 → 尝试延长 → 失败则缩短候选 → 保存新答案
```

由此自然得到下面的函数。构造时，比较之前始终有 `0 <= j < i`，
所以 pattern[j] 不会越界，也不会把整个当前字符串当成真前缀。
“j 等于模式串长度意味着完整匹配”的情形属于后面的 findAll，
不是 prefixFunction 构造过程中的状态。

#### Go 实现

```go
func prefixFunction(pattern string) []int {
    pi := make([]int, len(pattern))

    for i := 1; i < len(pattern); i++ {
        // 前一段的最长相等真前后缀长度。
        j := pi[i-1]

        // 新字符接不上，就尝试更短的候选。
        for j > 0 && pattern[i] != pattern[j] {
            j = pi[j-1]
        }

        // 能接上，长度增加 1。
        if pattern[i] == pattern[j] {
            j++
        }

        pi[i] = j
    }

    return pi
}
```

`pi[0]` 默认是 0，因为一个字符没有非空的相等真前后缀。
空串时循环不会执行，返回空表。

#### 2.1 先明确：每一轮到底在计算什么？

使用 `pattern = "ababaa"`，完整执行一次上面的函数。

```text
下标：   0 1 2 3 4 5
pattern: a b a b a a
```

每一轮有三个阶段，不要把它们混在一起：

1. **读取旧答案**：`j := pi[i-1]`，这是前一段的最长相等真前后缀长度。
2. **尝试延长**：用新字符 `pattern[i]` 与 `pattern[j]` 比较；接不上就缩短候选 j。
3. **保存新答案**：所有比较结束后，`pi[i] = j`。

注意：**刚进入这一轮时，j 还不是 pi[i] 的最终答案。**
它只是一个“准备尝试延长的旧长度”；成功接上后才增加 1。

对于当前候选长度 j，比较之前已经知道：

```text
pattern[:j] == pattern[i-j:i]
前面的 j 个字符 == 新字符之前、紧贴结尾的 j 个字符
```

左右两段都不包含新字符 `pattern[i]`。
如果再有 `pattern[j] == pattern[i]`，两边就可以同时延长一个字符：

```text
pattern[:j+1] == pattern[i-j:i+1]
```

这就是 `j++` 的依据，不是只因为两个字符相等就凭空得到一个长前后缀。

#### 2.2 初始化：pi[0] 为什么不用算？

执行：

```go
pi := make([]int, len(pattern))
```

Go 把数组元素初始化为 0：

```text
pi = [0, 0, 0, 0, 0, 0]
```

这时**只有 pi[0] 可以直接确定**：`"a"` 没有非空的相等真前后缀。
其他位置的 0 只是初始值，还没有经过计算。

下面用 `?` 表示“还没有计算”，区分它与真正算出的 0：

```text
已确定的结果：[0, ?, ?, ?, ?, ?]
```

外层循环从 `i = 1` 开始。

#### 2.3 第一轮：i = 1，计算 "ab" 的答案

读取旧答案：

```go
j := pi[0] // 0
```

j 为 0，说明没有已有的非空前后缀可延长，内层回退循环不执行。
但仍要比较新字符和模式串的第一个字符：

```text
pattern[i] = pattern[1] = 'b'
pattern[j] = pattern[0] = 'a'
```

不相等，`if` 不执行，j 保持 0。保存：

```go
pi[1] = 0
```

```text
当前字符串："ab"
相等真前后缀：无
已确定的结果：[0, 0, ?, ?, ?, ?]
```

#### 2.4 第二轮：i = 2，计算 "aba" 的答案

读取旧答案：

```go
j := pi[1] // 0
```

内层回退循环不执行。比较：

```text
pattern[i] = pattern[2] = 'a'
pattern[j] = pattern[0] = 'a'
```

相等，执行 `j++`，j 从 0 变为 1。保存 `pi[2] = 1`。

```text
当前字符串："aba"
前缀："a"，后缀："a"
已确定的结果：[0, 0, 1, ?, ?, ?]
```

这里从“没有保留字符”开始，匹配到了一个字符。

#### 2.5 第三轮：i = 3，计算 "abab" 的答案

读取旧答案：

```go
j := pi[2] // 1
```

旧字符串 `"aba"` 的开头和结尾已经有一个 `'a'` 相等：

```text
旧前缀：pattern[0:1] = "a"
旧后缀：pattern[2:3] = "a"
```

只需检查两边接下来的字符：

```text
pattern[j] = pattern[1] = 'b'  ← 前缀接下来的字符
pattern[i] = pattern[3] = 'b'  ← 新加入的字符
```

相等，内层回退循环不执行；`if` 中执行 `j++`，从 1 变为 2。

```text
旧前缀 "a" + 'b' = "ab"
旧后缀 "a" + 'b' = "ab"
```

保存 `pi[3] = 2`：

```text
已确定的结果：[0, 0, 1, 2, ?, ?]
```

注意：不需要重新逐个比较两段 `"ab"`，旧的 `'a'` 相等已经由 pi[2] 保证。

#### 2.6 第四轮：i = 4，计算 "ababa" 的答案

读取旧答案：

```go
j := pi[3] // 2
```

旧字符串 `"abab"` 的前缀和后缀都是 `"ab"`。比较新增字符：

```text
pattern[j] = pattern[2] = 'a'
pattern[i] = pattern[4] = 'a'
```

相等，不回退，执行 `j++`，从 2 变为 3：

```text
旧前缀 "ab" + 'a' = "aba"
旧后缀 "ab" + 'a' = "aba"
```

保存 `pi[4] = 3`：

```text
已确定的结果：[0, 0, 1, 2, 3, ?]
```

#### 2.7 第五轮：i = 5，计算 "ababaa" 的答案（重点）

读取旧答案：

```go
j := pi[4] // 3
```

旧字符串 `"ababa"` 的前缀和后缀都是 `"aba"`。
我们先尝试把它们延长成长度 4：

```text
前缀一侧："aba" + pattern[3] = "abab"
后缀一侧："aba" + pattern[5] = "abaa"
```

新字符 `'b'` 和 `'a'` 不同，不能延长。

**第一次回退：j 从 3 变为 1。**

```go
j = pi[j-1] // pi[2] = 1
```

查的是旧候选 `"aba"` 自己的最长相等真前后缀，得到 `"a"`。
所以可以保留旧字符串结尾的一个 `'a'`，尝试接上新字符：

```text
前缀一侧："a" + pattern[1] = "ab"
后缀一侧："a" + pattern[5] = "aa"
```

还是 `'b'` 与 `'a'` 不同，不能延长。

**第二次回退：j 从 1 变为 0。**

```go
j = pi[j-1] // pi[0] = 0
```

j 为 0，内层循环停止。此时没有旧字符可保留，
但仍要判断新字符自己能否与模式串开头匹配：

```text
pattern[j] = pattern[0] = 'a'
pattern[i] = pattern[5] = 'a'
```

相等，执行 `j++`，从 0 变为 1，最后保存：

```go
pi[5] = 1
```

```text
当前字符串："ababaa"
最长相等真前缀："a"
最长相等真后缀："a"
最终结果：[0, 0, 1, 2, 3, 1]
```

**这一轮 i 一直等于 5，只调整 j，反复拿同一个新字符 'a' 尝试接上。**
回退阶段的 `3 → 1 → 0` 是旧候选长度；最后的 `0 → 1` 才是新字符匹配成功后的延长。

#### 2.8 把所有轮次放在一起

| i | 当前考察字符串 | 初始 j = pi[i-1] | 比较与回退 | 最终 pi[i] |
| --- | --- | --- | --- | --- |
| 0 | `"a"` | 不适用 | 单字符，直接确定 | 0 |
| 1 | `"ab"` | 0 | `'b' != 'a'`，保持 0 | 0 |
| 2 | `"aba"` | 0 | `'a' == 'a'`，0 → 1 | 1 |
| 3 | `"abab"` | 1 | `'b' == 'b'`，1 → 2 | 2 |
| 4 | `"ababa"` | 2 | `'a' == 'a'`，2 → 3 | 3 |
| 5 | `"ababaa"` | 3 | 失配：3 → 1 → 0；匹配：0 → 1 | 1 |

#### 2.9 为什么这三个细节不能省？

- **为什么是 `for`，不是 `if`？** 最后一轮需要连续回退两次，只退一次还没有找到能接上的候选。
- **为什么 j == 0 后还要比较一次？** 0 表示没有旧字符可保留，不代表新字符不能与 pattern[0] 匹配。
- **为什么每轮最后才写 pi[i]？** 中途的 j 只是候选长度，回退和延长完成后才是当前答案。

再深入一点：假设新答案长度为 L > 0，去掉最后一个字符后，
剩余长度 L-1 必须是旧字符串的某个相等真前后缀（也可以是空串）。
因此不需要枚举所有长度，只要从旧的最长候选开始，沿 pi 链依次尝试更短候选。
候选从长到短，第一个能接上的就是当前最长答案。

### 3. Go 模板：查找全部匹配位置

#### 设计思路：j 的含义与失配回退

prefixFunction 已经生成了 pi，现在用它在文本中查找模式串。
这里的 i 扫描 text，j 表示当前已经匹配的模式串前缀长度。

例如 j = 3 时，已经匹配的是 pattern[:3]，下一个要比较的是 pattern[3]。
长度为 j 的前缀占据下标 0 到 j-1，因此 j 恰好也是下一个待匹配字符的下标。
这与 prefixFunction 中“候选长度决定下一位比较位置”的思路相同，
区别是：这里比较 text[i] 和 pattern[j]，构造前缀表时比较 pattern[i] 和 pattern[j]。

当 j == len(pattern) 时说明完整匹配，不能再访问 pattern[j]；
应先记录起点，再回退 j，为下一次匹配保留可复用的后缀。

下面先走一遍失配回退的例子，再看 Go 实现。

```text
模式串：a b a b c
pi：   0 0 1 2 0
```

要在文本 `"abababc"` 中找 `"ababc"`。

#### 第一步：前四个字符匹配成功

```text
文本： a b a b a b c
模式： a b a b c
       ✓ ✓ ✓ ✓
```

此时 `j = 4`，已经匹配 `"abab"`。接着比较：

```text
text[4]    = 'a'
pattern[4] = 'c'
```

失配了。

#### 第二步：之前的成功比较，能留下什么？

已经匹配的 `"abab"`，末尾是 `"ab"`；模式串的开头也是 `"ab"`。
所以可以把模式串移到这里：

```text
文本： a b a b a b c
模式：     a b a b c
           ✓ ✓
```

这两个字符不需要重新比较，因为已经知道它们相等。
因此不是回到 `j = 0`，而是保留长度为 2 的匹配，即 `j = 2`。

#### 第三步：这个 2 从哪里来？

刚刚已经匹配的内容是 `pattern[:4] = "abab"`。
它的最长相等真前后缀长度，就是 `pi[3] = 2`。
而刚才 `j = 4`，所以：

```go
j = pi[j-1] // pi[3] = 2
```

这句的实际意思是：

> 在已经匹配的那一段中，找一个最长的真后缀，让它同时也是模式串的前缀，保留下来继续用。

为什么是 `j-1`？因为已经匹配的长度为 j，这一段的最后一个下标就是 j-1。

为什么不是 `j--`？因为更短的候选必须仍然满足“前缀等于后缀”。
例如 `"abab"` 的长度 3 前缀为 `"aba"`、后缀为 `"bab"`，并不相等；
长度 2 才是下一个可用候选，pi 帮我们直接跳过去。

#### 第四步：为什么文本下标 i 不回退？

刚才失配的是 `text[4] = 'a'`。回退 j 后，仍然拿**同一个文本字符**比较：

```text
text[4]    = 'a'
pattern[2] = 'a'
```

这次相等，j 从 2 增加到 3。之后继续：

```text
text[5] = 'b'，与 pattern[3] 匹配 → j = 4
text[6] = 'c'，与 pattern[4] 匹配 → j = 5
```

模式串长度为 5，j == 5，找到完整匹配：

```text
结束下标 = 6
模式长度 = 5
起点     = 6 - 5 + 1 = 2
```

**文本不用往回看；模式串利用已有信息调整匹配位置。**
如果回退一次后仍然失配，就继续沿 `pi[j-1]` 回退，直到能接上或 j 为 0。

#### Go 实现：findAll

这里约定：**空模式串匹配文本的所有字节边界**，即 0 到 len(text)。
先处理空模式串，可以避免访问 pattern[0] 时越界。

```go
func findAll(text, pattern string) []int {
    positions := []int{}

    if len(pattern) == 0 {
        for i := 0; i <= len(text); i++ {
            positions = append(positions, i)
        }
        return positions
    }

    pi := prefixFunction(pattern)
    j := 0

    for i := 0; i < len(text); i++ {
        // 当前文本字符接不上，调整模式串的匹配长度。
        for j > 0 && text[i] != pattern[j] {
            j = pi[j-1]
        }

        if text[i] == pattern[j] {
            j++
        }

        if j == len(pattern) {
            positions = append(positions, i-len(pattern)+1)

            // 保留可以复用的后缀，继续找下一个匹配。
            j = pi[j-1]
        }
    }

    return positions
}
```

第 28 题只需第一个匹配位置，可以找到后立即返回；findAll 则收集所有匹配位置。

#### 完整匹配后，为什么还要回退？

例如：

```text
text    = "aaaa"
pattern = "aa"
pi      = [0, 1]
```

匹配完前两个 `'a'`：

```text
文本： a a a a
模式： a a
```

已经匹配的 `"aa"` 的末尾 `'a'`，也能作为下一次匹配的开头：

```text
文本： a a a a
模式：   a a
```

所以保留一个字符：

```go
j = pi[j-1] // 此时 j = 2，pi[1] = 1
```

下一个文本字符再匹配成功，就找到起点 1，依次得到 `[0, 1, 2]`。
如果每次完整匹配后都把 j 清零，就会漏掉中间起点 1 的重叠匹配。

### 4. 复杂度与字符单位

令 m 为模式串长度，n 为文本长度，k 为匹配结果数量。

- 构造前缀表：O(m) 时间，O(m) 空间。
- 使用前缀表扫描文本：O(n) 时间；写入结果需要 O(k) 时间和空间。
- findAll 总计：O(m+n+k) 时间，O(m+k) 空间，也适用于上述空模式串约定。

虽然代码有内层回退循环，但 j 每次回退都会减小，每次成功比较最多增加 1。
扫描过程中增加的总次数受扫描长度限制，因此总回退次数也是线性的，
并不是每个 i 都要从头比较一遍。

上面的 Go 模板按字节索引 string，长度与返回位置都是字节单位。
示例使用 ASCII 字符，字符数与字节数一致。
如需按 Unicode 码点匹配，应统一改为 rune 序列，并明确返回下标的含义。

### 5. 最核心的一句话

> 失配不代表之前匹配的信息全部作废。pi[j-1] 告诉我们，已经匹配的这段结尾，有多少字符还能作为模式串的开头继续使用。 

## 四、每道题要提炼的结论

- [28. 找出字符串中第一个匹配项的下标](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/)：把“构造 pi”和“扫描文本”分开，先掌握基础匹配流程。
- [1392. 最长快乐前缀](https://leetcode.com/problems/longest-happy-prefix/)：非空字符串的答案长度就是 pi[n-1]，返回 s[:pi[n-1]]。
- [459. 重复的子字符串](https://leetcode.com/problems/repeated-substring-pattern/)：令 p = n-pi[n-1]；当 pi[n-1] > 0 且 n%p == 0 时，字符串可由长度为 p 的子串重复构成。
- [686. 重复叠加字符串匹配](https://leetcode.com/problems/repeated-string-match/)：令 q = ceil(len(b)/len(a))，只需检查 a 重复 q 次和 q+1 次的文本；匹配可能跨拼接边界。也可以按下标取模扫描虚拟重复文本。
- [214. 最短回文串](https://leetcode.com/problems/shortest-palindrome/)：对 s + 分隔符 + reverse(s) 构造 pi，末项给出 s 的最长回文前缀长度；反转剩余后缀并补到开头。分隔符必须不出现在原串中。
- [3036. 匹配模式数组的子数组数目 II](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/)：把相邻数值比较结果编码为 -1、0、1，KMP 同样适用于整数序列；一次完整匹配对应一个合法子数组。
- [3008. 找出数组中的美丽下标 II](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/)：分别获取 a、b 的全部出现起点；对 a 的每个起点 i，检查 b 是否有起点落在 [i-k,i+k]，可使用二分或双指针。

## 五、扩展题单与分类总结

原主线 7 题加上下面新增的 11 题，共 18 题。
这些题都能联系到 KMP／前缀函数，但不意味着 KMP 是每题唯一或最简的解法。
先完成基础题，再按训练目标选择扩展，不必一次做完。

### A. 基础查找与循环文本

| 题目 | 如何转成模板问题 | 练完应掌握什么 |
| --- | --- | --- |
| [796. 旋转字符串](https://leetcode.com/problems/rotate-string/) | 先判断 s 与 goal 等长，再在 s+s 中查找 goal | 把循环位移变成普通子串匹配；是进入 686 前的热身题 |
| [1408. 数组中的字符串匹配](https://leetcode.com/problems/string-matching-in-an-array/) | 枚举不同单词，用一个作模式串，在另一个中查找 | 多次复用查找模板；数据较小，内置查找也足够，不是性能挑战题 |
| [1668. 最大重复子字符串](https://leetcode.com/problems/maximum-repeating-substring/) | 找到 word 在 sequence 中的所有出现起点，再连接间隔恰好为 len(word) 的匹配 | 出现多次不等于连续重复；可令 dp[i] = dp[i-m]+1（起点 i 存在匹配时，m 为 word 长度） |

建议顺序：28 → 796 → 686；1408、1668 作为可选巩固。
1668 的计数依赖相邻副本紧密衔接，不能只统计 findAll 的结果个数。

### B. 全部匹配位置与整数序列

| 题目 | 如何转成模板问题 | 练完应掌握什么 |
| --- | --- | --- |
| [3034. 匹配模式数组的子数组数目 I](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/) | 将相邻元素比较编码为 -1、0、1，再匹配 pattern | 作为 3036 的小规模热身，也可以先暴力写出对照答案 |
| [1764. 通过连接另一个数组的子数组得到一个数组](https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/) | 按顺序匹配每个 group，找到后从该匹配的末尾之后继续 | KMP 不限于字符串；每组贪心选最早匹配，为后续组留下最多空间 |
| [3006. 找出数组中的美丽下标 I](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/) | 分别找 a、b 的全部起点，再检查起点距离是否不超过 k | 作为 3008 的热身；先用简单方法合并结果，再升级为双指针或二分 |

整数序列版只需把比较对象改成整数元素，pi 构造和回退逻辑不变。
3034／3036 中，长度 m 的比较模式对应原数组长度 m+1 的子数组。
1764 是多个模式依次匹配，不能把 groups 直接拼成一个模式，因为组与组之间允许有间隔。

### C. 前缀表、border 链与周期

border 指一个字符串相等的真前缀与后缀。
非空串的全部非空 border 长度可以这样枚举：

```text
L = pi[n-1]
不断记录 L，并令 L = pi[L-1]，直到 L == 0
```

这条链除了用于失配回退，也可以用来枚举字符串的自重叠结构。

| 题目 | 如何转成前缀表问题 | 练完应掌握什么 |
| --- | --- | --- |
| [3029. 将单词恢复初始状态所需的最短时间 I](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/) | 操作 t 次且 t*k<n 时，需要 word[t*k:] == word[:n-t*k]，也就是 n-t*k 是一个 border 长度 | 把操作过程变成“后缀是否等于前缀”；先在小数据下直接比较 |
| [3031. 将单词恢复初始状态所需的最短时间 II](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/) | 沿 border 链寻找长度 L，使 (n-L)%k == 0；候选时间为 (n-L)/k | 用前缀函数优化 3029；若无可用非空 border，答案为 ceil(n/k) |

3031 中先取不依赖 border 的兜底答案 ceil(n/k)，再找更小的合法时间。
与 459 区分：459 要求整个字符串由子串完整重复，因此有整除条件；
3029／3031 允许操作时补入任意字符，重点是保留下来的后缀能否与原串前缀对齐。
例如 word = "ababa"、k = 2，保留 "aba" 就能一步恢复，不要求 n 能被 k 整除。

### D. KMP 状态与其他数据结构结合

| 题目 | 如何复用 KMP | 额外知识与易错点 |
| --- | --- | --- |
| [1910. 删除一个字符串中所有出现的给定子字符串](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/) | 从左到右压栈，每个栈位置保存对应的 KMP 匹配长度；完整匹配时弹出模式串长度个字符 | 删除后必须恢复新栈顶保存的 j，不能一律归零；这可处理删除后新形成的匹配 |
| [1397. 找到所有好字符串](https://leetcode.com/problems/find-all-good-strings/) | 将 j 看作已构造字符串末尾匹配 evil 前缀的长度；j 达到 len(evil) 的转移直接舍弃 | KMP 自动机 + 数位 DP，另外维护上下界限制；不是简单调用 findAll |
| [2851. 字符串转换](https://leetcode.com/problems/string-transformation/) | 在 s+s 中匹配 t，统计起点位于 [0,n) 的有效循环位移 | KMP + 计数 DP／矩阵快速幂；重复字符可能对应多个有效位移，起点 n 不能重复计数 |

1910 适合在掌握 findAll 后练“状态恢复”。
1397、2851 需要额外算法基础，现阶段可以跳过，不要把它们当作基础模板题。

### E. 看到题目时，如何选择模板？

| 题目特征 | 优先考虑的工具 | 代表题 |
| --- | --- | --- |
| 找一个或所有完整出现位置 | findAll 或找到后提前返回 | [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/)、[3008](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/) |
| 问字符串自己的相等前后缀 | prefixFunction，读取末项或沿 border 链回退 | [1392](https://leetcode.com/problems/longest-happy-prefix/)、[3031](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/) |
| 问完整重复或循环移位 | 前缀表判断周期，或在双倍／重复文本中匹配 | [459](https://leetcode.com/problems/repeated-substring-pattern/)、[796](https://leetcode.com/problems/rotate-string/) |
| 匹配的是大小关系而不是数值 | 先编码，再对整数序列匹配 | [3036](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/) |
| 要把回文变成前后缀关系 | 反转、加入不冲突的分隔符，再构造 pi | [214](https://leetcode.com/problems/shortest-palindrome/) |
| 删除、构造或计数过程中要持续维护匹配进度 | 把 j 当作状态，与栈或 DP 结合 | [1910](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/)、[1397](https://leetcode.com/problems/find-all-good-strings/) |

### F. 推荐练习顺序与阶段目标

1. **第一阶段：理解与手写模板。** [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) → [1392](https://leetcode.com/problems/longest-happy-prefix/) → [459](https://leetcode.com/problems/repeated-substring-pattern/)。能解释 j、pi[j-1]，不靠背代码。
2. **第二阶段：改变文本或序列表示。** [796](https://leetcode.com/problems/rotate-string/) → [686](https://leetcode.com/problems/repeated-string-match/) → [214](https://leetcode.com/problems/shortest-palindrome/) → [3034](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/) → [3036](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/)。学会先变换问题，再套模板。
3. **第三阶段：处理匹配结果。** [1764](https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/) → [3006](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/) → [3008](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/)。学会顺序、非重叠与距离约束。
4. **第四阶段：理解 border 链与状态恢复。** [3029](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/) → [3031](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/) → [1910](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/)。从“能匹配”进阶到“利用匹配结构”。
5. **可选热身与进阶。** [1408](https://leetcode.com/problems/string-matching-in-an-array/)、[1668](https://leetcode.com/problems/maximum-repeating-substring/) 可用于巩固；[1397](https://leetcode.com/problems/find-all-good-strings/)、[2851](https://leetcode.com/problems/string-transformation/) 等掌握数位 DP、矩阵快速幂后再做。

总结：真正复用的不是一份完整题解，而是“构造 pi、失配回退、保留匹配状态”这三件事。
基础题直接用模板；变形题先确定文本和模式分别是什么；综合题再决定如何处理匹配位置或状态。

## 六、练习检查清单

- [ ] 能手算 "abab"、"aaaa"、"ababc" 的前缀表。
- [ ] 能解释为什么失配时用 j = pi[j-1]，而不是 j--。
- [ ] 能独立写出 prefixFunction。
- [ ] 完成 [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/)，并覆盖无匹配、开头匹配、末尾匹配等情况。
- [ ] 完成 [1392](https://leetcode.com/problems/longest-happy-prefix/)，区分“长度”和“下标”。
- [ ] 完成 [459](https://leetcode.com/problems/repeated-substring-pattern/)，并解释周期公式与整除条件。
- [ ] 能写出 findAll，并用重叠匹配验证。
- [ ] 再推进 [686](https://leetcode.com/problems/repeated-string-match/)、[214](https://leetcode.com/problems/shortest-palindrome/)，最后做 [3036](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/)、[3008](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/)。
- [ ] 能用 [796](https://leetcode.com/problems/rotate-string/) 解释为什么循环移位可以变成双倍文本匹配。
- [ ] 能用 [1764](https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/) 把模板迁移到整数序列，并处理非重叠约束。
- [ ] 能用 [3031](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/) 枚举 border 链，而不仅仅将它用于失配回退。
- [ ] 能用 [1910](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/) 解释删除字符后为什么需要恢复栈顶的匹配状态。
