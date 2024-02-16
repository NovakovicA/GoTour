package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int, treeRoot bool) {
	if t.Left != nil {
		Walk(t.Left, ch, false)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch, false)
	}

	if treeRoot {
		close(ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var chan1, chan2 chan int
	chan1 = make(chan int)
	chan2 = make(chan int)

	go Walk(t1, chan1, true)
	go Walk(t2, chan2, true)

	for  {
		i1, open1 := <- chan1
		i2, open2 := <- chan2

		if (open1 && !open2) || (!open1 && open2) || (i1 != i2) {
			return false
		}

		if (!open1 && !open2) {
			return true;
		}
	}
}

func main() {
	var testChan1 chan int = make(chan int)

	go Walk(tree.New(1), testChan1, true)

	fmt.Println("Testing Walk func:")

	for i := range testChan1 {
		fmt.Println(i)
	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("Testing Same func:")
	fmt.Println("tree.New(1) == tree.New(1)", Same(tree.New(1), tree.New(1)))
	fmt.Println("tree.New(1) == tree.New(2)", Same(tree.New(1), tree.New(2)))
}
