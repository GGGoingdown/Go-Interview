# Trie

```go
const AlphabetSize = 26

type Node struct {
	isWord bool
	next   [AlphabetSize]*Node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(word string) {
	// Get the root node of the trie
	curNode := t.root
	// Iterate over the characters in the word to add
	for _, c := range word {
		// Get the index of the character in the alphabet
		idx := c - 'a'
		// If the node at the current index is nil, create a new node
		if curNode.next[idx] == nil {
			curNode.next[idx] = &Node{}
		}
		// Move to the next node
		curNode = curNode.next[idx]
	}
	// Mark the last node as a word
	curNode.isWord = true
}

// Search searches a word in the trie.
func (t *Trie) Search(word string) bool {
	// current node is the root node
	curNode := t.root
	// iterate through each character in the word
	for _, c := range word {
		// get the index of the current character
		idx := c - 'a'
		// if the current node doesn't have the current character
		// as a child, return false
		if curNode.next[idx] == nil {
			return false
		}
		// update the current node to the child node
		curNode = curNode.next[idx]
	}
	// after iterating through all characters in the word,
	// return true if the current node is a word
	return curNode.isWord
}
```