package l2530

import "leetcode/heap_priority_queue/internal/pqueue"

// maxKelements takes the current maximum and replaces it with ceil(x/3).
// O(n+k log n) time, O(n) space.
func maxKelements(nums []int,k int) int64 {
	q:=pqueue.New(func(a,b int) bool { return a>b },nums...)
	var score int64
	for ;k>0;k-- {
		x:=q.Pop(); score+=int64(x)
		y:=x/3; if x%3!=0 { y++ }; q.Push(y)
	}
	return score
}
