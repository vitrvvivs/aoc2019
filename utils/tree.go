package utils

type Tree struct {
	Index map[string]*Node
	Root  *Node
}

type Node struct {
	Parent   *Node
	Children []*Node
	Data     interface{}
}

func NewTree() (t *Tree) {
	t.Index = make(map[string]*Node)
	return t
}

func (t *Tree) Get(key string) *Node {
	n, found := t.Index[key]
	if !found {
		t.Index[key] = &Node{}
	}
	return n
}

func (t *Tree) IterateDepth() <-chan *Node {
	ch := make(chan *Node)

	var dfs func(*Node)
	dfs = func(n *Node) {
		ch <- n
		for _, child := range (*n).Children {
			dfs(child)
		}
	}

	go func() {
		dfs(t.Root)
		close(ch)
	}()

	return ch
}

func (t *Tree) IterateBreadth() <-chan *Node {
	ch := make(chan *Node)

	var bfs func(*Node)
	bfs = func(n *Node) {
		for _, child := range (*n).Children {
			ch <- child
		}
		for _, child := range (*n).Children {
			bfs(child)
		}
	}

	go func() {
		ch <- t.Root
		bfs(t.Root)
		close(ch)
	}()

	return ch
}
