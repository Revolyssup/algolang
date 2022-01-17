package main

import (
	"fmt"

	"github.com/Revolyssup/goalgo/tree"
)

func main() {
	bt := tree.NewBinary(1, 1)
	bt.Root.InsertLeft(12)
	bt.Root.InsertRight(12)
	bt.Root.Left.InsertLeft(15)
	bt.Root.Left.InsertRight(16)
	bt.Traverse(func(n *tree.Node) error {
		fmt.Println("val: ", n.Value)
		return nil
	}, tree.LRC) //LRC means Left Right Centre.(Postorder)
}
