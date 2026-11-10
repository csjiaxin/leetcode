package l857

import (
 "math"
 "sort"
 "leetcode/heap_priority_queue/internal/pqueue"
)

type worker struct { quality,wage int }
// mincostToHireWorkers scans increasing minimum wage/quality ratios.
// At each ratio a max heap retains k smallest qualities. O(n log n) time, O(n) space.
func mincostToHireWorkers(quality,wage []int,k int) float64 {
	workers:=make([]worker,len(quality))
	for i:=range quality { workers[i]=worker{quality[i],wage[i]} }
	sort.Slice(workers,func(i,j int) bool {
		return int64(workers[i].wage)*int64(workers[j].quality)<int64(workers[j].wage)*int64(workers[i].quality)
	})
	q:=pqueue.New(func(a,b int) bool { return a>b })
	var sum int64
	best:=math.Inf(1)
	for _,w:=range workers {
		q.Push(w.quality); sum+=int64(w.quality)
		if q.Len()>k { sum-=int64(q.Pop()) }
		if q.Len()==k { best=math.Min(best,float64(sum)*float64(w.wage)/float64(w.quality)) }
	}
	return best
}
