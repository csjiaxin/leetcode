package l621

import "leetcode/heap_priority_queue/internal/pqueue"

type cooling struct { ready,count int }
// leastInterval greedily runs the largest available frequency and queues cooldowns.
// O(T log A) time for T tasks and A distinct task kinds; O(A) space.
// Idle periods are jumped over rather than simulated one unit at a time.
func leastInterval(tasks []byte,n int) int {
	counts:=make(map[byte]int)
	for _,c:=range tasks { counts[c]++ }
	q:=pqueue.New(func(a,b int) bool { return a>b })
	for _,c:=range counts { q.Push(c) }
	wait:=make([]cooling,0,len(counts)); head,now:=0,0
	for q.Len()>0||head<len(wait) {
		if q.Len()==0 { now=max(now,wait[head].ready) }
		for head<len(wait)&&wait[head].ready<=now { q.Push(wait[head].count); head++ }
		remaining:=q.Pop()-1
		now++
		if remaining>0 { wait=append(wait,cooling{now+n,remaining}) }
		// Compact consumed cooldowns to keep storage bounded by active kinds.
		if head>0&&head*2>=len(wait) { wait=append(wait[:0],wait[head:]...); head=0 }
	}
	return now
}
