// 2020/4/6
package main

import "fmt"

func main() {
	trie := new(Trie)
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(!trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		index := c - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}
		curr = curr.children[index]
	}
	curr.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		index := c - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		index := c - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
