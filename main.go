package main

import (
	"fmt"
	"root/b_tree"
)

func main() {
	tree := b_tree.NewBTree(5, 5)

	tree.Insert([]byte("001"))
	tree.Insert([]byte("002"))
	tree.Insert([]byte("003"))
	tree.Insert([]byte("004"))
	tree.Insert([]byte("005"))
	tree.Insert([]byte("006"))
	tree.Insert([]byte("007"))
	tree.Insert([]byte("008"))
	tree.Insert([]byte("009"))
	tree.Insert([]byte("010"))
	tree.Insert([]byte("011"))
	tree.Insert([]byte("012"))
	tree.Insert([]byte("013"))

	// fmt.Println(tree.GetMemory()[0].GetKey(0))
	// fmt.Println(tree.GetMemory()[0].GetKey(1))
	// fmt.Println(tree.GetMemory()[0].GetKey(2))
	// fmt.Println(tree.GetMemory()[0].GetLink())
	// fmt.Println(tree.GetMemory()[0])

	// fmt.Println(tree.GetMemory()[0].GetKey(0))
	// fmt.Println(tree.GetMemory()[0].GetKey(1))
	// fmt.Println(tree.GetMemory()[0].GetKey(2))
	// fmt.Println(tree.GetMemory()[0].GetKey(3))
	// fmt.Println(tree.GetMemory()[0].GetLink())
	// fmt.Println(tree.GetMemory()[0])

	tree.Insert([]byte("014"))
	tree.Insert([]byte("015"))
	tree.Insert([]byte("016"))
	tree.Insert([]byte("017"))
	fmt.Println(tree.GetMemory()[0])
	fmt.Println("da", tree.GetMemory()[0].GetKey(0))
	fmt.Println(tree.GetMemory()[0].GetLink())
	fmt.Println("========================================")
	fmt.Println(tree.GetMemory()[0].GetKey(0).GetKey(0))
	fmt.Println(tree.GetMemory()[0].GetKey(0).GetKey(1))
	fmt.Println(tree.GetMemory()[0].GetKey(0).GetLink())

	fmt.Println(tree.GetMemory()[0].GetLink().GetKey(0))
	fmt.Println(tree.GetMemory()[0].GetLink().GetKey(1))
	fmt.Println(tree.GetMemory()[0].GetLink().GetLink())

	fmt.Println([]byte("6"))
}
