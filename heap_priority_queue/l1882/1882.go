package l1882

import "leetcode/heap_priority_queue/internal/pqueue"

type server struct {
	end        int64
	weight, id int
}

// assignTasks maintains idle weight/ID priority and busy end/weight/ID priority.
// O((n+m) log n) time, O(n+m) space including output; cumulative times use int64.
func assignTasks(servers, tasks []int) []int {
	byWeight := func(a, b server) bool {
		if a.weight != b.weight {
			return a.weight < b.weight
		}
		return a.id < b.id
	}
	idle := pqueue.New(byWeight)
	busy := pqueue.New(func(a, b server) bool {
		if a.end != b.end {
			return a.end < b.end
		}
		return byWeight(a, b)
	})
	for i, w := range servers {
		idle.Push(server{weight: w, id: i})
	}
	out := make([]int, len(tasks))
	var now int64
	for i, duration := range tasks {
		now = max(now, int64(i))
		if idle.Len() == 0 {
			now = max(now, busy.Peek().end)
		}
		// Release every simultaneous completion before choosing by weight.
		for busy.Len() > 0 && busy.Peek().end <= now {
			idle.Push(busy.Pop())
		}
		s := idle.Pop()
		out[i] = s.id
		s.end = now + int64(duration)
		busy.Push(s)
	}
	return out
}
