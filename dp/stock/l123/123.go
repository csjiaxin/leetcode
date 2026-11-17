package l123

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

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// 交易预算为两次。
func maxProfit(prices []int) int { return BoundedTrade(prices, 2) }
