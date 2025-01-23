package datastructures

type CharTrie struct {
	RootNodes map[uint8]*CharTrieNode
}

type CharTrieNode struct {
	Char      uint8
	Children  map[uint8]*CharTrieNode
	Parent    *CharTrieNode
	Depth     int
	IsWordEnd bool
}

func NewCharTrie() *CharTrie {
	return &CharTrie{}
}

func (t *CharTrie) Insert(word string) {
	if t.RootNodes == nil {
		t.RootNodes = make(map[uint8]*CharTrieNode)
	}

	if len(word) == 0 {
		return
	}

	if t.RootNodes[word[0]] == nil {
		isWordEnd := len(word) == 1
		t.RootNodes[word[0]] = &CharTrieNode{
			Char:      word[0],
			Parent:    nil,
			Depth:     1,
			IsWordEnd: isWordEnd,
		}
	}

	currentNode := t.RootNodes[word[0]]
	for i := 1; i < len(word); i++ {
		if currentNode.Children == nil {
			currentNode.Children = make(map[uint8]*CharTrieNode)
		}

		if currentNode.Children[word[i]] == nil {
			currentNode.Children[word[i]] = &CharTrieNode{
				Parent:    currentNode,
				Char:      word[i],
				Depth:     currentNode.Depth + 1,
				IsWordEnd: false,
			}
		}

		currentNode = currentNode.Children[word[i]]
	}

	currentNode.IsWordEnd = true
}

func (t *CharTrie) FindLastPossibleChar(word string) *CharTrieNode {
	if t.RootNodes == nil || len(word) == 0 {
		return nil
	}

	currentNode := t.RootNodes[word[0]]
	if currentNode == nil {
		return nil
	}

	for i := 1; i < len(word); i++ {
		node := currentNode.Children[word[i]]
		if node == nil {
			break
		}

		currentNode = node
	}

	for currentNode != nil && !currentNode.IsWordEnd {
		currentNode = currentNode.Parent
	}

	return currentNode
}
