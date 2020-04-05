package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	// Call Walk() function
	t := tree.New(1)
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
