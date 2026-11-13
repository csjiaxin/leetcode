package l648

import "testing"

func TestReplaceWords(t *testing.T) {
	tests := []struct {
		name           string
		dictionary     []string
		sentence, want string
	}{
		{"example", []string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{"single letter roots", []string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs", "a a b c"},
		{"shortest root despite insertion order", []string{"cattle", "cat", "c"}, "cattle cat catalog", "c c c"},
		{"prefix path is not a root", []string{"apple"}, "app apples application", "app apple application"},
		{"no match", []string{"cat", "dog"}, "bird fish", "bird fish"},
		{"exact word", []string{"cat"}, "cat", "cat"},
		{"repeated words", []string{"rat"}, "rattle rattle rat", "rat rat rat"},
		{"alphabet boundaries", []string{"z", "a"}, "zebra apple middle", "z a middle"},
		{"root only at beginning", []string{"cat"}, "bobcat scatter", "bobcat scatter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.dictionary, tt.sentence); got != tt.want {
				t.Errorf("replaceWords(%q, %q) = %q, want %q", tt.dictionary, tt.sentence, got, tt.want)
			}
		})
	}
}
