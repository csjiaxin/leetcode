package main

import "fmt"

type Trie struct {
	end    bool
	child  map[byte]*Trie
	c      byte
	parent *Trie
}

func Constructor() Trie {
	t := Trie{
		parent: nil,
		end:    false,
		child:  make(map[byte]*Trie),
	}
	return t
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if v, ok := node.child[c]; ok {
			node = v
		} else {
			tnode := Constructor()
			tnode.parent = node
			tnode.c = c
			node.child[c] = &tnode
			node = &tnode
		}
	}
	node.end = true
}
func (this *Trie) find(word string) *Trie {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := node.child[c]; !ok {
			return nil
		}
		node = node.child[c]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.find(word)
	if node != nil {
		return node.end
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.find(prefix)
	if node != nil {
		return true
	}
	return false
}

func (this *Trie) build() string {
	ret := ""
	n := this
	for n != nil {
		if n.c != 0 {
			ret = string(n.c) + ret
		}
		n = n.parent
	}
	return ret
}
func pre(p *Trie, suggestion int) int {
	if p != nil {
		for _, v := range p.child {
			suggestion = pre(v, suggestion)
		}
		if len(p.child) == 0 {
			v := p.build()
			suggestion += len(v) + 1
		}
	}
	return suggestion

}
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
func minimumLengthEncoding(words []string) int {
	obj := Constructor()
	for i := 0; i < len(words); i++ {
		p := reverse(words[i])
		obj.Insert(p)
	}

	return pre(&obj, 0)
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}
