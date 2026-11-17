# 字符串匹配：按匹配状态组织

本目录的已整理子系列为 [KMP](kmp/readme.md)，公共状态是“当前已经匹配的模式串前缀长度”。失配后沿前缀函数回退，文本指针保持单向推进；统计、周期、旋转匹配分别改变扫描对象或匹配成功后的处理。

```text
预处理模式串的前缀函数 pi
逐个读入文本字符 c：
    while 匹配长度 > 0 且下一个模式字符 != c：沿 pi 回退
    若字符相等：匹配长度加 1
    若匹配完整模式：记录结果，再按需要回退以允许重叠
```

[Trie](../trie/readme.md) 按词典前缀分支搜索；[双序列 DP](../dp/strings/readme.md) 枚举两个前缀的匹配决策。它们的状态和失配处理不同，因此保持独立系列，并通过链接选择算法。

## 原有资料

https://www.youtube.com/watch?v=5i7oKodCRJo
https://www.youtube.com/watch?v=KG44VoDtsAA


https://leetcode.com/problems/implement-strstr/

https://leetcode.com/tag/string-matching/
