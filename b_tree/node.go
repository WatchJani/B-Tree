package b_tree

import "fmt"

type Node struct {
	capacity int
	keys     []Key
	link     *Node // next node, last right link node -> [[][][][][x]]
	parent   *Node //parent node
}

func NewNode(degree int, next *Node) *Node {
	return &Node{
		keys: make([]Key, degree), //degree-1
		link: next,
	}
}

func (n *Node) IncreaseCapacity(value int) {
	n.capacity += value
}

func (n *Node) DecreaseCapacity(value int) {
	n.capacity -= value
}

// popraviti funkciju tako da vraca, node, value, node
func (n Node) SplitNode(degree int) ([]Key, Key, []Key) {
	splitting := degree / 2 //split node on three pears
	return n.keys[:splitting], n.keys[splitting], n.keys[splitting+1:]
}

func (n *Node) Insert(element []byte, index int) {
	ShiftAndInsert(&n.keys, []Key{NewKey(element)}, index)
	n.IncreaseCapacity(1)
}

func (n *Node) InsertKey(value []byte) int {
	index := SearchInsert(n.keys, value, n.capacity-1)
	n.Insert(value, index)

	return index
}

// naci univerzalniji nacin
func (n *Node) Split(tree *Tree) {
	rep := 0
	for {

		left, root, right := n.SplitNode(tree.degree)

		tree.memory = append(tree.memory, NewNode(tree.degree, nil)) //add new root node
		leftNode := tree.memory[len(tree.memory)-1]                  //my left Node

		leftNode.keys = Append(left)
		// leftNode.keys = left
		leftNode.IncreaseCapacity(tree.degree / 2) //append keys

		// n.keys = right //big bug, need to be capacity 5 not 2
		n.keys = Append(right) //big bug, need to be capacity 5 not 2
		n.DecreaseCapacity(tree.degree/2 + 1)

		if n.parent != nil {
			index := n.parent.InsertKey(root.value)
			n.parent.keys[index].link = leftNode
		} else {
			if rep == 1 {
				fmt.Println("yes")
				leftNode.link = root.link
			}
			tree.memory = append(tree.memory, NewNode(tree.degree, nil)) //add new root node
			tree.memory[len(tree.memory)-1].keys[0] = NewKey(root.value) //add new element in root node
			tree.memory[len(tree.memory)-1].IncreaseCapacity(1)
			tree.memory[len(tree.memory)-1].keys[0].link = leftNode //link parent with left node
			// tree.memory[len(tree.memory)-1], tree.memory[0] = tree.memory[0], tree.memory[len(tree.memory)-1] //need to be function
			tree.memory[len(tree.memory)-1].link = n
			tree.SwapRoot() //just in case when we need to create new root node!
			tree.memory[len(tree.memory)-1].parent = tree.memory[0]
		}
		leftNode.parent = tree.memory[0]

		if tree.memory[0].capacity < tree.degree {
			break
		} else {
			n = tree.memory[0]
			rep++
			fmt.Println("PONOVI")
		}
	}
}

func (n Node) GetLink() *Node {
	return n.link
}

func (n Node) GetKeys() *[]Key {
	return &n.keys
}

func (n Node) GetKey(index int) *Node {
	return n.keys[index].GetLink()
}
