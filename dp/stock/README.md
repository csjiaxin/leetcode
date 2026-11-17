# 股票 DP：昨天的状态与合法买入来源

[返回 DP 总目录](../README.md)

先定义“某天结束后是否持股”，再增加交易预算、手续费或冷冻约束。所有更新右侧只读取昨天或明确要求的前天状态。

| 子模板 | 不变量与参数 | 题目 |
| --- | --- | --- |
| BoundedTrade | cash[t]/hold[t] 为最多 t 次买入时的收益；买入来自昨天 cash[t-1] | 121 的 1 次、123 的 2 次、188 的 k 次 |
| UnlimitedTrade | cash/hold 是昨天的两种状态；手续费在卖出时扣一次 | 122、714 |
| UnlimitedTrade(cooldown=true) | 买入必须读取前天 cash，避开昨天卖出后的冷冻日 | 309 |

例如 714 一次交易只扣一次 fee，下面统一放到卖出动作。309 不能用同日更新的 cash，也不能直接读取昨天 cash。持股初值为足够小的负值；最终返回未持股的收益。

预算版按天分配新数组，避免 t 的遍历方向带来混淆，时间 O(nk)、辅助空间 O(k)。188 在 k>=n/2 时改用无限次模板，时间 O(n)。无限次模板 O(n) 时间、O(1) 空间。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### BoundedTrade

<!-- template: BoundedTrade -->
```go
// BoundedTrade 的 cash[t]/hold[t] 表示至多 t 次买入的现金/持股收益。
// 每天从昨天的整组数组转移，预算维度不会混入当天新状态。
func BoundedTrade(prices []int, budget int) int {
	if budget <= 0 || len(prices) < 2 {
		return 0
	}
	budget = min(budget, len(prices)/2)
	cash, hold := make([]int, budget+1), make([]int, budget+1)
	for t := range hold {
		hold[t] = -int(^uint(0)>>1) / 4
	}
	for _, price := range prices {
		nextCash := append([]int(nil), cash...)
		nextHold := append([]int(nil), hold...)
		for t := 1; t <= budget; t++ {
			nextCash[t] = max(cash[t], hold[t]+price)
			nextHold[t] = max(hold[t], cash[t-1]-price)
		}
		cash, hold = nextCash, nextHold
	}
	return cash[budget]
}
```

### UnlimitedTrade

<!-- template: UnlimitedTrade -->
```go
// UnlimitedTrade 支持卖出时扣手续费，及卖出后冷冻一天的独立限制。
func UnlimitedTrade(prices []int, fee int, cooldown bool) int {
	cash, hold := 0, -int(^uint(0)>>1)/4
	beforeYesterday := 0
	for _, price := range prices {
		oldCash, oldHold := cash, hold
		buyFrom := oldCash
		if cooldown {
			buyFrom = beforeYesterday
		}
		cash = max(oldCash, oldHold+price-fee)
		hold = max(oldHold, buyFrom-price)
		beforeYesterday = oldCash
	}
	return cash
}
```

## 逐题适配器

### 121

[代码](l121/121.go) · [测试](l121/121_test.go)

最多一次买入与卖出。

需要模板：`BoundedTrade`。

<!-- solution: dp/stock/l121/121.go; templates: BoundedTrade -->
```go
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// 最多一次买入与卖出。
func maxProfit(prices []int) int { return BoundedTrade(prices, 1) }
```

### 123

[代码](l123/123.go) · [测试](l123/123_test.go)

交易预算为两次。

需要模板：`BoundedTrade`。

<!-- solution: dp/stock/l123/123.go; templates: BoundedTrade -->
```go
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// 交易预算为两次。
func maxProfit(prices []int) int { return BoundedTrade(prices, 2) }
```

### 188

[代码](l188/188.go) · [测试](l188/188_test.go)

预算足够大时退化为无限次交易，避免 O(nk) 的大预算开销。

需要模板：`BoundedTrade`、`UnlimitedTrade`。

<!-- solution: dp/stock/l188/188.go; templates: BoundedTrade,UnlimitedTrade -->
```go
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// 预算足够大时退化为无限次交易，避免 O(nk) 的大预算开销。
func maxProfit(k int, prices []int) int {
	if k >= len(prices)/2 {
		return UnlimitedTrade(prices, 0, false)
	}
	return BoundedTrade(prices, k)
}
```

### 122

[代码](l122/122.go) · [测试](l122/122_test.go)

无限次交易，无手续费和冷冻。

需要模板：`UnlimitedTrade`。

<!-- solution: dp/stock/l122/122.go; templates: UnlimitedTrade -->
```go
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// 无限次交易，无手续费和冷冻。
func maxProfit(prices []int) int { return UnlimitedTrade(prices, 0, false) }
```

### 714

[代码](l714/714.go) · [测试](l714/714_test.go)

手续费每笔只扣一次，统一放在卖出时。

需要模板：`UnlimitedTrade`。

<!-- solution: dp/stock/l714/714.go; templates: UnlimitedTrade -->
```go
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// 手续费每笔只扣一次，统一放在卖出时。
func maxProfit(prices []int, fee int) int { return UnlimitedTrade(prices, fee, false) }
```

### 309

[代码](l309/309.go) · [测试](l309/309_test.go)

买入读取前天 cash，不能读取昨天刚卖出的收益。

需要模板：`UnlimitedTrade`。

<!-- solution: dp/stock/l309/309.go; templates: UnlimitedTrade -->
```go
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 买入读取前天 cash，不能读取昨天刚卖出的收益。
func maxProfit(prices []int) int { return UnlimitedTrade(prices, 0, true) }
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

```sh
go test ./dp/stock/...
```
