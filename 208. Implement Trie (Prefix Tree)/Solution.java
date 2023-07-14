class TrieNode {
    public TrieNode[] children;
    public boolean isEnd = false;
    public TrieNode() {
        children = new TrieNode[26];
    }
}

class Trie {

    private TrieNode root ;
    public Trie() {
        root = new TrieNode();
    }
    
    public void insert(String word) {
        var currentHead = root;
        for (int i = 0; i < word.length(); i++) {
            var idx = word.charAt(i) - 'a';
            if (currentHead.children[idx] == null) {
                currentHead.children[idx] = new TrieNode();
            } 
            currentHead = currentHead.children[idx];   
        }
        currentHead.isEnd = true;
        
    }
    
    public boolean search(String word) {
        var currentHead = root;
        for (int i = 0; i < word.length(); i++) {
            var idx = word.charAt(i) - 'a';
            if (currentHead.children[idx] == null) {
                return false;
            } 
            currentHead = currentHead.children[idx];   
        }
        return currentHead.isEnd;
    }
    
    public boolean startsWith(String word) {
        var currentHead = root;
        for (int i = 0; i < word.length(); i++) {
            var idx = word.charAt(i) - 'a';
            if (currentHead.children[idx] == null) {
                return false;
            } 
            currentHead = currentHead.children[idx];   
        }
        return true;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */