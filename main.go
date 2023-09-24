package main

import (
	"fmt"
	"strings"
)

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

// Insert a word into the Trie
func (t *Trie) Insert(word string) {
	word = stripWord(word)
	// insert starts from the root
	node := t.root
	// loop through each character in the word
	for _, char := range word {
		// if a the current character does not already exists as a child of the current node
		if _, exists := node.children[char]; !exists {
			// make the character a child and create a map to record its children
			node.children[char] = &Node{children: make(map[rune]*Node)}
		}
		// move down the tree
		node = node.children[char]
	}
	// outside the loop traversal, we have the leaf
	node.isEnd = true
}

// Search for a word in a Trie
func (t *Trie) Search(word string) bool {
	word = stripWord(word)
	// start from the root
	node := t.root
	// loop through each character in the word
	for _, char := range word {
		// if the current character is not a child in the trie
		if _, exists := node.children[char]; !exists {
			// the word cannot possibly be in the trie
			return false
		}
		// move one level down the trie
		node = node.children[char]
	}
	// if we get to a leaf, then the word exists in the trie
	return node.isEnd
}

func (t *Trie) Delete(word string) {
	word = stripWord(word)
}

func (t *Trie) Print() {
    t.printRecursive(t.root, "")
}

func (t *Trie) printRecursive(node *Node, currentPrefix string) {
    if node == nil {
        return
    }

    fmt.Printf("(%s, %v) -> ", currentPrefix, node.isEnd)

    for char, childNode := range node.children {
        t.printRecursive(childNode, currentPrefix+string(char))
		fmt.Println()
    }
}

func stripWord(word string) string {
	return strings.ToLower(strings.ReplaceAll(word, " ", ""))
}

func main() {
	trie := NewTrie()

	words := []string{"apple", "app", "banana", "bat", "cat"}
	for _, word := range words {
		trie.Insert(word)
	}

	// fmt.Println("apple: \t\t", trie.Search("apple"))
	// fmt.Println("app: \t\t",trie.Search("app"))
	// fmt.Println("ap: \t\t", trie.Search("ap"))
	// fmt.Println("banana: \t", trie.Search("banana"))
	// fmt.Println("ban: \t\t", trie.Search("ban"))
	// fmt.Println("cat: \t\t",trie.Search("cat"))
	// fmt.Println("dog: \t\t", trie.Search("dog"))

	trie.Print()
}
