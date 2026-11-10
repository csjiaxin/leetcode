package l264

import "leetcode/heap_priority_queue/internal/pqueue"

// nthUglyNumber expands unique multiples of 2,3,5 in sorted order.
// O(n log n) time, O(n) space. int64 intermediates avoid 32-bit overflow.
func nthUglyNumber(n int) int {
	q:=pqueue.New(func(a,b int64) bool { return a<b },int64(1))
	seen:=map[int64]bool{1:true}
	var x int64
	for count:=0;count<n;count++ {
		x=q.Pop()
		for _,factor:=range []int64{2,3,5} {
			y:=x*factor
			if !seen[y] { seen[y]=true; q.Push(y) }
		}
	}
	return int(x)
}
