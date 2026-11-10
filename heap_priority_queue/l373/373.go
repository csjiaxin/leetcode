package l373

import "leetcode/heap_priority_queue/internal/pqueue"

type pair struct { sum int64; i,j int }
// kSmallestPairs merges rows nums1[i]+nums2[j] without generating the product.
// O(k log min(k,len(nums1))) time, O(min(k,len(nums1))) heap space.
func kSmallestPairs(nums1,nums2 []int,k int) [][]int {
	out:=make([][]int,0)
	if len(nums1)==0 || len(nums2)==0 || k<=0 { return out }
	q:=pqueue.New(func(a,b pair) bool {
		if a.sum!=b.sum { return a.sum<b.sum }; if a.i!=b.i { return a.i<b.i }; return a.j<b.j
	})
	for i:=0;i<len(nums1)&&i<k;i++ { q.Push(pair{int64(nums1[i])+int64(nums2[0]),i,0}) }
	for len(out)<k && q.Len()>0 {
		x:=q.Pop(); out=append(out,[]int{nums1[x.i],nums2[x.j]})
		if x.j+1<len(nums2) { q.Push(pair{int64(nums1[x.i])+int64(nums2[x.j+1]),x.i,x.j+1}) }
	}
	return out
}
