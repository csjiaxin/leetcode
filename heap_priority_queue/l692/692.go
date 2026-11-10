package l692

import "leetcode/heap_priority_queue/internal/pqueue"

type entry struct { word string; count int }
// topKFrequent returns frequency descending, then lexicographic ascending.
// O(n+u log k+k log k) heap operations, O(u+k) space; string comparisons cost length.
func topKFrequent(words []string,k int) []string {
	counts:=make(map[string]int)
	for _,w:=range words { counts[w]++ }
	q:=pqueue.New(func(a,b entry) bool {
		if a.count!=b.count { return a.count<b.count }; return a.word>b.word
	})
	for w,c:=range counts { q.Push(entry{w,c}); if q.Len()>k { q.Pop() } }
	out:=make([]string,q.Len())
	for i:=len(out)-1;i>=0;i-- { out[i]=q.Pop().word }
	return out
}
