# 差分＋前缀扫描：批量区间变化

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [Go 通用骨架](#code)
- [推演与题目变化点](#mapping)
- [稀疏坐标、复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

批量给某些区间加减相同值，最后查询每个位置的最终值、最大重叠量或是否超过容量。把“整段变化”拆为两个端点事件，避免每次逐元素更新。

定义 diff[0]=a[0]，diff[i]=a[i]-a[i-1]。前缀扫描到 i 后，累计值恰好是 a[i]。

| 原始区间 | 端点更新 | 解释 |
| --- | --- | --- |
| 整数闭区间 [l,r] 加delta | diff[l]+=delta；diff[r+1]-=delta | r仍有效，r+1才取消 |
| 半开区间 [l,r) 加delta | diff[l]+=delta；diff[r]-=delta | r已经不再生效 |

**不变量**：扫描坐标 x 时，累计值已经包含 x 的全部净变化，表示在 x 生效的总量。同坐标加减必须先合并再检查，不应凭事件输入顺序判断超载。

适合离线批处理；若区间更新与实时查询交错，需要 Fenwick 树或线段树等结构。不是所有区间问题都只需差分。

<a id="code"></a>

## 2. Go 通用骨架

```go
package template

// 在长度 n 的零数组上执行合法的整数闭区间更新。
// Updates: [left,right,delta]；要求 0<=left<=right<n。
func RangeAdds(n int,updates [][3]int) []int64 {
    diff:=make([]int64,n+1) // 位置 n 是取消末尾变化的哨兵
    for _,u:=range updates {
        left,right,delta:=u[0],u[1],int64(u[2])
        diff[left]+=delta
        diff[right+1]-=delta
    }
    answer:=make([]int64,n)
    var current int64
    for i:=range answer {
        current+=diff[i]
        answer[i]=current
    }
    return answer
}

// LC1094: trips=[乘客数,上车站,下车站]；0<=from<to。
// 乘客在 [from,to) 占位；到达 to 后不再占位。
func CarPooling(trips [][]int,capacity int) bool {
    maxPosition:=0
    for _,t:=range trips { maxPosition=max(maxPosition,t[2]) }
    diff:=make([]int64,maxPosition+1)
    for _,t:=range trips {
        passengers:=int64(t[0])
        diff[t[1]]+=passengers
        diff[t[2]]-=passengers // 不是 to+1
    }
    var current int64
    for _,change:=range diff {
        current+=change
        if current>int64(capacity) { return false }
    }
    return true
}
```

初始数组不全为0时先构造它的差分，再做两端更新，最后前缀还原。n=0 且无更新时返回空数组；非法更新应由调用者验证，不能静默截断端点。

<a id="mapping"></a>

## 3. 推演与题目变化点

[1094. 拼车](https://leetcode.com/problems/car-pooling/)：trips=[[2,1,5],[3,5,7]]、capacity=3。

```text
diff[1]+=2
diff[5]-=2，同时 diff[5]+=3 → 同坐标净变化 +1
diff[7]-=3

位置1：车上2人
位置5：先合并净变化，车上3人，不是5人
位置7：车上0人
返回true
```

下车站点不占位，所以同一位置下车腾出的空间可以供新乘客使用。若把 to 当成闭区间并写 to+1，会误判重叠。

| 题目 | 区间含义 | 端点记录 | 汇总 |
| --- | --- | --- | --- |
| [1094. 拼车](https://leetcode.com/problems/car-pooling/) | [from,to) | from加人数，to减人数 | 任意位置载荷是否<=capacity |
| [1109. 航班预订统计](https://leetcode.com/problems/corporate-flight-bookings/) | 航班编号[first,last]，1-based | 转0-based后 left=first-1、right=last-1，right+1取消 | 返回每个航班座位数 |
| [1854. 人口最多的年份](https://leetcode.com/problems/maximum-population-year/) | [birth,death) | birth加1，death减1 | 找最大前缀；并列取最早，仅在严格变大时更新答案 |

固定骨架：端点记净变化 → 按坐标扫描前缀 → 汇总结果。变化点：端点是否包含、加多少、最终求完整数组还是最大值 / 可行性。

<a id="pitfalls"></a>

## 4. 稀疏坐标、复杂度与陷阱

- m 次更新、n 个位置：时间 O(m+n)，空间 O(n)。拼车用 U 表示最大坐标，时间 O(m+U)，空间 O(U)。
- 若坐标达10^9但只有少量事件，不要按最大坐标分配数组。用 map 聚合同坐标变化，再排序不同坐标扫描，时间 O(m log(m+1))、空间 O(m)。
- 坐标压缩保留顺序，适合最大覆盖量和容量检查；若计算覆盖总长度或面积，必须乘相邻原坐标间距，压缩下标差不能当实际距离。
- 半开区间的取消点为 to；整数闭区间的取消点为 r+1。实数闭区间不能机械用 +1，需专门定义同坐标事件顺序。
- 差分累计值和 delta 用 int64；先转型再运算，且检查总变化是否在类型范围内。
- 同坐标临时先加后减可能伪造超载。差分数组天然聚合；事件列表需先处理整组坐标或明确正确事件顺序。
- 只有预处理之后才输出结果是离线扫描，不要把它称为支持在线查询。

已有相关数据结构：[tree/segment](../../tree/segment/)。累计查询思想可对比 [前缀状态](../prefix_state_hash/README.md)。
