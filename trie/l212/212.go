package l212

func findWords(board [][]byte, words []string) []string {
	answer := []string{}
	if len(board) == 0 || len(board[0]) == 0 {
		return answer
	}
	root := NewTrieNode[string]()
	for _, word := range words {
		if word != "" {
			TrieWalk(root, word, true, nil).Data = word
		}
	}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var visit func(int, int, *TrieNode[string])
	visit = func(r, c int, node *TrieNode[string]) {
		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || visited[r][c] {
			return
		}
		node = node.Next[board[r][c]]
		if node == nil {
			return
		}
		if node.Data != "" {
			answer = append(answer, node.Data)
			node.Data = ""
		}
		visited[r][c] = true
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			visit(r+d[0], c+d[1], node)
		}
		visited[r][c] = false
	}
	for r := range board {
		for c := range board[r] {
			visit(r, c, root)
		}
	}
	return answer
}

// TrieNode 的 Data 保存题目需要的终止、计数或聚合信息。
type TrieNode[V any] struct {
	Next map[byte]*TrieNode[V]
	Data V
}

func NewTrieNode[V any]() *TrieNode[V] { return &TrieNode[V]{Next: make(map[byte]*TrieNode[V])} }

// TrieWalk 沿前缀走树；create 决定是否建点，onNode 可聚合沿途信息。
func TrieWalk[V any](node *TrieNode[V], key string, create bool, onNode func(*TrieNode[V])) *TrieNode[V] {
	if node == nil {
		return nil
	}
	if onNode != nil {
		onNode(node)
	}
	for i := 0; i < len(key); i++ {
		child := node.Next[key[i]]
		if child == nil {
			if !create {
				return nil
			}
			child = NewTrieNode[V]()
			node.Next[key[i]] = child
		}
		node = child
		if onNode != nil {
			onNode(node)
		}
	}
	return node
}
