package l188

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

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// 预算足够大时退化为无限次交易，避免 O(nk) 的大预算开销。
func maxProfit(k int, prices []int) int {
	if k >= len(prices)/2 {
		return UnlimitedTrade(prices, 0, false)
	}
	return BoundedTrade(prices, k)
}
