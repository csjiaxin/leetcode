package l901

// https://leetcode.com/problems/online-stock-span/
// Embedded subtemplate: see README.md for its invariant and tie policy.
type priceSpan struct{ price, span int }
type StockSpanner struct{ pending []priceSpan }

func collapseSpan(pending []priceSpan, price int) ([]priceSpan, int) {
	span := 1
	for len(pending) > 0 && pending[len(pending)-1].price <= price {
		span += pending[len(pending)-1].span
		pending = pending[:len(pending)-1]
	}
	return append(pending, priceSpan{price, span}), span
}

func Constructor() StockSpanner { return StockSpanner{} }
func (s *StockSpanner) Next(price int) int {
	var span int
	s.pending, span = collapseSpan(s.pending, price)
	return span
}
