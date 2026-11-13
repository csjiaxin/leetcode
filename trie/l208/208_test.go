package l208

import "testing"

func TestTrie(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		queries []struct {
			word           string
			search, prefix bool
		}
	}{
		{"empty dictionary", nil, []struct {
			word           string
			search, prefix bool
		}{{"a", false, false}}},
		{"word versus prefix", []string{"apple"}, []struct {
			word           string
			search, prefix bool
		}{
			{"apple", true, true}, {"app", false, true}, {"ap", false, true}, {"apples", false, false}, {"b", false, false},
		}},
		{"shared prefixes", []string{"app", "apple", "apt"}, []struct {
			word           string
			search, prefix bool
		}{
			{"app", true, true}, {"apple", true, true}, {"apt", true, true}, {"ap", false, true}, {"apz", false, false},
		}},
		{"duplicate insertion", []string{"cat", "cat"}, []struct {
			word           string
			search, prefix bool
		}{{"cat", true, true}, {"ca", false, true}}},
		{"alphabet boundaries", []string{"a", "z", "za"}, []struct {
			word           string
			search, prefix bool
		}{{"a", true, true}, {"z", true, true}, {"za", true, true}, {"zz", false, false}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, word := range tt.words {
				trie.Insert(word)
			}
			for _, q := range tt.queries {
				if got := trie.Search(q.word); got != q.search {
					t.Errorf("Search(%q) = %v, want %v", q.word, got, q.search)
				}
				if got := trie.StartsWith(q.word); got != q.prefix {
					t.Errorf("StartsWith(%q) = %v, want %v", q.word, got, q.prefix)
				}
			}
		})
	}
}

func TestTrieInterleavedOperations(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if trie.Search("app") || !trie.StartsWith("app") {
		t.Fatal("app should initially be only a prefix")
	}
	trie.Insert("app")
	if !trie.Search("app") || !trie.Search("apple") {
		t.Fatal("inserting a prefix must preserve its longer word")
	}
	trie.Insert("banana")
	if !trie.Search("banana") || !trie.Search("apple") {
		t.Fatal("inserting another branch must preserve existing words")
	}
	other := Constructor()
	if other.Search("apple") || other.StartsWith("app") {
		t.Fatal("instances must not share state")
	}
}
