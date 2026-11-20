# 哈希表：查找、分组、计数与状态映射

[返回总目录](../README.md)

本目录补充哈希专题的学习框架；下方是可复制的 Go 模板，不表示题单均已有独立题解。map 操作按平均 O(1) 分析，不能视为有序结构。

## 识别信号与分类

| 信号 | 建模 / 不变量 | 代表题 |
| --- | --- | --- |
| 找配对，避免重复遍历 | map 只保存已经扫描的元素 | 1 两数之和 |
| 对象等价但表示不同 | 为同一等价类设计相同的可比较 key | 49 字母异位词分组 |
| 判断连续序列 | 集合只从没有前驱的值开始扩展 | 128 最长连续序列 |
| 一一对应关系 | 同时维护两个方向的映射 | 205 同构字符串、290 单词规律 |
| 子数组和 / 状态相同 | 前缀状态的频次或最早位置 | 560、525，见[前缀状态](../prefix_sum/state_hash/readme.md) |
| 小值域映射到下标 | 原地放置或标记，不一定需要 map | 41 缺失的第一个正数 |

## 模板一：查后写，保证使用不同下标

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, x := range nums {
        if j, ok := seen[target-x]; ok { return []int{j, i} }
        seen[x] = i
    }
    return nil
}
```

处理 i 前，seen 仅代表 `[0,i)`。先写再查会错误地复用自己；`[3,3]`、target=6 则能正常匹配两个不同下标。时间平均 O(n)，空间 O(n)，算术按题目范围保证不溢出。

## 模板二：固定字母表用频次数组作为 key

```go
// 输入仅含小写 ASCII 字母。Go 数组可比较，切片不能作为 map key。
func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    for _, s := range strs {
        var key [26]int
        for i := 0; i < len(s); i++ { key[s[i]-'a']++ }
        groups[key] = append(groups[key], s)
    }
    result := make([][]string, 0, len(groups))
    for _, group := range groups { result = append(result, group) }
    return result
}
```

设字符串数 n、总字符数 L，平均时间 O(L+26n)，不计字符串原有存储的额外空间 O(26n+n)。结果组序不保证稳定；需要固定顺序时另存组首次出现顺序。Unicode 需要按 rune 重新设计 key，不能直接使用 26 个字母的假设。

## 通用解题步骤与易错点

1. 明确 key 对应单个值、频率、最早下标，还是完整状态。
2. 明确 map 代表已处理前缀还是整个输入，决定查写顺序。
3. 用 `value, ok := m[key]` 区分不存在和合法零值。
4. 计数用 `m[key]++`，不能用覆盖下标替代频次；最长距离一般保存最早位置。
5. 128 要先去重，只从 `x-1` 不存在的位置扩展，才能避免重复扫描；极值运算注意溢出。
6. map 不维护大小顺序，前驱 / 排名查询应考虑排序、二分、树状数组等。

练习：1 → 49 → 205 / 290 → 128 → 41；之后与 560、525 对照。
