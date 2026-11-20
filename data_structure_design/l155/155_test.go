package l155

import (
	"math/rand"
	"testing"
)

func TestMinStack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 {
		t.Fatal("minimum")
	}
	s.Pop()
	if s.Top() != 0 || s.GetMin() != -2 {
		t.Fatal("restore")
	}
}

func TestMinStackAgainstModel(t *testing.T) {
	s := Constructor()
	model := []int{}
	rng := rand.New(rand.NewSource(155))
	maxInt := int(^uint(0) >> 1)
	for _, value := range []int{maxInt, -maxInt - 1, -maxInt - 1, 0} {
		s.Push(value)
		model = append(model, value)
	}
	for step := 0; step < 5000; step++ {
		if len(model) == 0 || rng.Intn(2) == 0 {
			value := rng.Intn(21) - 10
			s.Push(value)
			model = append(model, value)
		} else {
			s.Pop()
			model = model[:len(model)-1]
		}
		if len(s.entries) != len(model) {
			t.Fatal("size")
		}
		if len(model) == 0 {
			continue
		}
		minimum := model[0]
		for _, value := range model {
			if value < minimum {
				minimum = value
			}
		}
		if s.Top() != model[len(model)-1] || s.GetMin() != minimum {
			t.Fatalf("step %d: state differs from %v", step, model)
		}
	}
	for len(model) > 0 {
		s.Pop()
		model = model[:len(model)-1]
	}
	s.Push(7)
	if s.GetMin() != 7 || s.Top() != 7 {
		t.Fatal("reuse empty stack")
	}
}
