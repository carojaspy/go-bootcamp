package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Walking From Left->Center ->Right
	// fmt.Println("Walk function")
	if t.Left != nil {
		// fmt.Println("moving to left node")
		Walk(t.Left, ch)
	}

	// Adding Current value
	ch <-t.Value
	//fmt.Println(t.Value)

	//  Trying to move to Right Node
	if t.Right != nil {
		// fmt.Println("moving to Right node")
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	values_channel := make(chan int)
	myTree := tree.New(1)
	fmt.Println(myTree)

	go Walk(myTree, values_channel)
	// values_channel <- 1

	// Infinity Loop
	for i:=0; i<10; i++ {
	// for {
		if m, ok := <-values_channel; ok {
		fmt.Println("from channel: ", m)
		} else {
			fmt.Println("Closing channel...")
			close(values_channel)
			break
		}//Endif
	}//EndFor
	close(values_channel)


}//End main


