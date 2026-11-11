package l1882

import (
	"slices"
	"testing"
)

func TestAssignTasks(t *testing.T) {
	for _, tt := range []struct {
		name                 string
		servers, tasks, want []int
	}{
		{"example one", []int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}, []int{2, 2, 0, 2, 1, 2}}, {"example two", []int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}, []int{1, 4, 1, 4, 1, 3, 2}}, {"index tie", []int{1, 1}, []int{5, 4, 1}, []int{0, 1, 0}}, {"simultaneous completion weight tie", []int{5, 1}, []int{5, 4, 1}, []int{1, 0, 1}}, {"single", []int{9}, []int{1, 1, 1}, []int{0, 0, 0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignTasks(tt.servers, tt.tasks); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
func TestAssignTasksWideClock(t *testing.T) {
	tasks := make([]int, 30000)
	for i := range tasks {
		tasks[i] = 200000
	}
	got := assignTasks([]int{1, 1}, tasks)
	for i, id := range got {
		if id != i%2 {
			t.Fatalf("%s: task %d got %d want %d", t.Name(), i, id, i%2)
		}
	}
}
