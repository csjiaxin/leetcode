package l232

import (
	"math/rand"
	"testing"
)

func TestMyQueue(t *testing.T) {
	q := Constructor()
	if !q.Empty() {
		t.Fatal("initially nonempty")
	}
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 || q.Peek() != 1 || q.Pop() != 1 {
		t.Fatal("front")
	}
	q.Push(3)
	if q.Pop() != 2 || q.Pop() != 3 || !q.Empty() {
		t.Fatal("FIFO across refill")
	}
	q.Push(4)
	if q.Pop() != 4 || !q.Empty() {
		t.Fatal("reuse")
	}
}

func TestQueueAgainstModel(t *testing.T) {
	q := Constructor()
	model := []int{}
	rng := rand.New(rand.NewSource(232))
	for step := 0; step < 5000; step++ {
		switch action := rng.Intn(3); {
		case len(model) == 0 || action == 0:
			value := rng.Intn(101) - 50
			q.Push(value)
			model = append(model, value)
		case action == 1:
			if got := q.Pop(); got != model[0] {
				t.Fatalf("step %d Pop=%d want %d", step, got, model[0])
			}
			model = model[1:]
		default:
			if got := q.Peek(); got != model[0] {
				t.Fatalf("step %d Peek=%d want %d", step, got, model[0])
			}
		}
		if q.Empty() != (len(model) == 0) || len(q.input)+len(q.output) != len(model) {
			t.Fatal("size invariant")
		}
		for i, want := range model {
			var got int
			if i < len(q.output) {
				got = q.output[len(q.output)-1-i]
			} else {
				got = q.input[i-len(q.output)]
			}
			if got != want {
				t.Fatal("internal FIFO invariant")
			}
		}
	}
}
