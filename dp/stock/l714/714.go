package l714

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

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// 手续费每笔只扣一次，统一放在卖出时。
func maxProfit(prices []int, fee int) int { return UnlimitedTrade(prices, fee, false) }
