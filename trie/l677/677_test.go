package l677

import "testing"

func TestMapSum(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			key    string
			value  int
			insert bool
			want   int
		}
	}{
		{"example and decreasing overwrite", []struct {
			key    string
			value  int
			insert bool
			want   int
		}{
			{"ap", 0, false, 0}, {"apple", 3, true, 0}, {"ap", 0, false, 3}, {"app", 2, true, 0}, {"ap", 0, false, 5},
			{"apple", 2, true, 0}, {"ap", 0, false, 4}, {"apple", 0, false, 2}, {"app", 0, false, 4},
		}},
		{"increasing and same value overwrite", []struct {
			key    string
			value  int
			insert bool
			want   int
		}{
			{"a", 2, true, 0}, {"a", 7, true, 0}, {"a", 0, false, 7}, {"a", 7, true, 0}, {"a", 0, false, 7},
		}},
		{"shared prefix and independent branches", []struct {
			key    string
			value  int
			insert bool
			want   int
		}{
			{"apple", 3, true, 0}, {"apt", 4, true, 0}, {"bat", 5, true, 0},
			{"a", 0, false, 7}, {"ap", 0, false, 7}, {"app", 0, false, 3}, {"b", 0, false, 5}, {"z", 0, false, 0}, {"apples", 0, false, 0},
		}},
		{"insert short key after long key", []struct {
			key    string
			value  int
			insert bool
			want   int
		}{
			{"ab", 10, true, 0}, {"a", 1, true, 0}, {"a", 0, false, 11}, {"ab", 0, false, 10}, {"a", 5, true, 0}, {"a", 0, false, 15},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Constructor()
			for i, op := range tt.operations {
				if op.insert {
					m.Insert(op.key, op.value)
					continue
				}
				if got := m.Sum(op.key); got != op.want {
					t.Errorf("operation %d: Sum(%q) = %d, want %d", i, op.key, got, op.want)
				}
			}
		})
	}
}

func TestMapSumIndependentInstances(t *testing.T) {
	first, second := Constructor(), Constructor()
	first.Insert("apple", 3)
	second.Insert("apple", 8)
	if first.Sum("ap") != 3 || second.Sum("ap") != 8 {
		t.Fatal("instances must not share values or trie nodes")
	}
}
