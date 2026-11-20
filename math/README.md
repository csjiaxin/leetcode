# 数学与数值算法

[返回总目录](../README.md)

本目录整理面试常用数学工具与 Go 模板，不把题单视作已完成题解，也不优先展开复杂数论。

## 分类与识别

| 信号 | 通用工具 | 代表题 |
| --- | --- | --- |
| 指数很大，反复乘法太慢 | 二进制快速幂 | 50 Pow(x,n) |
| 公共周期、整除、约分 | 欧几里得 GCD，LCM=a/gcd*b | 1071 字符串的最大公因子 |
| 查询范围内所有质数 | 埃氏筛 | 204 计数质数 |
| 数值以字符串保存、不能整体转整数 | 竖式加法、乘法和进位 | 415 字符串相加、43 字符串相乘 |
| 规则路径的方案数 | 组合计数与整除顺序 | 62 不同路径；与 DP 对照 |
| 数据流总长度未知，均匀抽样 | 蓄水池采样 | 382 链表随机节点、398 随机数索引 |

## 模板一：非负整数 GCD

```go
func GCD(a, b uint64) uint64 {
    for b != 0 { a, b = b, a%b }
    return a
}
```

利用 `gcd(a,b)=gcd(b,a%b)`，定义 gcd(0,0)=0。时间 O(log max(a,b))（按固定字长取模计算），空间 O(1)。LCM 应先除后乘，但仍可能溢出；不能直接对有符号最小值取绝对值再转成 uint64。

1071 先检查拼接兼容性 `s+t==t+s`，再取长度 GCD 的前缀；仅求长度 GCD 不足以证明字符串重复结构相同。

## 模板二：支持负指数的快速幂

```go
// n 可为 int64 最小值；x=0 且 n<0 不在数学定义域内。
func Pow(x float64, n int64) float64 {
    var exponent uint64
    if n < 0 {
        x = 1/x
        exponent = uint64(-(n+1))+1
    } else { exponent = uint64(n) }
    result := 1.0
    for exponent > 0 {
        if exponent&1 != 0 { result *= x }
        x *= x
        exponent >>= 1
    }
    return result
}
```

result 累积已处理的二进制位，x 每轮平方对应下一位。时间 O(log |n|)，空间 O(1)。`-(n+1)` 再转无符号加 1，避免直接 `-MinInt64` 溢出。浮点运算可能舍入、溢出为无穷或下溢；本模板遵循 float64 运算，0 的负指数不提供数学有效性保证。

## 模板三：字符串加法

```go
// a、b 为非空十进制数字串，无前导零（单独 "0" 除外）。
func AddStrings(a, b string) string {
    i, j, carry := len(a)-1, len(b)-1, 0
    reversed := make([]byte, 0, len(a)+len(b)+1)
    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        if i >= 0 { sum += int(a[i]-'0'); i-- }
        if j >= 0 { sum += int(b[j]-'0'); j-- }
        reversed = append(reversed, byte(sum%10)+'0')
        carry = sum/10
    }
    for l, r := 0, len(reversed)-1; l < r; l, r = l+1, r-1 {
        reversed[l], reversed[r] = reversed[r], reversed[l]
    }
    return string(reversed)
}
```

时间 O(m+n)，空间 O(m+n)。43 竖式乘法令第 i、j 位乘积累加到对应低位位置，再规范进位，基础版本时间 O(mn)；不能先把整个字符串解析成机器整数。

## 其他工具的正确边界

- **埃氏筛**：布尔数组标记合数；只从质数 p 的 p² 起标记，因为更小倍数已有更小质因子。时间 O(n log log n)，空间 O(n)。循环条件可用 `p<=n/p` 避免 p² 判断溢出，204 要统计严格小于 n 的质数。
- **组合数**：整数乘法中间值可能溢出，即使最终答案很小；模数不一定允许直接除法，需明确素数模与逆元条件。可先用 DP 或约分。
- **蓄水池 k=1**：第 i 个样本以 1/i 的概率替换已有样本，归纳可证扫描 i 项后各项等概率。要求随机选择均匀；需要在 API 层处理空数据流。

检查：负数、零、最小整数、模减法规范、先转换再乘法、浮点精度。数学公式减少状态，但不自动消除实现边界。

练习：50 → 1071 → 204 → 415 → 43 → 62 组合解 → 382 / 398。
