package l378

import "leetcode/heap_priority_queue/internal/pqueue"

type cell struct { value,row,col int }
// kthSmallest merges sorted rows. O(n+k log n) time, O(n) space.
func kthSmallest(matrix [][]int,k int) int {
	q:=pqueue.New(func(a,b cell) bool { return a.value<b.value })
	for r,row:=range matrix { q.Push(cell{row[0],r,0}) }
	for {
		x:=q.Pop(); k--
		if k==0 { return x.value }
		if x.col+1<len(matrix[x.row]) { q.Push(cell{matrix[x.row][x.col+1],x.row,x.col+1}) }
	}
}
