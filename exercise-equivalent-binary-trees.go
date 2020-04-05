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
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        Walk(t1, ch1)
        close(ch1)
    }()
    go func() {
        Walk(t2, ch2)
        close(ch2)
    }()

    for {
        v1, close1 := <- ch1
        v2, close2 := <- ch2
        if v1 != v2 {
            return false
        }
        if !close1 && !close2 {
            return true
        }
    }

    return false
}

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

       // Call Same() function
       fmt.Println("tree1 eq tree1 = ", Same(tree.New(1), tree.New(1)))
       fmt.Println("tree1 eq tree2 = ", Same(tree.New(1), tree.New(2)))
}
