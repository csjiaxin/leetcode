package l767

import "leetcode/heap_priority_queue/internal/pqueue"

type letter struct {
	char  byte
	count int
}

// reorganizeString withholds the previous character until a different one is used.
// O(n log A) time, O(n+A) space including output for A distinct characters.
func reorganizeString(s string) string {
	counts := make(map[byte]int)
	for i := range s {
		counts[s[i]]++
	}
	q := pqueue.New(func(a, b letter) bool {
		if a.count != b.count {
			return a.count > b.count
		}
		return a.char < b.char
	})
	for c, n := range counts {
		q.Push(letter{c, n})
	}
	out := make([]byte, 0, len(s))
	previous := letter{}
	for q.Len() > 0 {
		current := q.Pop()
		out = append(out, current.char)
		current.count--
		if previous.count > 0 {
			q.Push(previous)
		}
		previous = current
	}
	if len(out) != len(s) {
		return ""
	}
	return string(out)
}
