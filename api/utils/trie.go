package utils

type Trie struct {
	Children    map[byte]*Trie
	IsEndOfWord bool
}

func NewTrie() *Trie {
	return &Trie{
		Children:    make(map[byte]*Trie),
		IsEndOfWord: false,
	}
}

func (obj *Trie) Insert(word string) {

	var currNode *Trie = obj
	var currChar byte

	for i := 0; i < len(word); i++ {

		currChar = word[i]
		var child *Trie = currNode.getChild(currChar)

		if child == nil {
			var newNode *Trie = NewTrie()

			currNode.setChild(currChar, newNode)
			child = newNode

		}

		if i == len(word)-1 {
			child.IsEndOfWord = true
		}

		currNode = child

	}

}

func (obj *Trie) Check(word string) bool {

	var currNode *Trie = obj
	var currChar byte

	for i := 0; i < len(word); i++ {

		currChar = word[i]
		var child *Trie = currNode.getChild(currChar)

		if child == nil {
			return false
		}

		currNode = child

	}

	return true

}

func (obj *Trie) ValidateStringPrefix(word string) bool {

	var currNode *Trie = obj
	var currChar byte
	var isValid bool = false

	for i := 0; i < len(word); i++ {

		currChar = word[i]

		var child *Trie = currNode.getChild(currChar)

		if child == nil {
			break
		}

		currNode = child

		if currNode.IsEndOfWord {
			isValid = true
		}

	}

	return isValid

}

func (obj *Trie) getChild(char byte) *Trie {

	var child *Trie = nil
	var isChildExists bool = false

	child, isChildExists = obj.Children[char]

	if !isChildExists {
		return nil
	}

	return child

}

func (obj *Trie) setChild(char byte, child *Trie) {
	obj.Children[char] = child
}
