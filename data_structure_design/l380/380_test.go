package l380

import (
	"math/rand"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	s := Constructor()
	if !s.Insert(1) || s.Remove(2) || !s.Insert(2) {
		t.Fatal("insert/remove")
	}
	if got := s.GetRandom(); got != 1 && got != 2 {
		t.Fatal(got)
	}
	if !s.Remove(1) || s.Insert(2) || s.GetRandom() != 2 {
		t.Fatal("remaining singleton")
	}
	if !s.Remove(2) || s.Remove(2) || !s.Insert(2) || s.GetRandom() != 2 {
		t.Fatal("delete last/reinsert")
	}
}

func TestRandomizedSetAgainstModel(t *testing.T) {
	s := Constructor()
	model := map[int]bool{}
	rng := rand.New(rand.NewSource(380))
	for step := 0; step < 5000; step++ {
		value := rng.Intn(31) - 15
		if rng.Intn(2) == 0 {
			if got := s.Insert(value); got == model[value] {
				t.Fatalf("Insert(%d)=%v", value, got)
			}
			model[value] = true
		} else {
			if got := s.Remove(value); got != model[value] {
				t.Fatalf("Remove(%d)=%v", value, got)
			}
			delete(model, value)
		}
		if len(s.values) != len(model) || len(s.index) != len(model) {
			t.Fatal("size invariant")
		}
		for i, value := range s.values {
			index, exists := s.index[value]
			if !exists || index != i || !model[value] {
				t.Fatal("slice/map bijection")
			}
		}
		if len(model) > 0 {
			for sample := 0; sample < 3; sample++ {
				if value := s.GetRandom(); !model[value] {
					t.Fatalf("sampled removed value %d", value)
				}
			}
		}
	}
}
