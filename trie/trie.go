package trie

type Node struct {
    char     byte
    children *[26]*Node
    isEnd    bool
}

type Trie struct {
    root *Node
}

func NewTrie() *Trie {
    n := &Node{}
    t := &Trie{root:n}
    return t
}

func toEnd(n *Node) [][]byte {
    m := make([][]byte,0)
    for _, c := range n.children {
        if c == nil {
            continue
        }
        ends := c.toEnd()
        for _, e := range ends {
            e2 := make([]byte, len(e)+1, len(e)+1)
            e2[0] = n.char
            copy(e2[1:], e[0:])
            m = append(m, e2)
        }
    }
    return m
}

func (n *Node) toEnd() [][]byte {
    if n.children == nil {
        return [][]byte{[]byte{n.char}}
    }else if n.isEnd{
        m := make([][]byte,0)
        m = append(m,[]byte{n.char})
        ends := toEnd(n)
        newEnds := make([][]byte,len(ends) + 1,cap(ends) + 1)
        newEnds[0] = []byte{n.char}
        copy(newEnds[1:],ends[:])
        return newEnds
    }
    return toEnd(n)
}

func (t *Trie) Insert(s string) {
    n, p := len(s)-1, t.root

    for k, b := range s {
        if p.children == nil {
            p.children = &[26]*Node{}
        }

        if p.children[b-'a'] == nil {
            p.children[b-'a'] = &Node{
                char:  byte(b),
                isEnd: k == n,
            }
        }else if k == n {
            p.children[b - 'a'].isEnd = true
        }
        p = p.children[b-'a']
    }
}

func (t *Trie) FindBytes(prefix string) ([][]byte, bool) {
    if prefix == "" {
        return nil, false
    }

    p := *t.root
    for _, b := range prefix {
        if p.children[b-'a'] == nil {
            return nil, false
        }
        p = *p.children[b-'a']
    }

    ret := p.toEnd()
    for k, v := range ret {
        pre := prefix[:len(prefix)-1]
        np := len(pre)
        n := np + len(v)
        cp := make([]byte, n, n)
        copy(cp[:], []byte(pre))
        copy(cp[np:], v)
        ret[k] = cp
    }

    return ret, true
}

func (t *Trie) FindString(prefix string) ([]string, bool) {
    if bs, ok := t.FindBytes(prefix); ok {
        ret := make([]string, len(bs))
        for k, v := range bs {
            ret[k] = string(v)
        }
        return ret, true
    } else {
        return nil, false
    }
}
