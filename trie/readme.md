# 字典树 Trie（前缀树）：共享前缀、节点信息与搜索规则

[返回仓库目录](../README.md)

> 固定骨架：**从根出发，按字符沿边向下走。** 变化点：节点保存什么、什么时候停止、是否需要搜索多个分支。
>
> 本文先针对小写字母 `a..z`。README 提供参考实现，`l题号/` 中只保留空函数供练习，原有旧稿不改动。

## 目录

- [1. 解决什么问题？思路从哪里来？](#why)
- [2. 节点、前缀与完整单词](#node)
- [3. 基础模板：208. 实现 Trie](#lc208)
- [4. 最短词根：648. 单词替换](#lc648)
- [5. 通配符分支：211. 添加与搜索单词](#lc211)
- [6. 前缀聚合：677. 键值映射](#lc677)
- [7. 字典序候选：1268. 搜索推荐系统](#lc1268)
- [8. 网格与前缀剪枝：212. 单词搜索 II](#lc212)
- [9. 分类扩展题单](#practice)
- [10. 复杂度、易错点与复盘](#review)
- [原有资料与旧稿](#legacy)

<a id="why"></a>
## 1. 解决什么问题？思路从哪里来？

适合的场景：维护许多字符串，反复查询单词是否存在、某个前缀是否存在、前缀对应多少个词/多大权值，或者在搜索过程中不断判断“当前路径还有没有可能组成字典里的词”。

从朴素方法出发：有 `apple、app、apt、bat`，每次查询 `ap`，可以遍历所有单词逐个判断前缀。问题在于反复比较相同的 `a`、`p`，没有复用公共部分。

把公共前缀合并成一条路径：

```text
root（空前缀）
├── a
│   └── p
│       ├── p *          app 是完整词
│       │   └── l
│       │       └── e *  apple 是完整词
│       └── t *          apt 是完整词
└── b
    └── a
        └── t *          bat 是完整词
```

树中的**边表示一个字符，根到节点的路径表示一个前缀**。多个单词共享路径，不必每次从头比较每个单词。

这就是 Trie 的来源：不是按整个单词建立哈希索引，而是按“下一字符”逐层组织信息。查询长度为 L 的前缀，通常只沿 L 条边，不必扫描整个字典。

### 怎样理解“走到一个节点”？

已经匹配了 `ap`，当前指针就停在代表 `ap` 的节点。它的孩子表示后面允许出现的字符：`p` 或 `t`。若下一字符为 `z`，没有对应边，立刻知道**没有任何字典单词以 apz 开头**。

Trie 不只是存储完整词，还存储“哪些中间状态仍有希望”。这正是它与 DFS/回溯结合时能剪枝的原因。

### 不是所有字符串问题都需要 Trie

| 需求 | 优先考虑 |
| --- | --- |
| 只查完整字符串是否存在 | 哈希集合，通常更简单 |
| 频繁查询前缀、沿路径剪枝 | Trie |
| 静态词典中按前缀取候选 | Trie 或排序＋二分 |
| 求几个字符串的最长公共前缀 | 直接逐字符比较通常足够 |
| 大字符集且分支稀疏 | `map[rune]*Node` 或按实际 token 建边 |

Trie 的代价是节点和指针较多；不应为了套模板忽略更简单的解法。

<a id="node"></a>
## 2. 节点、前缀与完整单词

```go
type TrieNode struct {
    Children [26]*TrieNode
    IsEnd    bool
}
```

- `Children[c-'a']`：读取字符 c 后到达的下一个节点，nil 表示这条边不存在。
- `IsEnd`：是否有完整单词恰好在当前节点结束。
- 根不表示任何字符，代表空前缀；不需要额外保存当前字符。

只插入 `apple`：

```text
StartsWith("app") = true    路径存在
Search("app")     = false   路径存在，但不是完整词
Search("apple")   = true    路径存在，且结束标记为true
```

**节点存在不等于单词存在。** 插入时只把最后一个节点标为结束，中间节点不自动成为完整词。

本篇代码按字节遍历，约定输入字符符合各题的小写英文字母约束。中文不能直接用 `word[i]-'a'`；应改用 rune 遍历与 map 子节点。路径分隔符、数字和二进制位也需要对应的子节点结构。

<a id="lc208"></a>
## 3. 基础模板：208. 实现 Trie

[LeetCode 208](https://leetcode.com/problems/implement-trie-prefix-tree/) · [练习](l208/208.go)

**题意**：实现插入、完整单词查询、前缀查询。

**不变量**：处理完前 i 个字符后，node 表示 `word[:i]` 对应的节点。插入时缺边就创建；查询时缺边就失败。

下面是独立参考代码。后续各题也分别给出独立代码，不需要跨题复制类型，避免过度泛型或回调封装。

```go
package template

type TrieNode struct {
    Children [26]*TrieNode
    IsEnd    bool
}

type Trie struct { root *TrieNode }

func Constructor() Trie { return Trie{root: &TrieNode{}} }

func (t *Trie) Insert(word string) {
    node := t.root
    for i := 0; i < len(word); i++ {
        c := word[i] - 'a'
        if node.Children[c] == nil { node.Children[c] = &TrieNode{} }
        node = node.Children[c]
    }
    node.IsEnd = true
}

func (t *Trie) find(s string) *TrieNode {
    node := t.root
    for i := 0; i < len(s); i++ {
        node = node.Children[s[i]-'a']
        if node == nil { return nil }
    }
    return node
}

func (t *Trie) Search(word string) bool {
    node := t.find(word)
    return node != nil && node.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool { return t.find(prefix) != nil }
```

**复杂度**：插入/查询长度为 L 的字符串，时间 `O(L)`。累计插入字符总数为 S，固定26分支时节点空间 `O(S+1)`，但每个节点有26个指针，常数不小。

按本模板构造后，空前缀始终存在；只有显式插入空串，根才标为结束。208原题操作参数为非空字符串。调用前使用 Constructor 初始化 root。

<a id="lc648"></a>
## 4. 最短词根：648. 单词替换

[LeetCode 648](https://leetcode.com/problems/replace-words/) · [练习](l648/648.go)

**题意**：用字典中最短的词根替换句子中的单词。例：字典 `["cat","bat","rat"]`，`"the cattle was rattled"` 变为 `"the cat was rat"`。

**变化点**：插入不变；查询从左到右，一遇到 `IsEnd` 立即返回，保证选到最短词根。若路径中断或没有结束标记，保留原词。

```go
package template

import "strings"

type rootNode struct {
    children [26]*rootNode
    isEnd bool
}

func replaceWords(dictionary []string, sentence string) string {
    root := &rootNode{}
    for _, word := range dictionary {
        node := root
        for i := 0; i < len(word); i++ {
            c := word[i]-'a'
            if node.children[c] == nil { node.children[c] = &rootNode{} }
            node = node.children[c]
        }
        node.isEnd = true
    }
    words := strings.Split(sentence, " ")
    for j, word := range words {
        node := root
        for i := 0; i < len(word); i++ {
            node = node.children[word[i]-'a']
            if node == nil { break }
            if node.isEnd {
                words[j] = word[:i+1]
                break
            }
        }
    }
    return strings.Join(words, " ")
}
```

**复杂度**：字典总字符数 S，句子长度 C，时间 `O(S+C)`，额外空间 `O(S+C)`（含拆分与结果）。题设词根非空，句子单词以单个空格分隔。

<a id="lc211"></a>
## 5. 通配符分支：211. 添加与搜索单词

[LeetCode 211](https://leetcode.com/problems/design-add-and-search-words-data-structure/) · [练习](l211/211.go)

**题意**：添加单词；搜索模式中的 `.` 匹配任意**一个**字母。例如已加入 `bad`，则 `.ad`、`b..` 为true，`ba` 为false。

**变化点**：普通字符只走一条边；点号尝试所有非空孩子。搜索状态为 `(当前节点,模式下标)`，模式消耗完仍须检查结束标记。

```go
package template

type wordNode struct {
    children [26]*wordNode
    isEnd bool
}

type WordDictionary struct { root *wordNode }

func Constructor() WordDictionary { return WordDictionary{root: &wordNode{}} }

func (d *WordDictionary) AddWord(word string) {
    node := d.root
    for i := 0; i < len(word); i++ {
        c := word[i]-'a'
        if node.children[c] == nil { node.children[c] = &wordNode{} }
        node = node.children[c]
    }
    node.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
    var dfs func(*wordNode, int) bool
    dfs = func(node *wordNode, i int) bool {
        if node == nil { return false }
        if i == len(word) { return node.isEnd }
        if word[i] != '.' { return dfs(node.children[word[i]-'a'], i+1) }
        for _, child := range node.children {
            if child != nil && dfs(child, i+1) { return true }
        }
        return false
    }
    return dfs(d.root, 0)
}
```

**复杂度**：插入 `O(L)`；含 d 个点号的长度 L 搜索，粗略上界 `O(L·26^d)`，实际只访问已有分支，不能声称仍为 `O(L)`。递归栈 `O(L)`。点号不是正则里的 `.*`，不能跳过字符或匹配任意长度。

<a id="lc677"></a>
## 6. 前缀聚合：677. 键值映射

[LeetCode 677](https://leetcode.com/problems/map-sum-pairs/) · [练习](l677/677.go)

**题意**：`Insert(key,val)` 插入或覆盖键值；`Sum(prefix)` 返回以此前缀开头的所有键的值之和。

**节点信息变化**：增加 `sum`，表示当前前缀下所有完整键的总值。用 map 保存旧值，每次只传播 `delta=new-old`，而不是把新值重复加一遍。

```text
Insert("apple",3) → Sum("ap")=3
Insert("app",2)   → Sum("ap")=5
Insert("apple",2) → delta=-1，Sum("ap")=4
```

```go
package template

type sumNode struct {
    children [26]*sumNode
    sum int
}

type MapSum struct {
    root *sumNode
    values map[string]int
}

func Constructor() MapSum {
    return MapSum{root: &sumNode{}, values: map[string]int{}}
}

func (m *MapSum) Insert(key string, val int) {
    delta := val - m.values[key]
    m.values[key] = val
    node := m.root
    node.sum += delta
    for i := 0; i < len(key); i++ {
        c := key[i]-'a'
        if node.children[c] == nil { node.children[c] = &sumNode{} }
        node = node.children[c]
        node.sum += delta
    }
}

func (m *MapSum) Sum(prefix string) int {
    node := m.root
    for i := 0; i < len(prefix); i++ {
        node = node.children[prefix[i]-'a']
        if node == nil { return 0 }
    }
    return node.sum
}
```

**复杂度**：插入长度 L 的键为 `O(L)`，查询长度 P 的前缀为 `O(P)`；空间为Trie节点和旧值map，总计与所存不同键的字符量同阶。无需 `IsEnd`：此题不查询完整键存在性，旧值由map负责。

<a id="lc1268"></a>
## 7. 字典序候选：1268. 搜索推荐系统

[LeetCode 1268](https://leetcode.com/problems/search-suggestions-system/) · [练习](l1268/1268.go)

**题意**：输入搜索词的每个前缀后，返回字典序最小的至多3个商品。例：`["mobile","mouse","moneypot","monitor","mousepad"]`，前缀 `m` 的结果为 `["mobile","moneypot","monitor"]`，前缀 `mou` 为 `["mouse","mousepad"]`。

**节点信息变化**：在每个前缀节点缓存前三个词。先复制并排序商品，再按字典序插入；节点仅在缓存不足3个时添加当前词，就自然保留最小三个，不必每次搜索整棵子树。

```go
package template

import "sort"

type suggestionNode struct {
    children [26]*suggestionNode
    top []string
}

func suggestedProducts(products []string, searchWord string) [][]string {
    sorted := append([]string(nil), products...)
    sort.Strings(sorted)
    root := &suggestionNode{}
    for _, word := range sorted {
        node := root
        for i := 0; i < len(word); i++ {
            c := word[i]-'a'
            if node.children[c] == nil { node.children[c] = &suggestionNode{} }
            node = node.children[c]
            if len(node.top) < 3 { node.top = append(node.top, word) }
        }
    }
    answer := make([][]string, len(searchWord))
    node := root
    for i := 0; i < len(searchWord); i++ {
        if node != nil { node = node.children[searchWord[i]-'a'] }
        answer[i] = []string{}
        if node != nil { answer[i] = append(answer[i], node.top...) }
    }
    return answer
}
```

**复杂度**：n个商品，总字符数 S，最长商品长度 L，查询长度 Q。排序最坏 `O(n·L·log n)`，构建 `O(S)`，查询及输出 `O(Q)`（每次至多3个字符串引用），空间 `O(S+n+Q)`，不计原有字符串内容。商品在题设中互不相同；如果改为重复商品输入，要先明确是否去重。

另一条路线是查到前缀节点后按 `a..z` 做DFS，收集到3个立即停止；注意 map 遍历无序，且达到3个后必须停止继续搜。这两点不能仅靠结束标记保证。静态数据也可采用排序＋二分，不一定非用Trie。

<a id="lc212"></a>
## 8. 网格与前缀剪枝：212. 单词搜索 II

[LeetCode 212](https://leetcode.com/problems/word-search-ii/) · [练习](l212/212.go)

**题意**：在字母网格中，通过上下左右相邻格拼出字典中的单词。同一路径不能重复使用格子，同一答案只返回一次。

**核心来源**：若对每个单词分别做79的搜索，会重复探索相同前缀。把所有词放进Trie，一次网格回溯同时匹配整个字典；当Trie中不存在当前前缀，立即停止这条路径。

**状态**：`(row,col,当前前缀节点)`，加上当前路径占用的格子。节点保存完整词 `word`，找到后清空以避免重复输出。

```go
package template

type boardNode struct {
    children [26]*boardNode
    word string
}

func findWords(board [][]byte, words []string) []string {
    answer := []string{}
    if len(board) == 0 || len(board[0]) == 0 { return answer }
    root := &boardNode{}
    for _, word := range words {
        node := root
        for i := 0; i < len(word); i++ {
            c := word[i]-'a'
            if node.children[c] == nil { node.children[c] = &boardNode{} }
            node = node.children[c]
        }
        node.word = word
    }
    rows, cols := len(board), len(board[0])
    dirs := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(int, int, *boardNode)
    dfs = func(r, c int, parent *boardNode) {
        if r < 0 || r >= rows || c < 0 || c >= cols { return }
        ch := board[r][c]
        if ch == '#' { return }
        node := parent.children[ch-'a']
        if node == nil { return } // 没有任何单词能沿此前缀继续
        if node.word != "" {
            answer = append(answer, node.word)
            node.word = "" // 去重，但不能停止：可能还有更长的词
        }
        board[r][c] = '#'
        for _, d := range dirs { dfs(r+d[0], c+d[1], node) }
        board[r][c] = ch // 恢复，其他起点仍能使用这个格子
    }
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ { dfs(r,c,root) }
    }
    return answer
}
```

例：同一条路径可以先找到 `app` 再找到 `apple`，所以“找到完整词”不等于“整条分支完成”。清空结束信息只用于去重，不代表无条件删除其孩子。

**复杂度**：字典总字符数 S，网格格子数 B，最长词长 L。保守时间上界 `O(S+B·4^L)`，实际受前缀剪枝、不能重复格子及网格大小限制；不是线性算法。辅助空间 `O(S+min(L,B))`，不含答案。搜索时暂改网格，但每条路径返回时恢复，函数结束后网格不变。

<a id="practice"></a>
## 9. 分类扩展题单

### 基础路线：先掌握六个变化点

| 原题 | 本地练习 | 节点信息 / 搜索规则 |
| --- | --- | --- |
| [208. 实现 Trie](https://leetcode.com/problems/implement-trie-prefix-tree/) | [l208](l208/208.go) | Children＋IsEnd，插入/查词/查前缀 |
| [648. 单词替换](https://leetcode.com/problems/replace-words/) | [l648](l648/648.go) | 遇到第一个完整词根就返回 |
| [211. 添加与搜索单词](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | [l211](l211/211.go) | 点号分支搜索，模式耗尽再查结束 |
| [677. 键值映射](https://leetcode.com/problems/map-sum-pairs/) | [l677](l677/677.go) | 前缀sum，覆盖只传播差值 |
| [1268. 搜索推荐系统](https://leetcode.com/problems/search-suggestions-system/) | [l1268](l1268/1268.go) | 缓存字典序前3名，或按序DFS并及时停止 |
| [212. 单词搜索 II](https://leetcode.com/problems/word-search-ii/) | [l212](l212/212.go) | 网格回溯＋Trie前缀剪枝＋答案去重 |

### 前缀统计、词典路径与切分

| 原题 | 本地练习 | 迁移点与边界 |
| --- | --- | --- |
| [14. 最长公共前缀](https://leetcode.com/problems/longest-common-prefix/) | [l14](l14/14.go) | 仅一条孩子边且当前不是完整词时继续；空串立即终止。直接逐字符比较更简单 |
| [720. 词典中最长的单词](https://leetcode.com/problems/longest-word-in-dictionary/) | [l720](l720/720.go) | 只沿IsEnd为true的孩子向下搜索；同长度取字典序最小 |
| [1804. 实现 Trie II（前缀树）](https://leetcode.com/problems/implement-trie-ii-prefix-tree/) | [l1804](l1804/1804.go) | passCount/endCount处理重复插入、完整词计数、前缀计数和删除；会员题 |
| [2416. 字符串的前缀分数和](https://leetcode.com/problems/sum-of-prefix-scores-of-strings/) | [l2416](l2416/2416.go) | 每个非根节点记录经过次数，查询时累加各个非空前缀计数，重复词也贡献 |
| [139. 单词拆分](https://leetcode.com/problems/word-break/) | [l139](l139/139.go) | 从可达下标沿Trie向后走，遇完整词标记下一下标可达；Trie只加速匹配，还需要DP |
| [472. 连接词](https://leetcode.com/problems/concatenated-words/) | [l472](l472/472.go) | Trie＋切分DP/记忆化；要求至少两个词，不能用自身一次匹配就判断成功 |

### 后缀、路径 token 与混合查询

| 原题 | 本地练习 | 迁移点与边界 |
| --- | --- | --- |
| [820. 单词的压缩编码](https://leetcode.com/problems/short-encoding-of-words/) | [l820](l820/820.go) | 反向插入使后缀共享成为前缀共享；仅叶子词贡献长度+1，重复词不重复算 |
| [1233. 删除子文件夹](https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/) | [l1233](l1233/1233.go) | 以路径段而非字母建边，遇已结束父目录就不再向下；/a不是/ab的父目录 |
| [745. 前缀和后缀搜索](https://leetcode.com/problems/prefix-and-suffix-search/) | [l745](l745/745.go) | 可插入suffix+分隔符+word，在节点记录最大下标；构造每词O(L²)，分隔符需额外分支 |
| [1032. 字符流](https://leetcode.com/problems/stream-of-characters/) | [l1032](l1032/1032.go) | 反向Trie匹配最近字符的后缀，只保留最长词长的历史；不能只沿流向维护单个节点 |

### 二进制 Trie：把字符换成高位到低位的0/1

| 原题 | 本地练习 | 迁移点与边界 |
| --- | --- | --- |
| [421. 数组中两个数的最大异或值](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/) | [l421](l421/421.go) | children[2]，从高位到低位优先选相反位，高位收益支配低位 |
| [1707. 与数组中元素的最大异或值](https://leetcode.com/problems/maximum-xor-with-an-element-from-array/) | [l1707](l1707/1707.go) | 按查询上限离线排序，只插入不超过上限的数；Trie为空返回-1，恢复原查询顺序 |
| [1803. 统计异或值在范围内的数对有多少](https://leetcode.com/problems/count-pairs-with-xor-in-a-range/) | [l1803](l1803/1803.go) | 节点记数量，count(xor<high+1)-count(xor<low)，查询后插入避免自身配对；位宽覆盖high+1 |
| [2935. 找出强数对的最大异或值 II](https://leetcode.com/problems/maximum-strong-pair-xor-ii/) | [l2935](l2935/2935.go) | 排序滑窗满足max<=2*min，带计数Trie支持移出；计数为0的边不能查询 |

共20题。新增 `l题号/题号.go` 只放空函数声明，**不提供练习实现或新测试**；有返回值的空函数会产生missing return，完成后再运行该目录测试。上述六题参考代码各自独立，Constructor重名是因为来自不同题，不能拼在同一个包里编译。

<a id="review"></a>
## 10. 复杂度、易错点与复盘

### 固定骨架与少量变化点

```text
插入：从根开始 → 按字符找孩子 → 不存在则创建 → 更新节点信息
查询：从根开始 → 按字符找孩子 → 不存在则失败 → 按题意处理终点
```

| 目标 | 节点增加什么 | 查询如何变化 |
| --- | --- | --- |
| 完整词存在 | IsEnd | 走完还要检查结束标记 |
| 前缀存在 | 不需要额外信息 | 路径存在即可 |
| 重复词数量 | endCount | 返回结束次数 |
| 前缀下数量/权值 | passCount/sum | 走到前缀节点直接取统计值 |
| 返回候选词 | word/topK/单词下标 | 收集子树或读取缓存 |
| 单字符通配符 | 通常不用额外字段 | 遇点号向多个孩子搜索 |
| 后缀匹配 | 通常不用额外字段 | 反向插入，反向查询 |
| 异或最优值 | 两个孩子，可带计数 | 按位贪心，不是字典序字符搜索 |

### 常见错误

1. 把“前缀存在”当成“完整词存在”。
2. 插入重复键时无条件叠加新值，忘记覆盖差值。
3. 删除一个词就删除整条路径，破坏共享前缀；先检查通过计数和子树需求。
4. 遍历map时误以为自然有字典序。
5. 找到一个完整词就停止212的搜索，漏掉以它为前缀的长词。
6. 212中使用全局永久visited，导致其他路径无法再次使用格子。
7. 将任意字符直接映射到26分支，造成越界。
8. 把单路径查询O(L)推广到通配符、DFS或全部候选输出；这些操作需另算分支和输出规模。

### 推荐顺序

```text
主线：208 → 648 → 211 → 677 → 1268 → 212
统计：1804 → 2416
组合：139 → 472 → 1032 → 745
异或：421 → 1707 → 1803 → 2935
```

每题先回答：**边是什么？节点代表哪个前缀？保存什么信息？何时成功或剪枝？复杂度是否含分支和输出？**

<a id="legacy"></a>
## 原有资料与旧稿

保留现有文件，不覆盖历史实现；新版练习统一放在l题号目录。

- [14旧稿](14)、[1233路径分段旧稿](1233_ele)、[1268旧稿](1268)、[820旧稿](820.go)。
- [208原题](https://leetcode.com/problems/implement-trie-prefix-tree/)。
- [原有 Google OA 讨论资料](https://leetcode.com/discuss/interview-question/350363/Google-or-OA-2018-or-Max-Distance)。
