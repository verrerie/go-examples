// Question: Implement a trie supporting insertion, exact-word membership, and
// prefix membership for valid UTF-8 strings, without Unicode normalization.
// The zero value must be usable. Empty words may be inserted; the empty prefix
// always exists, while the empty word exists only after insertion.

package main

import "fmt"

type Trie struct {
	children map[rune]*Trie
	terminal bool
}

// Insert maintains one root-to-node path for each rune prefix.
// terminal distinguishes full words from prefixes. The zero-value Trie is usable.
// Time O(L) expected per operation; total storage O(total inserted runes).
// Strings must be valid UTF-8. No Unicode normalization is performed.
func (t *Trie) Insert(word string) {
	node := t
	for _, r := range word {
		if node.children == nil {
			node.children = map[rune]*Trie{}
		}
		if node.children[r] == nil {
			node.children[r] = &Trie{}
		}
		node = node.children[r]
	}
	node.terminal = true
}
func (t *Trie) find(s string) *Trie {
	node := t
	for _, r := range s {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node
}
func (t *Trie) Contains(word string) bool    { node := t.find(word); return node != nil && node.terminal }
func (t *Trie) HasPrefix(prefix string) bool { return t.find(prefix) != nil }
func main() {
	var t Trie
	t.Insert("tree")
	fmt.Println(t.Contains("tree"), t.Contains("tr"), t.HasPrefix("tr"))
}
