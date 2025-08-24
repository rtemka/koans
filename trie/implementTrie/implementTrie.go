// #208
// https://leetcode.com/problems/implement-trie-prefix-tree/description/
package leetcode

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func Constructor() Trie {
	return Trie{root: newTrieNode()}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		// For each char in the word, if it's not a child
		// of the current node, create a new TrieNode for that character.
		child, ok := node.children[c]
		if !ok {
			child = newTrieNode()
			node.children[c] = child
		}
		node = child
	}
	// Mark the last node as the end of a word.
	node.isWord = true
}

func (this *Trie) Search(word string) bool {
	node, ok := this.root, false
	for _, c := range word {
		// For each char in the word, if it's not a child of
		// the current node, the word doesn't exist in the Trie.
		node, ok = node.children[c]
		if !ok {
			return false
		}
	}
	// Return whether the current node is marked as the end of the word.
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node, ok := this.root, false
	for _, c := range prefix {
		node, ok = node.children[c]
		if !ok {
			return false
		}
	}
	// Once we've traversed the nodes corresponding to each character in the prefix, return true.
	return true
}
