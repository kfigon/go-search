package main

type node struct {
	children map[rune]*node
	isWord   bool
}

func newNode() *node {
	return &node{
		children: map[rune]*node{},
	}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) Add(word string) {
	currentNode := t.root

	for i, c := range word {
		if next, ok := currentNode.children[c]; ok {
			currentNode = next
		} else {
			new := newNode()
			currentNode.children[c] = new
			currentNode = new
		}

		if i == len(word)-1 {
			currentNode.isWord = true
		}
	}
}

func (t *Trie) FindPrefixed(prefix string) []string {
	if prefix == "" {
		return []string{}
	}

	prefixNode := t.root
	for _, c := range prefix {
		if next, ok := prefixNode.children[c]; ok {
			prefixNode = next
		} else {
			return []string{}
		}
	}

	return t.traverse(prefix, prefixNode)
}

func (t *Trie) findall() []string {
	return t.traverse("", t.root)
}

func (t *Trie) traverse(prefix string, prefixNode *node) []string {
	out := []string{}
	var dfs func(string, *node)
	dfs = func(prefixedWord string, n *node) {
		if n == nil {
			return
		} else if n.isWord {
			out = append(out, prefixedWord)
		}

		for c, next := range n.children {
			dfs(prefixedWord+string(c), next)
		}
	}

	dfs(prefix, prefixNode)
	return out
}
