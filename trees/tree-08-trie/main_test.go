package main

import "testing"

func TestTrie(t *testing.T) {
	var tr Trie
	if tr.Contains("") || !tr.HasPrefix("") {
		t.Fatal("empty trie")
	}
	for _, w := range []string{"tree", "trie", "树木", "🙂a", "tree", ""} {
		tr.Insert(w)
	}
	for _, w := range []string{"tree", "trie", "树木", "🙂a", ""} {
		if !tr.Contains(w) {
			t.Fatal(w)
		}
	}
	for _, w := range []string{"tr", "树", "🙂"} {
		if tr.Contains(w) || !tr.HasPrefix(w) {
			t.Fatal(w)
		}
	}
	if tr.HasPrefix("x") || tr.Contains("trees") {
		t.Fatal("missing word")
	}
}
