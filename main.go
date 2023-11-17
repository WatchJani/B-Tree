package main

import (
	"fmt"
	"root/b_tree"
)

func main() {
	tree := b_tree.NewBTree(5, 5)

	tree.Insert([]byte("1"))
	tree.Insert([]byte("2"))
	tree.Insert([]byte("3"))
	tree.Insert([]byte("4"))
	tree.Insert([]byte("5"))
	tree.Insert([]byte("6"))

	fmt.Println(tree.GetMemory()[0])
	fmt.Println(tree.GetMemory()[1])
	fmt.Println(tree.GetMemory()[2])
	// fmt.Println(tree.GetMemory()[0].GetLink())

	// node := b_tree.NewNode(5, nil)

	// b_tree.ShiftAndInsert(node.GetKeys(), []b_tree.Key{b_tree.NewKey([]byte("yes")), b_tree.NewKey([]byte("no"))}, 0)

	// fmt.Println(node)
}
