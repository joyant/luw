package trie

import (
    "fmt"
    "testing"
)

func loop(n *Node)  {
    if n.children != nil {
        for _,c := range n.children {
            if c != nil {
                loop(c)
            }
        }
    }
}

func TestTrie_Insert(t *testing.T) {
    n := &Node{}
    trie := Trie{root:n}

    trie.Insert("about")
    trie.Insert("above")
    trie.Insert("ability")
    trie.Insert("ab")
    trie.Insert("abo")
    trie.Insert("abou")

    r1,ok := trie.FindString("ab")
    fmt.Println(r1,ok)
}
