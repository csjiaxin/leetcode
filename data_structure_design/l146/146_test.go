package l146

import (
	"math/rand"
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if got := c.Get(1); got != 1 {
		t.Fatal(got)
	}
	c.Put(3, 3)
	if got := c.Get(2); got != -1 {
		t.Fatal(got)
	}
	c.Put(4, 4)
	for key, want := range map[int]int{1: -1, 3: 3, 4: 4} {
		if got := c.Get(key); got != want {
			t.Fatalf("Get(%d)=%d want %d", key, got, want)
		}
	}
}

func TestLRUAgainstModel(t *testing.T) {
	for capacity := 0; capacity <= 5; capacity++ {
		c := Constructor(capacity)
		values := map[int]int{}
		order := []int{} // Oldest first; deliberately linear independent model.
		touch := func(key int) {
			for i, k := range order {
				if k == key {
					order = append(order[:i], order[i+1:]...)
					break
				}
			}
			order = append(order, key)
		}
		rng := rand.New(rand.NewSource(int64(capacity)))
		for step := 0; step < 2000; step++ {
			key := rng.Intn(9) - 4
			if rng.Intn(2) == 0 {
				value := rng.Intn(21) - 10
				c.Put(key, value)
				values[key] = value
				touch(key)
				if len(order) > capacity {
					delete(values, order[0])
					order = order[1:]
				}
			} else {
				want, exists := values[key]
				if !exists {
					want = -1
				} else {
					touch(key)
				}
				if got := c.Get(key); got != want {
					t.Fatalf("capacity=%d step=%d Get(%d)=%d want %d", capacity, step, key, got, want)
				}
			}
			if len(c.index) != len(values) || c.order.Len() != len(order) {
				t.Fatal("size invariant")
			}
			i := len(order) - 1
			for element := c.order.Front(); element != nil; element = element.Next() {
				e := element.Value.(entry)
				if i < 0 || e.key != order[i] || e.value != values[e.key] || c.index[e.key] != element {
					t.Fatal("order/map invariant")
				}
				i--
			}
		}
	}
}

func TestNegativeCapacity(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	Constructor(-1)
}
