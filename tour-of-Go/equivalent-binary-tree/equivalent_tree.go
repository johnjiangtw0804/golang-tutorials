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
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch)
		close(ch)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	var path, path2 string

	// ready to receive
	for x := range ch {
		path = path + fmt.Sprint(x)
	}
	for x := range ch2 {
		path2 = path2 + fmt.Sprint(x)
	}
	return path == path2
}

func main() {
	new_tree := tree.New(1)
	ch := make(chan int)
	go func() {
		Walk(new_tree, ch)
		close(ch)
	}()
	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

	// should return true
	fmt.Println(Same(tree.New(1), tree.New(1)))

	// should return false
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
