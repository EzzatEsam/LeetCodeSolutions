package main

type Node struct {
	children [26]*Node
	isEnd bool
}

type Trie struct {
    head *Node
}


func Constructor() Trie {
    return Trie{head: &Node{}}
}


func (this *Trie) Insert(word string)  {
    currentHead := this.head
	for _, c := range word {
		if currentHead.children[c-'a'] == nil  {
			currentHead.children[c-'a'] = &Node{}
		}
		currentHead = currentHead.children[c-'a']
	}
	currentHead.isEnd = true
}


func (this *Trie) Search(word string) bool {
    currentHead := this.head
	for _, c := range word {
		if currentHead.children[c-'a'] == nil  {
			return false
		}
	}
	return currentHead.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    currentHead := this.head
	for _, c := range prefix {
		if currentHead.children[c-'a'] == nil  {
			return false
		}
		currentHead = currentHead.children[c-'a']
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