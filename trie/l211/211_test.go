package l211

import "testing"

func TestWordDictionarySearch(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		queries []struct {
			pattern string
			want    bool
		}
	}{
		{"empty dictionary", nil, []struct {
			pattern string
			want    bool
		}{{"a", false}, {".", false}}},
		{"example", []string{"bad", "dad", "mad"}, []struct {
			pattern string
			want    bool
		}{
			{"pad", false}, {"bad", true}, {".ad", true}, {"b..", true}, {"ba", false}, {"badz", false},
		}},
		{"prefix is not a word", []string{"apple"}, []struct {
			pattern string
			want    bool
		}{
			{"app", false}, {"a..", false}, {"app.e", true}, {"apple", true}, {"appl..", false},
		}},
		{"wildcard must try later branches", []string{"baa", "cat", "zap"}, []struct {
			pattern string
			want    bool
		}{
			{".at", true}, {"..p", true}, {".az", false}, {"c.t", true}, {"za.", true},
		}},
		{"single and double characters", []string{"a", "z", "ab", "zz"}, []struct {
			pattern string
			want    bool
		}{
			{".", true}, {"..", true}, {".b", true}, {"z.", true}, {"b.", false}, {"a", true},
		}},
		{"duplicates and shared endings", []string{"bad", "bad", "ba", "bake"}, []struct {
			pattern string
			want    bool
		}{
			{"ba", true}, {"b.", true}, {"ba..", true}, {"baker", false},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Constructor()
			for _, word := range tt.words {
				d.AddWord(word)
			}
			for _, q := range tt.queries {
				if got := d.Search(q.pattern); got != q.want {
					t.Errorf("Search(%q) = %v, want %v", q.pattern, got, q.want)
				}
			}
		})
	}
}

func TestWordDictionaryInterleavedOperations(t *testing.T) {
	d := Constructor()
	d.AddWord("apple")
	if d.Search("app") {
		t.Fatal("prefix must not be a complete word")
	}
	d.AddWord("app")
	if !d.Search("app") || !d.Search("app.e") {
		t.Fatal("insertion must preserve both words")
	}
	for i := 0; i < 3; i++ {
		if !d.Search("a.p") {
			t.Fatal("search must not consume words")
		}
	}
	other := Constructor()
	if other.Search("app") {
		t.Fatal("instances must be independent")
	}
}
