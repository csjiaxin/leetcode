package l122

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

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// 无限次交易，无手续费和冷冻。
func maxProfit(prices []int) int { return UnlimitedTrade(prices, 0, false) }
