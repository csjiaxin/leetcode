package l1882

import (
	"math/rand"
	"slices"
	"testing"
)

func TestAssignTasksRandom(t *testing.T) {
	r := rand.New(rand.NewSource(1882))
	for trial := 0; trial < 500; trial++ {
		n, m := r.Intn(6)+1, r.Intn(50)+1
		servers := make([]int, n)
		tasks := make([]int, m)
		for i := range servers {
			servers[i] = r.Intn(5) + 1
		}
		for i := range tasks {
			tasks[i] = r.Intn(20) + 1
		}
		end := make([]int64, n)
		want := make([]int, m)
		var now int64
		for i, duration := range tasks {
			now = max(now, int64(i))
			chosen := -1
			for chosen < 0 {
				for id := 0; id < n; id++ {
					if end[id] <= now && (chosen < 0 || servers[id] < servers[chosen] || servers[id] == servers[chosen] && id < chosen) {
						chosen = id
					}
				}
				if chosen < 0 {
					next := end[0]
					for _, e := range end {
						next = min(next, e)
					}
					now = next
				}
			}
			want[i] = chosen
			end[chosen] = now + int64(duration)
		}
		if got := assignTasks(servers, tasks); !slices.Equal(got, want) {
			t.Fatalf("%s: trial %d servers=%v tasks=%v got=%v want=%v", t.Name(), trial, servers, tasks, got, want)
		}
	}
}
