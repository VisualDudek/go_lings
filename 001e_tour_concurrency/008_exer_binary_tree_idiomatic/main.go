// Idiomatic usage of defer in Walk fn.
package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // Ensure channel is closed when Walk returns

	var walk func(*tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value // In-order traversal
		walk(t.Right)
	}

	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for v := range ch {
		println(v)
	}

}
